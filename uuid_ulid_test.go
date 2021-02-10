package main

import (
	"crypto/rand"
	"testing"

	googleUUID "github.com/google/uuid"
	oklogULID "github.com/oklog/ulid"
	rsXID "github.com/rs/xid"
	satoriUUID "github.com/satori/go.uuid"
	segmentioKSUID "github.com/segmentio/ksuid"
)

var (
	UUIDSatori    = satoriUUID.NewV4()
	UUIDSatoriStr = UUIDSatori.String()
	UUIDGoogle    = googleUUID.New()
	UUIDGoogleStr = UUIDGoogle.String()
	ULID          = oklogULID.MustNew(oklogULID.Now(), rand.Reader)
	ULIDStr       = ULID.String()
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
		_ = ULID.String()
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

func BenchmarkOklogUlidFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oklogULID.MustParse(ULIDStr)
	}
}

func BenchmarkRsXIDFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rsXID.FromString(XIDStr)
	}
}

func BenchmarkSegmentIOKSUIDFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		segmentioKSUID.Parse(KSUIDStr)
	}
}
