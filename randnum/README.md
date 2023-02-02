# Randnum

Utility to generate many lines of random numbers.

## Arguments

```text
$ randnum -h
Print space separated lines of random numbers
No args for now, fixed config
  -lines int
    	number of lines output (default 20)
  -maxE int
    	maximum value of elements per line (default 10)
  -maxV int
    	maximum value of element (default 999)
  -minE int
    	minimum number of elements per line (default 1)
  -minV int
    	minimum value of element (default 1)
```

## Example

```text
 -maxV 9 -maxE 5 -lines 10
6 3
1 9 1 4 8
7 1 6 2
9 9 3 5 4
3 1 8
4
4 1
8 1 3
2 2 8
3
```