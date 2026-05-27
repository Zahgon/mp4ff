package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

/*

subs definition according to ISO/IEC 14496-12 Section 8.7.7.2

aligned(8) class SubSampleInformationBox
    extends FullBox(‘subs’, version, flags) {
    unsigned int(32) entry_count;
	int i,j;
	for (i=0; i < entry_count; i++) {
		unsigned int(32) sample_delta;
		unsigned int(16) subsample_count;
		if (subsample_count > 0) {
			for (j=0; j < subsample_count; j++) {
				if(version == 1) {
					unsigned int(32) subsample_size;
				} else {
					unsigned int(16) subsample_size;
				}
				unsigned int(8) subsample_priority;
				unsigned int(8) discardable;
				unsigned int(32) codec_specific_parameters;
			}
		}
	}
}
*/

// SubsBox - SubSampleInformationBox
type SubsBox struct {
	Version byte
	Flags   uint32
	Entries []SubsEntry
}

// SubsEntry - entry in SubsBox
type SubsEntry struct {
	SampleDelta uint32
	SubSamples  []SubsSample
}

// SubsSample - sample in SubsEntry
type SubsSample struct {
	SubsampleSize           uint32
	CodecSpecificParameters uint32
	SubsamplePriority       uint8
	Discardable             uint8
}

// DecodeSubs - box-specific decode
func DecodeSubs(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSubsSR - box-specific decode
func DecodeSubsSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *SubsBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *SubsBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// FullBox + entry_count

// sample_delta + sub_sample_count
//  4 entries per subsample with different lengths for
// version 0 and 1

// Encode - write box to w
func (b *SubsBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *SubsBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - specificBoxLevels dump:1 gives details
func (b *SubsBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
