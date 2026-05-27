package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// SmhdBox - Sound Media Header Box (smhd - mandatory for sound tracks)
//
// Contained in : Media Information Box (minf)
type SmhdBox struct {
	Version byte
	Flags   uint32
	Balance uint16 // should be int16
}

// CreateSmhd - Create Sound Media Header Box (all is zero)
func CreateSmhd() *SmhdBox {
	_ = "STUB: not implemented"

	// DecodeSmhd - box-specific decode
	return nil
}

func DecodeSmhd(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSmhdSR - box-specific decode
func DecodeSmhdSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Reserved

// Type - box type
func (b *SmhdBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *SmhdBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *SmhdBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *SmhdBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Reserved

// Info - write box-specific information
func (b *SmhdBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
