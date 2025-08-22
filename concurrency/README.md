# concurrency

Concurrency utilities for Go

📦 `go get github.com/lif0/pkg/concurrency`  
🧪 Requires **Go 1.19+**

[![Go Reference](https://pkg.go.dev/badge/github.com/lif0/pkg.svg)](https://pkg.go.dev/github.com/lif0/pkg/concurrency)
![concurrency coverage](https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Flif0%2Fpkg%2Frefs%2Fheads%2Fmain%2F.github%2Fassets%2Fbadges%2Fcoverage-concurrency.json)

---

## Contents

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
  - [Semaphore](#semaphore)
- [Roadmap](#roadmap)
- [License](#license)

---

## Overview

This package provides lightweight concurrency primitives that avoid unnecessary allocations and focus on correctness and performance.

Currently available:

- `Semaphore` — counting semaphore built using buffered channels.

---

## Installation

```bash
go get github.com/lif0/pkg/concurrency@latest
```

---

## Usage

### Semaphore

A simple counting semaphore implemented.

```go
sem := concurrency.NewSemaphore(3)

sem.Acquire()
// critical section
sem.Release()
```

Unlimited semaphore

```go
sem := concurrency.NewSemaphore(0)

sem.Release() // no panic, because for unlimited no-op
sem.Acquire()
// critical section
sem.Release()
```

## Roadmap

- [] FanIn/FunOut
- [] Future/Promise
- [] MS Queue

Contributions and ideas are welcome! 🤗

---

## License

[MIT](./LICENSE)
