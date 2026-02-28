package main

import (
	"flag"
	"fmt"
	"math/rand"
	"slices"
	"sort"
	"strconv"
	"strings"
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

type numbers []int

func (nums numbers) String() string {
	var sb strings.Builder
	for i, v := range nums {
		sb.WriteString(strconv.Itoa(v))
		if i < len(nums)-1 {
			sb.WriteString(" ")
		}
	}

	return sb.String()
}

func (num numbers) hasDuplicate(val int) bool {
	return slices.Contains(num, val)
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
		fmt.Println(randNums(r1, cfg))
	}
}

func randNums(r1 *rand.Rand, cfg config) numbers {
	numElements := r1.Intn(cfg.maxElements-cfg.minElements+1) + cfg.minElements
	var nums numbers

	for range numElements {
		// try a few rand vals just in case there's a dup
		for range 5 {
			val := r1.Intn(cfg.maxNum-cfg.minNum+1) + cfg.minNum
			if nums.hasDuplicate(val) {
				continue
			}
			nums = append(nums, val)
			break
		}
	}

	if cfg.sort {
		sort.Ints(nums)
	}

	return nums
}
