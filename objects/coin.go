package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"lua-test/tiles"
)

type Coin struct {
	Position tiles.Vector
	tiles.Tileset
	InstanceId int
	Alive      bool
}

func (c *Coin) GetDirection() Direction {
	return None
}

func (c *Coin) IsAlive() bool {
	return c.Alive
}

func (c *Coin) GetPosition() tiles.Vector {
	return c.Position
}

func (c *Coin) GetInstanceId() int {
	return c.InstanceId
}

func (c *Coin) Update() error {
	return nil
}

func (c *Coin) Draw(screen *ebiten.Image) {
	c.Tileset.DrawTile(screen, "coin", c.Position.X, c.Position.Y)
}
