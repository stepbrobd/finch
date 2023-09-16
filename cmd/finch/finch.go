package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/stepbrobd/finch/genetic"

	tea "github.com/charmbracelet/bubbletea"
	ui "github.com/stepbrobd/finch/ui"
)

func ReadCSV(filename string) ([][]float32, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the header row
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	var data [][]float32
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		var row []float32
		for _, valueStr := range record {
			value, err := strconv.ParseFloat(valueStr, 32)
			if err != nil {
				return nil, err
			}
			row = append(row, float32(value))
		}

		data = append(data, row)
	}

	return data, nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var (
		input      = flag.Int("input", 0, "numbers of neurons in input layer\nexample: -input=1")
		output     = flag.Int("output", 0, "numbers of neurons in output layer\nexample: -output=1")
		hidden     = flag.String("hidden", "", "numbers of neurons in hidden layer\nexample: -hidden=1,1,1")
		population = flag.Int("population", 0, "numbers of individuals the population\nexample: -population=100")
		mutation   = flag.Float64("mutation", 0.0, "mutation rate, must be between 0.0 and 1.0\nexample: -mutation=0.01")
		example    = flag.String("example", "", "training example filename, a csv\nexample: -example=./example.csv")
		expected   = flag.String("expected", "", "training label filename, a csv\nexample: -expected=./expected.csv")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout,
			"usage: %s [options]\n"+
				"  -help\n"+
				"        print this help message\n",
			os.Args[0],
		)
		flag.PrintDefaults()
	}

	flag.Parse()

	required := []string{"input", "output", "hidden", "population", "mutation"}
	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			log.Fatalf("missing required -%s argument/flag\n", req)
		}
	}

	specs := make([]int, 0)
	if strings.Contains(*hidden, ",") {
		for _, c := range strings.Split(*hidden, ",") {
			v, err := strconv.Atoi(c)
			if err != nil {
				log.Fatalf("invalid hidden layer specification: %v", err)
			}
			specs = append(specs, v)
		}
	} else {
		v, err := strconv.Atoi(*hidden)
		if err != nil {
			log.Fatalf("invalid hidden layer specification: %v", err)
		}
		specs = append(specs, v)
	}
	specs = append([]int{*input}, append(specs, *output)...)

	trainExamples, err := ReadCSV(*example)
	if err != nil {
		log.Fatalf("error occurred when trying to read training examples: %v", err)
	}

	trainLabels, err := ReadCSV(*expected)
	if err != nil {
		log.Fatalf("error occurred when trying to read training labels: %v", err)
	}

	m := genetic.NewAlgo(
		float32(*mutation),
		*population,
		specs,
		trainExamples,
		trainLabels,
	)

	p := tea.NewProgram(ui.InitialModel(&m))
	if _, err := p.Run(); err != nil {
		log.Fatalf("error occurred when trying to start UI: %v", err)
	}
}
