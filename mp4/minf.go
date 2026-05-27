package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// MinfBox -  Media Information Box (minf - mandatory)
//
// Contained in : Media Box (mdia)
type MinfBox struct {
	Vmhd     *VmhdBox
	Smhd     *SmhdBox
	Sthd     *SthdBox
	Dinf     *DinfBox
	Stbl     *StblBox
	Children []Box
}

// NewMinfBox - Generate a new empty minf box
func NewMinfBox() *MinfBox {
	_ = "STUB: not implemented"

	// AddChild - Add a child box
	return nil
}

func (m *MinfBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// DecodeMinf - box-specific decode
func DecodeMinf(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeMinfSR - box-specific decode
func DecodeMinfSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box type
func (m *MinfBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (m *MinfBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren - list of child boxes
func (m *MinfBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Encode - write minf container to w
	return nil
}

func (m *MinfBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write minf container to sw
func (m *MinfBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (m *MinfBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
