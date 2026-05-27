package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// MdiaBox - Media Box (mdia)
//
// Contained in : Track Box (trak)
// Contains all information about the media data.
type MdiaBox struct {
	Mdhd     *MdhdBox
	Hdlr     *HdlrBox
	Elng     *ElngBox
	Minf     *MinfBox
	Children []Box
}

// NewMdiaBox - Generate a new empty mdia box
func NewMdiaBox() *MdiaBox {
	_ = "STUB: not implemented"

	// AddChild - Add a child box
	return nil
}

func (m *MdiaBox) AddChild(box Box) { _ = "STUB: not implemented"; return }

// DecodeMdia - box-specific decode
func DecodeMdia(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeMdiaSR - box-specific decode
func DecodeMdiaSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (m *MdiaBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (m *MdiaBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren - list of child boxes
func (m *MdiaBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// EncodeSW - write mdia container to w
	return nil
}

func (m *MdiaBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write mdia container via sw
func (m *MdiaBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (m *MdiaBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
