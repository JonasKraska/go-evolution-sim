package engine

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
