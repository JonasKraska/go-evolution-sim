package engine

type Nestable struct{
    children []any
}

type Nester interface {
    GetChildren() []any
    AddChild(other any)
}

func (n *Nestable) GetChildren() []any {
    return n.children
}

func (n *Nestable) AddChild(other any) {
    n.children = append(n.children, other)
}
