package engine

import "math"

type Position = Vector

type Placeable struct {
	position Position
}

type Placer interface {
	SetPosition(position Position)
	GetPosition() Position
}

func (p *Placeable) SetPosition(position Position) {
	p.position = position
}

func (p *Placeable) GetPosition() Position {
	return p.position
}

func (v Position) EqualsIgnoringDecimals(other Position) bool {
	return math.Floor(v.X) == math.Floor(other.X) && math.Floor(v.Y) == math.Floor(other.Y)
}
