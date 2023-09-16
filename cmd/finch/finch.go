package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/charmbracelet/log"
	ui "github.com/stepbrobd/finch/ui"
)

func main() {
	log.SetLevel(log.DebugLevel)

	var (
		input  = flag.Int("input", 0, "numbers of neurons in input layer\nexample: -input=1")
		output = flag.Int("output", 0, "numbers of neurons in output layer\nexample: -output=1")
		hidden = flag.String("hidden", "", "numbers of neurons in hidden layer\nexample: -hidden=1,1,1")
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
	flag.VisitAll(func(f *flag.Flag) {
		log.Debugf("%s: %v", f.Name, f.Value)
	})

	required := []string{"input", "output", "hidden"}
	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			fmt.Fprintf(os.Stderr, "missing required -%s argument/flag\n", req)
			flag.Usage()
			os.Exit(2)
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

	log.Debugf("input: %d", *input)
	log.Debugf("output: %d", *output)
	log.Debugf("hidden: %v", specs)

	p := tea.NewProgram(ui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error occurred when trying to start UI: %v", err)
		os.Exit(1)
	}
}
