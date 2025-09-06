# concurrency

Concurrency utilities for Go

📦 `go get github.com/lif0/pkg/concurrency@latest`  
🧪 Requires **Go 1.19+**

[![build](https://github.com/lif0/pkg/workflows/build/badge.svg)](https://github.com/lif0/pkg/workflows/build/badge.svg)
[![go reference](https://pkg.go.dev/badge/github.com/lif0/pkg.svg)](https://pkg.go.dev/github.com/lif0/pkg/concurrency)
![last version](https://img.shields.io/github/v/tag/lif0/pkg?label=latest&filter=concurrency/*)
[![concurrency coverage](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Flif0%2Fpkg%2Frefs%2Fheads%2Fmain%2F.github%2Fassets%2Fbadges%2Fcoverage-concurrency.json)](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Flif0%2Fpkg%2Frefs%2Fheads%2Fmain%2F.github%2Fassets%2Fbadges%2Fcoverage-concurrency.json)
[![concurrency report card](https://goreportcard.com/badge/github.com/lif0/pkg/concurrency)](https://goreportcard.com/report/github.com/lif0/pkg/concurrency)

---

## Contents

- [Overview](#-overview)
- [Requirements](#-requirements)
- [Installation](#-installation)
- [Features](#-features)
  - [Semaphore](#semaphore)
  - [WithLock](#withlock)
- [Roadmap](#roadmap)
- [License](#-license)

---

## 📋 Overview

The `concurrency` package provides lightweight, efficient concurrency primitives for Go, designed for correctness and performance with minimal memory allocations. It simplifies concurrent programming tasks in Go applications.

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

## 🗺️ Roadmap

- FanIn/FanOut patterns for channel-based concurrency.
- Future/Promise constructs for asynchronous programming.
- Michael-Scott Queue (MS Queue) for lock-free concurrent queues.

Contributions and feature suggestions are welcome 🤗.

---

## 📄 License

[MIT](./LICENSE)
