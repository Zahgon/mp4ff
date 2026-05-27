package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// MvexBox - MovieExtendsBox (mevx)
//
// Contained in : Movie Box (moov)
//
// Its presence signals a fragmented asset
type MvexBox struct {
	Mehd     *MehdBox
	Trex     *TrexBox
	Trexs    []*TrexBox
	Children []Box
}

// NewMvexBox - Generate a new empty mvex box
func NewMvexBox() *MvexBox {
	_ = "STUB: not implemented"

	// AddChild - Add a child box
	return nil
}

func (m *MvexBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// DecodeMvex - box-specific decode
func DecodeMvex(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeMvex - box-specific decode
func DecodeMvexSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (m *MvexBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (m *MvexBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren - list of child boxes
func (m *MvexBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Encode - write mvex container to w
	return nil
}

func (m *MvexBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write mvex container to sw
func (m *MvexBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (m *MvexBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetTrex - get trex box for trackID
func (m *MvexBox) GetTrex(trackID uint32) (trex *TrexBox, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}
