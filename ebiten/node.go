package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

type Processor struct {
	capacity        int
	currentEventOpt *Event
	face            *basicfont.Face
	name            string
	processing      int // Frames left to process the current event
	queue           []Event
	screen          *ebiten.Image
	x, y            float64
}

func (n Processor) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(n.x-25), float32(n.y-15), 50, 30, nodeColor, true)
	text.Draw(screen, n.name, n.face, int(n.x)-20, int(n.y)-40, textColor)
	status := "Idle"
	if n.currentEventOpt != nil {
		status = "Processing"
	}
	text.Draw(screen, status, n.face, int(n.x)-25, int(n.y)+45, textColor)
}
