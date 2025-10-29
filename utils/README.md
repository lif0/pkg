# package utils

[![build](https://github.com/lif0/pkg/workflows/build/badge.svg)](https://github.com/lif0/pkg/workflows/build/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/lif0/pkg.svg)](https://pkg.go.dev/github.com/lif0/pkg/utils)
![last version](https://img.shields.io/github/v/tag/lif0/pkg?label=latest&filter=utils/*)
[![utils coverage](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Flif0%2Fpkg%2Frefs%2Fheads%2Fmain%2F.github%2Fassets%2Fbadges%2Fcoverage-utils.json)](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Flif0%2Fpkg%2Frefs%2Fheads%2Fmain%2F.github%2Fassets%2Fbadges%2Fcoverage-utils.json)
[![utils report card](https://goreportcard.com/badge/github.com/lif0/pkg/utils)](https://goreportcard.com/report/github.com/lif0/pkg/utils)

## Contents

- [Overview](#-overview)
- [Requirements](#️-requirements)
- [Installation](#-installation)
- [Package: `reflect`](#-package-reflect)
  - [Function: `EstimatePayloadOf`](#function-estimatepayloadof)
    - [Supported Types](#-supported-types)
    - [Performance](#performance-notes)
    - [Use Case](#-use-case)
    - [Examples](#-examples)
- [Package: `errx`](#-package-errx)
  - [MultiError](#multierror)
- [Package: `structx`](#-package-structx)
  - [OrderedMap](#orderedmap)
- [Roadmap](#️-roadmap)
- [License](#-license)

---

## 📋 Overview

The `utils` module provides a set of lightweight, focused utility packages designed to extend Go's standard library with commonly needed functionality. These packages follow Go's philosophy of simplicity and efficiency, offering well-tested solutions for everyday development challenges.

For full documentation, see [https://pkg.go.dev/github.com/lif0/pkg/utils](https://pkg.go.dev/github.com/lif0/pkg/utils).

---

## ⚙️ Requirements

- **go 1.23 or higher**

## 📦 Installation

To add this package to your project, use `go get`:

```bash
go get github.com/lif0/pkg/utils@latest
```

Import the reflect extension in your code:

```go
import "github.com/lif0/pkg/utils/reflect"
```

---

## 📚 Package `reflect`

### Function: `EstimatePayloadOf`

```go
func EstimatePayloadOf(v any) int
```

Returns an **approximate payload size (in bytes)** of the given value. It is designed for efficiency and is allocation-free for supported types

⚠️ **Return Value and Errors**

- Returns the estimated size in bytes for supported types.
- Returns **`reflect.ErrFailEstimatePayload`** = `-1` if the type is not supported or cannot be estimated.

#### Performance Notes

This function performs **zero allocations** and runs with **0 B/op** for supported types. [See benchmark](/utils/reflect/estimate_payload_bench_out.txt)

- **No memory allocations** — `0 B/op` for all primitive types.
- **Reflection** is used only for arrays and structs:
  - Supports any primitive types in arrays.
  - Supports only `time.Time{}` for structs.
- For arrays `[N]T`, prefer passing `*[N]T` to avoid value copying.

Check:

- [Benchmarks](/utils/reflect/estimate_payload_bench_test.go)
- [Tests](/utils/reflect/estimate_payload_test.go)

#### ✅ Supported Types

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

If the type is not supported, the function returns `reflect.ErrFailEstimatePayload(-1)`.

#### 💡 Use Case

Calculate request/response sizes for logging or monitoring, such as writing to a span in a database provider library.

#### 🧪 Examples

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
// size == reflect.ErrFailEstimatePayload(-1), because it is a custom struct
```

Estimate an array (pass by pointer for performance):

```go
var arr [1000]int32
size := EstimatePayloadOf(&arr)
// size == 1000 * 4 == 4000
```

---





## 📚 Package `errx`

Provide additional feature for error.

### MultiError

MultiError is a slice of errors implementing the error interface.


| Item        | Signature                                 | Purpose                                          | Notes                                                                       |
| ----------- | ----------------------------------------- | ------------------------------------------------ | --------------------------------------------------------------------------- |
| Type        | `type MultiError []error`                 | Slice of `error` that itself implements `error`. | Zero-value is usable.                                                       |
| Append      | `func (m *MultiError) Append(err error)`  | Adds a non-nil error to the set.                 | No-op for `nil`; safe on zero value.                                        |
| MaybeUnwrap | `func (m MultiError) MaybeUnwrap() error` | Returns the simplest meaningful error.           | `len==0 → nil`, `len==1 → m[0]`, otherwise `m` itself.                      |
| Error       | `func (m MultiError) Error() string`      | Human-readable, counted, bulleted message.       | `""` when empty; format: `"<n> error(s) occurred:\n* <err1>\n* <err2>..."`. |
| IsEmpty     | `func (m MultiError) IsEmpty() bool`      | Quick emptiness check.                           | `true` when there are no errors.                                            |



#### Example

```go
var me errx.MultiError
me.Append(validateName(""))
me.Append(validateAge(-1))
if err := me.MaybeUnwrap(); err != nil {
	// 2 error(s) occurred:
	// * name must not be empty
	// * age must be >= 0
	fmt.Println(err)
}
```

#### Batch Example

```go
var me errx.MultiError
for _, job := range jobs {
	if err := job.Run(); err != nil {
		me.Append(fmt.Errorf("job %q: %w", job.Name, err))
	}
}
return me.MaybeUnwrap()
```

## 📚 Package `structx`

Provide additional golang type.

### OrderedMap

OrderedMap is a map[Type]Type1-like collection that preserves the order in which keys were inserted. It behaves like a regular map but allows deterministic iteration over its elements.

Useful:
Imagine you are making a closer or graceful shutdown lib, and you need to register/unregister some functions/service in it, and finally handle them in the order they were added. Use it structure. You are welcome🤗

The structure provide provice

#### API

| Func                                                                   | Complexity (time / mem)      |
| ---------------------------------------------------------------------- | ---------------------------- |
| `(m *OrderedMap[K, V]) Get(key K) (V, bool)`                           | O(1) / O(1)                  |
| `(m *OrderedMap[K, V]) Put(key K, value V)`                            | O(1) / O(1)                  |
| `(m *OrderedMap[K, V]) GetValues() []V`                                | O(N) / O(N)                  |
| `(m *OrderedMap[K, V]) Iter() []V`                                     | for k,v := range m.Iter() {} |
| `structx.Delete[K comparable, V any](m *OrderedMap[K, V], key K)`      | O(1) / O(1)                  |


#### Benchmarks: OrderedMap[Type, Type1] vs map[Type]Type1

Environment:

```text
goos: darwin
goarch: arm64
cpu: Apple M2
pkg: github.com/lif0/pkg/utils/structx
```

##### TL;DR

- Inserts (`put`): `map` is faster and uses less memory.
- Lookups (`get_hit`): `OrderedMap` is faster on string keys; a bit slower on int keys.
- Deletes (`delete`): almost the same.
- Iteration (`iterate_values`): OrderedMap is much faster and ordered.

---

##### Key/Value: `int, int`

| Operation      | ns/op (`OrderedMap`) | ns/op (`map`) | B/op (`OrderedMap`) | B/op (`map`) | allocs/op (`OrderedMap`) | allocs/op (`map`) |  time (`OrderedMap` vs `map`) |
| -------------- | --------------: | ------------: | -------------: | -----------: | ------------------: | ----------------: | -------------------------: |
| put            |         220,267 |       100,546 |        705,330 |      295,557 |                  39 |                33 | **+119.1%** (2.19× slower) |
| get_hit        |          74,626 |        65,668 |              0 |            0 |                   0 |                 0 |  **+13.6%** (1.14× slower) |
| delete         |          19,322 |        19,348 |              0 |            0 |                   0 |                 0 |         **−0.1%** (≈ same) |
| iterate_values |          11,131 |        61,998 |              0 |            0 |                   0 |                 0 |   **−82.0%** (5.6× faster) |

##### Key/Value: `string, []string`

| Operation      | ns/op (`OrderedMap`) | ns/op (`map`) | B/op (`OrderedMap`) | B/op (`map`) | allocs/op (`OrderedMap`) | allocs/op (`map`) | Δ time (Ordered vs `map`) |
| -------------- | --------------: | ------------: | -------------: | -----------: | ------------------: | ----------------: | ------------------------: |
| put            |         507,196 |       360,451 |      1,084,229 |      787,101 |                  40 |                33 | **+40.7%** (1.41× slower) |
| get_hit        |         136,184 |       193,829 |              0 |            0 |                   0 |                 0 | **−29.7%** (1.43× faster) |
| delete         |          20,713 |        20,758 |              0 |            0 |                   0 |                 0 |        **−0.2%** (≈ same) |
| iterate_values |          17,822 |        63,645 |              0 |            0 |                   0 |                 0 |  **−72.0%** (3.6× faster) |

##### Key/Value: `string, ComplexStruct`

| Operation      | ns/op (`OrderedMap`) | ns/op (`map`) | B/op (`OrderedMap`) | B/op (`map`) | allocs/op (`OrderedMap`) | allocs/op (`map`) | Δ time (Ordered vs `map`) |
| -------------- | --------------: | ------------: | -------------: | -----------: | ------------------: | ----------------: | ------------------------: |
| put            |         493,887 |       329,433 |      1,166,137 |      918,167 |                  40 |                33 | **+49.9%** (1.50× slower) |
| get_hit        |         117,553 |       174,528 |              0 |            0 |                   0 |                 0 | **−32.6%** (1.49× faster) |
| delete         |          20,471 |        20,420 |              0 |            0 |                   0 |                 0 |        **+0.2%** (≈ same) |
| iterate_values |          19,729 |        62,435 |              0 |            0 |                   0 |                 0 |  **−68.4%** (3.2× faster) |

##### How to run

```bash
go test -benchmem -run=^$ -bench ^Benchmark_OrderedMap -v github.com/lif0/pkg/utils/structx
```


#### Examples

```go
import "github.com/lif0/pkg/utils/structx"


func main() {
  m := structx.NewOrderedMap[string, int]()
  
  m.Put("key", 10)

  v, ok := m.Get("key") // v = 10
  
  structx.Delete(m, "key") // or build-in func m.Delete("key"), but prefer build-in function.

  for k,v := range m.Iter() {
    fmt.Println(k,v)
  }
}
```

## 🗺️ Roadmap

The future direction of this package is community-driven! Ideas and contributions are highly welcome.

☹️ No idea

Contributions and ideas are welcome! 🤗

**Contributions:**
Feel free to open an Issue to discuss a new idea or a Pull Request to implement it! 🤗

---

## 📄 License

[MIT](./LICENSE)
