# Grop

`grop` filters [gron](https://github.com/tomnomnom/gron) output, just printing blocks containing the specified key.

Contents of `test.json`:
```json
{
  "one": {
    "foo": "bar"
  },
  "two": {
    "foo": "baz"
  },
  "three": {
    "bar": "buz"
  }
}
```

```bash
$ gron test.json | grop foo
json.one = {};
json.one.foo = "bar";
json.two = {};
json.two.foo = "baz";
```

Note: In the example above, it is safer to specify the key as ".foo" in cases where "foo" may appear in the value.

## TBD

1. only search left hand of express for pattern
2. allow negative pattern match (`!source` or `-source` to find objects with a source key)