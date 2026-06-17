# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased] YYYY-MM-DD
### Added
### Fixed
### Changed

## [v1.0.0] - 2026-06-17
The repository is now a **single Go module** (`github.com/lif0/pkg`) instead of three
separate modules (`concurrency`, `sync`, `utils`). Packages were reorganized, the former
`concurrency` module was split into `async` (futures/promises) and `syncx` (locks), and a
couple of packages were renamed to stop shadowing the standard library.

### Changed (BREAKING — import paths changed)

| Old import | New import |
| --- | --- |
| `github.com/lif0/pkg/concurrency` — `Promise`, `Future`, `FutureAction` | `github.com/lif0/pkg/async` |
| `github.com/lif0/pkg/concurrency` — `Semaphore`, `SyncValue`, `WithLock` | `github.com/lif0/pkg/syncx` |
| `github.com/lif0/pkg/concurrency/chanx` | `github.com/lif0/pkg/chanx` |
| `github.com/lif0/pkg/sync` — `ReentrantMutex` | `github.com/lif0/pkg/syncx` |
| `github.com/lif0/pkg/utils/errx` | `github.com/lif0/pkg/errx` |
| `github.com/lif0/pkg/utils/structx` | `github.com/lif0/pkg/structx` |
| `github.com/lif0/pkg/utils/reflect` | `github.com/lif0/pkg/reflectx` |

- Minimum Go version is now **1.23** for the whole module.
- CI simplified to a single module: removed the per-module build/test matrix, the
  parallel Coveralls aggregation, and the badge-generating bot.

> Older tags (`concurrency/v*`, `sync/v*`, `utils/v*`) remain available, so existing pins
> keep building. New releases are published as plain `vX.Y.Z` on `github.com/lif0/pkg`.

---

## Legacy history (pre-unification)

### concurrency
- **v1.2.0** (2025-10-29) — [PKG-20](https://github.com/lif0/pkg/issues/28) Add `SyncValue[T]`
- **v1.1.0** — [PKG-1](https://github.com/lif0/pkg/issues/1) Add `chanx.FanIn` pattern
- **v1.0.1** — [PKG-11](https://github.com/lif0/pkg/issues/11) Introduce `FutureAction` pattern
- **v1.0.0** — [PKG-3](https://github.com/lif0/pkg/issues/3) Add `WithLock`; [PKG-4](https://github.com/lif0/pkg/issues/4) Add `Semaphore`

### sync
- **v1.2.0** (2025-10-29) — Set minimum go version as 1.19

### utils
- **v1.2.0** (2025-10-29) — [PKG-18](https://github.com/lif0/pkg/issues/18) Set minimum go version as 1.23; Add `OrderedMap[K]V`; [PKG-28](https://github.com/lif0/pkg/issues/28) Add `ObjectPool[T]`
- **v1.0.1** — [PKG-16] Add `errx.MultiError`
