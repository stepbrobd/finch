package genetic


func test() {
	gens := 1024
        popSize := 1024

        NetSizes := []int{2, 2, 1}

        exams := [][]float32{{0.0, 0.0}, {0.0, 1.0}, {1.0, 0.0}, {1.0, 1.0}}
        xorExpec := [][]float32{{0.0}, {1.0}, {1.0}, {0.0}}

        rate := float32(0.05)
        err := float64(0.1)

        genetic.Evolve(gens, popSize, NetSizes, exams, xorExpec, rate, err)
}
