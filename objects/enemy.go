package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"lua-test/tiles"
)

type Enemy struct {
	Position tiles.Vector
	Direction
	Tilemap    *tiles.Tilemap
	InstanceId int
}

func (e *Enemy) GetDirection() Direction {
	return e.Direction
}

func (e *Enemy) IsAlive() bool {
	return true
}

func (e *Enemy) GetPosition() tiles.Vector {
	return e.Position
}

func (e *Enemy) GetInstanceId() int {
	return e.InstanceId
}

func (e *Enemy) Update() error {
	switch e.Direction {
	case East:
		if e.Tilemap.GetTile(e.Position.X+1, e.Position.Y) == "wall" {
			e.Direction = West
			e.Position.X--
		} else {
			e.Position.X++
		}
	case North:
		if e.Tilemap.GetTile(e.Position.X, e.Position.Y-1) == "wall" {
			e.Direction = South
			e.Position.Y++
		} else {
			e.Position.Y--
		}
	case West:
		if e.Tilemap.GetTile(e.Position.X-1, e.Position.Y) == "wall" {
			e.Direction = East
			e.Position.X++
		} else {
			e.Position.X--
		}
	case South:
		if e.Tilemap.GetTile(e.Position.X, e.Position.Y+1) == "wall" {
			e.Direction = North
			e.Position.Y--
		} else {
			e.Position.Y++
		}
	}
	return nil
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	e.Tilemap.Tileset.DrawTile(screen, "enemy", e.Position.X, e.Position.Y)
}
