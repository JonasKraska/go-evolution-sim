package neuralnet

type ActivationFunction = func(sum float64) float64

type Neuron struct {
	input       float64
	output      float64
	connections []*Connection
	activation  ActivationFunction
}

func NewNeuron(activation ActivationFunction) *Neuron {
	neuron := &Neuron{
		activation:  activation,
		connections: make([]*Connection, 3),
	}

	neuron.Reset()

	return neuron
}

type Connection struct {
	from   *Neuron
	to     *Neuron
	weight float64
}

func (n *Neuron) Reset() *Neuron {
	n.input = 0.0
	n.output = 0.0

	return n
}

func (n *Neuron) ConnectTo(other *Neuron, weight float64) *Neuron {
	connection := &Connection{
		from:   n,
		to:     other,
		weight: weight,
	}

	n.connections = append(n.connections, connection)
	other.connections = append(other.connections, connection)

	return n
}

func (n *Neuron) Process() *Neuron {
	sum := 0.0

	for _, c := range n.connections {
		if c.to == n {
			sum += c.from.output * c.weight
		}
	}

	n.output = n.activation(sum)

	return n
}
