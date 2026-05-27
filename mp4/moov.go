package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// MoovBox - Movie Box (moov - mandatory)
//
// Contains all meta-data. To be able to stream a file, the moov box should be placed before the mdat box.
type MoovBox struct {
	Mvhd     *MvhdBox
	Trak     *TrakBox // The first trak box
	Traks    []*TrakBox
	Mvex     *MvexBox
	Pssh     *PsshBox
	Psshs    []*PsshBox
	Children []Box
	StartPos uint64
}

// NewMoovBox - Generate a new empty moov box
func NewMoovBox() *MoovBox {
	_ = "STUB: not implemented"

	// AddChild - Add a child box
	return nil
}

func (m *MoovBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// Possibly re-order to keep traks together on same
// side of mvex or similar. Put this trak box after last previous trak

// last one in middle

// DecodeMoov - box-specific decode
func DecodeMoov(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeMoovSR - box-specific decode
func DecodeMoovSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box type
func (m *MoovBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (m *MoovBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren - list of child boxes
func (m *MoovBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Encode - write moov container to w
	return nil
}

func (m *MoovBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write moov container to sw
func (m *MoovBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (m *MoovBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// RemovePsshs - remove and return all psshs children boxes
func (m *MoovBox) RemovePsshs() []*PsshBox { _ = "STUB: not implemented"; return nil }

func (m *MoovBox) GetSinf(trackID uint32) *SinfBox { _ = "STUB: not implemented"; return nil }

// Get first (and only)

// IsEncrypted returns true if SampleEntryBox is "encv" or "enca"
func (m *MoovBox) IsEncrypted(trackID uint32) bool { _ = "STUB: not implemented"; return false }

// Get first (and only)
