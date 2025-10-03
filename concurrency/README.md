# package concurrency

[![build](https://github.com/lif0/pkg/workflows/build/badge.svg)](https://github.com/lif0/pkg/workflows/build/badge.svg)
[![go reference](https://pkg.go.dev/badge/github.com/lif0/pkg.svg)](https://pkg.go.dev/github.com/lif0/pkg/concurrency)
![last version](https://img.shields.io/github/v/tag/lif0/pkg?label=latest&filter=concurrency/*)
[![concurrency coverage](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Flif0%2Fpkg%2Frefs%2Fheads%2Fmain%2F.github%2Fassets%2Fbadges%2Fcoverage-concurrency.json)](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Flif0%2Fpkg%2Frefs%2Fheads%2Fmain%2F.github%2Fassets%2Fbadges%2Fcoverage-concurrency.json)
[![concurrency report card](https://goreportcard.com/badge/github.com/lif0/pkg/concurrency)](https://goreportcard.com/report/github.com/lif0/pkg/concurrency)

## Contents

- [Overview](#-overview)
- [Requirements](#-requirements)
- [Installation](#-installation)
- [Features](#-features)
  - [Semaphore](#semaphore)
  - [WithLock](#withlock)
  - [FutureAction](#futureaction)
  - [Promise](#promise)
  - [chanx: Channel Extension](#chanx-channel-extension)
    - [FanIn](#fanin)
    - [ToRecvChans](#torecvchans)
    - [ToSendChans](#tosendchans)
- [Roadmap](#roadmap)
- [License](#-license)

---

## 📋 Overview

The `concurrency` package provides lightweight, efficient concurrency primitives for Go, designed for correctness and performance with minimal memory allocations. It simplifies concurrent programming tasks in Go applications.

For full documentation, see [https://pkg.go.dev/github.com/lif0/pkg/concurrency](https://pkg.go.dev/github.com/lif0/pkg/concurrency).

---

## ⚙️ Requirements

- **Go 1.19 or higher**

## 📦 Installation

To add this package to your project, use `go get`:

```bash
go get github.com/lif0/pkg/concurrency@latest
```

Import the reflect extension in your code:

```go
import "github.com/lif0/pkg/concurrency"
```

---

## ✨ Features

### Semaphore

The `Semaphore` type provides a counting semaphore to limit the number of concurrent holders of a shared resource. It supports both limited and unlimited capacity, with methods like `Acquire`, `AcquireContext`, `TryAcquire`, `Release`, `InUse`, and `Cap`.

#### Example: Limited Semaphore

```go
package main

import (
    "fmt"
    "github.com/lif0/pkg/concurrency"
)

func main() {
    // Create a semaphore with a capacity of 3
    sem := concurrency.NewSemaphore(3)

    // Acquire a slot
    sem.Acquire()
    fmt.Printf("Acquired a slot, in use: %d/%d\\n", sem.InUse(), sem.Cap())
    
    // Perform critical section work
    // ...

    // Release the slot
    sem.Release()
    fmt.Println("Released a slot")
}
```

#### Example: Unlimited Semaphore

```go
package main

import (
    "fmt"
    "github.com/lif0/pkg/concurrency"
)

func main() {
    // Create an unlimited semaphore
    sem := concurrency.NewSemaphore(0)

    // Acquire is a no-op for unlimited semaphores
    sem.Acquire()
    fmt.Printf("Acquired (no-op), in use: %d, cap: %d\\n", sem.InUse(), sem.Cap())

    // Perform work
    // ...

    // Release is a no-op for unlimited semaphores
    sem.Release()
    fmt.Println("Released (no-op)")
}
```

#### Example: Context-Aware Acquisition

```go
package main

import (
    "context"
    "fmt"
    "time"
    "github.com/lif0/pkg/concurrency"
)

func main() {
    // Create a semaphore with a capacity of 2
    sem := concurrency.NewSemaphore(2)

    // Create a context with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    // Attempt to acquire a slot with context
    if err := sem.AcquireContext(ctx); err != nil {
        fmt.Printf("Failed to acquire: %v\\n", err)
        return
    }
    fmt.Printf("Acquired a slot with context, in use: %d/%d\\n", sem.InUse(), sem.Cap())

    // Perform work
    // ...

    // Release the slot
    sem.Release()
    fmt.Println("Released a slot")
}
```

#### Example: Non-Blocking Acquisition

```go
package main

import (
    "fmt"
    "github.com/lif0/pkg/concurrency"
)

func main() {
    // Create a semaphore with a capacity of 1
    sem := concurrency.NewSemaphore(1)

    // Acquire the only slot
    sem.Acquire()
    fmt.Printf("Acquired a slot, in use: %d/%d\\n", sem.InUse(), sem.Cap())

    // Try to acquire another slot without blocking
    if sem.TryAcquire() {
        fmt.Println("Acquired another slot")
    } else {
        fmt.Println("Failed to acquire: no slots available")
    }

    // Release the slot
    sem.Release()
    fmt.Println("Released a slot")
}
```

### WithLock

`WithLock` is a helper function that executes an action while holding a lock.  
It guarantees that the lock will always be released, even if the action panics.

```go
import (
 "github.com/lif0/pkg/concurrency"
)

func main() {
 var mu sync.Mutex
 counter := 0
 var wg sync.WaitGroup

 for i := 0; i < 5; i++ {
  wg.Add(1)
        wg.Go(func() {
            for j := 0; j < 100; j++ {
    concurrency.WithLock(&mu, func() {
     counter++
    })
   }
        })
 }

 wg.Wait()
 fmt.Println("Final counter:", counter) // Always 500
}
```

### FutureAction

The `FutureAction` type provides an abstraction over a channel that models a task and its result. It allows executing a computation asynchronously in a goroutine and retrieving the result later via a blocking call. This is similar to the Future pattern in other languages, providing a simple way to handle asynchronous results without manual channel management.

The channel is closed after the result is sent, ensuring proper resource cleanup.

#### Example: Basic Usage

```go
package main

import (
    "fmt"
    "time"

    "github.com/lif0/pkg/concurrency"
)

func main() {
    callback := func() any {
        time.Sleep(time.Second)
        return "success"
    }

    future := concurrency.NewFutureAction(callback)
    result := future.Get()
    fmt.Println(result) // Output: success
}
```

#### Example: Generic Type Usage

```go
package main

import (
    "fmt"
    "time"
    "github.com/lif0/pkg/concurrency"
)

func main() {
    callback := func() int {
        time.Sleep(time.Second)
        return 42
    }

    future := concurrency.NewFutureAction(callback)
    result := future.Get()
    fmt.Printf("Result: %d\n", result) // Output: Result: 42
}
```

### Promise

The `Promise` type represents a writable, single-assignment container for a future value. It allows setting a value exactly once (subsequent sets are ignored) and provides a `Future` for reading the value asynchronously. This is similar to the Promise/Future pattern in other languages, enabling clean handling of asynchronous results with thread safety via atomic operations and mutexes.

The internal channel is buffered (capacity 1) and closed after setting the value. Aliases `PromiseError` and `FutureError` are provided for error handling.

#### Example: Basic Usage

```go
package main

import (
    "fmt"
    "time"
    "github.com/lif0/pkg/concurrency"
)

func main() {
    promise := concurrency.NewPromise[string]()
    go func() {
        time.Sleep(time.Second)
        promise.Set("Cake")
    }()

    future := promise.GetFuture()
    value := future.Get()
    fmt.Println(value) // Output: Cake
}
```

#### Example: Error Handling with PromiseError

```go
package main

import (
    "errors"
    "fmt"
    "time"
    "github.com/lif0/pkg/concurrency"
)

func main() {
    promise := concurrency.NewPromise[error]()
    go func() {
        time.Sleep(time.Second)
        promise.Set(errors.New("Something went wrong"))
    }()

    future := promise.GetFuture()
    err := future.Get()
    if err != nil {
        fmt.Printf("Error: %v\n", err) // Output: Error: Something went wrong
    }
}
```

#### Example: Wrapping Existing Channel with NewFuture

```go
package main

import (
    "fmt"
    "github.com/lif0/pkg/concurrency"
)

func main() {
    ch := make(chan int, 1)
    ch <- 42
    close(ch)

    future := concurrency.NewFuture(ch)
    value := future.Get()
    fmt.Printf("Value: %d\n", value) // Output: Value: 42
}
```

### chanx: Channel Extension

The `chanx` subpackage provides extension for working with channels.

#### FanIn

`FanIn` merges multiple input channels into a single output channel. It reads concurrently from each input channel and forwards values to the output. The output channel is closed when all input channels are closed or the context is canceled. This is a non-blocking, concurrent fan-in implementation that respects context cancellation. The order of values in the output is not guaranteed.

##### Example: Basic Usage

```go
package main

import (
	"context"
	"fmt"

	"github.com/lif0/pkg/concurrency/chanx"
)

func main() {
	ctx := context.Background()
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "from ch1"
		close(ch1)
	}()
	go func() {
		ch2 <- "from ch2"
		close(ch2)
	}()

	out := chanx.FanIn(ctx, ch1, ch2)
	for v := range out {
		fmt.Println(v) // Output may vary: from ch1, from ch2
	}
}
```

##### Example: Concurrent Usage with ToRecvChans

```go
package main

import (
	"context"
	"fmt"

	"github.com/lif0/pkg/concurrency/chanx"
)

func main() {
	ctx := context.Background()
	chans := make([]chan int, 10)

	for i := 0; i < len(chans); i++ {
		chans[i] = make(chan int)
		go func(ch chan int) {
			defer close(ch)
			ch <- 1
            ch <- 2
		}(chans[i])
	}

	out := chanx.FanIn(ctx, chanx.ToRecvChans(chans)...)
	sum := 0
	for v := range out {
		sum += v
	}
	fmt.Println("Sum:", sum) // Output: Sum: 30
}
```

#### ToRecvChans

`ToRecvChans` converts a slice of bidirectional channels to a slice of receive-only channels. This ensures safe passing to functions expecting read-only channels.

Complexity: time O(n), memory O(n).

##### Example

```go
package main

import (
	"fmt"

	"github.com/lif0/pkg/concurrency/chanx"
)

func main() {
	chans := []chan int{make(chan int), make(chan int)}
	recvChans := chanx.ToRecvChans(chans)
	fmt.Printf("Type: %T\n", recvChans) // Output: Type: []<-chan int
}
```

#### ToSendChans

`ToSendChans` converts a slice of bidirectional channels to a slice of send-only channels. This ensures safe passing to functions expecting write-only channels.

Complexity: time O(n), memory O(n).

##### Example

```go
package main

import (
	"fmt"

	"github.com/lif0/pkg/concurrency/chanx"
)

func main() {
	chans := []chan string{make(chan string), make(chan string)}
	sendChans := chanx.ToSendChans(chans)
	fmt.Printf("Type: %T\n", sendChans) // Output: Type: []chan<- string
}
```

## 🗺️ Roadmap

- [ ] Channel extensions
  - [x] FanIn pattern
  - [ ] FanOut pattern
  - [ ] Dynamic select
  - [ ] Transform pattern
- [x] Semaphore structure
- [x] WithLock extension
- [x] Future construct for asynchronous programming
- [x] Promise construct for asynchronous programming
- [ ] Michael-Scott Queue (MS Queue) for lock-free concurrent queues

---

Contributions and feature suggestions are welcome 🤗.

---

## 📄 License

[MIT](./LICENSE)
