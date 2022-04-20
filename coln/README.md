# COLN: Print Column

Print specified column from text file.  This replaces the [fish shell](https://fishshell.com/) 
function coln.fish which is slow for larger files.

```fish
function coln
    while read -l input
        echo $input | awk '{print $'$argv[1]'}'
    end
end
```
## Usage

```
Print specified column from STDIN or filename
USAGE: coln 3 filename.txt  # prints third column (1 based index)
USAGE: coln 0 filename.txt  # prints last column
  -avg 	Calculate the average of all numbers in column
  -d	Debug first line of text
  -map 	Count unique strings
  -q	Trim quotes
  -sum 	Calculate the sum of all numbers in column
```

## Examples

```
$ cat input.txt
one two three
1 2 3

$ echo "Use with Pipes"
$ cat input.txt | coln 2
two
2 

$ echo "Use with Files"
$ coln 3 input.txt
three
3

$ echo "One flag: -q Strip quotes"
$ cat input2.txt
a = "one";
b = "two";

$ coln -q 3 input2.txt
one
two

$ cat input3.txt
one 1
two 2
three 3

$ coln -sum 2 input3.txt
6

$ cat input4.txt
one   two three
three two one
zero  one zero
two   two two 

$ coln -map 2 input4.txt
one: 1
two: 3
```

## Install

Go makes it easy:

    go install
