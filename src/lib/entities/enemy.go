package entities

import (
	"rpg-game/src/lib/animation"
)

type Enemy struct {
	*animation.AnimatedHumanoid
	*Sprite
	*Character
	ShouldFollowPlayer bool
	IsLootable         bool
}
