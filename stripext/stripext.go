package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var r io.Reader
	if len(os.Args) == 1 {
		r = os.Stdin
	} else {
		r = strings.NewReader(strings.Join(os.Args[1:], " "))
	}

	var lines = bufio.NewScanner(r)
	lines.Split(bufio.ScanLines)
	for lines.Scan() {
		filename := lines.Text()
		fmt.Println(filename[:len(filename)-len(filepath.Ext(filename))])
	}
}
