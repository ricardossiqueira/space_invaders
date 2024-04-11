package spaceship

import (
	cli "space_invaders/Cli"
)

type Coord struct {
	X, Y int
}

type Spaceship struct {
	Lives  int
	Coord  cli.Coord
	Color  string
	sprite []string
}

func New(coord cli.Coord, color string) *Spaceship {
	ss := &Spaceship{
		Lives:  3,
		sprite: []string{" ⢀⣀⣾⣷⣀⡀ ", " ⣿⣿⣿⣿⣿⣿ "},
		Coord:  coord,
		Color:  color,
	}
	return ss
}

func (ss *Spaceship) Draw(c *cli.Cli) {
	c.ColorFprintfSprite(ss.sprite, ss.Color, ss.Coord)
}

func (ss *Spaceship) MoveLeft() {
	ss.Coord.X--
}

func (ss *Spaceship) MoveRight() {
	ss.Coord.X++
}
