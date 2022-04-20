# Erlang String Parser

For those occasions where Erlang is just printing a list of numbers when you really want to see the ASCII equivalent, use **erlnum**. Erlnum prints an Erlang list as with Text or a Hexdump (with the optional `-x` flag.)

## Usage

    erlnum file.txt
    cat file.txt | erlnum
    echo "[82,101,97,115,111,110,32,58,32]" | erlnum -x

## Example

```
$ cat testdata/input.txt
72,101,108,108,111,10
87,111,114,108,100,10

$ cat testdata/input.txt | erlnum
Hello.World.

$ cat testdata/input.txt | erlnum -x
00000000  48 65 6c 6c 6f 0a 57 6f 72 6c 64 0a    |Hello.World.|
```

## Install

Checkout the code and simply run the following command:

    go install
    