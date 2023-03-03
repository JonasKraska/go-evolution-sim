package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	"math"
)

const ConnectionByteSize int = 3
const ConnectionIdPadding = float64(256) / RegisteredNeuronCount

type Connection struct {
	from   uint8
	to     uint8
	weight uint8
}

func NewConnection() Connection {
	return Connection{
		from:   uint8(random.IntBetween(0, 255)),
		to:     uint8(random.IntBetween(0, 255)),
		weight: uint8(random.IntBetween(0, 255)),
	}
}

func DeserializeConnection(bytes []byte) Connection {
	return Connection{
		from:   bytes[0],
		to:     bytes[1],
		weight: bytes[2],
	}
}

func (c Connection) Serialize() []byte {
	return []byte{
		c.from,
		c.to,
		c.weight,
	}
}

func (c Connection) GetFrom() NeuronType {
	return NeuronType(float64(c.from) / ConnectionIdPadding)
}

func (c Connection) GetTo() NeuronType {
	return NeuronType(float64(c.to) / ConnectionIdPadding)
}

func (c Connection) GetWeight() float64 {
	return math.Tanh(float64(uint16(c.weight)-128)/64) * 4
}
