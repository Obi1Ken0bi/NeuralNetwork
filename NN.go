package main

func generateLayers(layerSizes []int, activators []func(float64) float64) []Layer {
	var layers []Layer
	var layer Layer = *NewLayer(layerSizes[0], activators[0], 0)
	layers = append(layers, layer)
	for i := 1; i < len(layerSizes); i++ {
		layer = *NewLayer(layerSizes[i], activators[i], layerSizes[i-1])
		layers = append(layers, layer)
	}
	return layers

}

type NN struct {
	layers []Layer
}

func NewNN(layerSizes []int, activators []func(float64) float64) *NN {
	return &NN{layers: generateLayers(layerSizes, activators)}
}

func (n NN) compute(params []float64) []float64 {
	n.layers[0] = *NewLayer(len(n.layers[0].neurons), n.layers[0].neurons[0].activator, len(params))
	n.layers[0].setInputs(params)
	output := n.layers[0].activate()
	for i := 1; i < len(n.layers); i++ {
		n.layers[i].setInputs(output)
		output = n.layers[i].activate()
	}
	return output
}
