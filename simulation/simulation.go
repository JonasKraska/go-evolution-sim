package simulation

import (
	"fmt"
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

const (
	MinOrganismCount = 25
)

type Simulation struct {
	engine.Game

	name       string
	world      *World
	generation int
}

type Config struct {
	Name string
}

func (s *Simulation) Update(delta time.Duration) {
	engine.DebugPrintln(fmt.Sprintf("Generation: %d", s.generation))
}

func (s *Simulation) Draw() *ebiten.Image {
	return s.world.Draw()
}

func New(worldConfig WorldConfig) *Simulation {
	s := &Simulation{
		name:       "simulation_" + time.Now().Format("2006-01-02") + "_" + time.Now().Format("15:04:05"),
		world:      NewWorld(worldConfig),
		generation: 1,
	}

	s.AddChild(s.world)

	return s
}

func (s *Simulation) GetSize() engine.Vector {
	return s.world.GetSize()
}

func (s *Simulation) Contains(position engine.Vector) bool {
	return s.world.Contains(position)
}
