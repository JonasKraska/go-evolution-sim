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
	WorldOrganismSpawned = engine.Hook("world.organism.spawned")
	WorldOrganismRemoved = engine.Hook("world.organism.removed")
)

type World struct {
	engine.Placeable
	engine.Game

	width  int
	height int

	organisms []*Organism
	foods     []*Food
}

var world *World

func NewWorld(width uint16, height uint16) *World {
	world = &World{
		width:     int(width),
		height:    int(height),
		organisms: make([]*Organism, 0),
		foods:     make([]*Food, 0),
		Game: engine.Game{
			Grid: *engine.NewGrid(
				engine.NewVector(float64(width), float64(height)),
				10,
			),
		},
	}

	return world
}

func (w *World) Update(delta time.Duration) {
	totalOrganismEnergy := 0.0
	for _, o := range w.organisms {
		totalOrganismEnergy += o.energy
	}

	averageOrganismEnergy := totalOrganismEnergy / float64(len(w.organisms))

	engine.DebugPrintln(fmt.Sprintf("Organisms: %d", len(w.organisms)))
	engine.DebugPrintln(fmt.Sprintf("Organisms Energy: %.2f", averageOrganismEnergy))
	engine.DebugPrintln(fmt.Sprintf("Food: %d", len(w.foods)))
}

func (w *World) Draw() *ebiten.Image {
	background := ebiten.NewImage(w.width, w.height)
	background.Fill(color.RGBA{R: 30, G: 30, B: 30, A: 255})

	return background
}

func (w *World) GetSize() engine.Vector {
	return engine.Vector{
		X: float64(w.width),
		Y: float64(w.height),
	}
}

func (w *World) Contains(position engine.Vector) bool {
	return position.X > 0 && position.Y > 0 && position.X <= float64(w.width) && position.Y <= float64(w.height)
}

func (w *World) spawnOrganism(position engine.Position, genome Genome, energy Energy) *Organism {
	organism := NewOrganism(position, genome, energy)

	w.organisms = append(w.organisms, organism)
	w.Grid.Add(organism)
	w.AddChild(organism)

	organism.Register(OrganismDeathHook, func() {
		w.organisms = engine.SliceRemoveUnordered(w.organisms, organism)
		w.onOrganismDeath(organism)
		w.Dispatch(WorldOrganismRemoved)
	})

	w.Dispatch(WorldOrganismSpawned)

	return organism
}

func (w *World) spawnFood(position engine.Position, energy Energy) *Food {
	food := NewFood(position, energy)

	w.foods = append(w.foods, food)
	w.Grid.Add(food)
	w.AddChild(food)

	food.Register(engine.NodeRemoveHook, func() {
		w.foods = engine.SliceRemoveUnordered(w.foods, food)
	})

	return food
}

func (w *World) onOrganismDeath(organism *Organism) {
	// @TODO make the hooks have context and move it to the simualtion
	w.spawnFood(w.randomPosition(), 10)
}

func (w *World) randomPosition() engine.Position {
	return engine.Position{
		X: random.FloatBetween(0, w.width),
		Y: random.FloatBetween(0, w.height),
	}
}
