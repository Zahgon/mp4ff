package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// StblBox - Sample Table Box (stbl - mandatory)
//
// Contained in : Media Information Box (minf)
//
// The table contains all information relevant to data samples (times, chunks, sizes, ...)
type StblBox struct {
	// Same order as in Table 1 in ISO/IEC 14496-12 Ed.6 2020
	Stsd  *StsdBox
	Stts  *SttsBox
	Ctts  *CttsBox
	Stsc  *StscBox
	Stsz  *StszBox
	Stss  *StssBox
	Stco  *StcoBox
	Co64  *Co64Box
	Sdtp  *SdtpBox
	Sbgp  *SbgpBox   // The first
	Sbgps []*SbgpBox // All
	Sgpd  *SgpdBox   // The first
	Sgpds []*SgpdBox // All
	Subs  *SubsBox
	Saio  *SaioBox
	Saiz  *SaizBox

	Children []Box
}

// NewStblBox - Generate a new empty stbl box
func NewStblBox() *StblBox {
	_ = "STUB: not implemented"

	// AddChild - Add a child box
	return nil
}

func (s *StblBox) AddChild(child Box) {
	_ = "STUB: not implemented"
	// Same order as in Table 1 in ISO/IEC 14496-12 Ed.6 2020
	return
}

// DecodeStbl - box-specific decode
func DecodeStbl(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeStblSR - box-specific decode
func DecodeStblSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (s *StblBox) Type() string {
	_ = "STUB: not implemented"

	// Size - box-specific size
	return ""
}

func (s *StblBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren - list of child boxes
func (s *StblBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Encode - write stbl container to w
	return nil
}

func (s *StblBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write stbl container to sw
func (b *StblBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (s *StblBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
