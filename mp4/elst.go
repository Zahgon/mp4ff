package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// ElstBox - Edit List Box (elst - optional)
//
// Contained in : Edit Box (edts)
type ElstBox struct {
	Version byte
	Flags   uint32
	Entries []ElstEntry
}

type ElstEntry struct {
	SegmentDuration   uint64
	MediaTime         int64
	MediaRateInteger  int16
	MediaRateFraction int16
}

// DecodeElst - box-specific decode
func DecodeElst(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeElstSR - box-specific decode
func DecodeElstSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box type
func (b *ElstBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *ElstBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// expectedSize - calculate size for a given entry count
func (b *ElstBox) expectedSize(entryCount uint32) uint64 { _ = "STUB: not implemented"; return 0 }

// 8 = version + flags + entryCount, 20 = uint64 + int64 + 2*int16

// 8 = version + flags + entryCount, 12 = uint32 + int32 + 2*int16

// Encode - write box to w
func (b *ElstBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *ElstBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *ElstBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
