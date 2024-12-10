package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"lua-test/objects"
	"lua-test/tiles"
	"strconv"
)

type Game struct {
	tilemap      *tiles.Tilemap
	screenWidth  int
	screenHeight int
	objects      []objects.Object
	frame        int
	gameOver     bool
	gameWon      bool
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	if g.gameOver || g.gameWon {
		return nil
	}

	for _, obj := range g.objects {
		if err := obj.Update(); err != nil {
			return err
		}
	}

	for _, obj := range g.objects {
		if player, ok := obj.(*objects.Player); ok {
			for _, other := range g.objects {
				if coin, ok := other.(*objects.Coin); ok {
					if player.Position.X == coin.Position.X && player.Position.Y == coin.Position.Y {
						coin.Alive = false
					}
				} else if enemy, ok := other.(*objects.Enemy); ok {
					if player.Position.X == enemy.Position.X && player.Position.Y == enemy.Position.Y {
						player.Alive = false
						g.gameOver = true
					}
				}
			}
		}
	}

	alive := []objects.Object{}
	coinsRemaining := 0
	for _, obj := range g.objects {
		if obj.IsAlive() {
			alive = append(alive, obj)
			if _, ok := obj.(*objects.Coin); ok {
				coinsRemaining++
			}
		}
	}
	g.objects = alive

	if coinsRemaining == 0 {
		g.gameWon = true
	}

	g.frame++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.tilemap.GetImage(), nil)
	for _, obj := range g.objects {
		obj.Draw(screen)
	}
	if g.gameOver {
		ebitenutil.DebugPrint(screen, "Frame: "+strconv.Itoa(g.frame)+" You died. Hold Escape to exit.")
	} else if g.gameWon {
		ebitenutil.DebugPrint(screen, "Frame: "+strconv.Itoa(g.frame)+" You win! Hold Escape to exit.")
	} else {
		ebitenutil.DebugPrint(screen, "Frame: "+strconv.Itoa(g.frame))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}
