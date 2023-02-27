package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"time"
)

const (
	OrganismDeathHook = engine.Hook("organism.death")
)

type Energy = float64

type Organism struct {
	engine.Node
	engine.Movable

	sprite *ebiten.Image

	brain  *Brain
	genome Genome
	energy Energy

	orientation engine.Vector
}

func NewOrganism(position engine.Position, genome Genome, energy Energy) *Organism {
	o := &Organism{
		brain:       NewBrain(),
		genome:      NewGenome(genome),
		energy:      energy,
		orientation: engine.RandomVectorOnUnitCircle(),
	}

	o.SetPosition(position)

	return o
}

func (o *Organism) Update(delta time.Duration) {
	o.burnEnergy(delta)

	if o.energy < 0 {
		o.die()
		return
	}

	o.brain.Process()

	o.move(delta)
}

func (o *Organism) Draw() *ebiten.Image {
	if o.sprite == nil {
		o.sprite = ebiten.NewImage(1, 1)
		o.sprite.Fill(o.genome.Color)
	}

	return o.sprite
}

func (o *Organism) Consume(energy Energy) {
	o.energy += energy
}

func (o *Organism) burnEnergy(delta time.Duration) {
	rate := 0.5*math.Sqrt(float64(o.genome.Speed)) + 1
	o.energy -= rate * delta.Seconds()
}

func (o *Organism) move(delta time.Duration) {
	directionChangeAngle := o.brain.GetDirectionChange() * 10
	o.orientation = o.orientation.Rotate(directionChangeAngle)

	speed := float64(o.genome.Speed)
	velocity := o.orientation.MulScalar(speed * delta.Seconds())

	o.SetVelocity(velocity)
}

func (o *Organism) die() {
	o.Dispatch(OrganismDeathHook)
	o.Remove()
}
