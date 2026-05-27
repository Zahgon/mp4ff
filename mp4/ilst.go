package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// IlstBox - iTunes Metadata Item List Atom (ilst)
// See https://developer.apple.com/library/archive/documentation/QuickTime/QTFF/Metadata/Metadata.html
type IlstBox struct {
	Children []Box
}

// AddChild - Add a child box and update SampleCount
func (b *IlstBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// DecodeIlstSR - box-specific decode
func DecodeIlstSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeIlst - box-specific decode
func DecodeIlst(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *IlstBox) Type() string {
	_ = "STUB: not implemented"

	// Size - box-specific type
	return ""
}

func (b *IlstBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren - list of child boxes
func (b *IlstBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Encode - write ilst container to w
	return nil
}

func (b *IlstBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write ilst container to sw
func (b *IlstBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *IlstBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
