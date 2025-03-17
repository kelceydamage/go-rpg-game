package world

import (
	"rpg-game/src/lib/components"
)

type World struct {
	NavPixelsPerTile int
	// Zone.Name
	CurrentZone string
	// Zone.Name -> Zone
	Zones     map[string]*Zone
	Colliders *components.Colliders
}
