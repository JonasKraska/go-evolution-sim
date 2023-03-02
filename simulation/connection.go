package simulation

import (
	"encoding/binary"
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	"math"
)

const ConnectionByteSize int = 10
const ConnectionIdPadding = 255 / RegisteredNeuronCount

type Connection struct {
	From   NeuronType
	To     NeuronType
	Weight float64
}

func NewConnection() Connection {
	return Connection{
		From:   NeuronType(random.IntBetween(0, RegisteredNeuronCount-1)),
		To:     NeuronType(random.IntBetween(0, RegisteredNeuronCount-1)),
		Weight: random.FloatBetween(-4, 4),
	}
}

func DeserializeConnection(bytes []byte) Connection {
	return Connection{
		From:   uint8(bytes[0]) / ConnectionIdPadding,
		To:     uint8(bytes[1]) / ConnectionIdPadding,
		Weight: math.Float64frombits(binary.BigEndian.Uint64(bytes[2:])),
	}
}

func (c Connection) Serialize() []byte {
	bytes := make([]byte, 2)

	bytes[0] = c.From * ConnectionIdPadding
	bytes[1] = c.To * ConnectionIdPadding

	weightBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(weightBytes[:], math.Float64bits(c.Weight))

	bytes = append(bytes, weightBytes...)

	return bytes
}
