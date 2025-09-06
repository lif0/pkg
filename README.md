<div align="center">
<img src=".github/assets/pkg_poster_round.png"  width="256" height="256" >
<h4 align="center">🚀A collection of low-level, dependency-free(mostly), high-performance Go packages🚀</h2>

<p align="center">
<!-- Build Status  -->
<a href="https://github.com/lif0/pkg/actions/">
<img src="https://github.com/lif0/pkg/workflows/build/badge.svg" />
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
            <th>badges</th>
        </tr>
        </thead>
        <tbody>
            <!-- Module concurrency -->
            <tr>
                <td>
                    <a href=".">
                        <img src="https://img.shields.io/github/v/tag/lif0/pkg?label=version&filter=concurrency/*"/>
                    </a>
                </td>
                <td>
                    <a href="https://github.com/lif0/pkg/tree/main/concurrency">
                        <img src="https://img.shields.io/badge/doc-concurrency-007d9c?logo=go&logoColor=white&style=platic" />
                    </a>
                </td>
                <td>
                    <p>Concurrency utilities for Go</p>
                </td>
                <td>
                    <a href="https://goreportcard.com/report/github.com/lif0/pkg/concurrency">
                        <img src="https://goreportcard.com/badge/github.com/lif0/pkg/concurrency" />
                    </a>
                    <a href="https://coveralls.io/github/lif0/pkg">
                    <img alt="concurrency coverage" src="https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Flif0%2Fpkg%2Frefs%2Fheads%2Fmain%2F.github%2Fassets%2Fbadges%2Fcoverage-concurrency.json">
                    </a>
                </td>
            </tr>
            <!-- Module utils -->
            <tr>
                <td>
                    <a href=".">
                        <img src="https://img.shields.io/github/v/tag/lif0/pkg?label=version&filter=utils/*"/>
                    </a>
                </td>
                <td>
                    <a href="https://github.com/lif0/pkg/tree/main/utils">
                        <img src="https://img.shields.io/badge/doc-utils-007d9c?logo=go&logoColor=white&style=platic" />
                    </a>
                </td>
                <td>
                    <p>Utils for Go</p>
                </td>
                <td>
                    <a href="https://goreportcard.com/report/github.com/lif0/pkg/utils">
                        <img src="https://goreportcard.com/badge/github.com/lif0/pkg/utils" />
                    </a>
                    <a href="https://coveralls.io/github/lif0/pkg">
                    <img alt="utils coverage" src="https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/lif0/pkg/refs/heads/main/.github/assets/badges/coverage-utils.json">
                    </a>
                </td>
            </tr>
            <!-- Module sync -->
            <tr>
                <td>
                    <a href=".">
                        <img src="https://img.shields.io/github/v/tag/lif0/pkg?label=version&filter=sync/*"/>
                    </a>
                </td>
                <td>
                    <a href="https://github.com/lif0/pkg/tree/main/sync">
                        <img src="https://img.shields.io/badge/doc-sync-007d9c?logo=go&logoColor=white&style=platic" />
                    </a>
                </td>
                <td>
                    <p>Extends the sync package.</p>
                </td>
                <td>
                    <a href="https://goreportcard.com/report/github.com/lif0/pkg/sync">
                        <img src="https://goreportcard.com/badge/github.com/lif0/pkg/sync" />
                    </a>
                    <a href="https://coveralls.io/github/lif0/pkg">
                    <img alt="sync coverage" src="https://img.shields.io/endpoint?url=https%3A%2F%2Fraw.githubusercontent.com%2Flif0%2Fpkg%2Frefs%2Fheads%2Fmain%2F.github%2Fassets%2Fbadges%2Fcoverage-sync.json">
                    </a>
                </td>
            </tr>
        </tbody>
</table>

</div>

---

## Purpose

This repository provides a set of **low-level**, production-grade packages for Go — designed for **maximum performance**, **zero allocations where possible**, and **clean, idiomatic APIs**.

The goal is to build a unified, reusable toolkit that can be safely used across personal and production systems — with strong guarantees on code quality, stability, and efficiency.

All packages are:

- Fully tested (95-100% coverage)
- Benchmarked
- API-stable and versioned (semver)
- Maintained and supported
- Free of any non-standard dependencies

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
