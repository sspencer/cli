package main

import (
	"flag"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sspencer/cli/jkc/cnt"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// print dotted path for each json key
func main() {
	rand.Seed(time.Now().UnixNano())
	var skipKeys, skipPartials arrayFlags

	flag.Var(&skipKeys, "v", "skip this \"key\" (invert match)")                    // can specify more than once
	flag.Var(&skipPartials, "p", "skip \"key\" with this substring (invert match)") // can specify more than once

	skipIds := flag.Bool("i", false, "skip keys that looks like push ids or uuids")
	csvOut := flag.Bool("csv", false, "output in CSV format")
	tsvOut := flag.Bool("tsv", false, "output in TSV format")
	minCnt := flag.Int("min", 0, "skip keys with too few items")

	flag.Usage = func() {
		w := flag.CommandLine.Output() // may be os.Stderr - but not necessarily
		fmt.Fprintf(w, "Count all unique JSON key paths in *.json files found in the specified directories\n\n")
		fmt.Fprintf(w, "USAGE:\n  %s <filename | dir> [<filename | <dir>...]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	counter := cnt.NewCounter(skipKeys, skipPartials, *skipIds, *minCnt)
	var allFiles []string

	// read multiple directories from input args
	for _, filename := range args {
		if filename[0] != os.PathSeparator {
			filename = filepath.Join(cwd, filename)
		}

		if stat, err := os.Stat(filename); err == nil {
			if stat.IsDir() {
				jsonFiles, err := walkDir(filename)
				if err != nil {
					fmt.Fprintf(os.Stderr, "ERROR walking directory %q: %s\n", filename, err)
					os.Exit(1)
				} else {
					allFiles = append(allFiles, jsonFiles...)
				}
			} else {
				allFiles = append(allFiles, filename)
			}
		} else {
			fmt.Fprintf(os.Stderr, "ERROR reading file %q\n", filename)
			os.Exit(1)
		}
	}

	// shuffle files so concurrent processes are more likely getting an equal mix of file sizes
	rand.Shuffle(len(allFiles), func(i, j int) {
		allFiles[i], allFiles[j] = allFiles[j], allFiles[i]
	})

	// concurrent, allFiles broken into chunks (one goroutine per cpu),
	counter.CountFiles(allFiles)

	out := os.Stdout
	if *csvOut {
		counter.CSVPrint(out, ',')
	} else if *tsvOut {
		counter.CSVPrint(out, '\t')
	} else {
		counter.PrettyPrint(out)
	}
}

// recursively walk directory looking for all *.json files
func walkDir(dir string) ([]string, error) {
	fileSystem := os.DirFS(dir)

	var fileNames []string
	err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if d.Type().IsRegular() && strings.HasSuffix(d.Name(), ".json") {
			fileNames = append(fileNames, filepath.Join(dir, path))
		}

		return nil
	})

	return fileNames, err
}
