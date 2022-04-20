package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {

	hexPtr := flag.Bool("x", false, "Output hexdump")
    flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		info, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}

		if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
			fmt.Println("Print erlang lists/strings, e.g. [72,101,108,108,111]")
			fmt.Println("Specify the [-x] flag to see output as a hexdump.")
			fmt.Println("The command works with pipes or files.")
			fmt.Println("Usage: cat file | erldump [-h]")
			fmt.Println("Usage: erldump file")
			return
		}

		parse(bufio.NewReader(os.Stdin), *hexPtr)
	} else {
		for _, fn := range args {
			if reader, err := os.Open(fn); err == nil {
				parse(reader, *hexPtr)
			}
		}
	}
}

func parse(reader io.Reader, hexdump bool) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	src := make([]byte, 0)
	rex := regexp.MustCompile(`[^0-9]+`)

	for scanner.Scan() {
		for _, n := range rex.Split(scanner.Text(), -1) {
			if i, err := strconv.Atoi(n); err == nil {
				if hexdump {
					src = append(src, uint8(i))
				} else {
					if i >= 32 && i < 127 {
						src = append(src, uint8(i))
					} else {
						src = append(src, 46) // 46 is "."
					}
				}
			}
		}
	}

	if hexdump {
		fmt.Println(hex.Dump(src))
	} else {
		fmt.Println(string(src))
	}
}
