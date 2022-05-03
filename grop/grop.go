package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: %s <pattern> [<file>]\n", os.Args[0])
		os.Exit(1)
	} else if len(os.Args) == 2 {
		pattern := os.Args[1]

		// PIPE
		run(os.Stdin, os.Stdout, pattern)

	} else {
		pattern := os.Args[1]
		filename := os.Args[2]
		fd, err := os.Open(filename)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		run(fd, os.Stdout, pattern)
		if err = fd.Close(); err != nil {
			panic(err)
		}
	}
}

func run(r io.Reader, w io.Writer, pattern string) {
	var lines = bufio.NewScanner(r)
	var sb strings.Builder
	hasPattern := false
	negativePattern := pattern[0] == '!' || pattern[0] == '^' || pattern[0] == '-' || pattern[0] == '~'
	if negativePattern {
		pattern = pattern[1:]
	}

	lines.Split(bufio.ScanLines)
	num := 0
	for lines.Scan() {
		line := lines.Text()
		if strings.HasSuffix(line, "= {};") {
			if num > 1 {
				if (!negativePattern && hasPattern) || (negativePattern && !hasPattern) {
					_, _ = fmt.Fprintf(w, "%s", sb.String())
				}
			}

			sb.Reset()
			num = 0
			hasPattern = false
		}

		sb.WriteString(line)
		sb.WriteString("\n")
		num++

		// grab last part of key path
		// "json["pushid"].source = 123;" -> "source
		equals := strings.Index(line, " = ")
		dot := strings.LastIndex(line[:equals], ".")
		if equals == -1 || dot == -1 {
			continue
		}

		if strings.Contains(line[dot+1:equals], pattern) {
			hasPattern = true
		}
	}

	if num > 1 {
		if (!negativePattern && hasPattern) || (negativePattern && !hasPattern) {
			_, _ = fmt.Fprintf(w, "%s", sb.String())
		}

		if negativePattern && !hasPattern {
			_, _ = fmt.Fprintf(w, "%s", sb.String())
		}
	}

}
