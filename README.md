# Go Slice Deduplicate using Generics

Remove duplicates for a slice with the use of generics

## Docs

[![Go Reference](https://pkg.go.dev/badge/github.com/lil5/go-slice-dedup.svg)](https://pkg.go.dev/github.com/lil5/go-slice-dedup)

```
DeDuplicate[T string | int | int32 | int64 | float32 | float64 | uint](list []T) []T
```

## Example

```go
list := []int{1, 2, 1, 3}
listUnique := DeDuplicate(list)
assert.Equal(t, []int{1, 2, 3}, listUnique, "should find 1,2,3 from deduplicated list of 1,2,1,3")
```

## License

MIT