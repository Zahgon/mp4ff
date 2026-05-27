package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// VmhdBox - Video Media Header Box (vhmd - mandatory for video tracks)
//
// Contained in : Media Information Box (minf)
type VmhdBox struct {
	Version      byte
	Flags        uint32
	GraphicsMode uint16
	OpColor      [3]uint16
}

// CreateVmhd - Create Video Media Header Box
func CreateVmhd() *VmhdBox {
	_ = "STUB: not implemented"
	// Flags should be 0x000001 according to ISO/IEC 14496-12 Sec.12.1.2.1
	return nil
}

// DecodeVmhd - box-specific decode
func DecodeVmhd(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeVmhdSR - box-specific decode
func DecodeVmhdSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *VmhdBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *VmhdBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *VmhdBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *VmhdBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *VmhdBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
