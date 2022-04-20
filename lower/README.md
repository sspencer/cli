# Lower

`lower` is a simple utility covert all arguments to lowercase.  While there are many ways to accomplish this in the shell, this is the simplest.

```bash
    $ lower "HELLO WORLD"
    hello world

    $ echo "HELLO WORLD" | lower"
    hello world
    
    # old approach
    $ echo "HELLO WORLD" | tr '[:upper:]' '[:lower:]'
    hello world
```
