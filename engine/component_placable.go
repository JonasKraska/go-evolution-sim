package engine

type Position = Vector

type Placeable struct {
	Node
	position Position
}

type Placer interface {
	Noder
	SetPosition(position Position)
	GetPosition() Position
}

func (p *Placeable) SetPosition(position Position) {
	p.position = position
}

func (p *Placeable) GetPosition() Position {
	return p.position
}
