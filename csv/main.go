package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	quote := flag.Bool("q", false, "double quote each value")
	flag.Parse()

	args := flag.Args()

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	first := true
	process := func(r io.Reader) {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			val := scanner.Text()
			if !first {
				fmt.Fprint(w, ",")
			}
			if *quote {
				fmt.Fprintf(w, "%q", val)
			} else {
				fmt.Fprint(w, val)
			}
			first = false
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

	if !first {
		fmt.Fprintln(w)
	}
}
