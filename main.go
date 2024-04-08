package main

import (
	"math/rand"
	cli "space_invaders/Cli"
	"time"
)

func main() {

	// init
	c, err := cli.New()
	if err != nil {
		panic("cli.New(): Failed to retrieve window dimensions")
	}

	// update
	c.HandleSIGTERM(func() { c.MoveCursor(cli.Coord{X: 0, Y: 0}) }, c.ShowCursor)

	c.ClearCli()
	c.HideCursor()

	c.ColorFprintfSprite([]string{" ⢀⣀⣾⣷⣀⡀ ", " ⣿⣿⣿⣿⣿⣿ "}, cli.Green, cli.Coord{X: 2, Y: 5})

	ticker := time.NewTicker(time.Millisecond * 200)

	for {

		select {
		case <-ticker.C:
			// draw
			c.ClearCli()
			c.Render()

			// update
			// t := time.Now()
			// c.ColorFprintfSprite([]string{t.Format(time.TimeOnly)}, cli.Green, cli.Coord{X: 1, Y: 1})
			for y := 0; y < c.Size.H; y++ {
				for x := 0; x < c.Size.W; x++ {
					c.ColorFprintfSprite([]string{"#"}, cli.Colors[rand.Intn(len(cli.Colors))], cli.Coord{X: x, Y: y})
				}
			}
		}
	}

	// end
	c.ShowCursor()

}
