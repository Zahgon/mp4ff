package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// NmhdBox - Null Media Header Box (nmhd - often used instead of sthd for subtitle tracks)
type NmhdBox struct {
	Version byte
	Flags   uint32
}

// DecodeNmhd - box-specific decode
func DecodeNmhd(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeNmhdSR - box-specific decode
func DecodeNmhdSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *NmhdBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *NmhdBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// FullBox

// Encode - write box to w
func (b *NmhdBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *NmhdBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *NmhdBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
