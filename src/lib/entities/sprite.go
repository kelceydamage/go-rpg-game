package entities

import (
	"image"
	"math"
	"rpg-game/src/lib/components"
	"rpg-game/src/lib/config"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image    *ebiten.Image
	X, Y     float64
	Collider components.Collider
}

func NewSprite(spritImage *ebiten.Image, x, y float64) *Sprite {
	return &Sprite{
		Image: spritImage,
		X:     x,
		Y:     y,
		Collider: image.Rect(
			int(x),
			int(y),
			int(x+config.DefaultTileSizeInPixels),
			int(y+config.DefaultTileSizeInPixels),
		),
	}
}

func (s *Sprite) GetDistanceFrom(object components.Collidable) float64 {
	return math.Sqrt(math.Pow(s.X-object.GetX()+(config.DefaultTileSizeInPixels/2), 2) + math.Pow(s.Y-object.GetY()+(config.DefaultTileSizeInPixels/2), 2))
}

func (s *Sprite) GetCollider() components.Collider {
	return s.Collider
}

func (s *Sprite) GetX() float64 {
	return s.X
}

func (s *Sprite) GetY() float64 {
	return s.Y
}

func (s *Sprite) SetX(value float64) {
	s.X = value
	s.UpdateCollider()
}

func (s *Sprite) SetY(value float64) {
	s.Y = value
	s.UpdateCollider()
}

func (s *Sprite) AddX(value float64) {
	s.X += value
	s.UpdateCollider()
}

func (s *Sprite) AddY(value float64) {
	s.Y += value
	s.UpdateCollider()
}

func (s *Sprite) UpdateCollider() {
	s.Collider = image.Rect(
		int(s.X),
		int(s.Y),
		int(s.X+config.DefaultTileSizeInPixels),
		int(s.Y+config.DefaultTileSizeInPixels),
	)
}
