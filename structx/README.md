# package structx

> Part of [**lif0/pkg**](../README.md) · [API reference](https://pkg.go.dev/github.com/lif0/pkg/structx)

Extra data structures for Go.

## Contents

- [Installation](#installation)
- [OrderedMap](#orderedmap)
  - [API](#api)
  - [Benchmarks](#benchmarks-orderedmap-vs-map)
  - [Examples](#examples)
- [ObjectPool](#objectpool)
- [License](#license)

---

## Installation

Requires **go 1.23+**.

```bash
go get github.com/lif0/pkg@latest
```

```go
import "github.com/lif0/pkg/structx"
```

---

## OrderedMap

`OrderedMap` is a `map[K]V`-like collection that preserves the order in which keys were inserted. It behaves like a regular map but allows deterministic iteration over its elements.

Handy when you need both quick key-based access and predictable iteration order, for example a graceful-shutdown registry where hooks must run in the order they were added.

### API

| Func                                                              | Complexity (time / mem)        |
| ----------------------------------------------------------------- | ------------------------------ |
| `(m *OrderedMap[K, V]) Get(key K) (V, bool)`                      | O(1) / O(1)                    |
| `(m *OrderedMap[K, V]) Put(key K, value V)`                       | O(1) / O(1)                    |
| `(m *OrderedMap[K, V]) Delete(key K)`                             | O(1) / O(1)                    |
| `(m *OrderedMap[K, V]) GetValues() []V`                           | O(N) / O(N)                    |
| `(m *OrderedMap[K, V]) Iter()`                                    | `for k, v := range m.Iter() {}`|
| `structx.Delete[K comparable, V any](m *OrderedMap[K, V], key K)` | O(1) / O(1)                    |

### Benchmarks: OrderedMap vs map

Environment:

```text
goos: darwin
goarch: arm64
cpu: Apple M2
pkg: github.com/lif0/pkg/structx
```

**TL;DR**

- Inserts (`put`): `map` is faster and uses less memory.
- Lookups (`get_hit`): `OrderedMap` is faster on string keys, a bit slower on int keys.
- Deletes (`delete`): about the same.
- Iteration (`iterate_values`): `OrderedMap` is much faster and ordered.

**Key/Value: `int, int`**

| Operation      | ns/op (`OrderedMap`) | ns/op (`map`) | time (`OrderedMap` vs `map`) |
| -------------- | --------------: | ------------: | -------------------------: |
| put            |         220,267 |       100,546 | **+119.1%** (2.19× slower) |
| get_hit        |          74,626 |        65,668 |  **+13.6%** (1.14× slower) |
| delete         |          19,322 |        19,348 |         **−0.1%** (≈ same) |
| iterate_values |          11,131 |        61,998 |   **−82.0%** (5.6× faster) |

**Key/Value: `string, []string`**

| Operation      | ns/op (`OrderedMap`) | ns/op (`map`) | time (Ordered vs `map`)   |
| -------------- | --------------: | ------------: | ------------------------: |
| put            |         507,196 |       360,451 | **+40.7%** (1.41× slower) |
| get_hit        |         136,184 |       193,829 | **−29.7%** (1.43× faster) |
| delete         |          20,713 |        20,758 |        **−0.2%** (≈ same) |
| iterate_values |          17,822 |        63,645 |  **−72.0%** (3.6× faster) |

**Key/Value: `string, ComplexStruct`**

| Operation      | ns/op (`OrderedMap`) | ns/op (`map`) | time (Ordered vs `map`)   |
| -------------- | --------------: | ------------: | ------------------------: |
| put            |         493,887 |       329,433 | **+49.9%** (1.50× slower) |
| get_hit        |         117,553 |       174,528 | **−32.6%** (1.49× faster) |
| delete         |          20,471 |        20,420 |        **+0.2%** (≈ same) |
| iterate_values |          19,729 |        62,435 |  **−68.4%** (3.2× faster) |

Run them yourself:

```bash
go test -benchmem -run=^$ -bench ^Benchmark_OrderedMap -v github.com/lif0/pkg/structx
```

### Examples

```go
import "github.com/lif0/pkg/structx"

func main() {
    m := structx.NewOrderedMap[string, int]()

    m.Put("key", 10)

    v, ok := m.Get("key") // v = 10, ok = true
    _ = ok

    structx.Delete(m, "key") // or the built-in method m.Delete("key")

    for k, v := range m.Iter() {
        fmt.Println(k, v)
    }
}
```

---

## ObjectPool

`ObjectPool[T]` is a simple free-list pool that hands out reusable `*T` values via `Get`, growing automatically as needed. It cuts allocations in hot paths.

> Note: `ObjectPool` is **not** safe for concurrent use — guard it yourself if it is shared across goroutines.

```go
import "github.com/lif0/pkg/structx"

func main() {
    pool := structx.NewObjectPool[MyType](16) // optionally preallocate 16 objects

    obj := pool.Get()
    // use obj ...
    _ = obj
}
```

---

## License

[MIT](../LICENSE)
