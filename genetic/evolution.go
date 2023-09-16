package genetic

import (
	"fmt"
)

func Evolve(gens, num int, sizes []int, examInputs, expecOutputs [][]float32, rate float32, err float64) {
	gen := 0
	term := false
	curErr := 0.0
	p := Init(num, sizes)
	for gen < gens && term == false {
		p.Crossover()
		p.Mutate(rate)
		p.FitEval(examInputs, expecOutputs)
		p.Terminate(err, &curErr, &term)
		fmt.Printf("Gen: %d, Err: %f\n", gen, curErr)
		gen++
	}
	fmt.Println("Ins:  EXP: ACT:")
	for i := 0; i < len(examInputs); i++ {
		act := p.Nets[0].ForProp(examInputs[i])
		fmt.Println(examInputs[i], expecOutputs[i], "", act[:1])
	}
}
