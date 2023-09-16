package genetic

type Algo struct {
	Generations     int
	ExampleInputs   [][]float32
	ExpectedOutputs [][]float32
	MutationRate    float32
	Population      Pop
}

// mutationRate: The Percent chance of a given weight/bias mutating. Must be between 0.0 and 1.0
// individuals: The number of individuals in the population per generation
// networkLayers: defines the structure of the networks (individuals). First index --> Number of neurons in input layer. Last index --> Number of neurons in output layer
// exampleInputs: The examples used to by the fitness function to calculate a given networks total error
// expectedOutputs: The corresponeding correct outputs to the given exampleInputs
// Returns: A initialized algorithm
func NewAlgo(mutationRate float32, individuals int, networkLayers []int, exampleInputs [][]float32, expectedOutputs [][]float32) Algo {
	var a Algo
	a.Generations = 0
	a.MutationRate = mutationRate
	a.Population = NewPop(individuals, networkLayers)
	a.ExampleInputs = exampleInputs
	a.ExpectedOutputs = expectedOutputs
	return a
}

// Returns: the number of generations that have been ran.
func (a *Algo) GetNumGens() int {
	return a.Generations
}

func (a *Algo) GetWeights() [][]float32 {
	network := a.Population.Nets[0]
	layers := make([][]float32, 0)
	valIdx := 0
	for idx := 1; idx < len(network.Sizes); idx++ {
		numberWeights := network.Sizes[idx-1] * network.Sizes[idx]
		layerWeights := make([]float32, numberWeights, numberWeights)
		copy(layerWeights, network.Values[valIdx:])
		layers = append(layers, layerWeights)
		numberBiases := network.Sizes[idx]
		valIdx += numberWeights + numberBiases
	}
	return layers
}

// number: The number of generations to simulate
// Returns: The lowest total error from the generations.
func (a *Algo) RunGens(number int) float64 {
	for generation := 0; generation < number; generation++ {
		a.Population.Crossover()
		a.Population.Mutate(a.MutationRate)
		a.Population.FitEval(a.ExampleInputs, a.ExpectedOutputs)
	}
	a.Generations += number
	return a.Population.Nets[0].Error
}
