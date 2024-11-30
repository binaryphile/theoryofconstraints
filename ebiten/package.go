package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"log"
)

// Colors
var backgroundColor = color.RGBA{R: 240, G: 240, B: 240, A: 255} // Light muted background
var eventColor = color.RGBA{R: 255, G: 165, B: 0, A: 255}        // Pastel orange for events
var nodeColor = color.RGBA{R: 144, G: 238, B: 144, A: 255}       // Light pastel green
var sinkColor = color.RGBA{R: 255, G: 182, B: 193, A: 255}       // Light pastel pink
var sourceColor = color.RGBA{R: 173, G: 216, B: 230, A: 255}     // Light pastel blue
var textColor = color.Black                                      // Darker text for better contrast
var fontFace = basicfont.Face7x13

func Run() {
	processors := []Processor{{
		capacity: 10,
		face:     fontFace,
		name:     "Step 1",
		x:        350, y: 300,
	}}

	sink := Terminal{
		color: sinkColor,
		face:  fontFace,
		name:  "Sink",
		x:     650, y: 300,
	}

	source := Terminal{
		color: sourceColor,
		face:  fontFace,
		name:  "Source",
		x:     150, y: 300,
	}

	system := &System{
		processors: processors,
		terminals:  []Terminal{source, sink},
	}

	// Set up the game window
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("ToC Compute Bottleneck Visualization (Event Processing)")
	if err := ebiten.RunGame(system); err != nil {
		log.Fatal(err)
	}
}
