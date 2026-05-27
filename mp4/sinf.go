package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// SinfBox -  Protection Scheme Information Box according to ISO/IEC 23001-7
type SinfBox struct {
	Frma     *FrmaBox // Mandatory
	Schm     *SchmBox // Optional
	Schi     *SchiBox // Optional
	Children []Box
}

// AddChild - Add a child box
func (b *SinfBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// DecodeSinf - box-specific decode
func DecodeSinf(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSinfSR - box-specific decode
func DecodeSinfSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box type
func (b *SinfBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *SinfBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren - list of child boxes
func (b *SinfBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Encode - write sinf container to w
	return nil
}

func (b *SinfBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write sinf container to sw
func (b *SinfBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *SinfBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
