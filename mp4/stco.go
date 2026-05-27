package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// StcoBox - Chunk Offset Box (stco - mandatory)
//
// Contained in : Sample Table box (stbl)
//
// The table contains the offsets (starting at the beginning of the file) for each chunk of data for the current track.
// A chunk contains samples, the table defining the allocation of samples to each chunk is stsc.
type StcoBox struct {
	Version     byte
	Flags       uint32
	ChunkOffset []uint32
}

// DecodeStco - box-specific decode
func DecodeStco(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeStcoSR - box-specific decode
func DecodeStcoSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *StcoBox) Type() string {
	_ = "STUB: not implemented"

	// Size - box-specific size
	return ""
}

func (b *StcoBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// expectedSize - calculate size for a given entry count
func (b *StcoBox) expectedSize(entryCount uint32) uint64 { _ = "STUB: not implemented"; return 0 }

// 8 = version + flags(4) + entryCount(4), 4 bytes per offset

// Encode - write box to w
func (b *StcoBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *StcoBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *StcoBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetOffset - get offset for 1-based chunkNr.
func (b *StcoBox) GetOffset(chunkNr int) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
