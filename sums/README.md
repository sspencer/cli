# sums

Find all combinations of three unique digits (1-9) that sum to a target value.

`sums` takes a target number between 6 and 24 and outputs all sets of three distinct digits from 1 to 9 that sum to it. This is particularly useful for solving Sudoku or other numerical logic puzzles.

## Usage

```text
sums [-x excluded_numbers] [-i included_numbers] <target>
```

- Target must be between 6 and 24 (inclusive), as 1+2+3=6 and 7+8+9=24.
- All combinations consist of exactly three unique digits from the set {1, 2, 3, 4, 5, 6, 7, 8, 9}.
- Output lists each combination and a summary of all unique digits that appear across the results.

## Flags

```text
-x    comma-separated numbers to exclude (e.g., 1,4)
-i    comma-separated numbers, one of which must be included (e.g., 5,7)
```

## Examples

```text
$ sums 7
1 + 2 + 4 = 7
Uniques: 1, 2, 4

$ sums 10
1 + 2 + 7 = 10
1 + 3 + 6 = 10
1 + 4 + 5 = 10
2 + 3 + 5 = 10
Uniques: 1, 2, 3, 4, 5, 6, 7

$ sums -x 1 10
2 + 3 + 5 = 10
Uniques: 2, 3, 5

$ sums -i 7 10
1 + 2 + 7 = 10
Uniques: 1, 2, 7
```

## Install

```text
go install
```
