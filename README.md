<div align="center">
<img src=".github/assets/pkg_poster_round.png"  width="256" height="256" > 
<h4 align="center">A collection of low-level, dependency-free(mostly), high-performance Go packages🚀</h2>


</div>




## Purpose

This repository provides a set of **low-level**, production-grade packages for Go — designed for **maximum performance**, **zero allocations where possible**, and **clean, idiomatic APIs**.

The goal is to build a unified, reusable toolkit that can be safely used across personal and production systems — with strong guarantees on code quality, stability, and efficiency.

All packages are:

- Fully tested (95–100% coverage)
- Benchmarked
- API-stable and versioned (semver)
- Maintained and supported
- Free of any non-standard dependencies

---

## Packages

### [concurrency](./concurrency/README.md)

Extensions for concurrent programming.

- **Semaphore** — traditional semaphore with capacity.

### [sync](./sync/README.md)

Synchronization primitives beyond the standard library.

- **ReentrantLock** — mutex-like structure that can be safely acquired multiple times by the same goroutine

### [semantic](./semantic/README.md)

Helpers for semantic runtime operations.

- **EstimatePayloadOf** — estimate size of a value in memory (bytes); supports scalars, strings, slices, arrays, pointers — with **0 B/op** ([see benchmarks](./semantic/estimate_payload_bench_out.txt))

---

## Stability

This repository follows [semantic versioning](https://semver.org/).  
All exported APIs are stable, and breaking changes will be reflected in the major version.

Releases are tagged and versioned. You can safely pin versions for use in production.

---

## Contribution Guidelines

Contributions are welcome!

To contribute a package, feature, or bugfix:

- Coverage must be ≥95%
- All code must be tested and benchmarked
- Very desirable use standard library imports are allowed
- APIs should be minimal, idiomatic, and efficient
- Conventional Commits [(see)](https://www.conventionalcommits.org/en/v1.0.0/)

To propose a new idea or package, please open an issue or discussion with:

- Motivation and use case
- Expected behavior and API shape
- Edge cases and potential risks

---

## License

[MIT](./LICENSE)
