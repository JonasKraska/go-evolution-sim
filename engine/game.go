package engine

type Game struct {
	Node
	Grid
}

type Gamer interface {
	Noder

	GetSize() Vector
	GetGrid() *Grid
	Contains(position Vector) bool
}

func (g *Game) GetGrid() *Grid {
	return &g.Grid
}

func (g *Game) GetParent() Noder {
	panic("The game is the engine root object and has no parents")
}
