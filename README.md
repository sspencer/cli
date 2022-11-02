# Command Line Utilities

This repository is a collection of command line utilities that make my shell experience better.  The utilities are written in [Go](https://go.dev/) and are available on [Github](https://github.com/sspencer/cli).  Some utilities are really simple, like `upper` and `lower`, but I find them easier to remember than the syntax to `awk` or `tr` to change the case of a string in the shell.  Other utilities like `stripext` and `stripfile` remove the annoying need to quote or escape filenames that have a hyppen as a prefix.

```bash
$ basename ./-N0F1689aweRQf395IVo.json .json
-N0F1689aweRQf395IVo

basename -- -N0F1689aweRQf395IVo.json .json
-N0F1689aweRQf395IVo

### OR ###
 
$ stripext -N0F1689aweRQf395IVo.json
-N0F1689aweRQf395IVo
```

spilling misteak

NOTE: [pushid](pushid/) comes from the [
themartorana pushID.go gist](https://gist.github.com/themartorana/8c8b704432c8be1fed9a).  All other code is original.

* [arr](arr/): print specified quoted column
* [coln](coln/): print specified column
* [erlnum](erlnum/): covert Erlang list of numbers to ascii
* [fire](fire/): retrieve firebase data
* [grop](grop/): filter objects from [gron](https://github.com/tomnomnom/gron) output
* [jkc](jkc/): count json keys in multiple files
* [lower](lower/): convert to lower case
* [pushid](pushid/): generate unique id
* [stripext](stripext/): strip file extension
* [stripfile](stripfile/): strip file name
* [upper](upper/): convert to upper case

To compile all executables, just `make` in the top level directory.  All source code is written in `Go` and the binaries will be installed in `$GOPATH/bin`.

```bash
$ make
cd arr && go install
cd coln && go install
cd erlnum && go install
cd fire && go install
cd group && go install
cd jkc && go install
cd lower && go install
cd pushid && go install
cd stripext && go install
cd stripfile && go install
cd upper && go install
```
