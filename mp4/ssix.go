package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

/*
Definition according to ISO/IEC 14496-12 Section 8.16.4.2
aligned(8) class SubsegmentIndexBox extends FullBox('ssix', 0, 0) {
  unsigned int(32) subsegment_count;
  for(i=1; i <= subsegment_count; i++){
    unsigned int(32) range_count;
    for (j=1; j <= range_count; j++) {
	  unsigned int(8) level;
	  unsigned int(24) range_size;
    }
  }
}
*/

// SsixBox - Subsegment Index Box according to ISO/IEC 14496-12 Section 8.16.4.2
type SsixBox struct {
	Version     byte
	Flags       uint32
	SubSegments []SubSegment
}

// SubSegment - subsegment data for SsixBox
type SubSegment struct {
	Ranges []SubSegmentRange
}

// SubSegmentRange - range data for SubSegment
type SubSegmentRange uint32

// Level - return level
func (s SubSegmentRange) Level() uint8 { _ = "STUB: not implemented"; return 0 }

// RangeSize - return range size
func (s SubSegmentRange) RangeSize() uint32 { _ = "STUB: not implemented"; return 0 }

// NewSubSegmentRange - create new SubSegmentRange
func NewSubSegmentRange(level uint8, rangeSize uint32) SubSegmentRange {
	_ = "STUB: not implemented"
	return *new(SubSegmentRange)
}

// DecodeSsix - box-specific decode
func DecodeSsix(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSsixSR - box-specific decode
func DecodeSsixSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *SsixBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *SsixBox) Size() uint64 {
	_ = "STUB: not implemented"
	// Add up all fields depending on version
	return 0
}

// Encode - write box to w
func (b *SsixBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *SsixBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - more info for level 1
func (b *SsixBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
