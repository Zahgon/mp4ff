package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// CttsBox - Composition Time to Sample Box (ctts - optional)
//
// Contained in: Sample Table Box (stbl)
type CttsBox struct {
	Version byte
	Flags   uint32
	// EndSampleNr - number (1-based) of last sample in chunk. Starts with 0 for index 0
	EndSampleNr []uint32
	// SampleOffeset - offset of first sample in chunk.
	SampleOffset []int32 // int32 to handle version 1
}

// DecodeCtts - box-specific decode
func DecodeCtts(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeCttsSR - box-specific decode
func DecodeCttsSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Adding sampleCount

// Type - box type
func (b *CttsBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *CttsBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// expectedSize - calculate size for a given entry count
func (b *CttsBox) expectedSize(entryCount uint32) uint64 { _ = "STUB: not implemented"; return 0 }

// 8 = version + flags + entryCount, 8 = sampleCount(4) + sampleOffset(4)

// Encode - write box to w
func (b *CttsBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *CttsBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// NrSampleCount - the number of SampleCount entries in box
func (b *CttsBox) NrSampleCount() int { _ = "STUB: not implemented"; return 0 }

// SampleCount - return sample count i (zero-based)
func (b *CttsBox) SampleCount(i int) uint32 { _ = "STUB: not implemented"; return 0 }

// AddSampleCountsAndOffsets - populate this box with data. Need the same number of entries in both
func (b *CttsBox) AddSampleCountsAndOffset(counts []uint32, offsets []int32) error {
	_ = "STUB: not implemented"
	return nil
}

// GetCompositionTimeOffset - composition time offset for (one-based) sampleNr in track timescale
func (b *CttsBox) GetCompositionTimeOffset(sampleNr uint32) int32 {
	_ = "STUB: not implemented"

	// This is bad index input. Should never happen
	return 0
}

// The following is essentially the sort.Search() code specialized to this case

// avoid overflow when computing h
// i ≤ h < j

// Info - get all info with specificBoxLevels ctts:1 or higher
func (b *CttsBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
