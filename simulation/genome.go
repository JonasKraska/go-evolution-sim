package simulation

import "image/color"

type Genome struct {
    Color color.RGBA
}

func NewGenome(g Genome) Genome {
    g.Color.A = 255

    return g
}

func FromBytes(bytes []byte) {

}

func (g Genome) ToBytes() []byte {
    bytes := make([]byte, 10)



    return bytes
}