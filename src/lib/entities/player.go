package entities

import (
	"rpg-game/src/lib/animation"
)

type Player struct {
	*animation.AnimatedHumanoid
	*Sprite
	*Character
}
