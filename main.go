package main

import (
	"fmt"
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/JonasKraska/go-evolution-sim/simulation"
	"image/color"
)

func main() {

	genome := simulation.Genome{
		Color: color.RGBA{
			R: 255, //uint8(random.IntBetween(50, 250)),
			G: 255, //uint8(random.IntBetween(50, 250)),
			B: 255, //uint8(random.IntBetween(50, 250)),
		},
		MetabolismRate: 100,
	}

	fmt.Println(genome.Serialize())

	game := simulation.NewWorld(simulation.WorldConfig{
		Width:  350,
		Height: 225,
		Food: []simulation.FoodCohort{
			{
				Count:  32,
				Energy: simulation.Energy(5),
			},
		},
		Organisms: []simulation.OrganismCohort{
			{
				Count:  256,
				Energy: simulation.Energy(100),
				Genome: genome,
			},
		},
	})

	simulation.New(
		simulation.Config{},
		game,
	)

	engine.
		New().
		SetZoom(3).
		SetTicksPerSecond(10).
		Run(game)
}
