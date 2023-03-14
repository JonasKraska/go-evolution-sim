package engine

type Point struct {
	// @TODO use float64 as well
	X, Y int
}

func (p Point) vectorTo(other Point) Vector {
	return Vector{
		X: float64(other.X - p.X),
		Y: float64(other.Y - p.Y),
	}
}

func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}
