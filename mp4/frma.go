package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// FrmaBox - Original Format Box
type FrmaBox struct {
	DataFormat string // uint32 - original box type
}

// DecodeFrma - box-specific decode
func DecodeFrma(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeFrmaSR - box-specific decode
func DecodeFrmaSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *FrmaBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *FrmaBox) Size() uint64 {
	_ = "STUB: not implemented"

	// Encode - write box to w
	return 0
}

func (b *FrmaBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *FrmaBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box info to w
func (b *FrmaBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
