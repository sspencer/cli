package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findCombinations(target int, available []int, current []int, results *[][]int) {
	if len(current) == 3 {
		if target == 0 {
			combination := make([]int, len(current))
			copy(combination, current)
			*results = append(*results, combination)
		}
		return
	}

	if target < 0 {
		return
	}

	for i, num := range available {
		findCombinations(target-num, available[i+1:], append(current, num), results)
	}
}

func main() {
	exclude := flag.String("x", "", "comma-separated numbers to exclude (e.g., 1,4)")
	include := flag.String("i", "", "comma-separated numbers, one of which must be included (e.g., 5,7)")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [-x numbers] [-i numbers] <number>\n", os.Args[0])
		os.Exit(1)
	}

	target, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid number '%s'\n", flag.Arg(0))
		os.Exit(1)
	}

	if target < 6 || target > 24 {
		fmt.Fprintf(os.Stderr, "Error: number must be between 6 and 24 (inclusive)\n")
		os.Exit(1)
	}

	// Build the available numbers list
	available := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if *exclude != "" {
		excluded := make(map[int]bool)
		parts := strings.SplitSeq(*exclude, ",")
		for part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil || num < 1 || num > 9 {
				fmt.Fprintf(os.Stderr, "Error: invalid exclusion number '%s'\n", part)
				os.Exit(1)
			}
			excluded[num] = true
		}

		// Filter out excluded numbers
		filtered := []int{}
		for _, num := range available {
			if !excluded[num] {
				filtered = append(filtered, num)
			}
		}
		available = filtered
	}

	// Parse include numbers if provided
	var includeSet map[int]bool
	if *include != "" {
		includeSet = make(map[int]bool)
		parts := strings.SplitSeq(*include, ",")
		for part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil || num < 1 || num > 9 {
				fmt.Fprintf(os.Stderr, "Error: invalid include number '%s'\n", part)
				os.Exit(1)
			}
			includeSet[num] = true
		}
	}

	var results [][]int

	findCombinations(target, available, []int{}, &results)

	// Filter results if include flag is set
	if includeSet != nil {
		filtered := [][]int{}
		for _, combination := range results {
			hasInclude := false
			for _, num := range combination {
				if includeSet[num] {
					hasInclude = true
					break
				}
			}
			if hasInclude {
				filtered = append(filtered, combination)
			}
		}
		results = filtered
	}

	// Collect all unique integers used across all results
	uniqueNums := make(map[int]bool)
	for _, combination := range results {
		for _, num := range combination {
			uniqueNums[num] = true
		}
	}

	for _, combination := range results {
		for i, num := range combination {
			if i > 0 {
				fmt.Print(" + ")
			}
			fmt.Print(num)
		}
		fmt.Printf(" = %d\n", target)
	}

	// Print unique integers line
	fmt.Print("Uniques: ")
	first := true
	for i := 1; i <= 9; i++ {
		if uniqueNums[i] {
			if !first {
				fmt.Print(", ")
			}
			fmt.Print(i)
			first = false
		}
	}
	fmt.Println()

}
