package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// TrefBox -  // TrackReferenceBox - ISO/IEC 14496-12 Ed. 9 Sec. 8.3
type TrefBox struct {
	Children []Box
}

// AddChild - Add a child box
func (b *TrefBox) AddChild(box Box) { _ = "STUB: not implemented"; return }

// DecodeTref - box-specific decode
func DecodeTref(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeTrefSR - box-specific decode
func DecodeTrefSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box type
func (b *TrefBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *TrefBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren - list of child boxes
func (b *TrefBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Encode - write minf container to w
	return nil
}

func (b *TrefBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write minf container to sw
func (b *TrefBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *TrefBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// TrefTypeBox - TrackReferenceTypeBox - ISO/IEC 14496-12 Ed. 9 Sec. 8.3
// Name can be one of hint, cdsc, font, hind, vdep, vplx, subt (ISO/IEC 14496-12)
// dpnd, ipir, mpod, sync (ISO/IEC 14496-14)
type TrefTypeBox struct {
	Name     string
	TrackIDs []uint32
}

// DecodeTrefType - box-specific decode
func DecodeTrefType(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeTrefTypeSR - box-specific decode
func DecodeTrefTypeSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box type
func (b *TrefTypeBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *TrefTypeBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (t *TrefTypeBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write box to sw
func (b *TrefTypeBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *TrefTypeBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
