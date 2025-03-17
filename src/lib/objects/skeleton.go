package objects

import (
	"log"
	"rpg-game/src/lib/animation"
	"rpg-game/src/lib/components"
	"rpg-game/src/lib/entities"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func NewSkeleton(spawnX, spawnY float64) *entities.Enemy {
	_image, _, err := ebitenutil.NewImageFromFile("src/assets/images/skeleton.png")
	if err != nil {
		log.Fatal(err)
	}
	return &entities.Enemy{
		ShouldFollowPlayer: true,
		IsLootable:         true,
		AnimatedHumanoid: &animation.AnimatedHumanoid{
			Animated: &animation.Animated{
				Animations: map[animation.SpriteMovementState]*animation.Animation{
					animation.Up:    animation.NewAnimation(5, 13, 4, 20.0),
					animation.Down:  animation.NewAnimation(4, 12, 4, 20.0),
					animation.Left:  animation.NewAnimation(6, 14, 4, 20.0),
					animation.Right: animation.NewAnimation(7, 15, 4, 20.0),
				},
				SpriteSheet: animation.NewSpriteSheet(4, 7, 16),
			},
		},
		Sprite: entities.NewSprite(_image, spawnX, spawnY),
		Character: &entities.Character{
			Moveable:      components.NewBasicMovement(),
			CombatCapable: components.NewEnemyCombat(10, 1, 30),
		},
	}
}
