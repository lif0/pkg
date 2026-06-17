# package errx

> Part of [**lif0/pkg**](../README.md) · [API reference](https://pkg.go.dev/github.com/lif0/pkg/errx)

Small error-handling helpers for Go.

## Contents

- [Installation](#installation)
- [MultiError](#multierror)
- [License](#license)

---

## Installation

Requires **go 1.23+**.

```bash
go get github.com/lif0/pkg@latest
```

```go
import "github.com/lif0/pkg/errx"
```

---

## MultiError

`MultiError` is a slice of errors that itself implements the `error` interface.

| Item        | Signature                                 | Purpose                                          | Notes                                                                       |
| ----------- | ----------------------------------------- | ------------------------------------------------ | --------------------------------------------------------------------------- |
| Type        | `type MultiError []error`                 | Slice of `error` that itself implements `error`. | Zero-value is usable.                                                       |
| Append      | `func (m *MultiError) Append(err error)`  | Adds a non-nil error to the set.                 | No-op for `nil`; safe on the zero value.                                    |
| MaybeUnwrap | `func (m MultiError) MaybeUnwrap() error` | Returns the simplest meaningful error.           | `len==0 → nil`, `len==1 → m[0]`, otherwise `m` itself.                       |
| Error       | `func (m MultiError) Error() string`      | Human-readable, counted, bulleted message.       | `""` when empty; format: `"<n> error(s) occurred:\n* <err1>\n* <err2>..."`. |
| IsEmpty     | `func (m MultiError) IsEmpty() bool`      | Quick emptiness check.                           | `true` when there are no errors.                                            |

### Example

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

### Batch Example

```go
var me errx.MultiError
for _, job := range jobs {
    if err := job.Run(); err != nil {
        me.Append(fmt.Errorf("job %q: %w", job.Name, err))
    }
}
return me.MaybeUnwrap()
```

---

## License

[MIT](../LICENSE)
