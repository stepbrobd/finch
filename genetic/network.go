package genetic


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
        Inputs  []float32
        Outputs []float32
}


func NewNet(sizes []int) Net {
        var n Net
        n.Error = 0.0
        n.Sizes = make([]int, len(sizes))
        copy(n.Sizes, sizes)
        // Calculate slice sizes bases on given layer sizes
        leng := 0
        max := sizes[0]
        for lay := 1; lay < len(sizes); lay++ {
                if sizes[lay] > max {
                        max = sizes[lay]
                }
                // Number of weights and biases needed for layer
                leng += sizes[lay-1]*sizes[lay] + sizes[lay]
        }
	// Allocate elements for weights and biases.
        n.Values = make([]float32, leng, leng)
        // Allocate elements for inputs and outputs.
        n.Inputs = make([]float32, max, max)
        n.Outputs = make([]float32, max, max)
        return n
}


func (n Net) ForProp(inputs []float32) []float32 {
        // First layer uses inputs as activations
        valIdx := 0
        valIdx += Mult(n.Values[valIdx:], inputs, n.Outputs, n.Sizes[1], n.Sizes[0])
        valIdx += Add(n.Values[valIdx:], n.Outputs, n.Inputs, n.Sizes[1])
        ReLU(n.Inputs, n.Sizes[1])
        for lay := 2; lay < len(n.Sizes); lay++ {
                // All other layers use prior layers activations
                valIdx += Mult(n.Values[valIdx:], n.Inputs, n.Outputs, n.Sizes[lay], n.Sizes[lay-1])
                valIdx += Add(n.Values[valIdx:], n.Outputs, n.Inputs, n.Sizes[lay])
                ReLU(n.Inputs, n.Sizes[lay])
        }
        copy(n.Outputs, n.Inputs)
        return n.Outputs
}


func (n Net) FitFunc(examInputs, expecOutputs [][]float32) float64 {
        n.Error = 0.0
        for num := 0; num < len(examInputs); num++ {
                outputs := n.ForProp(examInputs[num])
                for idx := 0; idx < n.Sizes[len(n.Sizes)-1]; idx++ {
                        n.Error += float64(Abs(expecOutputs[num][idx] - outputs[idx]))
                }
        }
	return n.Error
}

