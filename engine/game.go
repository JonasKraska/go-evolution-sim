package engine

type Game struct {
	Node
}

type Gamer interface {
	Noder

	Contains(Position) bool
}

func (g *Game) GetParent() Noder {
	panic("The game is the engine root object and has no parents")
}
