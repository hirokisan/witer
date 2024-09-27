[![Go Report Card](https://goreportcard.com/badge/github.com/hirokisan/witer)](https://goreportcard.com/report/github.com/hirokisan/witer)
[![test](https://github.com/hirokisan/witer/actions/workflows/test.yml/badge.svg)](https://github.com/hirokisan/witer/actions/workflows/test.yml)

# witer

witer is a wrapper around Go's iter package, designed to simplify and extend its functionality for more convenient iteration handling.

## Usage

### Filter

```golang
import "github.com/hirokisan/witer"

iter := witer.New(slices.Values([]int64{1, 2, 3}))
items := iter.Filter(func(v int64) bool {
    return v%2 == 0
}).Collect()
```

```golang
import "github.com/hirokisan/witer"

iter := witer.New2(maps.All(map[int64]string{
    0: "a",
    1: "b",
    2: "c",
}))
items := maps.Collect(iter.Filter(func(k int64, v string) bool {
    return k%2 == 0 && v == "a"
}).Iter())
```

### Apply

```golang
import "github.com/hirokisan/witer"

iter := witer.New(slices.Values([]int64{1, 2, 3}))
items := iter.Apply(func(v int64) int64 {
    return v * 10
}).Collect()
```

```golang
import "github.com/hirokisan/witer"

iter := witer.New2(maps.All(map[int64]string{
    0: "a",
    1: "b",
    2: "c",
}))
items := maps.Collect(iter.Apply(func(k int64, v string) (int64, string) {
    return k * 10, v
}).Iter())
```
