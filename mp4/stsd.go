package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// StsdBox - Sample Description Box (stsd - manatory)
// See ISO/IEC 14496-12 Section 8.5.2.2
// Full Box + SampleCount
// All Children are sampleEntries
type StsdBox struct {
	Version     byte
	Flags       uint32
	SampleCount uint32
	// AvcX is a pointer to box with name avc1 or avc3
	AvcX *VisualSampleEntryBox
	// HvcX is a pointer to a box with name hvc1 or hev1
	HvcX *VisualSampleEntryBox
	// VvcX is apointer to a box with name vvc1 or vvi1
	VvcX *VisualSampleEntryBox
	// Av01 is a pointer to a box with name av01
	Av01 *VisualSampleEntryBox
	// Avs3 is a pointer to a box with name avs3
	Avs3 *VisualSampleEntryBox
	// Encv is a pointer to a box with name encv
	Encv *VisualSampleEntryBox
	// VpXX is a pointer to a box with name vp08 or vp09 (VP8 or VP9 video)
	VpXX *VisualSampleEntryBox
	// Mp4a is a pointer to a box with name mp4a
	Mp4a *AudioSampleEntryBox
	// AC3 is a pointer to a box with name ac-3
	AC3 *AudioSampleEntryBox
	// EC3 is a pointer to a box with name ec-3
	EC3 *AudioSampleEntryBox
	// AC4 is a pointer to a box with name ac-4
	AC4 *AudioSampleEntryBox
	// Opus is a pointer to a box with name Opus
	Opus *AudioSampleEntryBox
	// Iamf is a pointer to a box with name iamf
	Iamf *AudioSampleEntryBox
	// MhXX is a pointer to an MPEG-H mha1, mha2, mh1, mh2 sample entry box
	MhXX *AudioSampleEntryBox
	// Enca is a pointer to a box with name enca
	Enca *AudioSampleEntryBox
	// Wvtt is a pointer to a WvttBox
	Wvtt *WvttBox
	// Stpp is a pointer to a StppBox
	Stpp *StppBox
	// Evte is a pointer to an EvteBox
	Evte     *EvteBox
	Children []Box
}

// NewStsdBox - Generate a new empty stsd box
func NewStsdBox() *StsdBox {
	_ = "STUB: not implemented"

	// AddChild - Add a child box, set relevant pointer, and update SampleCount
	return nil
}

func (s *StsdBox) AddChild(box Box) { _ = "STUB: not implemented"; return }

// GetSampleDescription - get one of multiple descriptions
func (s *StsdBox) GetSampleDescription(index int) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeStsd - box-specific decode
func DecodeStsd(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Note higher startPos below since not simple container

// DecodeStsdSR - box-specific decode
func DecodeStsdSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Note higher startPos below since not simple container

// set by  AddChild

// Type - box-specific type
func (s *StsdBox) Type() string {
	_ = "STUB: not implemented"

	// Size - box-specific type
	return ""
}

func (s *StsdBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - box-specific encode of stsd - not a usual container
func (s *StsdBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode of stsd - not a usual container
func (s *StsdBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (s *StsdBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetBtrt returns the first BtrtBox found in StsdBox children.
func (s *StsdBox) GetBtrt() *BtrtBox { _ = "STUB: not implemented"; return nil }
