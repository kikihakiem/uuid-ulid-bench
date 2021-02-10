# Bencmark various UUID and its variants (ULID, KSUID, XID)

```
BenchmarkSatoriUUIDToString-4           14531377                69.5 ns/op            48 B/op          1 allocs/op
BenchmarkGoogleUUIDToString-4           15627694                69.2 ns/op            48 B/op          1 allocs/op
BenchmarkOklogULIDToString-4            43120147                26.8 ns/op             0 B/op          0 allocs/op
BenchmarkRsXIDToString-4                47011206                25.9 ns/op             0 B/op          0 allocs/op
BenchmarkSegmentIOKSUIDToString-4        4008602               278 ns/op               0 B/op          0 allocs/op
BenchmarkSatoriUUIDFromString-4          8877264               132 ns/op              48 B/op          1 allocs/op
BenchmarkGoogleUUIDFromString-4         37557838                33.4 ns/op             0 B/op          0 allocs/op
BenchmarkOklogULIDFromString-4          41347821                25.4 ns/op             0 B/op          0 allocs/op
BenchmarkRsXIDFromString-4              28377471                39.1 ns/op             0 B/op          0 allocs/op
BenchmarkSegmentIOKSUIDFromString-4      5267553               227 ns/op               0 B/op          0 allocs/op
BenchmarkSatoriUUIDNew-4                12319255                87.2 ns/op            16 B/op          1 allocs/op
BenchmarkGoogleUUIDNew-4                13530266                91.6 ns/op            16 B/op          1 allocs/op
BenchmarkOklogULIDCryptoRandNew-4        7175775               169 ns/op              16 B/op          1 allocs/op
BenchmarkOklogULIDMathRandNew-4          7892816               155 ns/op              16 B/op          1 allocs/op
BenchmarkRsXIDNew-4                     12624160               100 ns/op               0 B/op          0 allocs/op
BenchmarkSegmentIOKSUIDNew-4             5395443               190 ns/op               0 B/op          0 allocs/op
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
