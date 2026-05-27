package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// EdtsBox - Edit Box (edts - optional)
//
// Contained in: Track Box ("trak")
//
// The edit box maps the presentation timeline to the media-time line
type EdtsBox struct {
	Elst     []*ElstBox
	Children []Box
}

// DecodeEdts - box-specific decode
func DecodeEdts(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeEdtsSR - box-specific decode
func DecodeEdtsSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// AddChild - Add a child box and update EntryCount
func (e *EdtsBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// Type - box type
func (b *EdtsBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *EdtsBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren - list of child boxes
func (b *EdtsBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Encode - write edts container to w
	return nil
}

func (b *EdtsBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - write edts container to sw
func (b *EdtsBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *EdtsBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
