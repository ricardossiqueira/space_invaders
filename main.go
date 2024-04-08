package main

import (
	"math/rand"
	cli "space_invaders/Cli"
	"time"
)

type gameLoop struct {
	delta       time.Duration
	lastUpdate  time.Time
	accumulator time.Duration
	slice       time.Duration
}

func stress(c *cli.Cli) {
	for y := 0; y < c.Size.H; y++ {
		for x := 0; x < c.Size.W; x++ {
			c.ColorFprintfSprite([]string{"#"}, cli.Colors[rand.Intn(len(cli.Colors))], cli.Coord{X: x, Y: y})
		}
	}
}

func main() {
	// init
	c, err := cli.New()
	if err != nil {
		panic("cli.New(): Failed to retrieve window dimensions")
	}
	gl := &gameLoop{lastUpdate: time.Now(), slice: time.Millisecond * 10} // update ratio
	ticker := time.NewTicker(time.Millisecond * 150)                      // render ratio

	// handle ctrl+c
	c.HandleSIGTERM(c.ShowCursor)

	// setup
	c.ClearCli()
	c.HideCursor()

	for {
		<-ticker.C

		gl.delta = time.Since(gl.lastUpdate)
		gl.lastUpdate = time.Now()

		// update
		for {
			if gl.accumulator > gl.slice {
				break
			}

			stress(c)
			// c.ColorFprintfSprite([]string{" ⢀⣀⣾⣷⣀⡀ ", " ⣿⣿⣿⣿⣿⣿ "}, cli.Green, cli.Coord{X: 2, Y: 5})

			// after update
			gl.accumulator -= gl.slice

		}

		// draw
		c.Render()
		c.ClearCli()
	}

	// end
	c.ShowCursor()

}
