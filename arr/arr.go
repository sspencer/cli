package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	ErrFileInvalid = errors.New("cannot open file")
)

func main() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintln(w, "Print nth (1 - based) quoted value, defaults to first")
		fmt.Fprintln(w, "USAGE: arr [1] filename.txt")
		fmt.Fprintln(w, "USAGE: cat filename.txt | arr [1]")
		flag.PrintDefaults()
	}

	flag.Parse()

	args := flag.Args()
	col := 1

	var err error

	if len(args) > 0 {
		if col, err = strconv.Atoi(args[0]); err != nil {
			col = 1
		} else {
			if col < 1 {
				col = 1
			}
			args = args[1:]
		}
	}

	if len(args) == 0 {
		// PIPE
		run(os.Stdin, os.Stdout, col)
	} else {
		// FILE(s)
		for _, fn := range args {
			fd, err := os.Open(fn)
			if err == nil {
				run(fd, os.Stdout, col)
				if err = fd.Close(); err != nil {
					panic(err)
				}

			} else {
				fmt.Fprintln(os.Stderr, fmt.Errorf("%s: %w", ErrFileInvalid.Error(), err))
				os.Exit(1)
			}
		}
	}
}

// print specified column and maybe perform sum/avg/map op
func run(r io.Reader, w io.Writer, col int) {
	var lines = bufio.NewScanner(r)

	lines.Split(bufio.ScanLines)

	for lines.Scan() {
		chunks := strings.Split(lines.Text(), "\"")
		l := len(chunks)
		if l >= col*2+1 {
			fmt.Fprintln(w, chunks[col*2-1])
		}
	}
}
