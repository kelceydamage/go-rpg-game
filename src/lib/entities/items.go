package entities

type Collectable interface {
	PickUp() any
}

type Item struct {
	FlaggedForRemoval bool
}

func (i *Item) PickUp() any {
	return i
}
