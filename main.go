package main

import (
	"fmt"
	"math/rand"
	cli "space_invaders/Cli"
	spaceship "space_invaders/Game"
	"time"
)

func stress(c *cli.Cli) {
	for y := 0; y < c.Size.H; y++ {
		for x := 0; x < c.Size.W; x++ {
			c.ColorFprintfSprite([]string{"#"}, cli.Colors[rand.Intn(len(cli.Colors))], cli.Coord{X: x, Y: y})
		}
	}
}

func main() {
	// init
	c, err := cli.New(5)
	if err != nil {
		panic(fmt.Sprintf("cli.New(): failed to init cli handler, reason: %s", err))
	}
	ticker := time.NewTicker(time.Millisecond * 10) // render ratio

	// handle ctrl+c
	c.HandleSIGTERM(func() { c.MoveCursor(cli.Coord{X: 0, Y: 0}) }, c.ShowCursor)

	// setup
	c.ClearCli()
	c.HideCursor()

	ss := spaceship.New(cli.Coord{X: c.Size.W / 2, Y: c.Size.H - 3}, cli.Green)

	// go handleInput(c.EventCh, *ticker)
	go c.HandleInput()

	for range ticker.C {
		<-ticker.C
		c.ClearCli()

		// update
		switch <-c.EventCh {
		case "right":
			ss.MoveRight()
		case "left":
			ss.MoveLeft()
		}

		// ss.MoveRight()

		ss.Draw(c)
		// draw
		c.Render()
	}

	// end
	c.ShowCursor()

}
