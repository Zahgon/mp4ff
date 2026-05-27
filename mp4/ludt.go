package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// LudtBox - Track loudness container
//
// Contained in : Udta Box (udta)
type LudtBox struct {
	Loudness      []*LoudnessBaseBox
	AlbumLoudness []*LoudnessBaseBox
	Children      []Box
}

// DecodeLudt - box-specific decode
func DecodeLudt(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeLudtSR - box-specific decode
func DecodeLudtSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// AddChild - add child box
func (b *LudtBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// Size - calculated size of box
func (b *LudtBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Type - return box type
func (b *LudtBox) Type() string {
	_ = "STUB: not implemented"

	// Encode - write ludt container to w
	return ""
}

func (b *LudtBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write ludt container to sw
func (b *LudtBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// GetChildren - list of child boxes
func (b *LudtBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Info - write box-specific information
	return nil
}

func (b *LudtBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
