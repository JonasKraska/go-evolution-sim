package engine

type Node struct {
	Nestable
	Hookable
}

type Noder interface {
	Nester
	Hooker
}
