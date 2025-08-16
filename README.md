# Golang pkg
Golang additional packages

## concurrency
Extensions for concurrent programming.

- **Semaphore64** — an optimized semaphore for up to 64 slots (based on bit operations over bytes; super-fast and memory-efficient).
- **Semaphore** — a regular semaphore implementation based on channels.  


## sync
Synchronization primitives beyond the standard library.

- **ReentrantLock** — lock implementation that can be safely acquired multiple times by the same goroutine.

## semantic
Helpers for primitive types

- **Estimate** — estimate payload size (in bytes) for built-in types and their pointers/slices/arrays ([zero allocations, 0 B/op](/semantic/estimate_payload_bench_out.txt))
