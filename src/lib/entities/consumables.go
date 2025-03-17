package entities

import "rpg-game/src/lib/components"

type Consumable interface {
	Consume() uint
}

type Food struct {
	*components.BasicItem
}

func (f *Food) Consume() {}

type Beverage struct {
	*components.BasicItem
}

func (b *Beverage) Consume() {}
