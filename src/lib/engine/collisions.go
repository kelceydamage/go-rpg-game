package engine

import (
	"math"
	"rpg-game/src/lib/components"
	"rpg-game/src/lib/config"
)

func CheckCollisionHorizontal(collidable components.Collidable, colliders *components.Colliders) {
	for _, collider := range colliders.Colliders {
		if collider.Overlaps(collidable.GetCollider()) {
			//fmt.Println("Collision detected")
			if collidable.GetChangeInX() > 0.0 {
				collidable.SetX(float64(collider.Min.X) - 16.0)
			} else if collidable.GetChangeInX() < 0.0 {
				collidable.SetX(float64(collider.Max.X))
			}
		}
	}
}

func CheckCollisionVertical(collidable components.Collidable, colliders *components.Colliders) {
	for _, collider := range colliders.Colliders {
		if collider.Overlaps(collidable.GetCollider()) {
			//fmt.Println("Collision detected")
			if collidable.GetChangeInY() > 0.0 {
				collidable.SetY(float64(collider.Min.Y) - 16.0)
			} else if collidable.GetChangeInY() < 0.0 {
				collidable.SetY(float64(collider.Max.Y))
			}
		}
	}
}

func GetDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2+(config.DefaultTileSizeInPixels/2), 2) + math.Pow(y1-y2+(config.DefaultTileSizeInPixels/2), 2))
}
