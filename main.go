package main

import (
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/JonasKraska/go-evolution-sim/simulation"
)

func main() {
	sim := simulation.New(simulation.Config{
		FoodCount:                  64,
		FoodEnergy:                 10,
		FoodGrowthRate:             1,
		FoodProliferationThreshold: 20,

		OrganismCount:                  128,
		OrganismGenes:                  4,
		OrganismEnergy:                 100,
		OrganismMotabolismRate:         5,
		OrganismProliferationThreshold: 150,

		ElitismThreeshold: 20,
	})

	engine.
		New().
		SetZoom(3).
		SetTicksPerSecond(600).
		Run(sim)
}
