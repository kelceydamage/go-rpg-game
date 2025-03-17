package render

import (
	"fmt"
	"image"
	"image/color"
	"rpg-game/src/lib/config"
	"rpg-game/src/lib/entities"
	"rpg-game/src/lib/tilemap"

	"rpg-game/src/lib/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Renderer struct {
	ScreenLayoutWidth  int
	ScreenLayoutHeight int
	Camera             *Camera
}

func (r *Renderer) LayoutScreen() (screenWidth, screenHeight int) {
	return r.ScreenLayoutWidth, r.ScreenLayoutHeight
}

func (r *Renderer) RenderScreen(screen *ebiten.Image, world *world.World, entities *entities.Entities) {
	screen.Fill(color.RGBA{120, 180, 255, 255})
	currentTilemap := *world.Zones[world.CurrentZone].Tilemap
	r.Camera.FollowTarget(
		entities.Player.X+float64(world.NavPixelsPerTile/2),
		entities.Player.Y+float64(world.NavPixelsPerTile/2),
		float64(r.ScreenLayoutWidth),
		float64(r.ScreenLayoutHeight),
	)
	r.Camera.Constrain(
		float64(currentTilemap.TilemapJSON.Layers[0].Width*world.NavPixelsPerTile),
		float64(currentTilemap.TilemapJSON.Layers[0].Height*world.NavPixelsPerTile),
		float64(r.ScreenLayoutWidth),
		float64(r.ScreenLayoutHeight),
	)

	r.drawTileMap(screen, currentTilemap)

	r.drawSprite(
		entities.Player.Sprite,
		screen,
		entities.Player.SpriteSheet.Rect(
			entities.Player.GetCurrentAnimation().Frame(),
		),
	)
	r.DrawCollider(screen, entities.Player.Sprite.Collider, color.RGBA{0, 255, 0, 255})

	ebitenutil.DebugPrint(
		screen,
		"Player Coordinates: "+fmt.Sprintf("%f", entities.Player.X)+", "+fmt.Sprintf("%f", entities.Player.Y),
	)

	for _, enemy := range entities.Enemies {
		r.drawSprite(
			enemy.Sprite,
			screen,
			enemy.SpriteSheet.Rect(
				enemy.GetCurrentAnimation().Frame(),
			),
		)
		r.DrawCollider(screen, enemy.Sprite.Collider, color.RGBA{255, 0, 0, 255})
	}

	for _, potion := range entities.Potions {
		r.drawSprite(
			potion.Sprite,
			screen,
			image.Rect(
				0,
				0,
				config.DefaultTileSizeInPixels,
				config.DefaultTileSizeInPixels,
			),
		)
	}

	// Debug draw colliders
	for _, collider := range world.Colliders.Colliders {
		r.DrawCollider(screen, collider, color.RGBA{0, 0, 255, 255})
	}
}

func (r *Renderer) DrawCollider(screen *ebiten.Image, collider image.Rectangle, color color.Color) {
	vector.StrokeRect(
		screen,
		float32(collider.Min.X)+float32(r.Camera.X),
		float32(collider.Min.Y)+float32(r.Camera.Y),
		float32(collider.Dx()),
		float32(collider.Dy()),
		1.0,
		color,
		true,
	)
}

func (r *Renderer) drawTile(image *ebiten.Image, x, y float64, screen *ebiten.Image, spriteBoundary image.Rectangle) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x, y)
	opts.GeoM.Translate(r.Camera.X, r.Camera.Y)
	screen.DrawImage(
		image.SubImage(spriteBoundary).(*ebiten.Image),
		&opts,
	)
	opts.GeoM.Reset()
}

func (r *Renderer) drawSprite(sprite *entities.Sprite, screen *ebiten.Image, spriteBoundary image.Rectangle) {
	r.drawTile(sprite.Image, sprite.X, sprite.Y, screen, spriteBoundary)
}

func (r *Renderer) drawTileMap(screen *ebiten.Image, tilemap tilemap.Tilemap) {
	for layerIndex, layer := range tilemap.TilemapJSON.Layers {
		for index, id := range layer.Data {

			if id == 0 {
				continue
			}

			x := index % layer.Width
			y := index / layer.Width
			x *= config.DefaultTileSizeInPixels
			y *= config.DefaultTileSizeInPixels

			image := tilemap.Tilesets[layerIndex].GetImage(id)

			opts := ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(x), float64(y))
			opts.GeoM.Translate(
				0.0,
				-(float64(image.Bounds().Dy()) + config.DefaultTileSizeInPixels),
			)
			opts.GeoM.Translate(r.Camera.X, r.Camera.Y)

			screen.DrawImage(image, &opts)
		}
	}
}
