<div align="center">
<img src=".github/assets/pkg_poster_round.png"  width="256" height="256" >
<h4 align="center">Low-level Go packages I use across my own projects. Fast, and mostly dependency-free.</h4>

<p align="center">
<!-- Build Status  -->
<a href="https://github.com/lif0/pkg/actions/workflows/build.yml">
<img src="https://github.com/lif0/pkg/actions/workflows/build.yml/badge.svg" />
</a>
<!-- Go Reference -->
<a href="https://pkg.go.dev/github.com/lif0/pkg">
<img src="https://pkg.go.dev/badge/github.com/lif0/pkg.svg" alt="Go Reference" />
</a>
<!-- Coverage -->
<a href="https://coveralls.io/github/lif0/pkg?branch=main">
<img src="https://coveralls.io/repos/github/lif0/pkg/badge.svg?branch=main" alt="Coverage Status" />
</a>
<!-- Go Report Card -->
<a href="https://goreportcard.com/report/github.com/lif0/pkg">
<img src="https://goreportcard.com/badge/github.com/lif0/pkg" />
</a>
<!-- Version -->
<a href="https://github.com/lif0/pkg/releases">
<img src="https://img.shields.io/github/v/tag/lif0/pkg?label=version&filter=v*" />
</a>
<!-- GitHub -->
<a href="https://github.com/lif0/pkg">
<img src="https://img.shields.io/github/last-commit/lif0/pkg.svg" />
</a>
</p>

<table align="center">
    <thead>
        <tr>
            <th>package</th>
            <th>doc</th>
            <th>about</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td><code>async</code></td>
            <td><a href="https://pkg.go.dev/github.com/lif0/pkg/async">go.dev</a></td>
            <td>Futures, promises, and deferred actions</td>
        </tr>
        <tr>
            <td><code>syncx</code></td>
            <td><a href="https://pkg.go.dev/github.com/lif0/pkg/syncx">go.dev</a></td>
            <td>Synchronization primitives: reentrant mutex, semaphore, sync value, <code>WithLock</code></td>
        </tr>
        <tr>
            <td><code>chanx</code></td>
            <td><a href="https://pkg.go.dev/github.com/lif0/pkg/chanx">go.dev</a></td>
            <td>Channel helpers: fan-in, send/receive conversions</td>
        </tr>
        <tr>
            <td><code>errx</code></td>
            <td><a href="https://pkg.go.dev/github.com/lif0/pkg/errx">go.dev</a></td>
            <td>Error utilities: <code>MultiError</code></td>
        </tr>
        <tr>
            <td><code>structx</code></td>
            <td><a href="https://pkg.go.dev/github.com/lif0/pkg/structx">go.dev</a></td>
            <td>Data structures: <code>OrderedMap</code>, <code>ObjectPool</code></td>
        </tr>
        <tr>
            <td><code>reflectx</code></td>
            <td><a href="https://pkg.go.dev/github.com/lif0/pkg/reflectx">go.dev</a></td>
            <td>Reflection helpers: payload-size estimation</td>
        </tr>
    </tbody>
</table>

</div>

---

## Install

```bash
go get github.com/lif0/pkg@latest
```

```go
import (
    "github.com/lif0/pkg/async"
    "github.com/lif0/pkg/syncx"
    "github.com/lif0/pkg/chanx"
    "github.com/lif0/pkg/errx"
    "github.com/lif0/pkg/structx"
    "github.com/lif0/pkg/reflectx"
)
```

> **Migrating from the old multi-module layout?** See [CHANGELOG.md](./CHANGELOG.md) for the
> import-path mapping. Old tags (`concurrency/v*`, `sync/v*`, `utils/v*`) still resolve, so
> existing pins keep working.

---

## Purpose

These are small, low-level Go packages I kept rewriting from one project to the next, so I put them in one place. Most of them just fill gaps I run into in the standard library.

What I try to hold to for each one:

- Decent test coverage (around 95%+) and benchmarks where speed matters
- Few or no external dependencies
- Small, predictable APIs
- Semver tags, so you can pin a version

---

## Stability

The module follows [semver](https://semver.org/). Exported APIs are meant to stay stable, and anything breaking lands in a new major version, so it's safe to pin a version and not worry about it.

---

## Contribution Guidelines

Contributions are welcome!

To contribute a package, feature, or bugfix:

- Coverage must be ≥95%
- All code must be tested and benchmarked
- Use of standard library imports is highly desirable and allowed; avoid external dependencies where possible
- APIs should be minimal, idiomatic, and efficient
- Follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)

To propose a new idea or package, please open an issue or discussion with:

- Motivation and use case
- Expected behavior and API shape
- Edge cases and potential risks

---

## License

[MIT](./LICENSE)
