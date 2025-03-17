package main

import (
	"log"

	"rpg-game/src/lib/sceneLoader"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Super RPG Game")
	ebiten.SetRunnableOnUnfocused(true)

	game := sceneLoader.NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
