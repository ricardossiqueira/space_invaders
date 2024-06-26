package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mattn/go-tty"
	"golang.org/x/term"
)

type Coord struct {
	X, Y int
}

type Size struct {
	W, H int
}

type Cli struct {
	canvas  *bufio.Writer
	Size    Size
	Update  *time.Ticker
	EventCh chan string
}

const (
	clearCli   = "\033[2J"
	hideCursor = "\033[?25l"
	showCursor = "\033[?25h"
	moveCursor = "\033[%d;%dH"
)

const (
	left  = 'D'
	right = 'C'
	up    = 'A'
	down  = 'B'
)

const (
	Black  = "\033[1;30m%s\033[0m"
	Red    = "\033[1;31m%s\033[0m"
	Green  = "\033[1;32m%s\033[0m"
	Yellow = "\033[1;32m%s\033[0m"
	Blue   = "\033[1;34m%s\033[0m"
	Purple = "\033[1;35m%s\033[0m"
	Cyan   = "\033[1;36m%s\033[0m"
	White  = "\033[1;37m%s\033[0m"
)

var Colors = []string{Black, Red, Green, Yellow, Blue, Purple, Cyan, White}

func New(tick time.Duration) (*Cli, error) {
	canvas := bufio.NewWriter(os.Stdout)
	w, h, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	update := time.NewTicker(time.Millisecond * tick)
	ch := make(chan string, 2)

	return &Cli{canvas: canvas, Size: Size{W: w, H: h}, Update: update, EventCh: ch}, nil
}

// Handle cursor and cli functions
func (c *Cli) ClearCli() {
	fmt.Fprint(c.canvas, clearCli)
}

func (c *Cli) HideCursor() {
	fmt.Fprint(c.canvas, hideCursor)
}

func (c *Cli) ShowCursor() {
	fmt.Fprint(c.canvas, showCursor)
}

func (c *Cli) MoveCursor(p Coord) {
	fmt.Fprintf(c.canvas, moveCursor, p.Y, p.X)
}

func (c *Cli) Render() {
	c.canvas.Flush()
}

func (c *Cli) HandleSIGTERM(f ...func()) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		for _, fn := range f {
			fn()
		}
		os.Exit(1)
	}()
}

// Print functions
func (c *Cli) Fprintf(p Coord, s string, a ...any) {
	c.MoveCursor(p)
	fmt.Fprintf(c.canvas, s, a...)
}

func (c *Cli) Fprint(p Coord, a ...any) {
	c.MoveCursor(p)
	fmt.Fprint(c.canvas, a...)
}

// Color print functions
func (c *Cli) ColorFprintf(p Coord, s string, color string, a ...any) {
	c.MoveCursor(p)
	fmt.Fprintf(c.canvas, color, fmt.Sprintf(s, a...))
}

func (c *Cli) ColorFprint(p Coord, s string, color string) {
	c.MoveCursor(p)
	fmt.Fprintf(c.canvas, color, s)
}

// Sprite print functions
func (c *Cli) ColorFprintfSprite(s []string, color string, p Coord) {
	for i, sp := range s {
		p.Y += i
		c.ColorFprint(p, sp, color)
	}
}

// Input functions
func (c *Cli) HandleInput() {
	tty, err := tty.Open()
	if err != nil {
		panic(err)
	}
	defer tty.Close()

	for {
		<-c.Update.C
		char, _ := tty.ReadRune()
		switch char {
		case left:
			c.EventCh <- "left"
		case right:
			c.EventCh <- "right"
		case up:
			c.EventCh <- "up"
		case down:
			c.EventCh <- "down"
		}
	}
}
