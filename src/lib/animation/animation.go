package animation

type SpriteMovementState uint8

const (
	Down SpriteMovementState = iota
	Up
	Left
	Right
)

type Animation struct {
	FirstFrame   int
	LastFrame    int
	CurrentFrame int
	Stepping     int
	SpeedInTicks float32
	FrameCounter float32
}

func NewAnimation(firstFrame, lastFrame, stepping int, speedInTicks float32) *Animation {
	return &Animation{
		FirstFrame:   firstFrame,
		LastFrame:    lastFrame,
		Stepping:     stepping,
		SpeedInTicks: speedInTicks,
		FrameCounter: speedInTicks,
		CurrentFrame: firstFrame,
	}
}

func (a *Animation) Update() {
	a.FrameCounter -= 1.0
	if a.FrameCounter <= 0.0 {
		a.FrameCounter = a.SpeedInTicks
		a.CurrentFrame += a.Stepping
		if a.CurrentFrame > a.LastFrame {
			a.CurrentFrame = a.FirstFrame
		}
	}
}

func (a *Animation) Frame() int {
	return a.CurrentFrame
}

type Animated struct {
	Animations       map[SpriteMovementState]*Animation
	CurrentAnimation SpriteMovementState
	SpriteSheet      *SpriteSheet
}

func (a *Animated) SetSpriteSheet(spriteSheet *SpriteSheet) {
	a.SpriteSheet = spriteSheet
}

func (a *Animated) SetAnimations(animations map[SpriteMovementState]*Animation) {
	a.Animations = animations
}

func (a *Animated) AddAnimation(triggerState SpriteMovementState, animation *Animation) {
	a.Animations[triggerState] = animation
}

func (a *Animated) GetCurrentAnimation() *Animation {
	return a.Animations[a.CurrentAnimation]
}

func (a *Animated) UpdateCurrentAnimation() {
	a.Animations[a.CurrentAnimation].Update()
}

type AnimatedHumanoid struct {
	*Animated
}

func (a *AnimatedHumanoid) SetCurrentAnimation(changeInX, changeInY int) {
	if changeInX > 0 {
		a.CurrentAnimation = Right
	}
	if changeInX < 0 {
		a.CurrentAnimation = Left
	}
	if changeInY > 0 {
		a.CurrentAnimation = Down
	}
	if changeInY < 0 {
		a.CurrentAnimation = Up
	}
}
