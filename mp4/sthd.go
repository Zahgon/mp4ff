package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// SthdBox - Subtitle Media Header Box (sthd - for subtitle tracks)
type SthdBox struct {
	Version byte
	Flags   uint32
}

// DecodeSthd - box-specific decode
func DecodeSthd(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSthdSR - box-specific decode
func DecodeSthdSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *SthdBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *SthdBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// FullBox

// Encode - write box to w
func (b *SthdBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *SthdBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *SthdBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
