package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// Fragment - MP4 Fragment ([prft] + moof + mdat)
type Fragment struct {
	Emsgs       []*EmsgBox
	Prft        *PrftBox
	Moof        *MoofBox
	Mdat        *MdatBox
	Children    []Box       // All top-level boxes in order
	nextTrunNr  uint32      // To handle multi-trun cases
	EncOptimize EncOptimize // Bit field with optimizations being done at encoding
	StartPos    uint64      // Start position in file added by parser
}

// NewFragment creates an empty MP4 Fragment.
func NewFragment() *Fragment {
	_ = "STUB: not implemented"

	// CreateFragment creates a single track fragment
	return nil
}

func CreateFragment(seqNumber uint32, trackID uint32) (*Fragment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Can only have error when adding second track

// Data will be provided by first sample

// CreateMultiTrackFragment creates a multi-track fragment without trun boxes.
func CreateMultiTrackFragment(seqNumber uint32, trackIDs []uint32) (*Fragment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Can only have error when adding second track

// Data will be provided by first sample

// Don't add trun, but let that happen in write order

// AddChild adds a top-level box to Fragment. Add in proper order.
func (f *Fragment) AddChild(b Box) { _ = "STUB: not implemented"; return }

// ParseSenc parses any deferred or heuristic-parsed senc boxes in the fragment
// using authoritative encryption info from the init segment.
// This is needed when the init segment is in a separate file and was not available during decoding.
// The priority for determining perSampleIVSize is:
// 1. seig sample group entry in the traf (checked inside ParseReadSenc)
// 2. tenc defaultPerSampleIVSize from the init segment
// 3. heuristic using saiz sample sizes (fallback, already applied during decode if possible)
func (f *Fragment) ParseSenc(init *InitSegment) error { _ = "STUB: not implemented"; return nil }

// AddEmsg inserts an emsg box at the end of a sequence of emsg boxes at the start of the fragment.
func (f *Fragment) AddEmsg(emsg *EmsgBox) { _ = "STUB: not implemented"; return }

// Size - return size of fragment including all boxes.
// Be aware that TrafBox.OptimizeTfhdTrun() can change size
func (f *Fragment) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetFullSamples - Get full samples including media and accumulated time
func (f *Fragment) GetFullSamples(trex *TrexBox) ([]FullSample, error) {
	_ = "STUB: not implemented"
	return nil, nil

	//seqNr := moof.Mfhd.SequenceNumber
}

// This trackID may not exist for this fragment

// The first one

// The default is moofStartPos according to Section 8.8.7.1

// len should be fine for 64-bit

// Next trun start after this

// AddFullSample - add a full sample to the first (and only) trun of a track
// AddFullSampleToTrack is the more general function
func (f *Fragment) AddFullSample(s FullSample) { _ = "STUB: not implemented"; return }

// AddFullSampleToTrack - allows for adding samples to any track
// New trun boxes will be created if latest trun of fragment is not in this track
func (f *Fragment) AddFullSampleToTrack(s FullSample, trackID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// AddSample - add a sample to the first (and only) trun of a track
// AddSampleToTrack is the more general function
func (f *Fragment) AddSample(s Sample, baseMediaDecodeTime uint64) {
	_ = "STUB: not implemented"
	return
}

// AddSamples - add a slice of Sample to the first (and only) trun of a track
func (f *Fragment) AddSamples(ss []Sample, baseMediaDecodeTime uint64) {
	_ = "STUB: not implemented"
	return
}

// AddSampleToTrack - allows for adding samples to any track
// New trun boxes will be created if latest trun of fragment is not in this track
// baseMediaDecodeTime will be used only for first sample in a trun
func (f *Fragment) AddSampleToTrack(s Sample, trackID uint32, baseMediaDecodeTime uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Create first trun if needed

// latest of this track

// We are not in the latest trun. Must make a new one

// Encode - write fragment via writer
func (f *Fragment) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - write fragment via SliceWriter
func (f *Fragment) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (f *Fragment) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetChildren - return children boxes
func (f *Fragment) GetChildren() []Box {
	_ = "STUB: not implemented"

	// SetTrunDataOffsets - if writeOrder available, sort and set dataOffset in truns
	return nil
}

func (f *Fragment) SetTrunDataOffsets() { _ = "STUB: not implemented"; return }

// GetSampleNrFromTime - look up sample number from a specified time. Return error if no matching time
func (f *Fragment) GetSampleNrFromTime(trex *TrexBox, sampleTime uint64) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetSampleInterval - get SampleInterval for a fragment with only one track
func (f *Fragment) GetSampleInterval(trex *TrexBox, startSampleNr, endSampleNr uint32) (SampleInterval, error) {
	_ = "STUB: not implemented"
	return *new(SampleInterval), nil
}

// AddSampleInterval - add SampleInterval for a fragment with only one track
func (f *Fragment) AddSampleInterval(sItvl SampleInterval) error {
	_ = "STUB: not implemented"
	return nil
}

// CommonSampleDuration returns a common non-zero sample duration for a track defined by trex if available.
func (f *Fragment) CommonSampleDuration(trex *TrexBox) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
