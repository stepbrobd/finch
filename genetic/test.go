package genetic

import "fmt"

func Test() {
	maxGens := 1024
	exams := [][]float32{{0.0, 0.0}, {0.0, 1.0}, {1.0, 0.0}, {1.0, 1.0}}
	//orExpec := [][]float32{{0.0}, {1.0}, {1.0}, {1.0}}
	norExpec := [][]float32{{1.0}, {0.0}, {0.0}, {0.0}}
	//xorExpec := [][]float32{{0.0}, {1.0}, {1.0}, {0.0}}

	a := NewAlgo(0.05, 1024, []int{2, 2, 1}, exams, norExpec)
	for a.GetNumGens() < maxGens {
		err := a.RunGens(16)
		fmt.Println("GEN:", a.GetNumGens(), "ERR:", err, "BIASES:")
		bis := a.GetBiases()
		for i := 0; i < len(bis); i++ {
			fmt.Println(i, "\t", bis[i])
		}
		fmt.Println("VALS:", a.Population.Nets[0].Values)
	}
	fmt.Println("Ins:  EXP: ACT:")
	for i := 0; i < len(exams); i++ {
		act := a.Population.Nets[0].ForProp(exams[i])
		fmt.Println(exams[i], norExpec[i], "", act[:1])
	}
}
