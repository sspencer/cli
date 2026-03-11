# rand

Print a random selection of lines from stdin or files.

`rand` reads newline-separated input and uses reservoir sampling to select lines with uniform probability, so every line has an equal chance of being chosen regardless of input size.

## Usage

```text
rand [-n count] [-count] [file ...]
cat file.txt | rand -10
```

- Input may come from one or more filename arguments or from stdin.
- If the input has fewer lines than the requested count, all lines are printed.
- Output order matches the order lines were encountered in the input.

## Flags

```text
-n count    number of random lines to output (default 10)
-count      shorthand for -n (e.g. -10, -25)
```

## Examples

```text
$ seq 1 100 | rand -10
42
7
83
19
61
34
77
5
90
28

$ seq 1 100 | rand -n 5
54
12
88
3
67

$ rand -20 words.txt

$ rand adjectives.txt nouns.txt
```

## Install

```text
go install
```
