package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	ErrInvalidSize = errors.New("size must be a positive integer")
	ErrInvalidFile = errors.New("cannot open file")
)

type config struct {
	size int
	tmpl string
	bits bool
	csv  bool
	nums bool
}

func main() {
	var cfg config

	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintln(w, "Convert numbers to bit string")
		fmt.Fprintln(w, "USAGE: bits 32 filename.txt")
		flag.PrintDefaults()
	}

	flag.IntVar(&cfg.size, "s", 32, "size of bit string")
	flag.StringVar(&cfg.tmpl, "t", "%s", "output template")
	flag.BoolVar(&cfg.bits, "bits", true, "print bit string in template")
	flag.BoolVar(&cfg.csv, "csv", false, "print comma separated in template")
	flag.BoolVar(&cfg.nums, "nums", false, "print space separated numbers in template")

	flag.Parse()

	args := flag.Args()

	var r io.Reader
	var err error

	if len(args) == 0 {
		// PIPE
		r = os.Stdin
	} else {
		// FILE
		r, err = os.Open(args[0])
		if err != nil {
			exit(fmt.Errorf("%s: %w", ErrInvalidFile.Error(), err))
		}
	}

	run(r, os.Stdout, cfg)
}

// print specified column and maybe perform sum/avg/map op
func run(r io.Reader, w io.Writer, cfg config) {
	var lines = bufio.NewScanner(r)

	lines.Split(bufio.ScanLines)
	lineNum := 1
	sep := []byte{' '}
	for lines.Scan() {
		lb := lines.Bytes()
		nums := make([]int, bytes.Count(lb, sep)+1)
		i := 0
		words := bufio.NewScanner(bytes.NewReader(lb))
		words.Split(bufio.ScanWords)

		for words.Scan() {
			s := words.Text()
			n, err := strconv.Atoi(s)
			if err != nil {
				exit(fmt.Errorf("line %d: value %q is not a number", lineNum, s))
			}

			if n > 0 && n <= cfg.size {
				nums[i] = n
			} else {
				exit(fmt.Errorf("line %d: value %d is out of bounds (max %d)", lineNum, n, cfg.size))
			}

			i++
		}

		var output string

		// get line of output
		if cfg.nums {
			output = printSep(nums, " ", cfg)
		} else if cfg.csv {
			output = printSep(nums, ", ", cfg)
		} else {
			output = printBits(nums, cfg)
		}

		fmt.Fprintf(w, cfg.tmpl+"\n", output)

		lineNum++
	}

}

func printSep(nums []int, sep string, cfg config) string {
	var b strings.Builder
	b.Grow(len(nums) * 4) // guess at string size
	b.WriteString(fmt.Sprintf("%d", nums[0]))
	for _, s := range nums[1:] {
		b.WriteString(sep)
		b.WriteString(fmt.Sprintf("%d", s))
	}
	return b.String()
}

func printBits(nums []int, cfg config) string {
	var one byte = '1'
	line := bytes.Repeat([]byte{'0'}, cfg.size)
	for _, v := range nums {
		line[v-1] = one
	}

	return string(line)
}

// print err or usage to stderr and exit
func exit(err error) {
	if err == nil {
		flag.Usage()
	} else {
		fmt.Fprintln(os.Stderr, err)
	}

	os.Exit(1)
}
