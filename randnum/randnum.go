package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type config struct {
	minNum      int
	maxNum      int
	minElements int
	maxElements int
	lines       int
}

func main() {
	var cfg = config{}
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintln(w, "Print space separated lines of random numbers")
		fmt.Fprintln(w, "No args for now, fixed config")
		flag.PrintDefaults()
	}

	flag.IntVar(&cfg.minNum, "minV", 1, "minimum value of element")
	flag.IntVar(&cfg.maxNum, "maxV", 999, "maximum value of element")
	flag.IntVar(&cfg.minElements, "minE", 1, "minimum number of elements per line")
	flag.IntVar(&cfg.maxElements, "maxE", 10, "maximum value of elements per line")
	flag.IntVar(&cfg.lines, "lines", 20, "number of lines output")

	flag.Parse()

	if cfg.minNum > cfg.maxNum {
		cfg.minNum, cfg.maxNum = cfg.maxNum, cfg.minNum
	}

	if cfg.minElements > cfg.maxElements {
		cfg.minElements, cfg.maxElements = cfg.maxElements, cfg.minElements
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < cfg.lines; i++ {
		elements := r1.Intn(cfg.maxElements-cfg.minElements+1) + cfg.minElements
		for j := 0; j < elements; j++ {
			val := r1.Intn(cfg.maxNum-cfg.minNum+1) + cfg.minNum
			if j < elements-1 {
				fmt.Printf("%d ", val)
			} else {
				fmt.Printf("%d\n", val)
			}
		}
	}
}
