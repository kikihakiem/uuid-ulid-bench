package main

import (
	cryptorand "crypto/rand"
	mathrand "math/rand"
	"testing"
	"time"

	googleUUID "github.com/google/uuid"
	oklogULID "github.com/oklog/ulid"
	rsXID "github.com/rs/xid"
	satoriUUID "github.com/satori/go.uuid"
	segmentioKSUID "github.com/segmentio/ksuid"
)

var (
	mathEntropy = mathrand.New(mathrand.NewSource(time.Now().UTC().UnixNano()))

	UUIDSatori    = satoriUUID.NewV4()
	UUIDSatoriStr = UUIDSatori.String()
	UUIDGoogle    = googleUUID.New()
	UUIDGoogleStr = UUIDGoogle.String()
	ULIDCrypto    = newULIDCryptoRand()
	ULIDCryptoStr = ULIDCrypto.String()
	XID           = rsXID.New()
	XIDStr        = XID.String()
	KSUID         = segmentioKSUID.New()
	KSUIDStr      = KSUID.String()
)

func BenchmarkSatoriUUIDToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = UUIDSatori.String()
	}
}

func BenchmarkGoogleUUIDToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = UUIDGoogle.String()
	}
}

func BenchmarkOklogULIDToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ULIDCrypto.String()
	}
}

func BenchmarkRsXIDToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = XID.String()
	}
}

func BenchmarkSegmentIOKSUIDToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = KSUID.String()
	}
}

func BenchmarkSatoriUUIDFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = satoriUUID.FromString(UUIDSatoriStr)
	}
}

func BenchmarkGoogleUUIDFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = googleUUID.Parse(UUIDGoogleStr)
	}
}

func BenchmarkOklogULIDFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = oklogULID.MustParse(ULIDCryptoStr)
	}
}

func BenchmarkRsXIDFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = rsXID.FromString(XIDStr)
	}
}

func BenchmarkSegmentIOKSUIDFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = segmentioKSUID.Parse(KSUIDStr)
	}
}

func BenchmarkSatoriUUIDNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = satoriUUID.NewV4()
	}
}

func BenchmarkGoogleUUIDNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = googleUUID.New()
	}
}

func BenchmarkOklogULIDCryptoRandNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = newULIDCryptoRand()
	}
}

func BenchmarkOklogULIDMathRandNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = newULIDMathRand()
	}
}

func BenchmarkRsXIDNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = rsXID.New()
	}
}

func BenchmarkSegmentIOKSUIDNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = segmentioKSUID.New()
	}
}

func newULIDCryptoRand() oklogULID.ULID {
	return oklogULID.MustNew(oklogULID.Now(), cryptorand.Reader)
}

func newULIDMathRand() oklogULID.ULID {
	return oklogULID.MustNew(oklogULID.Now(), mathEntropy)
}
