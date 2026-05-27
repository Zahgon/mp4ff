package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// DefaultTrakID - trakID used when generating new fragmented content
const DefaultTrakID = 1

// TrakBox - Track Box (tkhd - mandatory)
//
// Contained in : Movie Box (moov)
//
// A media file can contain one or more tracks.
type TrakBox struct {
	Tkhd     *TkhdBox
	Edts     *EdtsBox
	Mdia     *MdiaBox
	Children []Box
}

// NewTrakBox - Make a new empty TrakBox
func NewTrakBox() *TrakBox {
	_ = "STUB: not implemented"

	// AddChild - Add a child box
	return nil
}

func (t *TrakBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// DecodeTrak - box-specific decode
func DecodeTrak(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeTrakSR - box-specific decode
func DecodeTrakSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box type
func (t *TrakBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (t *TrakBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren - list of child boxes
func (t *TrakBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Encode - write trak container to w
	return nil
}

func (t *TrakBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write trak container to sw
func (b *TrakBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box info to w
func (t *TrakBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetNrSamples - get number of samples for this track defined in the parent moov box.
func (t *TrakBox) GetNrSamples() uint32 { _ = "STUB: not implemented"; return 0 }

// GetSampleData - get sample metadata for a specific interval of samples defined in moov.
// If going outside the range of available samples, an error is returned.
func (t *TrakBox) GetSampleData(startSampleNr, endSampleNr uint32) ([]Sample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createSampleFlagsFromProgressiveBoxes(stss *StssBox, sdtp *SdtpBox, sampleNr uint32) uint32 {
	_ = "STUB: not implemented"
	return 0
}

//2 = does not depend on others (I-picture). May be overridden by sdtp entry

// table starts at 0, but sampleNr is one-based

// DataRange is a range for sample data in a file relative to file start
type DataRange struct {
	Offset uint64
	Size   uint64
}

// GetRangesForSampleInterval - get ranges inside file for sample range [startSampleNr, endSampleNr]
func (t *TrakBox) GetRangesForSampleInterval(startSampleNr, endSampleNr uint32) ([]DataRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// First chunk, adapt startPoint
