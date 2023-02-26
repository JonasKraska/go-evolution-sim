package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"time"
)

type World struct {
	engine.Game

	width  int
	height int

	organisms []*Organism
	food      []*Food
}

type WorldConfig struct {
	Width     uint32
	Height    uint32
	Food      []FoodCohort
	Organisms []OrganismCohort
}

type OrganismCohort struct {
	Count  int
	Energy Energy
	Genome Genome
}

type FoodCohort struct {
	Count  int
	Energy Energy
}

func NewWorld(config WorldConfig) *World {
	if config.Width <= 0 {
		config.Width = 64
	}

	if config.Height <= 0 {
		config.Height = 64
	}

	w := &World{
		width:  int(config.Width),
		height: int(config.Height),
	}

	for _, cohort := range config.Food {
		for f := 1; f < cohort.Count; f++ {
			w.spawnFood(
				w.randomPosition(),
				cohort.Energy,
			)
		}
	}

	for _, cohort := range config.Organisms {
		for o := 1; o < cohort.Count; o++ {
			w.spawnOrganism(
				w.randomPosition(),
				cohort.Genome,
				cohort.Energy,
			)
		}
	}

	return w
}

func (w *World) Update(delta time.Duration) {
	for o := 1; o < len(w.organisms); o++ {
		for f := 1; f < len(w.food); f++ {
			organism := w.organisms[o]
			food := w.food[f]

			if organism.GetPosition().EqualsIgnoringDecimals(food.GetPosition()) {
				organism.Consume(food.Energy)
				food.Remove()
			}
		}
	}
}

func (w *World) Draw() *ebiten.Image {
	background := ebiten.NewImage(w.width, w.height)
	background.Fill(color.RGBA{R: 30, G: 30, B: 30, A: 255})

	return background
}

func (w *World) Contains(position engine.Vector) bool {
	return position.X > 0 && position.Y > 0 && position.X <= float64(w.width) && position.Y <= float64(w.height)
}

func (w *World) spawnOrganism(position engine.Position, genome Genome, energy Energy) *Organism {
	organism := NewOrganism(position, genome, energy)

	w.organisms = append(w.organisms, organism)
	w.AddChild(organism)

	organism.Register(OrganismDeathHook, func() {
		w.onOrganismDeath(organism)
	})

	return organism
}

func (w *World) spawnFood(position engine.Position, energy Energy) *Food {
	food := NewFood(position, energy)

	w.food = append(w.food, food)
	w.AddChild(food)

	food.Register(engine.NodeRemoveHook, func() {
		engine.PtrSliceRemovce(w.food, food)
	})

	return food
}

func (w *World) onOrganismDeath(organism *Organism) {
	engine.PtrSliceRemovce(w.organisms, organism)
	w.spawnFood(organism.GetPosition(), Energy(5))
}

func (w *World) randomPosition() engine.Position {
	return engine.Position{
		X: random.FloatBetween(0, w.width),
		Y: random.FloatBetween(0, w.height),
	}
}
