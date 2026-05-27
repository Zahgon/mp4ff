package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

const (
	sbgpInsideOffset = 65536
)

// SbgpBox - Sample To Group Box, ISO/IEC 14496-12 6'th edition 2020 Section 8.9.2
type SbgpBox struct {
	Version                 byte
	Flags                   uint32
	GroupingType            string // uint32, but takes values such as seig
	GroupingTypeParameter   uint32
	SampleCounts            []uint32
	GroupDescriptionIndices []uint32 // Starts at 65537 inside fragment, see Section 8.9.4
}

// DecodeSbgp - box-specific decode
func DecodeSbgp(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSbgpSR - box-specific decode
func DecodeSbgpSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *SbgpBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *SbgpBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// expectedSize - calculate size for a given entry count
func (b *SbgpBox) expectedSize(entryCount uint32) uint64 { _ = "STUB: not implemented"; return 0 }

// 12 = version + flags(4) + groupingType(4) + entryCount(4)

// GroupingTypeParameter

// 8 = SampleCount(4) + GroupDescriptionIndex(4)

// Encode - write box to w
func (b *SbgpBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *SbgpBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box info to w
func (b *SbgpBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
