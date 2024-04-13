package bullet

import (
	cli "space_invaders/Cli"
)

type Bullet struct {
	Coord  cli.Coord
	sprite []string
	Color  string
}

func New(pos cli.Coord) *Bullet {
	b := &Bullet{Coord: cli.Coord{X: pos.X, Y: pos.Y + 2}, sprite: []string{"⣿", "⡀"}, Color: cli.Green}
	return b
}

func (b *Bullet) Update(c *cli.Cli) {
	b.Coord.Y--
}

func (b *Bullet) Draw(c *cli.Cli) {
	c.ColorFprintfSprite(b.sprite, b.Color, b.Coord)
}

func (b *Bullet) Evaded(c *cli.Cli) bool {
	return b.Coord.Y < 1
}
