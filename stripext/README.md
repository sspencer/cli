# Strip Extension

`stripext` is a simple utility to strip the extension from a file name.  Unlike many shell utilities, `stripext` has no command line flags so it is not necessary to escape leading hyphens in file names.

```bash
$ stripext foo.bar
foo

$ stripext file with spaces.jpg
file with spaces

$ echo foo.bar | stripext
foo

$ stripext -LprG6oJ7RC9koBtQ303.json
LprG6oJ7RC9koBtQ303
```