# package async

> Part of [**lif0/pkg**](../README.md) · [API reference](https://pkg.go.dev/github.com/lif0/pkg/async)

Asynchronous orchestration primitives for Go: promises, futures, and deferred actions. Lightweight, with minimal allocations.

## Contents

- [Installation](#installation)
- [FutureAction](#futureaction)
- [Promise](#promise)
- [License](#license)

---

## Installation

Requires **go 1.23+**.

```bash
go get github.com/lif0/pkg@latest
```

```go
import "github.com/lif0/pkg/async"
```

---

## FutureAction

The `FutureAction` type is an abstraction over a channel that models a task and its result. It runs a computation asynchronously in a goroutine and lets you retrieve the result later via a blocking call, similar to the Future pattern in other languages, without manual channel management. The channel is closed after the result is sent.

### Example: Basic Usage

```go
package main

import (
    "fmt"
    "time"

    "github.com/lif0/pkg/async"
)

func main() {
    callback := func() any {
        time.Sleep(time.Second)
        return "success"
    }

    future := async.NewFutureAction(callback)
    result := future.Get()
    fmt.Println(result) // Output: success
}
```

### Example: Generic Type Usage

```go
package main

import (
    "fmt"
    "time"

    "github.com/lif0/pkg/async"
)

func main() {
    callback := func() int {
        time.Sleep(time.Second)
        return 42
    }

    future := async.NewFutureAction(callback)
    result := future.Get()
    fmt.Printf("Result: %d\n", result) // Output: Result: 42
}
```

---

## Promise

The `Promise` type is a writable, single-assignment container for a future value. You set the value exactly once (later sets are ignored) and hand out a `Future` for reading it asynchronously. It is thread-safe via atomic operations; the internal channel is buffered (capacity 1) and closed after the value is set. Aliases `PromiseError` and `FutureError` are provided for error handling.

### Example: Basic Usage

```go
package main

import (
    "fmt"
    "time"

    "github.com/lif0/pkg/async"
)

func main() {
    promise := async.NewPromise[string]()
    go func() {
        time.Sleep(time.Second)
        promise.Set("Cake")
    }()

    future := promise.GetFuture()
    value := future.Get()
    fmt.Println(value) // Output: Cake
}
```

### Example: Error Handling with PromiseError

```go
package main

import (
    "errors"
    "fmt"
    "time"

    "github.com/lif0/pkg/async"
)

func main() {
    promise := async.NewPromise[error]()
    go func() {
        time.Sleep(time.Second)
        promise.Set(errors.New("something went wrong"))
    }()

    future := promise.GetFuture()
    err := future.Get()
    if err != nil {
        fmt.Printf("Error: %v\n", err) // Output: Error: something went wrong
    }
}
```

### Example: Wrapping an Existing Channel with NewFuture

```go
package main

import (
    "fmt"

    "github.com/lif0/pkg/async"
)

func main() {
    ch := make(chan int, 1)
    ch <- 42
    close(ch)

    future := async.NewFuture(ch)
    value := future.Get()
    fmt.Printf("Value: %d\n", value) // Output: Value: 42
}
```

---

## License

[MIT](../LICENSE)
