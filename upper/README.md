# Lower

`upper` is a simple utility covert all arguments to uppercase.  While there are many ways to accomplish this in the shell, this is the simplest.

```bash
    $ upper "hello world"
    HELLO WORLD

    $ echo "hello world" | upper
    HELLO WORLD
    
    # old approach
    $ echo "hello world" | tr '[:lower:]' '[:upper:]'
    HELLO WORLD
```
