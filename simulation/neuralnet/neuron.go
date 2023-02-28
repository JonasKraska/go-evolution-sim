package neuralnet

import "errors"

type Neuron struct {
	layer       Layer
	input       float64
	output      float64
	connections []*Connection
	activation  ActivationFunction
}

func NewNeuron(layer Layer, activation ActivationFunction) *Neuron {
	neuron := &Neuron{
		layer:       layer,
		activation:  activation,
		connections: make([]*Connection, 0),
	}

	return neuron
}

type Neuroner interface {
	GetLayer() Layer
	GetOutput() float64
	ConnectTo(other Neuroner, weight float64) (Neuroner, error)
	Process() Neuroner

	connectFrom(connection *Connection) Neuroner
	reset() Neuroner
}

type Connection struct {
	from   Neuroner
	to     Neuroner
	weight float64
}

func (n *Neuron) GetLayer() Layer {
	return n.layer
}

func (n *Neuron) GetOutput() float64 {
	return n.output
}

func (n *Neuron) ConnectTo(other Neuroner, weight float64) (Neuroner, error) {
	if other.GetLayer() == LayerInput {
		return nil, errors.New("no neuron can connect to another input neuron")
	}

	if n.layer == LayerOutput {
		return nil, errors.New("output neuron can not connect to any other neuron")
	}

	connection := &Connection{
		from:   n,
		to:     other,
		weight: weight,
	}

	n.connections = append(n.connections, connection)
	other.connectFrom(connection)

	return n, nil
}

func (n *Neuron) Process() Neuroner {
	n.reset()
	sum := 0.0

	for _, c := range n.connections {
		if c.to == n {
			sum += c.from.GetOutput() * c.weight
		}
	}

	n.output = n.activation(sum)

	return n
}

func (n *Neuron) connectFrom(connection *Connection) Neuroner {
	n.connections = append(n.connections, connection)

	return n
}

func (n *Neuron) reset() Neuroner {
	n.input = 0.0
	n.output = 0.0

	return n
}
