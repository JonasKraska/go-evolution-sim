package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	"github.com/fogleman/gg"
	"image"
	"image/color"
)

type Energy float64

type Organism struct {
	engine.Movable

    sprite image.Image
	genome Genome
	energy Energy
}

type OrganismConfig struct {
	Color color.Color
}

func NewOrganism(config OrganismConfig, position engine.Position) *Organism {
	genome := Genome{
		Color: config.Color,
	}

	o := &Organism{
        genome: genome,
		energy: Energy(random.FloatBetween(1, 10)),
	}

	o.SetPosition(position)

	return o
}

func (o *Organism) Update() {
	position := o.GetPosition()

	position.X += random.IntBetween(-1, 1)
	position.Y += random.IntBetween(-1, 1)

	o.MoveTo(position)
}

func (o *Organism) Draw() image.Image {
	if o.sprite == nil {
		dc := gg.NewContext(1, 1)

		dc.DrawRectangle(0, 0, 1, 1)
		dc.SetColor(o.genome.Color)
		dc.Fill()

        o.sprite = dc.Image()
	}

    return o.sprite
}
