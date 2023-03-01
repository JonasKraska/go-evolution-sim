package engine

type Movable struct {
	Placeable
	velocity     Vector
	lastPosition Position
}

type Mover interface {
	Placer
	SetVelocity(velocity Vector)
	GetVelocity() Vector
	GetLastPosition() Position

	cancelMove()
	doMove()
}

func (m *Movable) SetVelocity(velocity Vector) {
	m.velocity = velocity
}

func (m *Movable) GetVelocity() Vector {
	return m.velocity
}

func (m *Movable) GetLastPosition() Position {
	return m.lastPosition
}

func (m *Movable) cancelMove() {
	m.position = m.lastPosition
}

func (m *Movable) doMove() {
	m.lastPosition = m.position
	m.position = m.position.Add(m.velocity)
}
