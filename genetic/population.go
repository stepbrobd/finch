package genetic

import (
	"os"
	"sort"
)

type Pop struct {
	Nets []Net
}

func NewPop(number int, sizes []int) Pop {
	var p Pop
	p.Nets = make([]Net, number)
	for num := 0; num < number; num++ {
		p.Nets[num] = NewNet(sizes)
	}
	return p
}

func Init(number int, sizes []int) Pop {
	if number%2 != 0 {
		os.Exit(2)
	}
	return NewPop(number, sizes)
}

func (p Pop) FitEval(examInputs, expecOutputs [][]float32) {
	for num := 0; num < len(p.Nets); num++ {
		p.Nets[num].Error = p.Nets[num].FitFunc(examInputs, expecOutputs)
	}
	sort.Slice(p.Nets, func(i, j int) bool {
		return p.Nets[i].Error < p.Nets[j].Error
	})
}

func (p Pop) SelectCrossReplace() {
	for idx := len(p.Nets) / 2; idx < len(p.Nets); idx++ {
		male := p.Nets[RandIntRange(0, len(p.Nets)/2)]
		female := p.Nets[RandIntRange(0, len(p.Nets)/2)]
		p.Nets[idx].Crossover(male, female)
	}
}

func (p Pop) Mutate(rate float32) {
	for idx := len(p.Nets) / 2; idx < len(p.Nets); idx++ {
		p.Nets[idx].Mutate(rate)
	}
}

func (p Pop) Terminate(err float64, curErr *float64, term *bool) {
	*curErr = p.Nets[0].Error
	*term = p.Nets[0].Error <= err
}
