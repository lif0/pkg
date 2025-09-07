# Contributing to pkg

Thank you for your interest in contributing to **pkg**! This repository is a collection of low-level, dependency-free (mostly), high-performance Go packages designed for maximum efficiency, zero allocations where possible, and clean, idiomatic APIs. We aim to maintain production-grade quality with strong guarantees on stability and code quality.

Contributions are welcome, whether it's reporting bugs, suggesting features, fixing issues, or adding new packages. All contributions help make this toolkit better for the Go community.

## Code of Conduct

This project adheres to the [Contributor Covenant Code of Conduct](https://www.contributor-covenant.org/version/2/0/code_of_conduct.html). By participating, you are expected to uphold this code. Please report unacceptable behavior to the project maintainers.

## How to Contribute

### Submitting Pull Requests

We encourage pull requests for bug fixes, features, documentation improvements, or new packages. Follow these steps:

1. **Fork the Repository**: Create your own fork of the repo.

2. **Create a Branch**: Branch names must follow the format `<package>/<type>/<description>`, where:
   - `<package>` is the affected package (e.g., `utils`, `concurrency`, `sync`).
   - `<type>` is the change type (e.g., `feature`, `fix`, `docs`, `test`).
   - `<description>` is a short, descriptive name (use hyphens for spaces).
   
   Examples:
   - `utils/feature/add-new-string-util`
   - `concurrency/fix/deadlock-in-semaphore`
   - `sync/docs/update-mutex-usage`
   - `utils/feature/PKG-123`

3. **Make Changes**:
   - Ensure your code aligns with the project's goals: low-level, high-performance, minimal dependencies (prefer standard library).
   - APIs should be idiomatic, efficient, and minimal.
   - Add or update tests to maintain ≥95% coverage.
   - Include benchmarks for performance-critical code.
   - Format code with `gofmt` and ensure it passes `go vet` and `golint` (or equivalent linters).
   - Update documentation (e.g., godoc comments) and README if needed.

4. **Commit Your Changes**: Use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for commit messages. Examples:
   - `feat(concurrency): add new semaphore implementation`
   - `[PKG-XX]feat(concurrency): add new semaphore implementation`
   - `fix(utils): resolve panic in edge case`
   - `docs: update README with new package details`
   - `test(sync): add benchmarks for mutex`

5. **Test Thoroughly**:
   - Run `go test ./...` to ensure all tests pass.
   - Aim for 95-100% test coverage (use `go test -cover`).
   - Benchmark with `go test -bench=.` and ensure no regressions.

6. **Push and Open a Pull Request**: Target your PR to the current open release branch (e.g., `release/vX.Y.Z`). Check the repository branches or recent PRs to identify the active release branch. Use the [Pull Request template](.github/pull_request_template.md) if available. Reference any related issues (e.g., "Closes #123").

7. **Review Process**: Maintainers will review your PR. Be responsive to feedback. Once approved, it will be merged.

## Style Guide

- Follow Go's standard conventions (effective Go).
- Use meaningful variable/function names.
- Comment code thoroughly for godoc.
- Avoid allocations in hot paths; profile if unsure.

## License

By contributing, you agree that your contributions will be licensed under the [MIT License](LICENSE). Ensure any new files include the license header if required.

Thank you for helping improve **pkg**! If you have questions, feel free to open an issue or discussion. 🚀
