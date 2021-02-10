# Bencmark various UUID and its variants (ULID, KSUID, XID)

```
BenchmarkSatoriUUIDToString-4           15283759                71.5 ns/op            48 B/op          1 allocs/op
BenchmarkGoogleUUIDToString-4           15020476                77.8 ns/op            48 B/op          1 allocs/op
BenchmarkOklogULIDToString-4            48481560                24.7 ns/op             0 B/op          0 allocs/op
BenchmarkRsXIDToString-4                45227292                27.6 ns/op             0 B/op          0 allocs/op
BenchmarkSegmentIOKSUIDToString-4        5158242               226 ns/op               0 B/op          0 allocs/op
BenchmarkSatoriUUIDFromString-4          7424193               154 ns/op              48 B/op          1 allocs/op
BenchmarkGoogleUUIDFromString-4         34607773                32.1 ns/op             0 B/op          0 allocs/op
BenchmarkOklogULIDFromString-4          43550916                25.1 ns/op             0 B/op          0 allocs/op
BenchmarkRsXIDFromString-4              33297958                32.2 ns/op             0 B/op          0 allocs/op
BenchmarkSegmentIOKSUIDFromString-4      5592534               213 ns/op               0 B/op          0 allocs/op
BenchmarkSatoriUUIDNew-4                13561297                87.2 ns/op            16 B/op          1 allocs/op
BenchmarkGoogleUUIDNew-4                13092903                86.6 ns/op            16 B/op          1 allocs/op
BenchmarkOklogULIDCryptoRandNew-4        7750065               160 ns/op              16 B/op          1 allocs/op
BenchmarkOklogULIDMathRandNew-4          7012248               167 ns/op              16 B/op          1 allocs/op
BenchmarkRsXIDNew-4                     12273880                96.4 ns/op             0 B/op          0 allocs/op
BenchmarkSegmentIOKSUIDNew-4             6461534               204 ns/op               0 B/op          0 allocs/op
```

## Run Benchmark

```sh
$ make bench
```

## Motivation

I wanted to compare performance and features of UUID and its variants.

| Package                    | Sortable |
|----------------------------|----------|
| github.com/google/uuid     | No       |
| github.com/satori/go.uuid  | No       |
| github.com/oklog/ulid      | Yes      |
| github.com/rs/xid          | Yes      |
| github.com/segmentio/ksuid | Yes      |
