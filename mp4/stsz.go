package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// StszBox - Sample Size Box (stsz - mandatory)
//
// Contained in : Sample Table box (stbl)
//
// For each track, either stsz of the more compact stz2 must be present. stz2 variant is not supported.
//
// This table lists the size of each sample. If all samples have the same size, it can be defined in the
// SampleUniformSize attribute.
type StszBox struct {
	Version           byte
	Flags             uint32
	SampleUniformSize uint32
	SampleNumber      uint32
	SampleSize        []uint32
}

// DecodeStsz - box-specific decode
func DecodeStsz(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeStszSR - box-specific decode
func DecodeStszSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *StszBox) Type() string {
	_ = "STUB: not implemented"

	// Size - box-specific size
	return ""
}

func (b *StszBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// expectedSize - calculate size based on SampleUniformSize and SampleNumber
func (b *StszBox) expectedSize() uint64 { _ = "STUB: not implemented"; return 0 }

// 12 = version + flags(4) + uniformSize(4) + sampleNumber(4)

// 4 bytes per sample size

// Encode - write box to w
func (b *StszBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *StszBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *StszBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// No samples

// GetNrSamples - get number of sampples
func (b *StszBox) GetNrSamples() uint32 { _ = "STUB: not implemented"; return 0 }

// GetSampleSize returns the size (in bytes) of a sample
func (b *StszBox) GetSampleSize(i int) uint32 { _ = "STUB: not implemented"; return 0 }

// One-based

// GetTotalSampleSize - get total size of a range [startNr, endNr] of samples
func (b *StszBox) GetTotalSampleSize(startNr, endNr uint32) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// 1-based numbers
