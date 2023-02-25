package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

const (
	OrganismDeathHook = engine.Hook("organism.death")
)

type Energy float64

type Organism struct {
	engine.Node
	engine.Movable

	sprite *ebiten.Image

	genome Genome
	energy Energy

	energyBurnRate float64
}

func NewOrganism(position engine.Position, genome Genome, energy Energy) *Organism {
	o := &Organism{
		genome: NewGenome(genome),
		energy: energy,
	}

	o.initEnergyBurnRate()
	o.SetPosition(position)

	return o
}

func (o *Organism) Update(delta time.Duration) {
	o.energy -= Energy(o.energyBurnRate * delta.Seconds())

	if o.energy < 0 {
		o.Dispatch(OrganismDeathHook)
		o.Remove()
		return
	}

	position := o.GetPosition()

	position.X += float64(random.IntBetween(-1, 1))
	position.Y += float64(random.IntBetween(-1, 1))

	o.MoveTo(position)
}

func (o *Organism) Draw() *ebiten.Image {
	if o.sprite == nil {
		o.sprite = ebiten.NewImage(1, 1)
		o.sprite.Fill(o.genome.Color)
	}

	return o.sprite
}

func (o *Organism) initEnergyBurnRate() {
	o.energyBurnRate = float64(o.genome.MetabolismRate) / 2
}
