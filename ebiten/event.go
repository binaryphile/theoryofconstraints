package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"math"
)

type Event struct {
	inProcess        bool
	screen           *ebiten.Image
	targetX, targetY float64
	x, y             float64
}

func (e Event) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(e.x), float32(e.y), 5, eventColor, true)
}

// MoveTowardTarget moves an event incrementally toward its target
func (e Event) MoveTowardTarget() {
	const speed = 2.0
	dx := e.targetX - e.x
	dy := e.targetY - e.y
	dist := math.Hypot(dx, dy)
	if dist > 0 {
		e.x += dx / dist * speed
		e.y += dy / dist * speed
	}
}
