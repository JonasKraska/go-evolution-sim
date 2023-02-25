package engine

const (
	NodeRemoveHook = Hook("node.remove")
)

type Hook string
type HookCallback func()

type Hookable struct {
	callbacks map[Hook][]HookCallback
}

type Hooker interface {
	Register(hook Hook, cb HookCallback)
	Dispatch(hook Hook)
}

func (n *Node) Register(hook Hook, cb HookCallback) {
	if n.callbacks == nil {
		n.callbacks = make(map[Hook][]HookCallback)
	}

	n.callbacks[hook] = append(n.callbacks[hook], cb)
}

func (n *Node) Dispatch(hook Hook) {
	for _, cb := range n.callbacks[hook] {
		cb()
	}
}
