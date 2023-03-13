package neuralnet

type Net struct {
	neurons map[Layer][]Neuroner
}

func New() *Net {
	return &Net{
		neurons: make(map[Layer][]Neuroner),
	}
}

func (n *Net) AddNeuron(neuron Neuroner) *Net {
	layer := neuron.GetLayer()

	if n.neurons[layer] == nil {
		n.neurons[layer] = make([]Neuroner, 0)
	}

	n.neurons[layer] = append(n.neurons[layer], neuron)

	return n
}

func (n *Net) Prune() *Net {
	// @TODO remove all uneccessary connections that not contribute to any output to speed up `Process()`
	return n
}

func (n *Net) Process() *Net {

	for _, input := range n.neurons[LayerInput] {
		input.Process()
	}

	// @TODO internal neurons can have conenctions to themselves
	for _, internal := range n.neurons[LayerInternal] {
		internal.Process()
	}

	for _, output := range n.neurons[LayerOutput] {
		output.Process()
	}

	return n
}
