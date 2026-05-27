package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// MediaSegment is an MP4 Media Segment with one or more Fragments.
type MediaSegment struct {
	Styp        *StypBox
	Sidx        *SidxBox   // The first sidx box in a segment
	Sidxs       []*SidxBox // All sidx boxes in a segment
	Fragments   []*Fragment
	EncOptimize EncOptimize
	StartPos    uint64 // Start position in file
}

// NewMediaSegment - create empty MediaSegment with CMAF styp box
func NewMediaSegment() *MediaSegment { _ = "STUB: not implemented"; return nil }

// NewMediaSegmentWithStyp - create empty MediaSegment with styp box
func NewMediaSegmentWithStyp(styp *StypBox) *MediaSegment { _ = "STUB: not implemented"; return nil }

// NewMediaSegmentWithoutStyp - create empty media segment with no styp box
func NewMediaSegmentWithoutStyp() *MediaSegment { _ = "STUB: not implemented"; return nil }

// AddSidx adds a sidx box to the MediaSegment.
func (s *MediaSegment) AddSidx(sidx *SidxBox) { _ = "STUB: not implemented"; return }

// AddFragment - Add a fragment to a MediaSegment
func (s *MediaSegment) AddFragment(f *Fragment) { _ = "STUB: not implemented"; return }

// LastFragment returns the currently last fragment, or nil if no fragments.
func (s *MediaSegment) LastFragment() *Fragment { _ = "STUB: not implemented"; return nil }

// ParseSenc parses any deferred senc boxes in all fragments using encryption info from the init segment.
// This is needed when the init segment is in a separate file and was not available during decoding.
func (s *MediaSegment) ParseSenc(init *InitSegment) error { _ = "STUB: not implemented"; return nil }

// Size - return size of media segment
func (s *MediaSegment) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - Write MediaSegment via writer
func (s *MediaSegment) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - Write MediaSegment via SliceWriter
func (s *MediaSegment) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box tree with indent for each level
func (s *MediaSegment) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// Fragmentify - Split into multiple fragments. Assume single mdat and trun for now
func (s *MediaSegment) Fragmentify(timescale uint64, trex *TrexBox, duration uint32) ([]*Fragment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//of.AddFullSample(s)

// fmt.Printf("Wrote fragment with duration %d\n", cumDur)

// CommonSampleDuration returns a common non-zero sample duration for a track defined by trex if available.
func (s *MediaSegment) CommonSampleDuration(trex *TrexBox) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// FirstBox returns the first box in the segment, or an error if no boxes are found.
func (s *MediaSegment) FirstBox() (Box, error) { _ = "STUB: not implemented"; return *new(Box), nil }
