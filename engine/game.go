package engine

type Game struct {
	Node
}

type Gamer interface {
	Noder

	Contains(position Vector) bool
}

func (g *Game) GetParent() Noder {
	panic("The game is the engine root object and has no parents")
}
