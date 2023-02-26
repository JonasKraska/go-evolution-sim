package neuralnet

type Net struct {
	neurons []*Neuron
}

func New() *Net {
	return &Net{
		neurons: make([]*Neuron, 8),
	}
}

func (n *Net) AddNeuron() *Neuron {
	neuron := &Neuron{}

	n.neurons = append(n.neurons, neuron)

	return neuron
}
