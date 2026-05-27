package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// BtrtBox - BitRateBox - ISO/IEC 14496-12 Section 8.5.2.2
type BtrtBox struct {
	BufferSizeDB uint32
	MaxBitrate   uint32
	AvgBitrate   uint32
}

// DecodeBtrt - box-specific decode
func DecodeBtrt(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeBtrtSR - box-specific decode
func DecodeBtrtSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *BtrtBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *BtrtBox) Size() uint64 {
	_ = "STUB: not implemented"

	// Encode - write box to w
	return 0
}

func (b *BtrtBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *BtrtBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *BtrtBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
