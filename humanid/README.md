# Human ID

Generate readable IDs built from word lists (nouns, adjectives, verbs, adverbs, conjunctions),
or generate random character IDs.

## Usage

```text
humanid [flags] [count]
```

- `count` is optional and must be `>= 1`.
- If `count` is omitted, it defaults to `4`.

## Flags

```text
-i    generate random character ID instead of a human-readable ID
-s    separator string for human-readable IDs (default "-")
```

Notes:
- `-i` uses `count` as character length.
- `-s` affects only human-readable IDs.

## Human ID Patterns

Human IDs are composed from patterns keyed by word count:

1. noun
2. adjective + noun
3. adverb + adjective + noun
4. adjective + noun + verb + adverb
5. adverb + adjective + noun + verb + adverb
6. adjective + noun + verb + adverb + conjunction

For `count >= 6`, generation cycles by applying pattern `6` first, then restarting with the
remainder patterns.

## Examples

```text
$ humanid
steady-bird-build-quietly

$ humanid 2
brisk-otter

$ humanid -s "_" 4
bright_badger_walk_rapidly

$ humanid -i 10
Bk92fQ7mXa
```

