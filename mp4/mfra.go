package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// MfraBox - Movie Fragment Random Access Box (mfra)
// Container for TfraBox(es) that can be used to find sync samples
type MfraBox struct {
	Tfra     *TfraBox
	Tfras    []*TfraBox
	Mfro     *MfroBox
	Children []Box
	StartPos uint64
}

// DecodeMfra - box-specific decode
func DecodeMfra(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeMfraSR - box-specific decode
func DecodeMfraSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// AddChild - add child box
func (m *MfraBox) AddChild(child Box) error { _ = "STUB: not implemented"; return nil }

// Type - returns box type
func (m *MfraBox) Type() string {
	_ = "STUB: not implemented"

	// Size - returns calculated size
	return ""
}

func (m *MfraBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write mfra container to w
func (m *MfraBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW- write mfra container via sw
func (m *MfraBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// GetChildren - list of child boxes
func (m *MfraBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Info - write box-specific information
	return nil
}

func (m *MfraBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// FindEntry - find tfra entry for given moof start offset and trackID. Return nil if not found.
func (m *MfraBox) FindEntry(moofStart uint64, trackID uint32) *TfraEntry {
	_ = "STUB: not implemented"
	return nil
}
