# package chanx

> Part of [**lif0/pkg**](../README.md) · [API reference](https://pkg.go.dev/github.com/lif0/pkg/chanx)

Channel helpers for Go: fan-in and safe send/receive conversions.

## Contents

- [Installation](#installation)
- [FanIn](#fanin)
- [ToRecvChans](#torecvchans)
- [ToSendChans](#tosendchans)
- [License](#license)

---

## Installation

Requires **go 1.23+**.

```bash
go get github.com/lif0/pkg@latest
```

```go
import "github.com/lif0/pkg/chanx"
```

---

## FanIn

`FanIn` merges multiple input channels into a single output channel. It reads concurrently from each input and forwards values to the output. The output channel is closed once all inputs are closed or the context is canceled. This is a non-blocking, concurrent fan-in that respects context cancellation. The order of values in the output is not guaranteed.

### Example: Basic Usage

```go
package main

import (
    "context"
    "fmt"

    "github.com/lif0/pkg/chanx"
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

### Example: Concurrent Usage with ToRecvChans

```go
package main

import (
    "context"
    "fmt"

    "github.com/lif0/pkg/chanx"
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

---

## ToRecvChans

`ToRecvChans` converts a slice of bidirectional channels into a slice of receive-only channels, so they can be safely passed to functions expecting read-only channels.

Complexity: time O(n), memory O(n).

### Example

```go
package main

import (
    "fmt"

    "github.com/lif0/pkg/chanx"
)

func main() {
    chans := []chan int{make(chan int), make(chan int)}
    recvChans := chanx.ToRecvChans(chans)
    fmt.Printf("Type: %T\n", recvChans) // Output: Type: []<-chan int
}
```

---

## ToSendChans

`ToSendChans` converts a slice of bidirectional channels into a slice of send-only channels, so they can be safely passed to functions expecting write-only channels.

Complexity: time O(n), memory O(n).

### Example

```go
package main

import (
    "fmt"

    "github.com/lif0/pkg/chanx"
)

func main() {
    chans := []chan string{make(chan string), make(chan string)}
    sendChans := chanx.ToSendChans(chans)
    fmt.Printf("Type: %T\n", sendChans) // Output: Type: []chan<- string
}
```

---

## License

[MIT](../LICENSE)
