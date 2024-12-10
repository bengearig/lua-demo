package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"lua-test/tiles"
)

type Object interface {
	GetInstanceId() int
	GetPosition() tiles.Vector
	Update() error
	Draw(screen *ebiten.Image)
	IsAlive() bool
	GetDirection() Direction
}
