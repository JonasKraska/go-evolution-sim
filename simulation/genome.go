package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	"image/color"
)

const GenomeStaticByteSize int = 3

type Genome struct {
	Color       color.RGBA
	Connections []Connection
}

func NewGenome(genes uint8) Genome {
	genome := Genome{
		Color: color.RGBA{
			R: uint8(random.IntBetween(50, 200)),
			G: uint8(random.IntBetween(50, 200)),
			B: uint8(random.IntBetween(50, 200)),
			A: 255,
		},
		Connections: make([]Connection, genes),
	}

	for c := 0; c < int(genes); c++ {
		genome.Connections[c] = NewConnection()
	}

	return genome
}

func DeserializeGenome(bytes []byte) Genome {
	genes := (len(bytes) - GenomeStaticByteSize) / ConnectionByteSize

	genome := Genome{
		Color: color.RGBA{
			R: bytes[0],
			G: bytes[1],
			B: bytes[2],
			A: 255,
		},
		Connections: make([]Connection, genes),
	}

	for g := 0; g < genes; g++ {
		from := GenomeStaticByteSize + (g * ConnectionByteSize)
		to := from + ConnectionByteSize

		genome.Connections[g] = DeserializeConnection(bytes[from:to])
	}

	return genome
}

func (g Genome) Copy() Genome {
	return g
}

func (g Genome) Serialize() []byte {
	bytes := make([]byte, GenomeStaticByteSize)

	bytes[0] = g.Color.R
	bytes[1] = g.Color.G
	bytes[2] = g.Color.B

	for _, c := range g.Connections {
		bytes = append(bytes, c.Serialize()...)
	}

	return bytes[:]
}

func (g Genome) PointMutation() Genome {
	bytes := g.Serialize()

	byteIndex := random.IntBetween(0, len(bytes)-1)
	bytes[byteIndex] ^= 1 << random.IntBetween(0, 7)

	return DeserializeGenome(bytes)
}
