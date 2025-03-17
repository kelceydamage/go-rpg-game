package objects

import (
	"log"
	"rpg-game/src/lib/entities"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func NewPotion(SpawnX, SpawnY float64) *entities.Potion {
	image, _, err := ebitenutil.NewImageFromFile("src/assets/images/potion.png")
	if err != nil {
		log.Fatal(err)
	}
	return &entities.Potion{
		Sprite: &entities.Sprite{
			Image: image,
			X:     SpawnX,
			Y:     SpawnY,
		},
		HealAmount: 10,
	}
}
