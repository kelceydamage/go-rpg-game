package tilemap

import (
	"encoding/json"
	"os"
	"path"
	"rpg-game/src/lib/tileset"
)

type Tilemap struct {
	*TilemapJSON
	Tilesets []tileset.Tileset
}

type TilemapLayerJSON struct {
	Data   []int  `json:"data"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Name   string `json:"name"`
}

type TilemapJSON struct {
	Layers   []TilemapLayerJSON `json:"layers"`
	Tilesets []map[string]any   `json:"tilesets"`
}

func (t *TilemapJSON) GenerateTilesets() ([]tileset.Tileset, error) {
	tilesets := make([]tileset.Tileset, 0)
	for _, tilesetData := range t.Tilesets {
		tilesetPath := path.Join("src/assets/maps", tilesetData["source"].(string))
		tileset, err := tileset.NewTileset(tilesetPath, int(tilesetData["firstgid"].(float64)))
		if err != nil {
			return nil, err
		}

		tilesets = append(tilesets, tileset)
	}
	return tilesets, nil
}

func LoadTilemapJSONFromDisk(filepath string) (*TilemapJSON, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	tilemapJSON := TilemapJSON{}
	err = json.Unmarshal(content, &tilemapJSON)
	if err != nil {
		return nil, err
	}

	return &tilemapJSON, err
}
