package entities

type Entities struct {
	Player  *Player
	Enemies map[uint16]*Enemy
	Potions map[uint16]*Potion
}
