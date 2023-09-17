package genetic

import "sync"

type Net struct {
	// Holds total error of network
	Error float64
	// Holds sizes of layer
	// First index holds the number of neurons in input layer
	// Last index holds the number of neurons in output layer
	Sizes []int
	// Holds weights and biases in the following linear order:
	// 1ST_HIDDEN_WEIGHTS 1ST_HIDDEN_BIASES 2ND_HIDDEN_WEIGHTS 2ND_HIDDEN_BIASES ... OUTPUT_WEIGHTS OUTPUT>
	Values []float32
	// Holds intermediate calculations for forward propagation
	Outputs []float32
}

func NewNet(sizes []int) Net {
	var n Net
	n.Error = 0.0
	n.Sizes = make([]int, len(sizes))
	copy(n.Sizes, sizes)
	// Calculate slice sizes bases on given layer sizes
	lenVals := 0
	lenOuts := 0
	for lay := 1; lay < len(sizes); lay++ {
		// Number of Outputs need for layer
		lenOuts += sizes[lay]
		// Number of weights and biases needed for layer
		lenVals += sizes[lay-1]*sizes[lay] + sizes[lay]
	}
	// Allocate elements for weights and biases.
	n.Values = make([]float32, lenVals, lenVals)
	for idx := 0; idx < lenVals; idx++ {
		n.Values[idx] = RandFloatRange(-0.1, 0.1)
	}
	// Allocate elements for outputs.
	n.Outputs = make([]float32, lenOuts, lenOuts)
	return n
}

func (n *Net) ForProp(inputs []float32) []float32 {
	// First layer uses inputs as activations
	valIdx := 0
	insIdx := 0
	outIdx := 0
	valIdx += Mult(n.Values[valIdx:], inputs, n.Outputs, n.Sizes[1], n.Sizes[0])
	valIdx += Add(n.Values[valIdx:], n.Outputs, n.Outputs, n.Sizes[1])
	outIdx += ReLU(n.Outputs, n.Sizes[1])
	for lay := 2; lay < len(n.Sizes); lay++ {
		// All other layers use prior layers activations
		valIdx += Mult(n.Values[valIdx:], n.Outputs[insIdx:], n.Outputs[outIdx:], n.Sizes[lay], n.Sizes[lay-1])
		valIdx += Add(n.Values[valIdx:], n.Outputs[outIdx:], n.Outputs[outIdx:], n.Sizes[lay])
		insIdx = outIdx
		outIdx += ReLU(n.Outputs[outIdx:], n.Sizes[lay])
	}
	return n.Outputs[insIdx:]
}

func (n *Net) FitFunc(examInputs, expecOutputs [][]float32, wg *sync.WaitGroup) {
	n.Error = 0.0
	for num := 0; num < len(examInputs); num++ {
		outputs := n.ForProp(examInputs[num])
		for idx := 0; idx < n.Sizes[len(n.Sizes)-1]; idx++ {
			n.Error += float64(Abs(expecOutputs[num][idx] - outputs[idx]))
		}
	}
	wg.Done()
}

func (n *Net) Crossover(male, female Net, wg *sync.WaitGroup) {
	copy(n.Values, female.Values)
	copy(n.Values, male.Values[:RandIntRange(1, len(female.Values)-1)])
	wg.Done()
}

func (n *Net) Mutate(rate float32, wg *sync.WaitGroup) {
	for idx := 0; idx < len(n.Values); idx++ {
		if RandFloatRange(0.0, 1.0) <= rate {
			n.Values[idx] += RandFloatRange(-0.25, 0.25)
		}
	}
	wg.Done()
}
