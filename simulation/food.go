package simulation

import (
    "github.com/JonasKraska/go-evolution-sim/engine"
    "github.com/fogleman/gg"
    "image"
    "image/color"
)

type Food struct {
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

func (f *Food) Update() {
    // @TODO: energy decay?
}

var foodSprite image.Image

func (f *Food) Draw() image.Image {
    if foodSprite == nil {
		dc := gg.NewContext(1, 1)

		dc.DrawRectangle(0, 0, 1, 1)
		dc.SetColor(color.RGBA{R: 219, G: 93, B: 81, A: 255})
		dc.Fill()

        foodSprite = dc.Image()
	}

    return foodSprite
}