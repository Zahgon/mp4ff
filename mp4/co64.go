package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// Co64Box - Chunk Large Offset Box
//
// Contained in : Sample Table box (stbl)
//
// 64-bit version of StcoBox
type Co64Box struct {
	Version     byte
	Flags       uint32
	ChunkOffset []uint64
}

// DecodeCo64 - box-specific decode
func DecodeCo64(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeCo64 - box-specific decode
func DecodeCo64SR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *Co64Box) Type() string { _ = "STUB: not implemented"; return "" }

func (b *Co64Box) expectedSize(nrEntries uint32) uint64 { _ = "STUB: not implemented"; return 0 }

// Size - box-specific size
func (b *Co64Box) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *Co64Box) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *Co64Box) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *Co64Box) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetOffset - get offset for 1-based chunkNr.
func (b *Co64Box) GetOffset(chunkNr int) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
