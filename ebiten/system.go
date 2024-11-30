package ebiten

import (
	"github.com/binaryphile/fluentfp/fluent"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type System struct {
	terminals  fluent.SliceOf[Terminal]
	events     fluent.SliceOf[Event]
	processors fluent.SliceOf[Processor]
}

type Drawer interface {
	Draw(*ebiten.Image)
}

func Draw[T Drawer](screen *ebiten.Image) func(T) {
	return func(t T) {
		t.Draw(screen)
	}
}

// Draw renders the system, steps, and events
func (s System) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	s.terminals.Each(Draw[Terminal](screen))
	s.processors.Each(Draw[Processor](screen))
	s.events.Each(Draw[Event](screen))
}

func (s System) Layout(_, _ int) (int, int) {
	return 800, 600
}

// Update handles the game logic, event flow, and processing simulation
func (s System) Update() error {
	// Generate a new event periodically
	source := s.terminals[0]
	sink := s.terminals[1]
	if len(s.events) < 10 && (len(s.events) == 0 || s.events[len(s.events)-1].x > source.x+30) {
		newEvent := Event{
			x: source.x, y: source.y,
			targetX: s.processors[0].x, targetY: s.processors[0].y, // Initially target the first processor
		}
		s.events = append(s.events, newEvent)
	}

	// Process events in nodes
	for i, processor := range s.processors {
		// If the processor is processing an event, decrement its processing time
		if processor.currentEventOpt != nil {
			processor.processing--
			if processor.processing <= 0 {
				// Processing complete; move the event to the next step or sink
				processor.currentEventOpt.inProcess = false
				if i+1 < len(s.processors) {
					processor.currentEventOpt.targetX = s.processors[i+1].x
					processor.currentEventOpt.targetY = s.processors[i+1].y
				} else {
					processor.currentEventOpt.targetX = sink.x
					processor.currentEventOpt.targetY = sink.y
				}
				processor.currentEventOpt = nil
			}
		} else if len(processor.queue) > 0 {
			// Start processing the next event in the queue
			processor.currentEventOpt = &processor.queue[0]
			processor.queue = processor.queue[1:]
			processor.processing = 60 // Example: 60 frames for processing
			processor.currentEventOpt.inProcess = true
		}
	}

	// Move events
	for _, event := range s.events {
		if !event.inProcess {
			event.MoveTowardTarget()
			if math.Hypot(event.x-event.targetX, event.y-event.targetY) < 5 {
				// Check if the event has reached its target
				for _, processor := range s.processors {
					if event.targetX == processor.x && event.targetY == processor.y {
						processor.queue = append(processor.queue, event)
						event.inProcess = true
						break
					}
				}
			}
		}
	}

	return nil
}
