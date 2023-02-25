package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	"image/color"
)

type Genome struct {
	Color          color.RGBA
	MetabolismRate uint8
}

func NewGenome(g Genome) Genome {
	g.Color.A = 255

	return g
}

func DeserializeGenome(bytes []byte) Genome {
	return NewGenome(Genome{
		Color: color.RGBA{
			R: bytes[0],
			G: bytes[1],
			B: bytes[2],
		},
		MetabolismRate: bytes[3],
	})
}

func (g Genome) Serialize() []byte {
	bytes := make([]byte, 4)

	bytes[0] = g.Color.R
	bytes[1] = g.Color.G
	bytes[2] = g.Color.B

	bytes[3] = g.MetabolismRate

	return bytes
}

func (g Genome) PointMutation() Genome {
	bytes := g.Serialize()

	byteIndex := random.IntBetween(0, len(bytes)-1)
	bytes[byteIndex] ^= 1 << random.IntBetween(0, 7)

	return DeserializeGenome(bytes)
}
