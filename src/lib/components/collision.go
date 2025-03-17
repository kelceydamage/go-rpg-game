package components

import "image"

type Collidable interface {
	GetCollider() image.Rectangle
	GetChangeInX() float64
	GetChangeInY() float64
	GetX() float64
	GetY() float64
	SetX(float64)
	SetY(float64)
	GetDistanceFrom(object Collidable) float64
}

type Collider = image.Rectangle

type Colliders struct {
	Colliders []Collider
}
