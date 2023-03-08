package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine"
	NN "github.com/JonasKraska/go-evolution-sim/simulation/neuralnet"
)

type NeuronType = uint8

const (
	InputNeuronFoodDirection NeuronType = iota
	InputNeuronFoodDistance
	InternalNeuron
	OutputNeuronDirection

	RegisteredNeuronCount = 4
)

type Brain struct {
	NN.Net
	neurons []*NN.Neuron
}

func NewBrain() *Brain {
	b := &Brain{
		Net:     *NN.New(),
		neurons: make([]*NN.Neuron, RegisteredNeuronCount),
	}

	// @TODO add input neuron for distance to obstancle (world boundaries for now)
	b.neurons[InputNeuronFoodDirection] = NN.NewNeuron(NN.LayerInput, NN.ActivationRandom)
	b.neurons[InputNeuronFoodDistance] = NN.NewNeuron(NN.LayerInput, NN.ActivationRandom)
	b.neurons[InternalNeuron] = NN.NewNeuron(NN.LayerInternal, NN.ActivationTanh)
	b.neurons[OutputNeuronDirection] = NN.NewNeuron(NN.LayerOutput, NN.ActivationTanh)

	for _, n := range b.neurons {
		b.AddNeuron(n)
	}

	return b
}

func (b *Brain) Connection(from NeuronType, to NeuronType, weight float64) *Brain {
	neuronFrom := b.neurons[from]
	neuronTo := b.neurons[to]

	neuronFrom.ConnectTo(neuronTo, weight)

	return b
}

func (b *Brain) Prune() {
	b.Net.Prune()
}

func (b *Brain) Process() {
	b.Net.Process()
}

func (b *Brain) SetFoodDistance(distance float64) {
	b.neurons[InputNeuronFoodDistance].SetActivation(func(sum float64) float64 {
		return distance / OrganismViewRange
	})
}

func (b *Brain) SetFoodAngle(angle engine.Angle) {
	b.neurons[InputNeuronFoodDirection].SetActivation(func(sum float64) float64 {
		return angle.GetDeg() / OrganismFieldOfView / 2
	})
}

func (b *Brain) GetDirectionAngle() engine.Angle {
	directionChangeAngle := b.neurons[OutputNeuronDirection].GetOutput() * OrganismMaxTurnDeg
	return engine.NewAngleDeg(directionChangeAngle)
}

func NeuronTypeToString(neuronType NeuronType) string {
	switch neuronType {
	case InputNeuronFoodDirection:
		return "InputNeuronFoodDirection"
	case InputNeuronFoodDistance:
		return "InputNeuronFoodDistance"
	case InternalNeuron:
		return "InternalNeuron"
	case OutputNeuronDirection:
		return "OutputNeuronDirection"
	}

	return "unkown"
}
