package engine

type Nestable struct {
	parent           Noder
	children         []Noder
	markedForRemoval bool
}

type Nester interface {
	GetParent() Noder
	GetChildren() []Noder
	AddChild(child Noder)
	IsMarkedForRemoval() bool
	Remove()

	setParent(parent Noder)
	removeChildren()
}

func (n *Node) GetParent() Noder {
	return n.parent
}

func (n *Node) GetChildren() []Noder {
	return n.children
}

func (n *Node) AddChild(child Noder) {
	child.setParent(n)
	n.children = append(n.children, child)
}

func (n *Node) IsMarkedForRemoval() bool {
	return n.markedForRemoval
}

func (n *Node) Remove() {
	n.markedForRemoval = true
}

func (n *Node) setParent(parent Noder) {
	n.parent = parent
}

func (n *Node) removeChildren() {
	index := 0
	for _, c := range n.children {
		if c.IsMarkedForRemoval() {
			if hooker, ok := c.(Hooker); ok {
				hooker.Dispatch(NodeRemoveHook)
			}

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
