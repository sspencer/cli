package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/sspencer/cli/internal/sumslib"
)

func main() {
	exclude := flag.String("x", "", "comma-separated numbers to exclude (e.g., 1,4)")
	include := flag.String("i", "", "comma-separated numbers, one of which must be included (e.g., 5,7)")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [-x numbers] [-i numbers] <number>\n", os.Args[0])
		os.Exit(1)
	}

	target, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid number '%s'\n", flag.Arg(0))
		os.Exit(1)
	}

	if target < 6 || target > 24 {
		fmt.Fprintf(os.Stderr, "Error: number must be between 6 and 24 (inclusive)\n")
		os.Exit(1)
	}

	var excludeNums []int
	if *exclude != "" {
		if excludeNums, err = sumslib.ParseNumList(*exclude); err != nil {
			fmt.Fprintf(os.Stderr, "Error: invalid exclusion number '%s'\n", err)
			os.Exit(1)
		}
	}

	var includeNums []int
	if *include != "" {
		if includeNums, err = sumslib.ParseNumList(*include); err != nil {
			fmt.Fprintf(os.Stderr, "Error: invalid include number '%s'\n", err)
			os.Exit(1)
		}
	}

	sumslib.FindCombinations(os.Stdout, target, includeNums, excludeNums)
}
