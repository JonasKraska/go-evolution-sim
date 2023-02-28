package engine

type Point struct {
	X, Y int
}

func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}
