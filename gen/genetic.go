package gen


import "fmt"
import "os"
import "sort"
import "crypto/rand"
import "math/big"

type Net struct {
    // Holds total error of network
    Error   float64
    // Holds sizes of layer
    // First index holds the number of neurons in input layer
    // Last index holds the number of neurons in output layer
    Sizes   []int
    // Holds weights and biases in the following linear order:
    // 1ST_HIDDEN_WEIGHTS 1ST_HIDDEN_BIASES 2ND_HIDDEN_WEIGHTS 2ND_HIDDEN_BIASES ... OUTPUT_WEIGHTS OUTPUT_BIAS
    Values  []float32
    // Holds intermediate calculations for forward propagation
    Inputs  []float32
    Outputs []float32
}


type Pop struct {
    Nets []Net
}


func NewPop(number int, sizes []int) Pop {
    var p Pop;
    p.Nets = make([]Net, number);
    for num := 0; num < number; num++ {
        p.Nets[num] = NewNet(sizes);
    }
    return p;
}


func NewNet(sizes []int) Net {
    var n Net;
    n.Error = 0.0;
    n.Sizes = make([]int, len(sizes));
    copy(n.Sizes, sizes);
    // Calculate slice sizes bases on given layer sizes
    leng := 0;
    max := sizes[0];
    for lay := 1; lay < len(sizes); lay++ {
        if sizes[lay] > max {
            max = sizes[lay];
        }
        // Number of weights and biases needed for layer
        leng += sizes[lay-1] * sizes[lay] + sizes[lay];
    }
    // Allocate elements for weights and biases.
    n.Values = make([]float32, leng, leng);
    // Allocate elements for inputs and outputs.
    n.Inputs = make([]float32, max, max);
    n.Outputs = make([]float32, max, max);
    return n;
}


func Mult(matrixA, vectorB, vectorC []float32, n, m int) int {
    rowCol := 0;
    for row := 0; row < n; row++ {
        vectorC[row] = 0.0
        for col := 0; col < m; col++ {
            vectorC[row] += matrixA[rowCol] * vectorB[col];
            rowCol++;
        }
    }
    return n * m;
}


func Add(vectorA, vectorB, vectorC []float32, n int) int {
    for row := 0; row < n; row++ {
        vectorC[row] = vectorA[row] + vectorB[row];
    }
    return n;
}


func ReLU(vectorA []float32, n int) {
    for row := 0; row < n; row++ {
        if vectorA[row] < 0.0 {
            vectorA[row] = 0.0;
        }
    }
}


func (n Net) ForProp(inputs []float32) []float32 {
    // First layer uses inputs as activations
    valIdx := 0
    valIdx += Mult(n.Values[valIdx:], inputs, n.Outputs, n.Sizes[1], n.Sizes[0]);
    valIdx += Add(n.Values[valIdx:], n.Outputs, n.Inputs, n.Sizes[1]);
    ReLU(n.Inputs, n.Sizes[1]);
    for lay := 2; lay < len(n.Sizes); lay++ {
        // All other layers use prior layers activations
        valIdx += Mult(n.Values[valIdx:], n.Inputs, n.Outputs, n.Sizes[lay], n.Sizes[lay-1]);
        valIdx += Add(n.Values[valIdx:], n.Outputs, n.Inputs, n.Sizes[lay]);
        ReLU(n.Inputs, n.Sizes[lay]);
    }
    copy(n.Outputs, n.Inputs)
    return n.Outputs;
}


func Abs(value float32) float32 {
    if value < 0.0 {
        value *= -1.0;
    }
    return value;
}


func (n Net) FitFunc(examInputs, expecOutputs [][]float32) float64 {
    n.Error = 0.0;
    for num := 0; num < len(examInputs); num++ {
        outputs := n.ForProp(examInputs[num]);
        for idx := 0; idx < n.Sizes[len(n.Sizes)-1]; idx++ {
            n.Error += float64(Abs(expecOutputs[num][idx] - outputs[idx]));
        }
    }
    return n.Error;
}


func Init(number int, sizes []int) Pop {
    if number % 2 != 0 {
        os.Exit(2);
    }
    return NewPop(number, sizes);
}


func (p Pop) FitEval(examInputs, expecOutputs [][]float32) {
    for num := 0; num < len(p.Nets); num++ {
        p.Nets[num].Error = p.Nets[num].FitFunc(examInputs, expecOutputs);
    }
    sort.Slice(p.Nets, func(i, j int) bool {
        return p.Nets[i].Error < p.Nets[j].Error;
    });
}


func RandIntRange(min, max int) int {
    nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(max+1-min)))
    n := nBig.Int64()
    return int(n) + min
}


func (p Pop) SelectCrossReplace() {
    for idx := len(p.Nets) / 2; idx < len(p.Nets); idx++ {
        male := p.Nets[RandIntRange(0, len(p.Nets) / 2)];
        female := p.Nets[RandIntRange(0, len(p.Nets) / 2)];
        copy(p.Nets[idx].Values, female.Values);
        copy(p.Nets[idx].Values, male.Values[:RandIntRange(1, len(female.Values) - 1)])
    }
}


func RandFloatRange(min, max float32) float32 {
	minInt := int(min * 1000000);
        maxInt := int(max * 1000000);
	return float32(float64(RandIntRange(minInt, maxInt)) / 1000000);
}


func (p Pop) Mutate(rate float32) {
    for idx := len(p.Nets) / 2; idx < len(p.Nets); idx++ {
        for idy := 0; idy < len(p.Nets[idx].Values); idy++ {
            if(RandFloatRange(0.0, 1.0) <= rate) {
               p.Nets[idx].Values[idy] += RandFloatRange(-0.25, 0.25);
            }
        }
    }
}


func (p Pop) Terminate(err float64, curErr *float64, term *bool) {
    *curErr = p.Nets[0].Error;
    *term = p.Nets[0].Error <= err;
}


func GenAlgo(gens, num int, sizes []int, examInputs, expecOutputs [][]float32, rate float32, err float64) {
    gen := 0;
    term := false;
    curErr := 0.0;
    p := Init(num, sizes);
    for gen < gens && term == false {
        p.SelectCrossReplace();
        p.Mutate(rate);
        p.FitEval(examInputs, expecOutputs);
        p.Terminate(err, &curErr, &term);
        fmt.Printf("Gen: %d, Err: %f\n", gen, curErr);
        gen++;
    }
    fmt.Println("Ins:  EXP: ACT:")
    for i := 0; i < len(examInputs); i++ {
        act := p.Nets[0].ForProp(examInputs[i]);
        fmt.Println(examInputs[i], expecOutputs[i], "", act[:1])
    }
}


func main() {
    gens := 1024;
    popSize := 1024;

    NetSizes := []int{2, 2, 1};

    exams := [][]float32{[]float32{0.0, 0.0}, []float32{0.0, 1.0}, []float32{1.0, 0.0}, []float32{1.0, 1.0}};
    xorExpec := [][]float32{[]float32{0.0}, []float32{1.0}, []float32{1.0}, []float32{0.0}};

    rate := float32(0.05);
    err := float64(0.1);

    GenAlgo(gens, popSize, NetSizes, exams, xorExpec, rate, err);
}
