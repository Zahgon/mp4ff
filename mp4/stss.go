package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// StssBox - Sync Sample Box (stss - optional)
//
// Contained in : Sample Table box (stbl)
//
// This lists all sync samples (key frames for video tracks) in the data. If absent, all samples are sync samples.
type StssBox struct {
	Version      byte
	Flags        uint32
	SampleNumber []uint32
}

// DecodeStss - box-specific decode
func DecodeStss(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeStssSR - box-specific decode
func DecodeStssSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// EntryCount - number of sync samples
func (b *StssBox) EntryCount() uint32 { _ = "STUB: not implemented"; return 0 }

// Type - box-specific type
func (b *StssBox) Type() string {
	_ = "STUB: not implemented"

	// Size - box-specific size
	return ""
}

func (b *StssBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// expectedSize - calculate size for a given entry count
func (b *StssBox) expectedSize(entryCount uint32) uint64 { _ = "STUB: not implemented"; return 0 }

// IsSyncSample - check if sample (one-based) sampleNr is a sync sample
func (b *StssBox) IsSyncSample(sampleNr uint32) (isSync bool) {
	_ = "STUB: not implemented"
	// Based on a binary search algorithm from the Go standard library code.
	// i will be the lowest index such that b.SampleNumber[i] >= sampleNr
	// or len(b.SampleNumber) if not possible.
	return false
}

// i ≤ h < j

// Encode - write box to w
func (b *StssBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *StssBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *StssBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
