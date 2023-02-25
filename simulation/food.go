package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
    "time"
)

type Food struct {
    engine.Node
	engine.Placeable

	Energy Energy
}

func NewFood(position engine.Position, energy Energy) *Food {
	f := &Food{
		Energy: energy,
	}

	f.SetPosition(position)

	return f
}

func (f *Food) Update(delta time.Duration) {
	// @TODO: energy decay?
}

var foodSprite *ebiten.Image

func (f *Food) Draw() *ebiten.Image {
	if foodSprite == nil {
		foodSprite = ebiten.NewImage(1, 1)
		foodSprite.Fill(color.RGBA{R: 219, G: 93, B: 81, A: 255})
	}

	return foodSprite
}
