package simulation

import (
	"github.com/JonasKraska/go-evolution-sim/engine"
	"github.com/JonasKraska/go-evolution-sim/engine/random"
	NN "github.com/JonasKraska/go-evolution-sim/simulation/neuralnet"
)

type Brain struct {
	NN.Net

	inputNeuronFoodDirection *NN.Neuron
	inputNeuronFoodDistance  *NN.Neuron

	internalNeuron *NN.Neuron

	outputNeuronDirection *NN.Neuron
}

func NewBrain() *Brain {
	b := &Brain{
		Net: *NN.New(),

		inputNeuronFoodDirection: NN.NewNeuron(NN.LayerInput, NN.ActivationRandom),
		inputNeuronFoodDistance:  NN.NewNeuron(NN.LayerInput, NN.ActivationRandom),

		internalNeuron: NN.NewNeuron(NN.LayerInternal, NN.ActivationTanh),

		outputNeuronDirection: NN.NewNeuron(NN.LayerOutput, NN.ActivationTanh),
	}

	b.AddNeuron(b.inputNeuronFoodDirection)
	b.AddNeuron(b.inputNeuronFoodDistance)
	b.AddNeuron(b.internalNeuron)
	b.AddNeuron(b.outputNeuronDirection)

	b.Connection(b.inputNeuronFoodDirection, b.internalNeuron, random.FloatBetween(-4, 4))
	b.Connection(b.inputNeuronFoodDistance, b.internalNeuron, random.FloatBetween(-4, 4))
	b.Connection(b.internalNeuron, b.outputNeuronDirection, random.FloatBetween(-4, 4))

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

func (b *Brain) SetFoodDistance(distance float64) {
	b.inputNeuronFoodDirection.SetActivation(func(sum float64) float64 {
		return distance / OrganismViewRange
	})
}

func (b *Brain) SetFoodAngle(angle engine.Angle) {
	b.inputNeuronFoodDirection.SetActivation(func(sum float64) float64 {
		return angle.GetDeg() / OrganismFieldOfView / 2
	})
}

func (b *Brain) GetDirectionAngle() engine.Angle {
	directionChangeAngle := b.outputNeuronDirection.GetOutput() * OrganismMaxTurnDeg
	return engine.NewAngleDeg(directionChangeAngle)
}
