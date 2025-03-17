package components

type Moveable interface {
	GetChangeInX() float64
	GetChangeInY() float64
	SetChangeInX(float64)
	SetChangeInY(float64)
	ResetVelocity()
}

type Movement struct {
	ChangeInX float64
	ChangeInY float64
}

func (m *Movement) GetChangeInX() float64 {
	return m.ChangeInX
}

func (m *Movement) GetChangeInY() float64 {
	return m.ChangeInY
}

func (m *Movement) SetChangeInX(changeInX float64) {
	m.ChangeInX = changeInX
}

func (m *Movement) SetChangeInY(changeInY float64) {
	m.ChangeInY = changeInY
}

func (m *Movement) ResetVelocity() {
	m.ChangeInX = 0
	m.ChangeInY = 0
}

type BasicMovement struct {
	*Movement
}

func NewBasicMovement() *BasicMovement {
	return &BasicMovement{
		Movement: &Movement{
			ChangeInX: 0,
			ChangeInY: 0,
		},
	}
}
