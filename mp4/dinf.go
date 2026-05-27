package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// DinfBox - Data Information Box (dinf - mandatory)
//
// Contained in : Media Information Box (minf) or Meta Box (meta)
type DinfBox struct {
	Dref     *DrefBox
	Children []Box
}

// AddChild - Add a child box
func (d *DinfBox) AddChild(box Box) { _ = "STUB: not implemented"; return }

// DecodeDinf - box-specific decode
func DecodeDinf(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeDinfSR - box-specific decode
func DecodeDinfSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (d *DinfBox) Type() string {
	_ = "STUB: not implemented"

	// Size - box-specific size
	return ""
}

func (d *DinfBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren - list of child boxes
func (d *DinfBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Encode - write dinf container to w
	return nil
}

func (d *DinfBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - write container using slice writer
func (d *DinfBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box info to w
func (d *DinfBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
