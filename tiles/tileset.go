package tiles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Tileset struct {
	Image   *ebiten.Image
	Size    int
	Aliases map[string]Vector
}

func (t *Tileset) DrawTile(img *ebiten.Image, tile string, x, y int) {
	srcX := t.Aliases[tile].X * t.Size
	srcY := t.Aliases[tile].Y * t.Size
	rect := image.Rect(srcX, srcY, srcX+t.Size, srcY+t.Size)
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x*t.Size), float64(y*t.Size))
	img.DrawImage(t.Image.SubImage(rect).(*ebiten.Image), options)
}
