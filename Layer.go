package main

type Layer struct {
	neurons []Neuron
}

func NewLayer(size int, activator func(float64) float64, neuronSize int) *Layer {
	return &Layer{neurons: generateNeurons(size, activator, neuronSize)}
}

func (layer *Layer) setInputs(inputs []float64) {
	for i := 0; i < len(layer.neurons); i++ {
		layer.neurons[i].setInputs(inputs)
	}
}

func (layer *Layer) activate() []float64 {
	var activated []float64
	for _, neuron := range layer.neurons {
		activated = append(activated, neuron.activate())

	}
	return activated
}

func generateNeurons(size int, activator func(float64) float64, neuronSize int) []Neuron {
	var neurons []Neuron
	for i := 0; i < size; i++ {
		neurons = append(neurons, *newNeuron(activator, neuronSize))
	}
	return neurons
}
