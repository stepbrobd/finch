package genetic

import "fmt"

func Test() {
	maxGens := 1024
	exams := [][]float32{{0.0, 0.0}, {0.0, 1.0}, {1.0, 0.0}, {1.0, 1.0}}
	xorExpec := [][]float32{{0.0}, {1.0}, {1.0}, {0.0}}

	a := NewAlgo(0.05, 1024, []int{2, 2, 1}, exams, xorExpec)
	for a.GetNumGens() < maxGens {
		err := a.RunGens(16)
		fmt.Println("GEN:", a.GetNumGens(), "ERR:", err, "WEIGHTS:")
		wts := a.GetWeights()
		for i := 0; i < len(wts); i++ {
			fmt.Println(i, "\t", wts[i])
		}
		fmt.Println("VALS:", a.Population.Nets[0].Values)
	}
	fmt.Println("Ins:  EXP: ACT:")
	for i := 0; i < len(exams); i++ {
		act := a.Population.Nets[0].ForProp(exams[i])
		fmt.Println(exams[i], xorExpec[i], "", act[:1])
	}
}
