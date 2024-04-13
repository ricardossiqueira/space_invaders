package main

import (
	"fmt"
	"math/rand"
	cli "space_invaders/Cli"
	spaceship "space_invaders/Spaceship"
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
	ticker := time.NewTicker(time.Millisecond * 5) // render ratio

	// handle ctrl+c
	c.HandleSIGTERM(func() { c.MoveCursor(cli.Coord{X: 0, Y: 0}) }, c.ShowCursor)

	// setup
	c.ClearCli()
	c.HideCursor()

	ss := spaceship.New(
		cli.Coord{X: (c.Size.W / 2), Y: c.Size.H - 3},
		[]string{" ⢀⣀⣾⣷⣀⡀ ", " ⣿⣿⣿⣿⣿⣿ "},
		cli.Green,
		800,
	)

	var enemies = []*spaceship.Spaceship{}
	for i := 0; i < 5; i++ {
		enemies = append(enemies, []*spaceship.Spaceship{spaceship.New(
			cli.Coord{X: c.Size.W / 5 * i, Y: 2},
			[]string{" ⢀⣀⣾⣷⣀⡀ ", " ⣿⣿⣿⣿⣿⣿ "},
			cli.Red,
			500,
		)}...)
	}

	go c.HandleInput()

	for range ticker.C {

		select {
		case <-ticker.C:
			c.ClearCli()

			for i, b := range ss.Bullets {
				b.Update(c)
				if b.Evaded(c) {
					ss.RemoveBullet(i)
				}
			}

			// update
			ss.ColideWall(c)
			for i, e := range enemies {
				for _, b := range ss.Bullets {
					colided := e.ColideBullet(b)
					if colided {
						enemies = e.RemoveEnemie(i, enemies)
					}
				}
			}

			// draw
			// ss.PrintLifes(c)
			for _, e := range enemies {
				e.Draw(c)
			}
			ss.Draw(c)
			for _, b := range ss.Bullets {
				b.Draw(c)
			}

			c.Render()

		case evt := <-c.EventCh:
			switch evt {
			case "right":
				ss.MoveRight()
			case "left":
				ss.MoveLeft()
			}

		case <-ss.ShootFreq.C:
			ss.Shoot(c)
		}

	}

	// end
	c.ShowCursor()

}
