package tileset

import (
	"encoding/json"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"rpg-game/src/lib/config"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Tileset interface {
	GetImage(id int) *ebiten.Image
}

type UniformTilesetJSON struct {
	Path string `json:"image"`
}

type UniformTileset struct {
	image       *ebiten.Image
	gid         int
	TilesPerRow int
}

func (u *UniformTileset) GetImage(id int) *ebiten.Image {
	u.TilesPerRow = 22
	id -= u.gid
	srcX := id % u.TilesPerRow
	srcY := id / u.TilesPerRow

	srcX *= config.DefaultTileSizeInPixels
	srcY *= config.DefaultTileSizeInPixels
	return u.image.SubImage(
		image.Rect(
			srcX,
			srcY,
			srcX+config.DefaultTileSizeInPixels,
			srcY+config.DefaultTileSizeInPixels,
		),
	).(*ebiten.Image)
}

type TileJSON struct {
	Id     int    `json:"id"`
	Path   string `json:"image"`
	Width  int    `json:"imagewidth"`
	Height int    `json:"imageheight"`
}

type DynamicTilesetJSON struct {
	Tiles []*TileJSON `json:"tiles"`
}

type DynamicTileset struct {
	Images []*ebiten.Image
	gid    int
}

func (u *DynamicTileset) GetImage(id int) *ebiten.Image {
	id -= u.gid
	return u.Images[id]
}

func NewTileset(path string, gid int) (Tileset, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if strings.Contains(path, "buildings") {
		var dynamicTilesetJSON DynamicTilesetJSON
		err := json.Unmarshal(contents, &dynamicTilesetJSON)
		if err != nil {
			return nil, err
		}

		dynamicTileset := DynamicTileset{}
		dynamicTileset.gid = gid
		dynamicTileset.Images = make([]*ebiten.Image, 0)

		for _, tileJSON := range dynamicTilesetJSON.Tiles {

			// refactor this out
			tileJSONPath := tileJSON.Path
			tileJSONPath = filepath.Clean(tileJSONPath)
			tileJSONPath = strings.ReplaceAll(tileJSONPath, "\\", "/")
			tileJSONPath = strings.TrimPrefix(tileJSONPath, "../")
			tileJSONPath = strings.TrimPrefix(tileJSONPath, "../")
			tileJSONPath = filepath.Join("src/assets/", tileJSONPath)

			image, _, err := ebitenutil.NewImageFromFile(tileJSONPath)
			if err != nil {
				return nil, err
			}
			fmt.Println("loaded", tileJSONPath)
			dynamicTileset.Images = append(dynamicTileset.Images, image)
		}

		return &dynamicTileset, nil
	}
	var uniformTilesetJSON UniformTilesetJSON
	err = json.Unmarshal(contents, &uniformTilesetJSON)
	if err != nil {
		return nil, err
	}

	uniformTileset := UniformTileset{}

	tileJSONPath := uniformTilesetJSON.Path
	tileJSONPath = filepath.Clean(tileJSONPath)
	tileJSONPath = strings.ReplaceAll(tileJSONPath, "\\", "/")
	tileJSONPath = strings.TrimPrefix(tileJSONPath, "../")
	tileJSONPath = strings.TrimPrefix(tileJSONPath, "../")
	tileJSONPath = filepath.Join("src/assets/", tileJSONPath)

	image, _, err := ebitenutil.NewImageFromFile(tileJSONPath)
	if err != nil {
		return nil, err
	}
	fmt.Println("loaded", tileJSONPath)
	uniformTileset.image = image
	uniformTileset.gid = gid

	return &uniformTileset, nil
}
