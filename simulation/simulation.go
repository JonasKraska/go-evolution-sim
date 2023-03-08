package simulation

import (
	"fmt"
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

const (
	MinOrganismCount  = 25
	FoodSpawnInterval = 1.5 // seconds
)

type Simulation struct {
	engine.Game

	config            Config
	world             *World
	generation        int
	lastFoodSpawnedAt time.Time
}

func New(config Config) *Simulation {
	config = NewConfig(config)
	world := NewWorld(config.Width, config.Height)

	s := &Simulation{
		config:            config,
		world:             world,
		lastFoodSpawnedAt: time.Now(),
	}

	s.AddChild(s.world)
	s.spawnGeneration()

	world.Register(WorldOrganismRemoved, func() {
		if len(world.organisms) <= MinOrganismCount {
			s.spawnGeneration()
		}
	})

	return s
}

func (s *Simulation) Update(delta time.Duration) {
	engine.DebugPrintln(fmt.Sprintf("Generation: %d", s.generation))

	if time.Since(s.lastFoodSpawnedAt).Seconds() > FoodSpawnInterval {
		s.world.spawnFood(world.randomPosition(), s.config.FoodEnergy)
		s.lastFoodSpawnedAt = time.Now()
	}
}

func (s *Simulation) Draw() *ebiten.Image {
	return s.world.Draw()
}

func (s *Simulation) GetSize() engine.Vector {
	return s.world.GetSize()
}

func (s *Simulation) Contains(position engine.Vector) bool {
	return s.world.Contains(position)
}

func (s *Simulation) spawnGeneration() {
	// @TODO copy the most successful organisms from old generation and mutate them at a certain chance

	// replenish existing foods
	for _, ef := range s.world.foods {
		ef.energy = s.config.FoodEnergy
	}

	// replenish existing organisms
	for _, eo := range s.world.organisms {
		eo.energy = s.config.OrganismEnergy
	}

	// spawn new foods until config is reached
	for nf := len(s.world.foods); nf < int(s.config.FoodCount); nf++ {
		s.world.spawnFood(
			s.world.randomPosition(),
			s.config.FoodEnergy,
		)
	}

	// spawn new organims until config is reached
	for no := len(s.world.organisms); no < int(s.config.OrganismCount); no++ {
		s.world.spawnOrganism(
			world.randomPosition(),
			NewGenome(s.config.OrganismGenes),
			s.config.OrganismEnergy,
		)
	}

	s.generation++
}
