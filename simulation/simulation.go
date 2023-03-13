package simulation

import (
	"fmt"
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Simulation struct {
	engine.Game

	config Config
	world  *World

	generation        int
	lastFoodSpawnedAt time.Time
}

var simulation *Simulation

func New(config Config) *Simulation {
	config = NewConfig(config)
	world := NewWorld(config.Width, config.Height)

	simulation = &Simulation{
		config:            config,
		world:             world,
		lastFoodSpawnedAt: time.Now(),
	}

	simulation.AddChild(simulation.world)
	simulation.spawnGeneration()

	world.Register(WorldOrganismRemoved, func() {
		if len(world.organisms) <= int(config.ElitismThreeshold) {
			simulation.spawnGeneration()
		}
	})

	return simulation
}

func (s *Simulation) Update(delta time.Duration) {
	engine.DebugPrintln(fmt.Sprintf("Generation: %d", s.generation))
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
		ef.SetPosition(s.world.randomPosition())
		ef.energy = s.config.FoodEnergy
	}

	// replenish existing organisms
	for _, eo := range s.world.organisms {
		eo.SetPosition(s.world.randomPosition())
		eo.energy = s.config.OrganismEnergy
	}

	// spawn new foods until config is reached
	for nf := len(s.world.foods); nf < int(s.config.FoodCount); nf++ {
		s.world.spawnFood(
			s.world.randomPosition(),
			s.config.FoodEnergy,
		)
	}

	// despawn over supply on foods until config is reached
	for of := len(s.world.foods) - 1; of >= int(s.config.FoodCount); of-- {
		s.world.foods[of].Remove()
	}

	// copy most successful organims
	existingOrganismCount := len(s.world.organisms)
	for co := 0; co < existingOrganismCount; co++ {
		s.world.spawnOrganism(
			world.randomPosition(),
			s.world.organisms[co].genome.PointMutation(),
			s.config.OrganismEnergy,
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

	// despawn over populated organisms until config is reached
	//	for oo := len(s.world.organisms) - 1; oo >= int(s.config.OrganismCount); oo-- {
	//		s.world.organisms[oo].Remove()
	//	}

	s.generation++
}
