package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	ErrInvalidColumn = errors.New("invalid column number")
	ErrFileInvalid   = errors.New("cannot open file")
)

type config struct {
	column int
	trim   bool
	mop    bool
	dbg    bool
	max    int
	min    int
}

func main() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintln(w, "Print specified column from STDIN or filename")
		fmt.Fprintln(w, "USAGE: coln 3 filename.txt  # prints third column (1 based index)")
		fmt.Fprintln(w, "USAGE: coln 0 filename.txt  # prints last column")
		flag.PrintDefaults()
	}

	mop := flag.Bool("map", false, "Count unique strings")
	max := flag.Int("max", -1, "Filter all strings with more than (n) chars")
	min := flag.Int("min", -1, "Filter all strings with less than (n) chars")
	dbg := flag.Bool("d", false, "Debug first line of text")
	trim := flag.Bool("q", false, "Trim quotes")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		exit(nil)
	}

	column, err := strconv.Atoi(args[0])
	if err != nil {
		exit(fmt.Errorf("%w: %q", ErrInvalidColumn, args[0]))
	}

	cfg := config{
		column: column,
		trim:   *trim,
		mop:    *mop,
		dbg:    *dbg,
		max:    *max,
		min:    *min,
	}

	var r io.Reader

	if len(args) == 1 {
		// PIPE
		r = os.Stdin
	} else {
		// FILE
		r, err = os.Open(args[1])
		if err != nil {
			exit(fmt.Errorf("%s: %w", ErrFileInvalid.Error(), err))
		}
	}

	run(r, os.Stdout, cfg)
}

// print specified column and maybe perform sum/avg/map op
func run(r io.Reader, w io.Writer, cfg config) {
	var lines = bufio.NewScanner(r)
	uniqs := make(map[string]int)

	lines.Split(bufio.ScanLines)

	for lines.Scan() {
		words := bufio.NewScanner(bytes.NewReader(lines.Bytes()))
		words.Split(scanOptionallyQuotedWords)

		if cfg.dbg {
			debugPrint(w, words, cfg)
			return
		}

		word := scanWord(words, cfg)
		if word != "" {
			if cfg.trim {
				word = strings.Trim(word, "\";'")
				word = strings.Replace(word, "\":", "", 1)
			}

			if cfg.max > 0 && len(word) > cfg.max {
				continue
			}

			if cfg.min > 0 && len(word) < cfg.min {
				continue
			}

			if cfg.mop {
				uniqs[word]++
			} else {
				fmt.Fprintln(w, word)
			}
		}
	}

	if cfg.mop {
		mapPrint(w, uniqs)
	}
}

// return text from specified column or empty string if column doesn't exist
func scanWord(words *bufio.Scanner, cfg config) string {
	c := 0
	word := ""
	for words.Scan() {
		word = words.Text()
		c++ // 1 based index so increment before check
		if cfg.column == c {
			return word
		}
	}

	// return last word
	if cfg.column == 0 {
		return word
	}

	return ""
}

// scanOptionallyQuotedWords is a split function for a Scanner that returns each
// space-separated or quoted word of text, with surrounding spaces deleted. It will
// never return an empty string. The definition of space is set by
// unicode.IsSpace.
func scanOptionallyQuotedWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isSpace(r) {
			break
		}
	}
	// Scan until space, marking end of word.
	hasQuote := false
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])

		if i == 0 && isQuote(r) {
			// skip leading quote
			hasQuote = true
			continue
		}

		if hasQuote && isQuote(r) {
			return i + width, data[start+1 : i], nil
		} else if !hasQuote && isSpace(r) {
			return i + width, data[start:i], nil

		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

func isQuote(r rune) bool {
	switch r {
	case '"':
		return true
	}

	return false
}

// isSpace reports whether the character is a Unicode white space character.
// We avoid dependency on the unicode package, but check validity of the implementation
// in the tests.
func isSpace(r rune) bool {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}

// sort and print word count map
func mapPrint(w io.Writer, m map[string]int) {
	var maxLenKey int
	keys := make([]string, 0, len(m))
	for k := range m {
		if len(k) > maxLenKey {
			maxLenKey = len(k)
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := m[k]
		fmt.Fprintf(w, "%s:%s %d\n", k, strings.Repeat(" ", maxLenKey-len(k)), v)
	}
}

// enumerate each column from first row of input
func debugPrint(w io.Writer, row *bufio.Scanner, cfg config) {
	var words []string
	for row.Scan() {
		words = append(words, row.Text())
	}

	col := cfg.column
	if cfg.column == 0 {
		col = len(words)
	}

	match := ""

	for c, word := range words {
		if c == col-1 {
			match = "**"
		} else {
			match = "  "
		}
		fmt.Fprintf(w, "%s %3d:  %s\n", match, c+1, word)
	}
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
