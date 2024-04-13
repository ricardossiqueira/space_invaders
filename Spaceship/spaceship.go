package spaceship

import (
	"fmt"
	bullet "space_invaders/Bullet"
	cli "space_invaders/Cli"
	"time"
)

type Spaceship struct {
	Lives     int
	Coord     cli.Coord
	Color     string
	sprite    []string
	Bullets   []*bullet.Bullet
	ShootFreq *time.Ticker
}

func New(coord cli.Coord, sprite []string, color string, f time.Duration) *Spaceship {
	tk := time.NewTicker(time.Millisecond * f) // player shooting freq

	ss := &Spaceship{
		Lives:     3,
		sprite:    sprite,
		Coord:     cli.Coord{X: coord.X + len(sprite)/2, Y: coord.Y},
		Color:     color,
		ShootFreq: tk,
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

func (ss *Spaceship) ColideWall(c *cli.Cli) {
	for ss.Coord.X <= 0 {
		ss.MoveRight()
	}
	for ss.Coord.X >= c.Size.W-ss.spriteW() {
		ss.MoveLeft()
	}
}

func (ss *Spaceship) Shoot(c *cli.Cli) {
	b := bullet.New(cli.Coord{Y: ss.Coord.Y - 3, X: (ss.Coord.X + ss.spriteW()/2)})
	ss.Bullets = append(ss.Bullets, b)
}

func (ss *Spaceship) RemoveBullet(i int) {
	ss.Bullets = append(ss.Bullets[:i], ss.Bullets[i+1:]...)
}

func (ss *Spaceship) ColideBullet(b *bullet.Bullet) bool {
	return b.Coord.X > ss.Coord.X &&
		b.Coord.X < ss.Coord.X+ss.spriteW() &&
		b.Coord.Y < ss.Coord.Y+2
}

func (ss *Spaceship) RemoveEnemie(i int, e []*Spaceship) []*Spaceship {
	e = append(e[:i], e[i+1:]...)
	return e
}

func (ss *Spaceship) spriteW() int {
	return len(ss.sprite[0]) / 2
}

func (ss *Spaceship) PrintLifes(c *cli.Cli) {
	c.ColorFprintf(cli.Coord{X: 1, Y: c.Size.H - 1}, cli.Blue, "Lifes: %v", fmt.Sprint(ss.Lives))
}
