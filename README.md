# fastrand

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

:warning: API is not stable.

## Performance

Tests run on a `Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz` with Turbo Boost disabled:

### `GOMAXPROCS=1`

| PRNG         | Plain  | Atomic              | Sharded |
| ------------ | ------ | ------------------- | ------- |
| SplitMix64   | 2.02ns | 8.72ns              | 7.33ns  |
| PCG-XSH-RR   | 3.17ns | 11.90ns             | 7.33ns  |
| Xoshiro256** | 4.57ns | -                   | 12.40ns |
| math/rand    | 7.06ns | 24.20ns<sup>1</sup> | -       |

### `GOMAXPROCS=8`

| PRNG         | Plain  | Atomic              | Sharded |
| ------------ | ------ | ------------------- | ------- |
| SplitMix64   | 0.29ns | 26.20ns             | 1.33ns  |
| PCG-XSH-RR   | 0.41ns | 13.20ns             | 1.34ns  |
| Xoshiro256** | 0.81ns | -                   | 2.12ns  |
| math/rand    | 1.19ns | 72.40ns<sup>1</sup> | -       |

<sup>1</sup>: the `math/rand` atomic variant is not a pure non-locking implementation, since it is implemented by guarding a `rand.Rand` using a `sync.Mutex`.

