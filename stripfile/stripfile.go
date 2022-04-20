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
		ext := filepath.Ext(filename)

		if len(ext) > 1 {
			fmt.Println(filename[len(filename)-len(filepath.Ext(filename))+1:])
		} else {
			fmt.Println("<NO EXT>")
		}
	}
}
