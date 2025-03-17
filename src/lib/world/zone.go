package world

import "rpg-game/src/lib/tilemap"

type Zone struct {
	Name    string
	Tilemap *tilemap.Tilemap
}
