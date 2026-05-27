package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
	"github.com/Eyevinn/mp4ff/vvc"
)

// VvcCBox - VVC Configuration Box (ISO/IEC 14496-15)
// Contains one VVCDecoderConfigurationRecord
type VvcCBox struct {
	Version byte
	Flags   uint32
	vvc.DecConfRec
}

// CreateVvcC creates a VvcC box
func CreateVvcC(naluArrays []vvc.NaluArray) (*VvcCBox, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 4 bytes

// DecodeVvcC - box-specific decode
func DecodeVvcC(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeVvcCSR - box-specific decode
func DecodeVvcCSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *VvcCBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *VvcCBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// 4 bytes for version and flags

// Encode - write box to w
func (b *VvcCBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - write box to sw
func (b *VvcCBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - box-specific Info
func (b *VvcCBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
