package entities

type Potion struct {
	*Sprite
	*Beverage
	HealAmount uint
}

func (p *Potion) PickUp() *Potion {
	return p
}

func (p *Potion) Consume() uint {
	p.SetFlaggedForRemoval(true)
	return p.HealAmount
}
