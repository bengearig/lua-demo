package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	lua "github.com/yuin/gopher-lua"
	"lua-test/tiles"
)

type Player struct {
	Position tiles.Vector
	Direction
	Tilemap    *tiles.Tilemap
	InstanceId int
	L          *lua.LState
	Alive      bool
}

func (p *Player) GetDirection() Direction {
	return p.Direction
}

func (p *Player) IsAlive() bool {
	return p.Alive
}

func (p *Player) GetPosition() tiles.Vector {
	return p.Position
}

func (p *Player) GetInstanceId() int {
	return p.InstanceId
}

func (p *Player) Update() error {
	if err := p.L.DoFile("scripts/update.lua"); err != nil {
		panic(err)
	}

	switch p.Direction {
	case East:
		if p.Tilemap.GetTile(p.Position.X+1, p.Position.Y) == "none" {
			p.Position.X++
		}
	case North:
		if p.Tilemap.GetTile(p.Position.X, p.Position.Y-1) == "none" {
			p.Position.Y--
		}
	case West:
		if p.Tilemap.GetTile(p.Position.X-1, p.Position.Y) == "none" {
			p.Position.X--
		}
	case South:
		if p.Tilemap.GetTile(p.Position.X, p.Position.Y+1) == "none" {
			p.Position.Y++
		}
	case None:
	default:
		break
	}

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Tilemap.Tileset.DrawTile(screen, "player", p.Position.X, p.Position.Y)
}
