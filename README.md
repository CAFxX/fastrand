# fastrand

[![Go Report](https://goreportcard.com/badge/github.com/CAFxX/fastrand)](https://goreportcard.com/report/github.com/CAFxX/fastrand) [![Go Reference](https://pkg.go.dev/badge/github.com/CAFxX/fastrand.svg)](https://pkg.go.dev/github.com/CAFxX/fastrand) :warning: API is not stable yet.

Some fast, non-cryptographic PRNG sources, in three variants:

- **Plain** - the basic implementation. Fastest, but can not be used concurrently.
- **Atomic** - implementation using atomic operations. Non-locking, can be used concurrently, but a bit slower (especially at high concurrency).
- **Sharded** - implementation using per-thread (P) sharding. Non-locking, can be used concurrently, fast (even at high concurrency), but does not support explicit seeding.

PRNGs currently implemented:

| Name                                                         | State (bits) | Output (bits) | Period            | Variants               |
| ------------------------------------------------------------ | ------------ | ------------- | ----------------- | ---------------------- |
| [SplitMix64](https://dl.acm.org/doi/10.1145/2714064.2660195) | 64           | 64            | 2<sup>64</sup>    | Plain, Atomic, Sharded |
| [PCG-XSH-RR](https://www.pcg-random.org/)                    | 64           | 32            | 2<sup>64</sup>    | Plain, Atomic, Sharded |
| [Xoshiro256**](http://prng.di.unimi.it/)                     | 256          | 64            | 2<sup>256</sup>-1 | Plain, Sharded         |

## Performance

Tests run on a `Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz` with Turbo Boost disabled. Lower is better.

### `GOMAXPROCS=1`

| PRNG         |  Plain |              Atomic | Sharded |
| ------------ | -----: | ------------------: | ------: |
| SplitMix64   | 2.02ns |              8.72ns |  7.33ns |
| PCG-XSH-RR   | 3.17ns |             11.90ns |  7.33ns |
| Xoshiro256** | 4.57ns |       -<sup>1</sup> | 12.40ns |
| math/rand    | 7.06ns | 24.20ns<sup>2</sup> |       - |

### `GOMAXPROCS=8`

| PRNG         |  Plain |              Atomic | Sharded |
| ------------ | -----: | ------------------: | ------: |
| SplitMix64   | 0.29ns |             26.20ns |  1.33ns |
| PCG-XSH-RR   | 0.41ns |             13.20ns |  1.34ns |
| Xoshiro256** | 0.81ns |       -<sup>1</sup> |  2.12ns |
| math/rand    | 1.19ns | 72.40ns<sup>2</sup> |       - |

## Usage notes

### Atomic variants

The atomic variant currently relies on `unsafe` to improve the performance of its CAS loops. It does so by calling the unexported `procyield` function in package `runtime`. This dependency will be removed in a future release. Usage of `unsafe` can be disabled by setting the `fastrand_nounsafe` build tag, at the cost of lower performance.

The state of the atomic variants is not padded/aligned to fill the cacheline: if needed users should pad the structure to avoid false sharing of the cacheline.

### Sharded variants

Sharded variants rely on `unsafe` to implement sharding. They do so by calling the unexported `procPin` and `procUnpin` functions in package `runtime`. These functions are used by other packages (e.g. `sync`) for the same purpose, so they are unlikely to disappear/change. Usage of `unsafe` can be disabled by setting the `fastrand_nounsafe` build tag, at the cost of lower performance.

Sharded variants detect the value of `GOMAXPROCS` when they are instantiated (with `NewShardedXxx`). If `GOMAXPROCS` is increased after a sharded PRNG is instantiated it will yield suboptimal performance, as it may dynamically fallback to the corresponding atomic variant.

Sharded variants use more memory for the state than the other variants: in general they use at least `GOMAXPROCS * 64` bytes. This is done to avoid false sharing of cachelines between shards.

Sharded variants do not allow explicit seeding since there is no easy way for a user to obtain a deterministic sequence from these variants (because, in general, goroutines can migrate between threads at any time).

## License

[MIT](LICENSE)

---

<sup>1</sup> there is no atomic variant for Xoshiro256** because its large state is not amenable to a performant atomic implementation.

<sup>2</sup> the `math/rand` atomic variant is not a pure non-locking implementation, since it is implemented by guarding a `rand.Rand` using a `sync.Mutex`.
