# Command Line Utilities

This repository is a collection of command line utilities that I've found useful.

* [arr](arr/): print specified quoted column
* [coln](coln/): print specified column
* [erlnum](erlnum/): covert Erlang list of numbers to ascii
* [fire](fire/): retrieve firebase data
* [jkc](jkc/): count json keys in multiple files
* [lower](lower/): convert to lower case
* [stripext](stripext/): strip file extension
* [stripfile](stripfile/): strip file name
* [upper](upper/): convert to upper case

To compile all executables, just `make` in the top level directory.  All source code is written in `Go` and the binaries will be installed in `$GOPATH/bin`.

```bash
$ make
cd arr/. && go install
cd coln/. && go install
cd erlnum/. && go install
cd fire/. && go install
cd jkc/. && go install
cd lower/. && go install
cd stripext/. && go install
cd stripfile/. && go install
cd upper/. && go install
```
