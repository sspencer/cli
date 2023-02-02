package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type config struct {
	minNum      int
	maxNum      int
	minElements int
	maxElements int
	lines       int
	sort        bool
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
	flag.BoolVar(&cfg.sort, "sort", false, "sort each line")

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
		nums := randNums(r1, cfg)
		for j, v := range nums {
			if j < len(nums)-1 {
				fmt.Printf("%d ", v)
			} else {
				fmt.Printf("%d\n", v)
			}
		}
	}
}

func getKeys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func randNums(r1 *rand.Rand, cfg config) []int {
	numElements := r1.Intn(cfg.maxElements-cfg.minElements+1) + cfg.minElements
	keys := make(map[int]bool)

	// iterate twice just in case there are dups
	for i := 0; i < numElements*2; i++ {
		val := r1.Intn(cfg.maxNum-cfg.minNum+1) + cfg.minNum
		keys[val] = true
	}

	nums := getKeys(keys)

	if cfg.sort {
		sort.Ints(nums)
	}

	return nums
}
