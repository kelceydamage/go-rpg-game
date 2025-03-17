package scenes

import (
	"rpg-game/src/lib/engine"
	"rpg-game/src/lib/entities"
	"rpg-game/src/lib/objects"
	"rpg-game/src/lib/render"
	"rpg-game/src/lib/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameScene struct {
	*render.Renderer
	*engine.Controller
	*entities.Entities
	*world.World
	*engine.Controls
	loaded bool
}

func NewGameScene() *GameScene {
	return &GameScene{
		Renderer: &render.Renderer{
			ScreenLayoutWidth:  320,
			ScreenLayoutHeight: 240,
			Camera:             render.NewCamera(0, 0),
		},
		Entities: &entities.Entities{
			Player:  nil,
			Enemies: make(map[uint16]*entities.Enemy, 0),
			Potions: make(map[uint16]*entities.Potion, 0),
		},
		World:    nil,
		Controls: &engine.Controls{},
		loaded:   false,
	}
}

func (g *GameScene) IsLoaded() bool {
	return g.loaded
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	g.Renderer.RenderScreen(screen, g.World, g.Entities)
}

func (g *GameScene) FirstLoad() {
	// load the image from file
	g.Entities.Player = objects.NewPlayer(100, 100)
	g.Entities.Enemies = map[uint16]*entities.Enemy{
		0: objects.NewSkeleton(200, 200),
		1: objects.NewSkeleton(300, 300),
		2: objects.NewSkeleton(400, 400),
	}
	g.Entities.Potions = map[uint16]*entities.Potion{
		0: objects.NewPotion(80, 80),
		1: objects.NewPotion(160, 160),
		2: objects.NewPotion(240, 240),
	}
	g.World = objects.NewWorld()
	g.loaded = true
}

func (g *GameScene) OnEnter() {
}

func (g *GameScene) OnExit() {
}

func (g *GameScene) Update() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ExitSceneId
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return PauseSceneId
	}

	g.Controller.MovePlayer(
		g.Player,
		g.World.Colliders,
	)
	g.Controller.MoveEnemies(
		g.Entities.Player,
		g.Entities.Enemies,
		g.World.Colliders,
		g.Controls,
		g.Renderer.Camera,
	)
	g.Controller.PickupPotion(
		g.Entities.Player,
		g.Entities.Potions,
	)
	return GameSceneId
}
