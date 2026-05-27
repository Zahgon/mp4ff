package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// ClapBox - Clean Aperture Box, ISO/IEC 14496-12 2020 Sec. 12.1.4
type ClapBox struct {
	CleanApertureWidthN  uint32
	CleanApertureWidthD  uint32
	CleanApertureHeightN uint32
	CleanApertureHeightD uint32
	HorizOffN            uint32
	HorizOffD            uint32
	VertOffN             uint32
	VertOffD             uint32
}

// DecodeClap - box-specific decode
func DecodeClap(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeClapSR - box-specific decode
func DecodeClapSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box type
func (b *ClapBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *ClapBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *ClapBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *ClapBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *ClapBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
