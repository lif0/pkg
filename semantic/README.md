# semantic

Semantic operations in Go.

📦 `go get github.com/lif0/pkg/semantic@latest`  
🧪 Requires **Go 1.19+**

[![build](https://github.com/lif0/pkg/workflows/build/badge.svg)](https://github.com/lif0/pkg/workflows/build/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/lif0/pkg.svg)](https://pkg.go.dev/github.com/lif0/pkg/semantic)
[![semantic coverage](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Flif0%2Fpkg%2Frefs%2Fheads%2Fmain%2F.github%2Fassets%2Fbadges%2Fcoverage-semantic.json)](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Flif0%2Fpkg%2Frefs%2Fheads%2Fmain%2F.github%2Fassets%2Fbadges%2Fcoverage-semantic.json)
[![semantic report card](https://goreportcard.com/badge/github.com/lif0/pkg/semantic)](https://goreportcard.com/report/github.com/lif0/pkg/semantic)

---

## Contents

- [Overview](#overview)
- [Requirements](#requirements)
- [Installation](#installation)
- [API Reference](#api-reference)
  - [EstimatePayloadOf](#function-estimatepayloadof)
    - [Supported Types](#supported-types)
    - [Performance Notes](#performance-notes)
    - [Use Case](#use-case)
    - [Examples](#examples)
- [Roadmap](#roadmap)
- [License](#license)

---

## Overview

This package provides tools for semantic-level operations on Go values.  

---

## Requirements

- **Go 1.19 or higher**

## Installation

To install the package, run:

```bash
go get github.com/lif0/pkg/semantic@latest
```

---

## API Reference

### Function: EstimatePayloadOf

```go
func EstimatePayloadOf(v any) int
```

Returns an **approximate payload size (in bytes)** of the given value.

#### Performance Notes

This function performs **zero allocations** and runs with **0 B/op** for supported types. [See benchmark](/semantic/estimate_payload_bench_out.txt)

- **No memory allocations** — `0 B/op` for all primitive types.
- **Reflection** is used only for arrays and structs:
  - Supports any primitive types in arrays.
  - Supports only `time.Time{}` for structs.
- For arrays `[N]T`, prefer passing `*[N]T` to avoid value copying.

Check:

- [Benchmarks](/semantic/estimate_payload_bench_test.go)
- [Tests](/semantic/estimate_payload_test.go)

#### Supported Types

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

#### Use Case

Calculate request/response sizes for logging or monitoring, such as writing to a span in a database provider library.

#### Examples

Estimate an int value:

```go
var v int = 42
n := EstimatePayloadOf(v)
// n == 8 on 64-bit systems
```

Estimate a string:

```go
s := EstimatePayloadOf("hello")
// s == 5 (len("hello"))
```

Estimate a slice of strings:

```go
names := []string{"John", "Doe"}
size := EstimatePayloadOf(names)
// size == len("John") + len("Doe") == 4 + 3 == 7
```

Estimate a slice with nil strings:

```go
names := []string{"John", nil}
size := EstimatePayloadOf(names)
// size == len("John") + len("") == 4 + 0 == 4
```

Estimate a struct:

```go
type User struct {
    ID   int64
    Name string
}
u := User{ID: 123, Name: "Alice"}
size := EstimatePayloadOf(u)
// size == semantic.ErrFailEstimatePayload(-1), because it is a custom struct
```

Estimate an array (pass by pointer for performance):

```go
var arr [1000]int32
size := EstimatePayloadOf(&arr)
// size == 1000 * 4 == 4000
```

---

## Roadmap

☹️ No idea

Contributions and ideas are welcome! 🤗

---

## License

[MIT](./LICENSE)