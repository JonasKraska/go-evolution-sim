package main

import (
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/JonasKraska/go-evolution-sim/simulation"
)

func main() {
	simulation := simulation.New(simulation.WorldConfig{
		Width:  350,
		Height: 225,

		FoodCount:  512,
		FoodEnergy: 50,

		OrganismCount:  256,
		OrganismGenes:  3,
		OrganismEnergy: 100,
	})

	engine.
		New().
		SetZoom(3).
		SetTicksPerSecond(180).
		Run(simulation)
}
