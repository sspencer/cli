package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// extractCount scans args for a -<digits> style argument (e.g. -10), removes
// it from the slice, and returns the parsed value along with the remaining args.
// If multiple such arguments are present the last one wins.
func extractCount(args []string) (int, []string) {
	count := 0
	out := args[:0:0]
	for _, a := range args {
		if strings.HasPrefix(a, "-") {
			if v, err := strconv.Atoi(a[1:]); err == nil && v > 0 {
				count = v
				continue
			}
		}
		out = append(out, a)
	}
	return count, out
}

func main() {
	shortCount, filteredArgs := extractCount(os.Args[1:])
	os.Args = append(os.Args[:1], filteredArgs...)

	n := flag.Int("n", 10, "number of random lines to output")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: rand [-n count] [-count] [file ...]\n\n")
		fmt.Fprintf(os.Stderr, "Randomly select lines from stdin or files and print them to stdout.\n")
		fmt.Fprintf(os.Stderr, "Uses reservoir sampling so each line has an equal chance of being selected.\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "  -<count>\n    \tnumber of random lines to output (e.g. -10)\n")
	}
	flag.Parse()

	if shortCount > 0 {
		*n = shortCount
	}

	if *n <= 0 {
		fmt.Fprintln(os.Stderr, "error: count must be a positive integer")
		os.Exit(1)
	}

	args := flag.Args()

	reservoir := make([]string, 0, *n)
	count := 0

	process := func(r io.Reader) {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			line := scanner.Text()
			count++
			if len(reservoir) < *n {
				reservoir = append(reservoir, line)
			} else {
				j := rand.Intn(count)
				if j < *n {
					reservoir[j] = line
				}
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
			os.Exit(1)
		}
	}

	if len(args) == 0 {
		process(os.Stdin)
	} else {
		for _, arg := range args {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error opening file %s: %v\n", arg, err)
				os.Exit(1)
			}
			process(f)
			f.Close()
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for _, line := range reservoir {
		fmt.Fprintln(w, line)
	}
}
