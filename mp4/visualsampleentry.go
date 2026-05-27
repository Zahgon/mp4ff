package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// VisualSampleEntryBox Video Sample Description box (avc1/avc3/hvc1/hev1...)
type VisualSampleEntryBox struct {
	name               string
	DataReferenceIndex uint16
	Width              uint16
	Height             uint16
	Horizresolution    uint32
	Vertresolution     uint32
	FrameCount         uint16
	CompressorName     string
	AvcC               *AvcCBox
	HvcC               *HvcCBox
	Av1C               *Av1CBox
	Av3c               *Av3cBox
	VvcC               *VvcCBox
	VppC               *VppCBox
	Btrt               *BtrtBox
	Clap               *ClapBox
	Pasp               *PaspBox
	Sinf               *SinfBox
	SmDm               *SmDmBox
	CoLL               *CoLLBox
	Children           []Box
	TrailingBytes      []byte
}

// NewVisualSampleEntryBox creates new empty box with an appropriate name such as avc1
func NewVisualSampleEntryBox(name string) *VisualSampleEntryBox {
	_ = "STUB: not implemented"
	return nil
}

// CreateVisualSampleEntryBox creates a new VisualSampleEntry such as avc1, avc3, hev1, hvc1, vvc1, vvi1
func CreateVisualSampleEntryBox(name string, width, height uint16, sampleEntry Box) *VisualSampleEntryBox {
	_ = "STUB: not implemented"
	return nil
}

// 72dpi
// 72dpi

// AddChild adds a child box and sets pointer to common types
func (b *VisualSampleEntryBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// DecodeVisualSampleEntry decodes avc1/avc3/hvc1/hev1/vvc1/vvi1 box
func DecodeVisualSampleEntry(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeVisualSampleEntrySR decodes avc1/avc3/hvc1/hev1/vvc1/vvi1 box
func DecodeVisualSampleEntrySR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// 14496-12 8.5.2.2 Sample entry (8 bytes)
// Skip 6 reserved bytes

// 14496-12 12.1.3.2 Visual Sample entry (70 bytes)

// pre_defined and reserved == 0
// 3 x 32 bits pre_defined == 0

// reserved
// Should be 1

// Skip depth
// pre_defined == -1

// Now there may be clap, pasp, btrt and other boxes
// 14496-15  5.4.2.1.2 avcC should be inside avc1, avc3 box
// Size of all previous data

// This should not happen, but was reported in issue #444

// Type returns box type
func (b *VisualSampleEntryBox) Type() string {
	_ = "STUB: not implemented"

	// SetType sets the type (name) of the box
	return ""
}

func (b *VisualSampleEntryBox) SetType(name string) {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return
}

func (b *VisualSampleEntryBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode writes box to w
func (b *VisualSampleEntryBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// pre_defined and reserved

//36 bytes

//50 bytes

// depth
// pre_defined == -1  //86 bytes

// Only write  written bytes

// Next output child boxes in order

// EncodeSW writes box to sw
func (b *VisualSampleEntryBox) EncodeSW(sw bits.SliceWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// pre_defined and reserved

//36 bytes

//50 bytes

// depth
// pre_defined == -1  //86 bytes

// Next output child boxes in order

// Info writes box-specific information
func (b *VisualSampleEntryBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveEncryption removes sinf box and set type to unencrypted type
func (b *VisualSampleEntryBox) RemoveEncryption() (*SinfBox, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConvertHev1ToHvc1 converts visual sample entry box type and insert VPS, SPS, and PPS parameter sets
func (b *VisualSampleEntryBox) ConvertHev1ToHvc1(vpss [][]byte, spss [][]byte, ppss [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

// ConvertAvc3ToAvc1 converts visual sample entry box type and insert SPS and PPS parameter sets
func (b *VisualSampleEntryBox) ConvertAvc3ToAvc1(spss [][]byte, ppss [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}
