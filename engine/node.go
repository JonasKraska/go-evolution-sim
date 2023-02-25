package engine

type Node struct {
	parent   Noder
	children []Noder
}

type Noder interface {
	GetChildren() []Noder
	AddChild(child Noder)
	RemoveChild(child Noder)

	GetParent() *Noder
	setParent(parent Noder)
}

func (n *Node) GetChildren() []Noder {
	return n.children
}

func (n *Node) AddChild(child Noder) {
	child.setParent(n)
	n.children = append(n.children, child)
}

func (n *Node) RemoveChild(child Noder) {
	index := 0
	for _, c := range n.children {
		if c == child {
			continue
		}

		n.children[index] = c
		index++
	}

	// clean up to prevent memory leaks
	for i := index; i < len(n.children); i++ {
		n.children[i] = nil
	}

	// resize resize slice to remove nil pointers
	n.children = n.children[:index]
}

func (n *Node) GetParent() *Noder {
	return &n.parent
}

func (n *Node) setParent(parent Noder) {
	n.parent = parent
}
