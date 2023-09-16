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
                copy(p.Nets[idx].Values, female.Values)
                copy(p.Nets[idx].Values, male.Values[:RandIntRange(1, len(female.Values)-1)])
        }
}


func (p Pop) Mutate(rate float32) {
        for idx := len(p.Nets) / 2; idx < len(p.Nets); idx++ {
                for idy := 0; idy < len(p.Nets[idx].Values); idy++ {
                        if RandFloatRange(0.0, 1.0) <= rate {
                                p.Nets[idx].Values[idy] += RandFloatRange(-0.25, 0.25)
                        }
                }
        }
}


func (p Pop) Terminate(err float64, curErr *float64, term *bool) {
        *curErr = p.Nets[0].Error
        *term = p.Nets[0].Error <= err
}
