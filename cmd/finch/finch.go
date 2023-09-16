package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	log "github.com/charmbracelet/log"
	ui "github.com/stepbrobd/finch/ui"
)

func main() {
	var (
		input  = flag.Int("input", 0, "numbers of neurons in input layer")
		hidden = flag.Int("hidden", 0, "numbers of neurons in hidden layer")
		output = flag.Int("output", 0, "numbers of neurons in output layer")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout,
			"usage: %s [options]\n"+
				"  -help\n"+
				"	print this help message\n",
			os.Args[0],
		)
		flag.PrintDefaults()
	}

	flag.Parse()

	required := []string{"input", "hidden", "output"}
	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			fmt.Fprintf(os.Stderr, "missing required -%s argument/flag\n", req)
			flag.Usage()
			os.Exit(2)
		}
	}

	log.Debugf("input layner has %d neurons", *input)
	log.Debugf("output layer has %d neurons", *output)
	log.Debugf("%d hidden layer specified", *hidden)

	specs := make(map[int]int)
	for i := 0; i < *hidden; i++ {
		var n int
		fmt.Printf("neurons in hidden layer %d: ", i)
		fmt.Scanf("%d", &n)
		specs[i] = n
		log.Debugf("hidden layer %d has %d neurons", i, n)
	}

	p := tea.NewProgram(ui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error occurred when trying to start UI: %v", err)
		os.Exit(1)
	}
}
