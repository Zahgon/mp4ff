package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// UdtaBox - User Data Box is a container for User Data
//
// Contained in : moov, trak, moof, or traf
type UdtaBox struct {
	Children []Box
}

// AddChild - Add a child box
func (b *UdtaBox) AddChild(box Box) { _ = "STUB: not implemented"; return }

// DecodeUdta - box-specific decode
func DecodeUdta(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeUdtaSR - box-specific decode
func DecodeUdtaSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box type
func (b *UdtaBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *UdtaBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren - list of child boxes
func (b *UdtaBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Encode - write udta container to w
	return nil
}

func (b *UdtaBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write udta container to sw
func (b *UdtaBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *UdtaBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
