package sumslib

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ParseNumList parses a comma-separated string of integers, each in [1,9].
// Returns the bad token as the error message if any value is invalid.
func ParseNumList(s string) ([]int, error) {
	var nums []int
	for part := range strings.SplitSeq(s, ",") {
		part = strings.TrimSpace(part)
		num, err := strconv.Atoi(part)
		if err != nil || num < 1 || num > 9 {
			return nil, fmt.Errorf("%s", part)
		}
		nums = append(nums, num)
	}
	return nums, nil
}

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

// FindCombinationsConv is a string-input wrapper around FindCombinations that
// validates the target and parses include/exclude comma-separated number lists.
func FindCombinationsConv(w io.Writer, targetStr, include, exclude string) error {
	target, err := strconv.Atoi(targetStr)
	if err != nil {
		return fmt.Errorf("invalid number %q", targetStr)
	}

	if target < 6 || target > 24 {
		return errors.New("number must be between 6 and 24 (inclusive)")
	}

	includeNums, err := parseOptionalList(include, "include")
	if err != nil {
		return err
	}

	excludeNums, err := parseOptionalList(exclude, "exclude")
	if err != nil {
		return err
	}

	FindCombinations(w, target, includeNums, excludeNums)
	return nil
}

func parseOptionalList(s, name string) ([]int, error) {
	if s == "" {
		return nil, nil
	}
	nums, err := ParseNumList(s)
	if err != nil {
		return nil, fmt.Errorf("invalid %s number %q", name, err)
	}
	return nums, nil
}

// FindCombinations finds all 3-number combinations from 1-9 that sum to target,
// optionally excluding numbers in exclude and requiring at least one number from include.
// Results are written to w. Inputs are assumed to be already validated.
func FindCombinations(w io.Writer, target int, include, exclude []int) {
	available := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if len(exclude) > 0 {
		excluded := make(map[int]bool, len(exclude))
		for _, num := range exclude {
			excluded[num] = true
		}
		var filtered []int
		for _, num := range available {
			if !excluded[num] {
				filtered = append(filtered, num)
			}
		}
		available = filtered
	}

	var results [][]int
	findCombinations(target, available, []int{}, &results)

	if len(include) > 0 {
		includeSet := make(map[int]bool, len(include))
		for _, num := range include {
			includeSet[num] = true
		}
		var filtered [][]int
		for _, combination := range results {
			for _, num := range combination {
				if includeSet[num] {
					filtered = append(filtered, combination)
					break
				}
			}
		}
		results = filtered
	}

	uniqueNums := make(map[int]bool)
	for _, combination := range results {
		for _, num := range combination {
			uniqueNums[num] = true
		}
	}

	var sb strings.Builder
	for _, combination := range results {
		for i, num := range combination {
			if i > 0 {
				sb.WriteString(" + ")
			}
			sb.WriteString(strconv.Itoa(num))
		}
		fmt.Fprintf(&sb, " = %d\n", target)
	}

	sb.WriteString("Uniques: ")
	first := true
	for i := 1; i <= 9; i++ {
		if uniqueNums[i] {
			if !first {
				sb.WriteString(", ")
			}
			sb.WriteString(strconv.Itoa(i))
			first = false
		}
	}
	sb.WriteString("\n")

	fmt.Fprint(w, sb.String())
}
