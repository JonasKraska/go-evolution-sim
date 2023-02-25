package random

import (
	"math/rand"
	"time"
)

var generator *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	generator = rand.New(source)
}

func IntBetween(min, max int) int {
	return generator.Intn(max-min+1) + min
}

func FloatBetween(min, max int) float64 {
	return float64(IntBetween(min, max-1)) + generator.Float64()
}
