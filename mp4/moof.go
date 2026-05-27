package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// MoofBox -  Movie Fragment Box (moof)
//
// Contains all meta-data. To be able to stream a file, the moov box should be placed before the mdat box.
type MoofBox struct {
	Mfhd     *MfhdBox
	Traf     *TrafBox // The first traf child box
	Trafs    []*TrafBox
	Pssh     *PsshBox
	Psshs    []*PsshBox
	Children []Box
	StartPos uint64
}

// DecodeMoof - box-specific decode
func DecodeMoof(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeMoofSR - box-specific decode
func DecodeMoofSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// AddChild - add child box
func (m *MoofBox) AddChild(child Box) error { _ = "STUB: not implemented"; return nil }

// Type - returns box type
func (m *MoofBox) Type() string {
	_ = "STUB: not implemented"

	// Size - returns calculated size
	return ""
}

func (m *MoofBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write moof after updating trun dataoffset
func (m *MoofBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write moof after updating trun dataoffset
func (m *MoofBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// GetChildren - list of child boxes
func (m *MoofBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Info - write box-specific information
	return nil
}

func (m *MoofBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// RemovePsshs - remove and return all psshs children boxes
func (m *MoofBox) RemovePsshs() (psshs []*PsshBox, totalSize uint64) {
	_ = "STUB: not implemented"
	return nil, 0
}
