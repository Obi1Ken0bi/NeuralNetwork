package main

type Neuron struct {
	activator func(float64) float64
	weights   []float64
	bias      float64
	input     []float64
}

func newNeuron(activator func(float64) float64, size int) *Neuron {
	n := new(Neuron)
	n.activator = activator
	n.weights = randomizeWeights(size)
	n.bias = randomizeBias()
	return n
}

func randomizeWeights(size int) []float64 {
	var w []float64
	for i := 0; i < size; i++ {
		w = append(w, 2) //rand.Float64()+1)
	}
	return w
}

func randomizeBias() float64 {
	b := float64(5) //rand.Float64()
	return b
}

func (n *Neuron) sum() float64 {
	var sum float64
	for i := 0; i < len(n.input); i++ {
		sum += n.input[i] * n.weights[i]
	}
	sum += n.bias
	return sum
}
func (n *Neuron) activate() float64 {
	res := n.sum()
	return n.activator(res)
}

func (n *Neuron) setInputs(input []float64) {
	n.input = input
}
