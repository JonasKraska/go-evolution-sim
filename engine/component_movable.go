package engine

type Movable struct {
	Placeable
	nextPosition Position
}

type Mover interface {
	Placer
	MoveTo(position Position)

	getNextPosition() Position
	cancelMove()
	doMove()
}

func (m *Movable) MoveTo(position Position) {
	m.nextPosition = position
}

func (m *Movable) getNextPosition() Position {
	return m.nextPosition
}

func (m *Movable) cancelMove() {
	m.nextPosition = m.position
}

func (m *Movable) doMove() {
	m.position = m.nextPosition
}
