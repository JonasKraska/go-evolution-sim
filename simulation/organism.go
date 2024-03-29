package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

const (
	OrganismDeathHook   = engine.Hook("organism.death")
	OrganismMaxTurnDeg  = 10.0
	OrganismFieldOfView = 60.0
	OrganismViewRange   = 10.0
)

type Energy = float64

type Organism struct {
	engine.Node
	engine.Movable

	sprite *ebiten.Image

	brain  *Brain
	genome Genome
	energy Energy
}

func NewOrganism(position engine.Position, genome Genome, energy Energy) *Organism {
	o := &Organism{
		brain:  NewBrain(),
		genome: genome,
		energy: energy,
	}

	for _, c := range o.genome.Connections {
		o.brain.Connection(c.GetFrom(), c.GetTo(), c.GetWeight())
	}

	o.brain.Prune()

	o.SetPosition(position)
	o.SetVelocity(engine.RandomVectorOnUnitCircle())

	return o
}

func (o *Organism) Update(delta time.Duration) {
	o.burnEnergy(delta)
	o.consumeFood()

	if o.energy < 0 {
		o.die()
		return
	}

	o.proliferate()

	_, foodDistance, foodAngle := o.detectClosestFood()
	o.brain.SetFoodDistance(foodDistance)
	o.brain.SetFoodAngle(foodAngle)

	obstacleDetected := o.detectObstacle()
	if obstacleDetected == true {
		o.brain.SetObstacleDistance(OrganismViewRange)
	} else {
		o.brain.SetObstacleDistance(OrganismViewRange * -1)
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
	o.energy -= simulation.config.OrganismMotabolismRate * delta.Seconds()
}

func (o *Organism) consumeFood() {
	nodes, _ := world.GetGrid().GetNodesInCellOf(o)

	for _, n := range nodes {
		if food, ok := n.(*Food); ok {
			if o.GetPosition().ToPoint().Equals(food.GetPosition().ToPoint()) {
				if err := food.Remove(); err == nil {
					o.Consume(food.energy)
				}
			}
		}
	}
}

func (o *Organism) proliferate() {
	if o.energy >= float64(simulation.config.OrganismProliferationThreshold) {
		o.energy = o.energy / 2
		world.spawnOrganism(o.GetPosition(), o.genome, o.energy)
	}
}

func (o *Organism) detectClosestFood() (*Food, float64, engine.Angle) {
	nodes, _ := world.GetGrid().GetNodesInCellOf(o, 1)

	var (
		closestFood         *Food
		closestFoodDistance float64
		closestFoodAngle    engine.Angle
	)

	for _, n := range nodes {
		if food, ok := n.(*Food); ok {
			foodDirection := food.GetPosition().Sub(o.GetPosition())
			foodDistance := o.GetPosition().Distance(food.GetPosition())
			foodAngle := o.GetVelocity().AngleBetween(foodDirection)

			isInViewRange := foodDistance < OrganismViewRange
			isInFieldOfView := foodAngle.GetDeg() < OrganismFieldOfView/2

			if isInViewRange && isInFieldOfView && (closestFoodDistance == 0 || closestFoodDistance > foodDistance) {
				closestFood = food
				closestFoodDistance = foodDistance
				closestFoodAngle = foodAngle
			}
		}
	}

	return closestFood, closestFoodDistance, closestFoodAngle
}

func (o *Organism) detectObstacle() bool {
	viewPoint := o.GetVelocity().Normalize().MulScalar(OrganismViewRange)
	return !world.Contains(viewPoint)
}

func (o *Organism) move(delta time.Duration) {
	currentVelocity := o.GetVelocity().Normalize()
	newOrientation := currentVelocity.Rotate(o.brain.GetDirectionAngle())
	velocity := newOrientation.MulScalar(17.5 * delta.Seconds())

	o.SetVelocity(velocity)
}

func (o *Organism) die() {
	o.Dispatch(OrganismDeathHook)
	o.Remove()
}
