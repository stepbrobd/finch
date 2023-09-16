package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/stepbrobd/finch/genetic"

	tea "github.com/charmbracelet/bubbletea"
	ui "github.com/stepbrobd/finch/ui"
)

func main() {
	var (
		input      = flag.Int("input", 0, "numbers of neurons in input layer\nexample: -input=1")
		output     = flag.Int("output", 0, "numbers of neurons in output layer\nexample: -output=1")
		hidden     = flag.String("hidden", "", "numbers of neurons in hidden layer\nexample: -hidden=1,1,1")
		population = flag.Int("population", 0, "numbers of individuals the population\nexample: -population=100")
		mutation   = flag.Float64("mutation", 0.0, "mutation rate, must be between 0.0 and 1.0\nexample: -mutation=0.01")
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
	for _, c := range strings.Split(*hidden, ",") {
		v, err := strconv.Atoi(c)
		if err != nil {
			log.Fatalf("invalid hidden layer specification: %v", err)
		}
		specs = append(specs, v)
	}
	specs = append([]int{*input}, append(specs, *output)...)

	m := genetic.NewAlgo(
		float32(*mutation),
		*population,
		specs,
		[][]float32{{0.0}, {1.0}},
		[][]float32{{1.0}, {0.0}},
	)

	p := tea.NewProgram(ui.InitialModel(&m), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("error occurred when trying to start UI: %v", err)
	}
}
