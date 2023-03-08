package simulation

import (
	"fmt"
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"time"
)

const (
	WorldFoodSpawnInterval = 1.5 // seconds
)

type World struct {
	engine.Placeable
	engine.Game

	config            WorldConfig
	lastFoodSpawnedAt time.Time
	organisms         []*Organism
}

type WorldConfig struct {
	Width          int
	Height         int
	FoodCount      uint16
	FoodEnergy     Energy
	OrganismCount  uint16
	OrganismGenes  uint8
	OrganismEnergy Energy
}

var world *World

func NewWorld(config WorldConfig) *World {
	if config.Width <= 0 {
		config.Width = 64
	}

	if config.Height <= 0 {
		config.Height = 64
	}

	world = &World{
		config:            config,
		lastFoodSpawnedAt: time.Now(),
		organisms:         make([]*Organism, 0),
		Game: engine.Game{
			Grid: *engine.NewGrid(
				engine.NewVector(float64(config.Width), float64(config.Height)),
				10,
			),
		},
	}

	for f := 0; f < int(config.FoodCount); f++ {
		world.spawnFood(
			world.randomPosition(),
			config.FoodEnergy,
		)
	}

	for o := 0; o < int(config.OrganismCount); o++ {
		world.spawnOrganism(
			world.randomPosition(),
			NewGenome(config.OrganismGenes),
			config.OrganismEnergy,
		)
	}

	return world
}

func (w *World) Update(delta time.Duration) {
	engine.DebugPrintln(fmt.Sprintf("Organisms: %d", len(w.organisms)))

	if time.Since(w.lastFoodSpawnedAt).Seconds() > WorldFoodSpawnInterval {
		w.spawnFood(world.randomPosition(), w.config.FoodEnergy)
		w.lastFoodSpawnedAt = time.Now()
	}
}

func (w *World) Draw() *ebiten.Image {
	background := ebiten.NewImage(w.config.Width, w.config.Height)
	background.Fill(color.RGBA{R: 30, G: 30, B: 30, A: 255})

	return background
}

func (w *World) GetSize() engine.Vector {
	return engine.Vector{
		X: float64(w.config.Width),
		Y: float64(w.config.Height),
	}
}

func (w *World) Contains(position engine.Vector) bool {
	return position.X > 0 && position.Y > 0 && position.X <= float64(w.config.Width) && position.Y <= float64(w.config.Height)
}

func (w *World) spawnOrganism(position engine.Position, genome Genome, energy Energy) *Organism {
	organism := NewOrganism(position, genome, energy)

	w.organisms = append(w.organisms, organism)
	w.Grid.Add(organism)
	w.AddChild(organism)

	organism.Register(OrganismDeathHook, func() {
		w.organisms = engine.SliceRemoveUnordered(w.organisms, organism)
		w.onOrganismDeath(organism)
	})

	return organism
}

func (w *World) spawnFood(position engine.Position, energy Energy) *Food {
	food := NewFood(position, energy)

	w.Grid.Add(food)
	w.AddChild(food)

	return food
}

func (w *World) onOrganismDeath(organism *Organism) {
	w.spawnFood(w.randomPosition(), w.config.FoodEnergy)
}

func (w *World) randomPosition() engine.Position {
	return engine.Position{
		X: random.FloatBetween(0, w.config.Width),
		Y: random.FloatBetween(0, w.config.Height),
	}
}
