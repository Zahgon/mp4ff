package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// SgpdBox - Sample Group Description Box, ISO/IEC 14496-12 6'th edition 2020 Section 8.9.3
// Version 0 is deprecated
type SgpdBox struct {
	Version                      byte
	Flags                        uint32
	GroupingType                 string // uint32, but takes values such as seig
	DefaultLength                uint32
	DefaultGroupDescriptionIndex uint32
	DescriptionLengths           []uint32
	SampleGroupEntries           []SampleGroupEntry
}

// DecodeSgpd - box-specific decode
func DecodeSgpd(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSgpdSR - box-specific decode
func DecodeSgpdSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *SgpdBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *SgpdBox) Size() uint64 {
	_ = "STUB: not implemented"
	// Version + Flags:4
	// GroupingType: 4
	// (v>=11) DefaultLength: 4
	// (v>=2) DefaultGroupDescriptionIndex
	// EntryCount: 4
	// SampleCount + GroupDescriptionIndex : 8
	// DescriptionLength: 4
	// SampleGroupEntries: default or individual lengths
	return 0
}

// DefaultLength

// DefaultGroupDescriptionIndex

// Encode - write box to w
func (b *SgpdBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *SgpdBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box info to w
func (b *SgpdBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
