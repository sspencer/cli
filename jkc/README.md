# JSON Key Counter

JSON Key counter is a command line utility that recursively walks specified 
directories to count unique key paths found in *.json files.

## Features 

1. Specify one or more filenames or directories for `jkc` to recursively descend into
2. Directories are walked recursively and all *.json files are scanned
3. Array elements are rolled up into `[*]`
4. Keys may be excluded to reduce noise with `-v key` flag
5. IDs (UUID and PushIds) may be rolled up with the `-i` flag which reduces them to `<id>`
6. Partial keys can be ignored with the `-p` flag
7. The type of each key is reported (string, number, boolean, unknown).  Objects and arrays are flattened or rolled up. 
8. Reports with directory crawled may be printed with -csv or -tsv flags.

## Usage 

```
USAGE:
  jkc <filename | dir> [<filename | dir> ...]
  -csv      output in CSV format
  -i        skip keys that looks like push ids or uuids
  -p value  skip "key" with this substring (invert match)
  -tsv      output in TSV format
  -v value  skip this "key" (invert match)```
```

```
$ jkc -i -v type testdata
pets[*].<id>.animal.name  string      2
pets[*].<id>.animal.type  string      2
pets[*].<id>.type         string      2
pets[*].bark              boolean(*)  6
pets[*].meow              boolean     1
pets[*].name              string      7
pets[*].speak             boolean     3
pets[*].type              string(*)   7

```

## Notes

* You may skip more than one key: `jkc -v input -v output dir1 dir2 dir3`
* If a type ends with (*), it means that the field may have more than one type.


