package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// AudioSampleEntryBox according to ISO/IEC 14496-12
type AudioSampleEntryBox struct {
	name               string
	DataReferenceIndex uint16
	ChannelCount       uint16
	SampleSize         uint16
	SampleRate         uint16 // Integer part
	Esds               *EsdsBox
	Dac3               *Dac3Box
	Dac4               *Dac4Box
	Dec3               *Dec3Box
	DfLa               *DfLaBox
	Dops               *DopsBox
	Iacb               *IacbBox
	MhaC               *MhaCBox
	Btrt               *BtrtBox
	Sinf               *SinfBox
	Children           []Box
}

// NewAudioSampleEntryBox - Create new empty mp4a box
func NewAudioSampleEntryBox(name string) *AudioSampleEntryBox {
	_ = "STUB: not implemented"
	return nil
}

func makeFixed32Uint(nr uint16) uint32 { _ = "STUB: not implemented"; return 0 }

func makeUint16FromFixed32(nr uint32) uint16 { _ = "STUB: not implemented"; return 0 }

// CreateAudioSampleEntryBox - Create new AudioSampleEntry such as mp4
func CreateAudioSampleEntryBox(name string, nrChannels, sampleSize, sampleRate uint16, child Box) *AudioSampleEntryBox {
	_ = "STUB: not implemented"
	return nil
}

// AddChild - add a child box (avcC normally, but clap and pasp could be part of visual entry)
func (a *AudioSampleEntryBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

const nrAudioSampleBytesBeforeChildren = 36

// DecodeAudioSampleEntry - decode mp4a... box
func DecodeAudioSampleEntry(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// 14496-12 8.5.2.2 Sample entry (8 bytes)
// Skip 6 reserved bytes

// 14496-12 12.2.3.2 Audio Sample entry (20 bytes)

//  reserved == 0

// Predefined + reserved

// Size of all previous data

// DecodeAudioSampleEntry - decode mp4a... box
func DecodeAudioSampleEntrySR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// 14496-12 8.5.2.2 Sample entry (8 bytes)
// Skip 6 reserved bytes

// 14496-12 12.2.3.2 Audio Sample entry (20 bytes)

//  reserved == 0

// Predefined + reserved

// Size of all previous data

// Type - return box type
func (a *AudioSampleEntryBox) Type() string {
	_ = "STUB: not implemented"

	// SetType sets the type (name) of the box
	return ""
}

func (a *AudioSampleEntryBox) SetType(name string) {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return
}

func (a *AudioSampleEntryBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (a *AudioSampleEntryBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// pre_defined and reserved

// Pre-defined and reserved
// nrAudioSampleBytesBeforeChildren bytes this far

// Only write written bytes

// Next output child boxes in order

// Encode - write box to sw
func (a *AudioSampleEntryBox) EncodeSW(sw bits.SliceWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// pre_defined and reserved

// Pre-defined and reserved
// nrAudioSampleBytesBeforeChildren bytes this far

// Next output child boxes in order

// Info - write box info to w
func (a *AudioSampleEntryBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveEncryption - remove sinf box and set type to unencrypted type
func (a *AudioSampleEntryBox) RemoveEncryption() (*SinfBox, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
