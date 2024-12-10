package tiles

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Tilemap struct {
	Tileset *Tileset
	tiles   [][]string
	image   *ebiten.Image
}

func (t *Tilemap) SetTiles(w int, h int, tiles [][]string) {
	t.image = ebiten.NewImage(w*t.Tileset.Size, h*t.Tileset.Size)
	t.tiles = tiles
	for y, row := range t.tiles {
		for x, tile := range row {
			t.Tileset.DrawTile(t.image, tile, x, y)
		}
	}
}

func (t *Tilemap) GetTile(x, y int) string {
	return t.tiles[y][x]
}

func (t *Tilemap) GetImage() *ebiten.Image {
	return t.image
}
