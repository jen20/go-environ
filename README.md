# Go - `os.Environ` as a map

This library simply converts the output of `os.Environ()` into a key-value delimited map, taking care to match key/value pairs correctly with a regular expression. I have no idea why this isn't part of the standard library.

Usage:

```
env := environ.Map()

fmt.Println(env["COLUMNS"])
```
