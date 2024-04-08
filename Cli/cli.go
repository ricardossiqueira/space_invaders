package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

type Coord struct {
	X, Y int
}

type Size struct {
	W, H int
}

type Cli struct {
	canvas *bufio.Writer
	Size   Size
}

const (
	clearCli   = "\033[2J"
	hideCursor = "\033[?25l"
	showCursor = "\033[?25h"
	moveCursor = "\033[%d;%dH"
)

const (
	Green = "\033[1;32m%s\033[0m"
)

func New() (*Cli, error) {
	canvas := bufio.NewWriter(os.Stdout)
	w, h, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	return &Cli{canvas: canvas, Size: Size{
		W: w,
		H: h,
	}}, nil
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
func (c *Cli) Fprintf(s string, a ...any) {
	fmt.Fprintf(c.canvas, s, a...)
}

func (c *Cli) Fprint(a ...any) {
	fmt.Fprint(c.canvas, a...)
}

// Color print functions
func (c *Cli) ColorFprintf(s string, color string, a ...any) {
	args := append([]any{s}, a...)
	fmt.Fprintf(c.canvas, color, args)
}

func (c *Cli) ColorFprint(s string, color string) {
	fmt.Fprintf(c.canvas, color, s)
}

// Sprite print functions
func (c *Cli) ColorFprintfSprite(s []string, color string, p Coord) {
	for i, sp := range s {
		p.Y += i
		c.MoveCursor(p)
		c.ColorFprint(sp, color)
	}
}
