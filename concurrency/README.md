# concurrency

Concurrency utilities for Go

📦 `go get github.com/lif0/pkg/concurrency`  
🧪 Requires **Go 1.19+**

---

## Contents

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
  - [Semaphore](#semaphore)
  - [Semaphore64](#semaphore64)
- [Roadmap](#roadmap)
- [License](#license)

---

## Overview

This package provides lightweight concurrency primitives that avoid unnecessary allocations and focus on correctness and performance.

Currently available:

- `Semaphore` — counting semaphore built using buffered channels.
- `Semaphore64` — zero-allocation, atomic-based semaphore using bit operations.

---

## Installation

```bash
go get github.com/lif0/pkg/concurrency@latest
```

---

## Usage

### Semaphore

A simple counting semaphore implemented via buffered channels.

```go
sem := concurrency.NewSemaphore(3)

sem.Acquire()
// critical section
sem.Release()
```

### Semaphore64

A zero-allocation semaphore using atomic and bit manipulation. Suitable for high-performance scenarios. But it has only 64 slots;

```go
var sem concurrency.Semaphore64

ok := sem.TryAcquire()
if ok {
    // critical section
    sem.Release()
}
```

---

## Roadmap

- [] FanIn/FunOut
- [] Future/Promise
- [] MS Queue

Contributions and ideas are welcome!

---

## License

[MIT](./LICENSE)
