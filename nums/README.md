# nums

Extract integers from text and print them as a single comma‑separated list.

`nums` scans input (from STDIN or a file) for positive integers and outputs them in the order found, joined by commas.

## Usage

```text
nums [filename]
cat file.txt | nums
```

- Input may come from a filename argument or from STDIN.
- Output is a single line of comma‑separated positive integers, with no trailing newline.
- Only strictly positive integers are included (values > 0). Zeros and negative numbers are ignored.

## Examples

```text
$ echo "abc 12 def 7 xyz" | nums
12,7

$ cat sample.txt
id=42 user=alice
count=3 retries=0

$ nums sample.txt
42,3
```

Notes:
- Non‑numeric content is ignored.
- `0` is excluded by design; negative values (like `-5`) are not included.

## Install

Go makes it easy:

    go install
