package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

/*

Definition according to ISO/IEC 14496-12 Section 8.16.3.2
aligned(8) class SegmentIndexBox extends FullBox(‘sidx’, version, 0) {
	unsigned int(32) reference_ID;
	unsigned int(32) timescale;
	if (version==0) {
		unsigned int(32) earliest_presentation_time;
		unsigned int(32) first_offset;
	} else {
		unsigned int(64) earliest_presentation_time;
		unsigned int(64) first_offset;
	}
	unsigned int(16) reserved = 0;
	unsigned int(16) reference_count;
	for(i=1; i <= reference_count; i++) {
		bit (1)           reference_type;
		unsigned int(31)  referenced_size;
		unsigned int(32)  subsegment_duration;
		bit(1)            starts_with_SAP;
		unsigned int(3)   SAP_type;
		unsigned int(28)  SAP_delta_time;
    }
}
*/

// SidxBox - SegmentIndexBox
type SidxBox struct {
	Version                  byte
	Flags                    uint32
	ReferenceID              uint32
	Timescale                uint32
	EarliestPresentationTime uint64
	// FirstOffset is offset of first media segment relative to AnchorPoint
	FirstOffset uint64
	// AnchorPoint is first byte offset after SidxBox
	AnchorPoint uint64
	SidxRefs    []SidxRef
}

// SidxRef - reference as used inside SidxBox
type SidxRef struct {
	ReferencedSize     uint32
	SubSegmentDuration uint32
	SAPDeltaTime       uint32
	ReferenceType      uint8 // 1-bit
	StartsWithSAP      uint8 // 1-bit
	SAPType            uint8
}

// DecodeSidx - box-specific decode
func DecodeSidx(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSidxSR - box-specific decode
func DecodeSidxSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// CreateSidx - Create a new TfdtBox with baseMediaDecodeTime
func CreateSidx(baseMediaDecodeTime uint64) *SidxBox { _ = "STUB: not implemented"; return nil }

// Type - return box type
func (b *SidxBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *SidxBox) Size() uint64 {
	_ = "STUB: not implemented"
	// Add up all fields depending on version
	return 0
}

// Encode - write box to w
func (b *SidxBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *SidxBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Reserved

// Info - more info for level 1
func (b *SidxBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
