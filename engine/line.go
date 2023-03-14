package engine

import (
	"errors"
	"math"
)

type Line struct {
	from, to Point
}

func (l Line) Intersect(other Line) (Point, error) {
	v := l.from.vectorTo(other.from)
	v1 := l.toVector()
	v2 := other.toVector()

	cp := Cross(v1, v2)
	if cp == 0 {
		return Point{}, errors.New("no intersection found")
	}

	cp1 := Cross(v, v1)
	cp2 := Cross(v, v2)

	t1 := cp2 / cp
	t2 := cp1 / cp

	e := math.SmallestNonzeroFloat64
	if t1+e < 0 || t1-e > 1 || t2+e < 0 || t2-e > 1 {
		return Point{}, errors.New("no intersection found")
	}

	factor := v1.MulScalar(t1)

	return Point{
		X: l.from.X + int(factor.X),
		Y: l.from.X + int(factor.Y),
	}, nil
}

func (l Line) toVector() Vector {
	return l.from.vectorTo(l.to)
}
