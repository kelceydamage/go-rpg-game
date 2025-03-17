package animation

import (
	"image"
)

type SpriteSheet struct {
	WidthInTiles     int
	HeightInTiles    int
	TileSizeInPixels int
}

func (s *SpriteSheet) Rect(index int) image.Rectangle {
	x := (index % s.WidthInTiles) * s.TileSizeInPixels
	y := (index / s.WidthInTiles) * s.TileSizeInPixels
	return image.Rect(
		x,
		y,
		x+s.TileSizeInPixels,
		y+s.TileSizeInPixels,
	)
}

func NewSpriteSheet(widthInTiles, heightInTiles, tileSizeInPixels int) *SpriteSheet {
	return &SpriteSheet{
		WidthInTiles:     widthInTiles,
		HeightInTiles:    heightInTiles,
		TileSizeInPixels: tileSizeInPixels,
	}
}
