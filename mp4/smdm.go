package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// SmDmBox - Sample Mastering Display Metadata Box (smdm)
// Can be used for VP9 codec in vp09 box (VisualSampleEntryBox).
// Defined in [WebM Project].
//
// [WebM Project]: https://www.webmproject.org/vp9/mp4/
type SmDmBox struct {
	Version                 byte
	Flags                   uint32
	PrimaryRChromaticityX   uint16
	PrimaryRChromaticityY   uint16
	PrimaryGChromaticityX   uint16
	PrimaryGChromaticityY   uint16
	PrimaryBChromaticityX   uint16
	PrimaryBChromaticityY   uint16
	WhitePointChromaticityX uint16
	WhitePointChromaticityY uint16
	LuminanceMax            uint32
	LuminanceMin            uint32
}

// CreateSmDmBox - Create a new SmDmBox with specified values
func CreateSmDmBox(primaryRX, primaryRY, primaryGX, primaryGY, primaryBX, primaryBY, whitePointX, whitePointY uint16,
	luminanceMax, luminanceMin uint32) *SmDmBox {
	_ = "STUB: not implemented"
	return nil
}

const smDmBoxSize = boxHeaderSize + 4 + 8*2 + 2*4 // Header + version/flags + 8 uint16s + 2 uint32s

// DecodeSmDm - box-specific decode
func DecodeSmDm(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	// Only allow header size of 8 and correct total box size
	return *new(Box), nil
}

// DecodeSmDmSR - decode box from SliceReader
func DecodeSmDmSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	// Only allow header size of 8 and correct total box size
	return *new(Box), nil
}

// Type - box type
func (b *SmDmBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *SmDmBox) Size() uint64 {
	_ = "STUB: not implemented"

	// Encode - write box to w
	return 0
}

func (b *SmDmBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *SmDmBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *SmDmBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
