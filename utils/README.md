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
- [Roadmap](#️-roadmap)
- [License](#-license)

---

## 📋 Overview

The `utils` module provides a set of lightweight, focused utility packages designed to extend Go's standard library with commonly needed functionality. These packages follow Go's philosophy of simplicity and efficiency, offering well-tested solutions for everyday development challenges.

For full documentation, see [https://pkg.go.dev/github.com/lif0/pkg/utils](https://pkg.go.dev/github.com/lif0/pkg/utils).

---

## ⚙️ Requirements

- **go 1.22 or higher**

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

## 🗺️ Roadmap

The future direction of this package is community-driven! Ideas and contributions are highly welcome.

☹️ No idea

Contributions and ideas are welcome! 🤗

**Contributions:**
Feel free to open an Issue to discuss a new idea or a Pull Request to implement it! 🤗

---

## 📄 License

[MIT](./LICENSE)
