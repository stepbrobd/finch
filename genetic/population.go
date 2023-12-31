package genetic

import (
	"os"
	"sort"
	"sync"
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
	var wg sync.WaitGroup
	for num := 0; num < len(p.Nets); num++ {
		wg.Add(1)
		go p.Nets[num].FitFunc(examInputs, expecOutputs, &wg)
	}
	wg.Wait()
	sort.Slice(p.Nets, func(i, j int) bool {
		return p.Nets[i].Error < p.Nets[j].Error
	})
}

func (p Pop) Crossover() {
	var wg sync.WaitGroup
	for idx := len(p.Nets) / 2; idx < len(p.Nets); idx++ {
		male := p.Nets[RandIntRange(0, len(p.Nets)/2)]
		female := p.Nets[RandIntRange(0, len(p.Nets)/2)]
		wg.Add(1)
		go p.Nets[idx].Crossover(male, female, &wg)
	}
	wg.Wait()
}

func (p Pop) Mutate(rate float32) {
	var wg sync.WaitGroup
	for idx := len(p.Nets) / 2; idx < len(p.Nets); idx++ {
		wg.Add(1)
		go p.Nets[idx].Mutate(rate, &wg)
	}
	wg.Wait()
}

func (p Pop) Terminate(err float64, curErr *float64, term *bool) {
	*curErr = p.Nets[0].Error
	*term = p.Nets[0].Error <= err
}
