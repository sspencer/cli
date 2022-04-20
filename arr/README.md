# ARR Command Line

Simple command line utility to print the specified "doubled quoted" item on each line.  Nothing is output for lines that do not contain a double-quoted entry.  By default, the first item is printed.

File:
```
$ cat data.txt
json.collateral = {};
json.collateral["-K4MyCNlbYCgXx12G7CP"] = "hello";
json.collateral["-L4Pla5OTQPV8tb2V7sr"] = "world";

# Use STDIN (if no number is specified, defaults to print first quoted item)

$ cat data.txt | arr
-K4MyCNlbYCgXx12G7CP
-L4Pla5OTQPV8tb2V7sr

# OR specify the filename (print the second quoted item)

$ arr 2 data.txt
hello
world
```

