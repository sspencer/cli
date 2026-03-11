# csv

Join newline-separated input into a single comma-separated line.

`csv` reads lines from stdin or files and outputs them joined by commas on a single line, optionally double-quoting each value.

## Usage

```text
csv [-q] [file ...]
cat file.txt | csv
```

- Input may come from one or more filename arguments or from stdin.
- Multiple files are concatenated before joining.
- Output is a single comma-separated line followed by a newline.

## Flags

```text
-q    double-quote each value
```

## Examples

```text
$ printf "apple\nbanana\ncherry" | csv
apple,banana,cherry

$ printf "apple\nbanana\ncherry" | csv -q
"apple","banana","cherry"

$ cat words.txt
foo
bar
baz

$ csv words.txt
foo,bar,baz

$ csv -q words.txt
"foo","bar","baz"
```

## Install

```text
go install
```
