package main

import "fmt"

func main() {
	var fns []func(float642 float64) float64
	fns = append(fns, nothingFunc)
	fns = append(fns, nothingFunc)
	fns = append(fns, nothingFunc)
	nn := NewNN([]int{5, 6, 7}, fns)
	result := nn.compute([]float64{5, 4, 3, 4, 5})
	fmt.Printf("%v", result)

}

func nothingFunc(x float64) float64 {
	return x
}
