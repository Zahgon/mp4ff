package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// MetaBox is MPEG-4 Meta box or QuickTime meta Atom (without version and flags)

// MPEG box defined in ISO/IEC 14496-12 Ed. 6 2020 Section 8.11
//
// Note. QuickTime meta atom has no version and flags field.
// https://developer.apple.com/library/archive/documentation/QuickTime/QTFF/Metadata/Metadata.html#//apple_ref/doc/uid/TP40000939-CH1-SW10
type MetaBox struct {
	Version     byte
	Flags       uint32
	Hdlr        *HdlrBox
	Children    []Box
	isQuickTime bool // Has no version and flags
}

// IsQuickTime returns true if box is QuickTime compatible (has no version and flags)
func (m *MetaBox) IsQuickTime() bool { _ = "STUB: not implemented"; return false }

// CreateMetaBox creates a new MetaBox
func CreateMetaBox(version byte, hdlr *HdlrBox) *MetaBox { _ = "STUB: not implemented"; return nil }

// AddChild adds a child box
func (b *MetaBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// DecodeMeta decodes a MetaBox in either MPEG or QuickTime version
func DecodeMeta(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeMetaSR decodes a MetaBox in either MPEG or QuickTime version
func DecodeMetaSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

//Note larger offset below since not simple container

// Type returns box type
func (b *MetaBox) Type() string {
	_ = "STUB: not implemented"

	// Size calculates size of box
	return ""
}

func (b *MetaBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren lists child boxes
func (b *MetaBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Encode writes minf container to w
	return nil
}

func (b *MetaBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode writes minf container to sw
func (b *MetaBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info writes box-specific info
func (b *MetaBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
