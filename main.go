package main

import (
	cli "space_invaders/Cli"
)

func main() {
	c, err := cli.New()
	if err != nil {
		panic("cli.New(): Failed to retrieve window dimensions")
	}
	c.HandleSIGTERM(c.ShowCursor, func() { c.MoveCursor(cli.Coord{X: 0, Y: 0}) })

	c.ClearCli()
	c.HideCursor()

	c.ColorFprintfSprite([]string{" ⢀⣀⣾⣷⣀⡀ ", " ⣿⣿⣿⣿⣿⣿ "}, cli.Green, cli.Coord{X: 2, Y: 5})

	c.ShowCursor()
	c.Render()
}
