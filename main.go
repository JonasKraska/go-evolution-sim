package main

import (
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	"github.com/JonasKraska/go-evolution-sim/simulation"
	"image/color"
)

func main() {
	game := simulation.NewWorld(simulation.WorldConfig{
		Width:        350,
		Height:       225,
		NumberOfFood: 32,
		Organisms: []simulation.OrganismCohort{
			{
				Count: 256,
				Config: simulation.OrganismConfig{
					Color: color.RGBA{
						R: uint8(random.IntBetween(50, 250)),
						G: uint8(random.IntBetween(50, 250)),
						B: uint8(random.IntBetween(50, 250)),
					},
				},
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
