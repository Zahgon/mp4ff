package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// StypBox - Segment Type Box (styp)
type StypBox struct {
	data []byte
}

// Copy - deep copy of Styp box.
func (b *StypBox) Copy() *StypBox { _ = "STUB: not implemented"; return nil }

// MajorBrand - major brand (4 chars)
func (b *StypBox) MajorBrand() string { _ = "STUB: not implemented"; return "" }

// MinorVersion - minor version
func (b *StypBox) MinorVersion() uint32 { _ = "STUB: not implemented"; return 0 }

// AddCompatibleBrands adds new compatible brands to Styp box.
func (b *StypBox) AddCompatibleBrands(compatibleBrands []string) { _ = "STUB: not implemented"; return }

// CompatibleBrands - slice of compatible brands (4 chars each)
func (b *StypBox) CompatibleBrands() []string { _ = "STUB: not implemented"; return nil }

// CreateStyp - Create an Styp box suitable for DASH/CMAF
func CreateStyp() *StypBox { _ = "STUB: not implemented"; return nil }

// NewStyp - new styp box with parameters
func NewStyp(majorBrand string, minorVersion uint32, compatibleBrands []string) *StypBox {
	_ = "STUB: not implemented"
	return nil
}

// DecodeStyp - box-specific decode
func DecodeStyp(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeStypSR - box-specific decode
func DecodeStypSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *StypBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *StypBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *StypBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *StypBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write specific box info to w
func (b *StypBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
