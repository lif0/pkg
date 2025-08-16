# semantic

Helpers for semantic operations in Go.

📦 `go get github.com/lif0/pkg/semantic`  
🧪 Requires **Go 1.19+**

[![Go Reference](https://pkg.go.dev/badge/github.com/lif0/pkg.svg)](https://pkg.go.dev/github.com/lif0/pkg/semantic)
[![semantic](https://coveralls.io/repos/github/lif0/pkg/badge.svg?flag=semantic&branch=main)](https://coveralls.io/github/lif0/pkg?flag=semantic&branch=main)


---

## Contents

- [Overview](#overview)
- [Installation](#installation)
- [Function: EstimatePayloadOf](#function-estimatepayloadof)
  - [Supported Types](#supported-types)
  - [Performance Notes](#performance-notes)
  - [Use case](#use-case)
  - [Examples](#examples)
- [License](#license)

---

## Overview

This package provides tools for semantic-level operations on Go values.  
Currently, it includes a single function:

- `EstimatePayloadOf` — returns an approximate size in bytes of the given value, without allocations.

---

## Installation

```bash
go get github.com/lif0/pkg/semantic
```

---

## Function: EstimatePayloadOf

```go
func EstimatePayloadOf(v any) int
```

Returns an **approximate payload size (in bytes)** of the given value.

## Performance Notes

This function performs **zero allocations** and runs with **0 B/op**. [See benchmark](/semantic/estimate_payload_bench_out.txt)

- **No memory allocations** — `0 B/op`, for all primitive types.
- **Reflection** is used only for arrays and structs
- - support any primitive types like array
- - support only time.Time{} like struct
- For arrays `[N]T`, prefer passing `*[N]T` to avoid value copy.

Check:

- [Benchmarks](/semantic/estimate_payload_bench_test.go)
- [Tests](/semantic/estimate_payload_test.go)

---

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

For pointers and slices, `nil` is treated as zero-size.  
For `string` and `[]string`, the actual content size is summed.

If the type is not supported, the function returns `semantic.ErrFailEstimatePayload(-1)`.

---

### Use case

Real-world case: Calculating request/response size and writing it into the span(in database provider library)

### Examples

Estimate an int value:

```go
var v int = 42
n, err := EstimatePayloadOf(v)
// n == 8 on 64-bit systems
```

Estimate a string:

```go
s, err := EstimatePayloadOf("hello")
// s == 5 len("hello")
```

Estimate a slice of strings:

```go
names := []string{"John", "Doe"}
size, err := EstimatePayloadOf(names)
// size == len("John") + len("Doe") == 4 + 3 == 7
```

```go
names := []string{"John", nil}
size, err := EstimatePayloadOf(names)
// size == len("John") + len("") == 4 + 1 == 5
```

Estimate a struct:

```go
type User struct {
	ID   int64
	Name string
}

u := User{ID: 123, Name: "Alice"}
size := EstimatePayloadOf(u)
// size == semantic.ErrFailEstimatePayload(-1), because it is custom structure
```

Estimate an array (pass by pointer for performance):

```go
var arr [1000]int32
size, err := EstimatePayloadOf(&arr)
// size = 1000*4 = 4000. It happen because array pre-allocate all line(in some way slice do that also)
```

---

## Roadmap

☹️ No idea

Contributions and ideas are welcome! 🤗

---

## License

[MIT](./LICENSE)
