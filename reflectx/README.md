# package reflectx

> Part of [**lif0/pkg**](../README.md) · [API reference](https://pkg.go.dev/github.com/lif0/pkg/reflectx)

Reflection-based helpers that extend the standard `reflect` package.

## Contents

- [Installation](#installation)
- [EstimatePayloadOf](#estimatepayloadof)
  - [Supported Types](#supported-types)
  - [Performance Notes](#performance-notes)
  - [Examples](#examples)
- [License](#license)

---

## Installation

Requires **go 1.23+**.

```bash
go get github.com/lif0/pkg@latest
```

```go
import "github.com/lif0/pkg/reflectx"
```

---

## EstimatePayloadOf

```go
func EstimatePayloadOf(v any) int
```

Returns an **approximate payload size (in bytes)** of the given value. It is designed for efficiency and is allocation-free for supported types.

**Return value and errors:**

- Returns the estimated size in bytes for supported types.
- Returns `reflectx.ErrFailEstimatePayload` (`-1`) if the type is not supported or cannot be estimated.

### Supported Types

Scalar types (and pointers to them):

- `int`, `int8`, `int16`, `int32` (`rune`), `int64`
- `uint`, `uint8` (`byte`), `uint16`, `uint32`, `uint64`, `uintptr`
- `float32`, `float64`
- `complex64`, `complex128`
- `bool`
- `string`
- `time.Time`, `time.Duration`

Containers:

- `[]T`, `[]*T`, `*[]T`, `*[]*T` — slices
- `[N]T`, `*[N]T`, `[N]*T`, `*[N]*T` — arrays (via reflection)

For pointers and slices, `nil` is treated as zero size. For `string` and `[]string`, the actual content size is summed. If the type is not supported, the function returns `reflectx.ErrFailEstimatePayload` (`-1`).

### Performance Notes

This function performs **zero allocations** and runs with **0 B/op** for supported types.

- No memory allocations — `0 B/op` for all primitive types.
- Reflection is used only for arrays and structs:
  - any primitive type is supported inside arrays;
  - only `time.Time{}` is supported for structs.
- For arrays `[N]T`, prefer passing `*[N]T` to avoid copying the array by value.

See [benchmarks](./estimate_payload_bench_test.go) and [tests](./estimate_payload_test.go).

### Use Case

Estimate request/response sizes for logging or monitoring, e.g. writing a size attribute to a tracing span.

### Examples

Estimate an `int`:

```go
n := reflectx.EstimatePayloadOf(42)
// n == 8 on 64-bit systems
```

Estimate a `string`:

```go
s := reflectx.EstimatePayloadOf("hello")
// s == 5 (len("hello"))
```

Estimate a slice of strings:

```go
names := []string{"John", "Doe"}
size := reflectx.EstimatePayloadOf(names)
// size == len("John") + len("Doe") == 4 + 3 == 7
```

Estimate a custom struct (not supported):

```go
type User struct {
    ID   int64
    Name string
}
u := User{ID: 123, Name: "Alice"}
size := reflectx.EstimatePayloadOf(u)
// size == reflectx.ErrFailEstimatePayload (-1): custom structs are not supported
```

Estimate an array (pass a pointer for performance):

```go
var arr [1000]int32
size := reflectx.EstimatePayloadOf(&arr)
// size == 1000 * 4 == 4000
```

---

## License

[MIT](../LICENSE)
