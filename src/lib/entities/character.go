package entities

import "rpg-game/src/lib/components"

type Character struct {
	Skills    []uint
	Inventory components.Container
	components.Moveable
	components.CombatCapable
}
