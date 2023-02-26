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
		Speed: 10,
	}

	fmt.Println(genome.Serialize())

	game := simulation.NewWorld(simulation.WorldConfig{
		Width:  350,
		Height: 225,
		Food: []simulation.FoodCohort{
			{
				Count:  64,
				Energy: 50,
			},
		},
		Organisms: []simulation.OrganismCohort{
			{
				Count:  256,
				Energy: 100,
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
		SetTicksPerSecond(60).
		Run(game)
}
