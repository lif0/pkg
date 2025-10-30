# package sync

[![build](https://github.com/lif0/pkg/workflows/build/badge.svg)](https://github.com/lif0/pkg/workflows/build/badge.svg)
[![go reference](https://pkg.go.dev/badge/github.com/lif0/pkg.svg)](https://pkg.go.dev/github.com/lif0/pkg/sync)
![last version](https://img.shields.io/github/v/tag/lif0/pkg?label=latest&filter=sync/*)
[![sync coverage](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Flif0%2Fpkg%2Frefs%2Fheads%2Fmain%2F.github%2Fassets%2Fbadges%2Fcoverage-sync.json)](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Flif0%2Fpkg%2Frefs%2Fheads%2Fmain%2F.github%2Fassets%2Fbadges%2Fcoverage-sync.json)
[![sync report card](https://goreportcard.com/badge/github.com/lif0/pkg/sync)](https://goreportcard.com/report/github.com/lif0/pkg/sync)

## Contents

- [Overview](#-overview)
- [Requirements](#️-requirements)
- [Installation](#-installation)
- [Features](#-features)
  - [ReentrantMutex](#reentrantmutex)
    - [Performance](#performance)
    - [Example](#example)
- [Roadmap](#️-roadmap)
- [License](#-license)

---

## 📋 Overview

The `sync` package provides advanced synchronization primitives for Go, designed for correctness, performance, and ease of use in concurrent programming.

For full documentation, see [https://pkg.go.dev/github.com/lif0/pkg/sync](https://pkg.go.dev/github.com/lif0/pkg/sync).

---

## ⚙️ Requirements

- **go 1.19 or higher**

## 📦 Installation

To add this package to your project, use `go get`:

```bash
go get github.com/lif0/pkg/sync@latest
```

Import the sync extension in your code:

```go
import "github.com/lif0/pkg/sync"
```

---

## ✨ Features

### ReentrantMutex

The `ReentrantMutex` type provides a reentrant mutual exclusion lock that allows the same goroutine to acquire the lock multiple times without deadlocking. It supports recursive locking, ownership tracking, and methods like `Lock` and `Unlock`.

Important details:

- Panics on unlocking an unlocked mutex.
- Panics on unlocking from a different goroutine than the owner.
- Panics if recursion count goes negative.

#### Performance

```bash
goos: darwin
goarch: arm64
pkg: github.com/lif0/pkg/sync
cpu: Apple M2
BenchmarkMutexes/sync.Mutex-8         	                    155_743_212	         7.684 ns/op       0 B/op	       0 allocs/op
BenchmarkMutexes/sync.RWMutex-8       	                    71_501_988	        16.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutexes/sync.ReentrantMutex-8         	            65_857_020	        17.82 ns/op	       0 B/op	       0 allocs/op

BenchmarkMutexesParallel/sync.MutexParallel-8               13_736_034          73.02 ns/op        0 B/op          0 allocs/op
BenchmarkMutexesParallel/sync.RWMutexParallel-8             14_190_777          84.19 ns/op        0 B/op          0 allocs/op
BenchmarkMutexesParallel/sync.ReentrantMutexParallel-8      35_037_007          34.84 ns/op        0 B/op          0 allocs/op
```

#### Example

##### Example: Basic Usage

```go
package main

import (
    "fmt"

    "github.com/lif0/pkg/sync"
)

func main() {
    // Create a new reentrant mutex
    rm := sync.ReentrantMutex{}

    // Acquire the lock
    rm.Lock()
    fmt.Println("Acquired lock")

    // Perform critical section work
    // ...

    // Release the lock
    rm.Unlock()
    fmt.Println("Released lock")
}
```

##### Example: Recursive Locking

```go
package main

import (
    "fmt"

    "github.com/lif0/pkg/sync"
)

func recursiveFunction(rm *sync.ReentrantMutex, depth int) {
    rm.Lock()
    defer rm.Unlock()

    fmt.Printf("Recursion depth: %d\n", depth)

    if depth > 0 {
        recursiveFunction(rm, depth-1)
    }
}

func main() {
    rm := sync.ReentrantMutex{}

    recursiveFunction(rm, 3)
}
```

##### Example: Contention Handling

```go
package main

import (
    "fmt"
    "sync"
    "time"

    "github.com/lif0/pkg/sync"
)

func main() {
    rm := sync.NewReentrantMutex()

    rm.Lock()
    fmt.Println("Main goroutine acquired lock")

    wg := sync.WaitGroup{}
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

## 🗺️ Roadmap

The future direction of this package is community-driven! Ideas and contributions are highly welcome.

☹️ No idea

Contributions and feature suggestions are welcome 🤗.

---

## 📄 License

[MIT](./LICENSE)