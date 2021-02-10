# Bencmark various UUID and its variants

```
BenchmarkSatoriUUIDToString-4           19175822                62.8 ns/op
BenchmarkGoogleUUIDToString-4           15320474                67.4 ns/op
BenchmarkOklogULIDToString-4            47046517                25.7 ns/op
BenchmarkRsXIDToString-4                41172187                25.9 ns/op
BenchmarkSegmentIOKSUIDToString-4        3526119               298 ns/op
BenchmarkSatoriUUIDFromString-4          8826212               136 ns/op
BenchmarkGoogleUUIDFromString-4         34026153                34.6 ns/op
BenchmarkOklogUlidFromString-4          42967579                25.6 ns/op
BenchmarkRsXIDFromString-4              28064146                50.5 ns/op
BenchmarkSegmentIOKSUIDFromString-4      5655710               210 ns/op
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
