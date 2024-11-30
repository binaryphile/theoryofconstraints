package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

type Terminal struct {
	color color.RGBA
	face  *basicfont.Face
	name  string
	x, y  float64
}

func (t Terminal) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(t.x), float32(t.y), 20, t.color, true)
	text.Draw(screen, t.name, t.face, int(t.x)-20, int(t.y)+40, textColor)
}
