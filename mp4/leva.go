package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

/*
Definition according to ISO/IEC 14496-12 Section 8.8.13.2
aligned(8) class LevelAssignmentBox extends FullBox('leva', 0, 0) {
  unsigned int(8) level_count;
  for (j=1; j <= level_count; j++) {
    unsigned int(32) track_ID;
    unsigned int(1) padding_flag;
    unsigned int(7) assignment_type;
    if (assignment_type == 0) {
      unsigned int(32) grouping_type;
    } else if (assignment_type == 1) {
      unsigned int(32) grouping_type;
      unsigned int(32) grouping_type_parameter;
    } else if (assignment_type == 2) {
      // no further syntax elements needed
    } else if (assignment_type == 3) {
      // no further syntax elements needed
	} else if (assignment_type == 4) {
      unsigned int(32) sub_track_ID;
    }
    // other assignment_type values are reserved
  }
}
*/

// LevaBox - Subsegment Index Box according to ISO/IEC 14496-12 Section 8.8.13.2.
type LevaBox struct {
	Version byte
	Flags   uint32
	Levels  []LevaLevel
}

// LevaLevel - level data for LevaBox
type LevaLevel struct {
	TrackID                  uint32
	GroupingType             uint32
	GroupingTypeParameter    uint32
	SubTrackID               uint32
	paddingAndAssignmentType byte
}

// PaddingFlag - return padding flag.
func (l LevaLevel) PaddingFlag() bool { _ = "STUB: not implemented"; return false }

// AssignmentType - return assignment type.
func (l LevaLevel) AssignmentType() byte { _ = "STUB: not implemented"; return 0 }

// Size - return calculated size.
func (l LevaLevel) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// NewLevaLevel - create new level for LevaBox.
func NewLevaLevel(trackID uint32, paddingFlag bool, assignmentType byte,
	groupingType, groupingTypeParameter, subTrackID uint32) (LevaLevel, error) {
	_ = "STUB: not implemented"
	return *new(LevaLevel), nil
}

// DecodeLeva - box-specific decode
func DecodeLeva(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeLevaSR - box-specific decode
func DecodeLevaSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *LevaBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *LevaBox) Size() uint64 {
	_ = "STUB: not implemented"
	// Add up all fields depending on version
	return 0
}

// Encode - write box to w
func (b *LevaBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *LevaBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - more info for level 1
func (b *LevaBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
