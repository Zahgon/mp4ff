package iamf

import (
	"errors"

	"github.com/Eyevinn/mp4ff/bits"
)

var (
	ErrInvalidLeb = errors.New("invalid Leb128")
)

// ReadLeb128 reads an unsigned Leb128 (Little Endian Base 128) encoded integer
func ReadLeb128(sr bits.SliceReader) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteLeb128 writes an unsigned Leb128 (Little Endian Base 128) encoded integer
func WriteLeb128(sw bits.SliceWriter, value uint64) { _ = "STUB: not implemented"; return }

// Leb128Size calculates the size in bytes of a Leb128 (Little Endian Base 128) encoded integer
func Leb128Size(value uint64) int { _ = "STUB: not implemented"; return 0 }
