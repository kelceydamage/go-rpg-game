package engine

import (
	"rpg-game/src/lib/render"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Keyboard struct{}

type Mouse struct {
}

func (m *Mouse) GetMousePosition() (int, int) {
	return ebiten.CursorPosition()
}

func (m *Mouse) GetMousePositionRelativeToCamera(camera *render.Camera) (int, int) {
	x, y := m.GetMousePosition()
	x -= int(camera.X)
	y -= int(camera.Y)
	return x, y
}

func (m *Mouse) GetMouseButtonState(button ebiten.MouseButton) bool {
	return inpututil.IsMouseButtonJustPressed(button)
}

type Controls struct {
	Keyboard *Keyboard
	Mouse    *Mouse
}
