package main

import (
	"flag"
	"fmt"
	"os"

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

	err := sumslib.FindCombinationsConv(os.Stdout, flag.Arg(0), *include, *exclude)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
