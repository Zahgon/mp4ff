package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// FtypBox - File Type Box (ftyp - mandatory in full file/init segment)
type FtypBox struct {
	data []byte
}

// Copy - deep copy of Ftyp box.
func (b *FtypBox) Copy() *FtypBox { _ = "STUB: not implemented"; return nil }

// MajorBrand - major brand (4 chars)
func (b *FtypBox) MajorBrand() string { _ = "STUB: not implemented"; return "" }

// MinorVersion - minor version
func (b *FtypBox) MinorVersion() uint32 { _ = "STUB: not implemented"; return 0 }

// AddCompatibleBrands adds new compatible brands to Ftyp box.
func (b *FtypBox) AddCompatibleBrands(compatibleBrands []string) { _ = "STUB: not implemented"; return }

// CompatibleBrands - slice of compatible brands (4 chars each)
func (b *FtypBox) CompatibleBrands() []string { _ = "STUB: not implemented"; return nil }

// CreateFtyp - Create an Ftyp box suitable for DASH/CMAF
func CreateFtyp() *FtypBox { _ = "STUB: not implemented"; return nil }

// NewFtyp - new ftyp box with parameters
func NewFtyp(majorBrand string, minorVersion uint32, compatibleBrands []string) *FtypBox {
	_ = "STUB: not implemented"
	return nil
}

// DecodeFtyp - box-specific decode
func DecodeFtyp(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeFtypSR - box-specific decode
func DecodeFtypSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *FtypBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *FtypBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *FtypBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *FtypBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write specific box info to w
func (b *FtypBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
