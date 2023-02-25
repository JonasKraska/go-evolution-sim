package simulation

import (
	"time"
)

type Simulation struct {
	name  string
	world *World
}

type Config struct {
	Name string
}

func New(config Config, world *World) *Simulation {
	if config.Name == "" {
		config.Name = "simulation_" + time.Now().Format("2006-01-02") + "_" + time.Now().Format("15:04:05")
	}

	return &Simulation{
		name:  config.Name,
		world: world,
	}
}
