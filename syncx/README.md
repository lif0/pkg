# package syncx

> Part of [**lif0/pkg**](../README.md) · [API reference](https://pkg.go.dev/github.com/lif0/pkg/syncx)

Synchronization primitives for Go that extend the standard `sync` package: a reentrant mutex, a counting semaphore, a `WithLock` helper, and a mutex-guarded value.

## Contents

- [Installation](#installation)
- [ReentrantMutex](#reentrantmutex)
- [Semaphore](#semaphore)
- [WithLock](#withlock)
- [SyncValue](#syncvalue)
- [License](#license)

---

## Installation

Requires **go 1.23+**.

```bash
go get github.com/lif0/pkg@latest
```

```go
import "github.com/lif0/pkg/syncx"
```

---

## ReentrantMutex

The `ReentrantMutex` type is a reentrant mutual-exclusion lock that lets the same goroutine acquire the lock multiple times without deadlocking. It supports recursive locking and ownership tracking via `Lock` and `Unlock`.

Important details:

- Panics on unlocking an unlocked mutex.
- Panics on unlocking from a goroutine other than the owner.
- Panics if the recursion count goes negative.

### Performance

```text
goos: darwin
goarch: arm64
pkg: github.com/lif0/pkg/syncx
cpu: Apple M2
BenchmarkMutexes/sync.Mutex-8                              155_743_212     7.684 ns/op     0 B/op     0 allocs/op
BenchmarkMutexes/sync.RWMutex-8                            71_501_988     16.49 ns/op      0 B/op     0 allocs/op
BenchmarkMutexes/ReentrantMutex-8                          65_857_020     17.82 ns/op      0 B/op     0 allocs/op

BenchmarkMutexesParallel/sync.MutexParallel-8              13_736_034     73.02 ns/op      0 B/op     0 allocs/op
BenchmarkMutexesParallel/sync.RWMutexParallel-8            14_190_777     84.19 ns/op      0 B/op     0 allocs/op
BenchmarkMutexesParallel/ReentrantMutexParallel-8          35_037_007     34.84 ns/op      0 B/op     0 allocs/op
```

### Example: Basic Usage

```go
package main

import (
    "fmt"

    "github.com/lif0/pkg/syncx"
)

func main() {
    rm := syncx.NewReentrantMutex()

    rm.Lock()
    fmt.Println("Acquired lock")

    // Perform critical section work
    // ...

    rm.Unlock()
    fmt.Println("Released lock")
}
```

### Example: Recursive Locking

```go
package main

import (
    "fmt"

    "github.com/lif0/pkg/syncx"
)

func recursiveFunction(rm *syncx.ReentrantMutex, depth int) {
    rm.Lock()
    defer rm.Unlock()

    fmt.Printf("Recursion depth: %d\n", depth)

    if depth > 0 {
        recursiveFunction(rm, depth-1)
    }
}

func main() {
    rm := syncx.NewReentrantMutex()
    recursiveFunction(rm, 3)
}
```

### Example: Contention Handling

```go
package main

import (
    "fmt"
    "sync"
    "time"

    "github.com/lif0/pkg/syncx"
)

func main() {
    rm := syncx.NewReentrantMutex()

    rm.Lock()
    fmt.Println("Main goroutine acquired lock")

    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()

        fmt.Println("Worker trying to acquire lock...")
        rm.Lock()
        fmt.Println("Worker acquired lock")
        rm.Unlock()
    }()

    time.Sleep(1 * time.Second)
    rm.Unlock()
    fmt.Println("Main goroutine released lock")

    wg.Wait()
}
```

---

## Semaphore

The `Semaphore` type is a counting semaphore that limits the number of concurrent holders of a shared resource. It supports both limited and unlimited capacity via `Acquire`, `AcquireContext`, `TryAcquire`, `Release`, `InUse`, and `Cap`.

### Example: Limited Semaphore

```go
package main

import (
    "fmt"

    "github.com/lif0/pkg/syncx"
)

func main() {
    sem := syncx.NewSemaphore(3) // capacity of 3

    sem.Acquire()
    fmt.Printf("Acquired a slot, in use: %d/%d\n", sem.InUse(), sem.Cap())

    // critical section
    // ...

    sem.Release()
    fmt.Println("Released a slot")
}
```

### Example: Unlimited Semaphore

```go
package main

import (
    "fmt"

    "github.com/lif0/pkg/syncx"
)

func main() {
    sem := syncx.NewSemaphore(0) // 0 -> unlimited

    sem.Acquire() // no-op for unlimited semaphores
    fmt.Printf("Acquired (no-op), in use: %d, cap: %d\n", sem.InUse(), sem.Cap())

    sem.Release() // no-op for unlimited semaphores
    fmt.Println("Released (no-op)")
}
```

### Example: Context-Aware Acquisition

```go
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/lif0/pkg/syncx"
)

func main() {
    sem := syncx.NewSemaphore(2)

    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    if err := sem.AcquireContext(ctx); err != nil {
        fmt.Printf("Failed to acquire: %v\n", err)
        return
    }
    fmt.Printf("Acquired a slot with context, in use: %d/%d\n", sem.InUse(), sem.Cap())

    sem.Release()
    fmt.Println("Released a slot")
}
```

### Example: Non-Blocking Acquisition

```go
package main

import (
    "fmt"

    "github.com/lif0/pkg/syncx"
)

func main() {
    sem := syncx.NewSemaphore(1)

    sem.Acquire()
    fmt.Printf("Acquired a slot, in use: %d/%d\n", sem.InUse(), sem.Cap())

    if sem.TryAcquire() {
        fmt.Println("Acquired another slot")
    } else {
        fmt.Println("Failed to acquire: no slots available")
    }

    sem.Release()
    fmt.Println("Released a slot")
}
```

---

## WithLock

`WithLock` executes an action while holding the given lock. It guarantees the lock is released even if the action panics.

```go
package main

import (
    "fmt"
    "sync"

    "github.com/lif0/pkg/syncx"
)

func main() {
    var mu sync.Mutex
    counter := 0

    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 100; j++ {
                syncx.WithLock(&mu, func() {
                    counter++
                })
            }
        }()
    }

    wg.Wait()
    fmt.Println("Final counter:", counter) // Always 500
}
```

---

## SyncValue

`SyncValue` is a generic wrapper around a value of any type `T` that allows concurrent access with safe read and write operations protected by an `RWMutex`. Use it whenever you reach for `any + sync.Mutex`.

It provides two methods:

- `MutateValue` gives exclusive write access (`mu.Lock()` / `mu.Unlock()`).
- `ReadValue` gives shared read access (`mu.RLock()` / `mu.RUnlock()`).

Rules for both callbacks:

- The pointer is only safe to use **within** the callback. Do not store it beyond the callback.
- Avoid calling other `SyncValue` methods from inside a callback (Go mutexes are not reentrant).
- If `T` (or its fields) contains reference types (slices/maps), copying `*v` is shallow. To use the value after the callback returns, make a defensive deep copy.

### Benchmark

```text
goos: darwin
goarch: arm64
pkg: github.com/lif0/pkg/syncx
cpu: Apple M2
```

**SyncValue[int64] vs Atomic.Int64:**

```text
Benchmark_Int64_Mixed/SyncValue-8        36694729     32.53 ns/op     0 B/op         0 allocs/op
Benchmark_Int64_Mixed/Atomic.Int64-8    471930460      2.649 ns/op    0 B/op         0 allocs/op
```

**SyncValue[complex] vs Atomic.Value:**

```text
Benchmark_Complex_Mixed/SyncValue-8      14479608     82.42 ns/op     22 B/op        0 allocs/op
Benchmark_Complex_Mixed/Atomic.Value-8    1000000      1247 ns/op   28952 B/op       0 allocs/op
```

### Example: simple type

```go
import "github.com/lif0/pkg/syncx"

func main() {
    sv := syncx.NewSyncValue[int]()

    sv.MutateValue(func(v *int) { *v++ })

    var out int
    sv.ReadValue(func(v *int) { out = *v }) // safe copy
    fmt.Println(out) // 1
}
```

### Example: slice

```go
import "github.com/lif0/pkg/syncx"

func main() {
    sv := syncx.NewSyncValue[[]int]([]int{})

    sv.MutateValue(func(v *[]int) { *v = append(*v, 10) })

    var out []int
    sv.ReadValue(func(v *[]int) {
        out = make([]int, len(*v)) // allocate: a shallow copy would share the backing array
        copy(out, *v)
    })
    fmt.Println(out) // [10]
}
```

### Example: complex type

```go
import "github.com/lif0/pkg/syncx"

type User struct {
    ID    string
    Roles []string
}

type State struct {
    Users map[string]User
}

func main() {
    sv := syncx.NewSyncValue[State](State{
        Users: make(map[string]User),
    })

    sv.MutateValue(func(s *State) {
        s.Users["u1"] = User{ID: "u1", Roles: []string{"reader"}}
        s.Users["u2"] = User{ID: "u2", Roles: []string{"writer", "admin"}}
    })

    var usersSnap map[string]User
    sv.ReadValue(func(s *State) {
        usersSnap = make(map[string]User, len(s.Users))
        for id, u := range s.Users {
            // deep-copy Roles, otherwise the backing array would be shared
            usersSnap[id] = User{
                ID:    u.ID,
                Roles: append([]string(nil), u.Roles...),
            }
        }
    })
    fmt.Println("users snapshot:", usersSnap)
}
```

### Example: race (incorrect — do not do this)

```go
import "github.com/lif0/pkg/syncx"

func main() {
    sv := syncx.NewSyncValue([]int{1, 2, 3})

    // BAD: stores a shallow copy of the slice header outside the callback.
    // After the callback returns, `shared` still points to the same backing
    // array as the value inside SyncValue, so concurrent mutations race.
    var shared []int
    sv.ReadValue(func(v *[]int) {
        shared = *v // SHALLOW COPY: ptr/len/cap copied, array shared
    })
    _ = shared
}
```

---

## License

[MIT](../LICENSE)
