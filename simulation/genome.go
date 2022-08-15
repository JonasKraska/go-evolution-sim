package simulation

import "image/color"

type Genome struct {
    Color color.Color
}

func FromBytes(bytes []byte) {

}

func (g *Genome) ToBytes() []byte {
    bytes := make([]byte, 10)
    return bytes
}