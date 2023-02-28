package neuralnet

import (
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	"math"
)

type ActivationFunction = func(sum float64) float64

func ActivationTanh(sum float64) float64 {
	return math.Tanh(sum)
}

func ActivationRandom(sum float64) float64 {
	return random.FloatBetween(-1, 1)
}
