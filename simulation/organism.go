package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/JonasKraska/go-evolution-sim/engine/random"
    "github.com/hajimehoshi/ebiten/v2"
)

type Energy float64

type Organism struct {
	engine.Movable

	sprite *ebiten.Image
	genome Genome
	energy Energy
}

func NewOrganism(genome Genome, position engine.Position) *Organism {
	o := &Organism{
		genome: NewGenome(genome),
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

func (o *Organism) Draw() *ebiten.Image {
	if o.sprite == nil {
        o.sprite = ebiten.NewImage(1, 1)
        o.sprite.Fill(o.genome.Color)
	}

	return o.sprite
}
