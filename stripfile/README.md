# Strip Filename

`stripfile` outputs a file's extension, stripping the filenmae.  Unlike many shell utilities, `stripfile` has no command line flags so it is not necessary to escape leading hyphens in file names.

```bash
$ stripext -LprG6oJ7RC9koBtQ303.json
json

$ stripfile foo.bar
bar

$ stripfile file with spaces.jpg
jpg

$ echo foo | stripfile
<NO EXT>
```