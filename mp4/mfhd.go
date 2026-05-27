package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// MfhdBox - Media Fragment Header Box (mfhd)
//
// Contained in : Movie Fragment box (moof))
type MfhdBox struct {
	Version        byte
	Flags          uint32
	SequenceNumber uint32
}

// DecodeMfhd - box-specific decode
func DecodeMfhd(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeMfhdSR - box-specific decode
func DecodeMfhdSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// CreateMfhd - create an MfhdBox
func CreateMfhd(sequenceNumber uint32) *MfhdBox { _ = "STUB: not implemented"; return nil }

// Type - box type
func (m *MfhdBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (m *MfhdBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (m *MfhdBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (m *MfhdBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (m *MfhdBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
