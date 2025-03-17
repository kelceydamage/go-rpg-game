package objects

import (
	"image"
	"log"
	"rpg-game/src/lib/components"
	"rpg-game/src/lib/tilemap"
	"rpg-game/src/lib/world"
)

func NewWorld() *world.World {
	tilemapJSON, err := tilemap.LoadTilemapJSONFromDisk("src/assets/maps/spawn.json")
	if err != nil {
		log.Fatal(err)
	}

	Tilesets, err := tilemapJSON.GenerateTilesets()
	if err != nil {
		log.Fatal(err)
	}
	return &world.World{
		CurrentZone:      "startingZone",
		NavPixelsPerTile: 16,
		Zones: map[string]*world.Zone{
			"startingZone": {
				Name: "startingZone",
				Tilemap: &tilemap.Tilemap{
					TilemapJSON: tilemapJSON,
					Tilesets:    Tilesets,
				},
			},
		},
		Colliders: &components.Colliders{
			Colliders: []components.Collider{
				image.Rect(144, 0, 208, 32),
				image.Rect(224, 240, 288, 272),
			},
		},
	}
}
