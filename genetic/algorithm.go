package genetic

type algo struct {
	NumberGenerations  int
	CurrentGenerations int
	ExampleInputs      [][]float32
	ExpectedOutputs    [][]float32
	MutationRate       float32
	TargetTotalError   float64
	Population         Pop
}

func NewAlgo(Generations, Individuals int, MutationRate float32, TargetTotalError float64, NetworkLayers []int, ExampleInputs, ExpectedOutputs [][]float32) {
}
