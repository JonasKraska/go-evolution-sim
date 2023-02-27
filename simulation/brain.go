package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	NN "github.com/JonasKraska/go-evolution-sim/simulation/neuralnet"
)

type Brain struct {
	NN.Net

	foodDirectionInputNeuron *NN.Neuron
	foodDistanceInputNeuron  *NN.Neuron
	internalNeuron           *NN.Neuron
	directionOutputNeuron    *NN.Neuron
}

func NewBrain() *Brain {
	b := &Brain{
		Net: *NN.New(),

		foodDirectionInputNeuron: NN.NewNeuron(NN.LayerInput, NN.ActivationRandom),
		foodDistanceInputNeuron:  NN.NewNeuron(NN.LayerInput, NN.ActivationRandom),
		internalNeuron:           NN.NewNeuron(NN.LayerInternal, NN.ActivationTanh),
		directionOutputNeuron:    NN.NewNeuron(NN.LayerOutput, NN.ActivationTanh),
	}

	b.AddNeuron(b.foodDirectionInputNeuron)
	b.AddNeuron(b.foodDistanceInputNeuron)
	b.AddNeuron(b.internalNeuron)
	b.AddNeuron(b.directionOutputNeuron)

	b.Connection(b.foodDirectionInputNeuron, b.internalNeuron, random.FloatBetween(-4, 4))
	b.Connection(b.foodDistanceInputNeuron, b.internalNeuron, random.FloatBetween(-4, 4))
	b.Connection(b.internalNeuron, b.directionOutputNeuron, random.FloatBetween(-4, 4))

	return b
}

func (b *Brain) Connection(from, to NN.Neuroner, weight float64) *Brain {
	if _, err := from.ConnectTo(to, weight); err != nil {
		panic(err)
	}

	return b
}

func (b *Brain) Process() *Brain {
	b.Net.Process()

	return b
}

func (b *Brain) GetDirectionChange() float64 {
	return b.directionOutputNeuron.GetOutput()
}
