package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

const TfhdBaseDataOffsetPresentFlag uint32 = 0x000001
const TfhdSampleDescriptionIndexPresentFlag uint32 = 0x000002
const TfhdDefaultSampleDurationPresentFlag uint32 = 0x000008
const TfhdDefaultSampleSizePresentFlag uint32 = 0x000010
const TfhdDefaultSampleFlagsPresentFlag uint32 = 0x000020
const TfhdDurationIsEmptyFlag uint32 = 0x010000
const TfhdDefaultBaseIsMoofFlag uint32 = 0x020000

// TfhdBox - Track Fragment Header Box (tfhd)
//
// Contained in : Track Fragment box (traf))
type TfhdBox struct {
	Version                byte
	Flags                  uint32
	TrackID                uint32
	BaseDataOffset         uint64
	SampleDescriptionIndex uint32
	DefaultSampleDuration  uint32
	DefaultSampleSize      uint32
	DefaultSampleFlags     uint32
}

// DecodeTfhd - box-specific decode
func DecodeTfhd(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeTfhdSR - box-specific decode
func DecodeTfhdSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// CreateTfhd - Create a new TfdtBox with baseMediaDecodeTime
func CreateTfhd(trackID uint32) *TfhdBox {
	_ = "STUB: not implemented"
	// The only flag set is defaultBaseIsMoof
	return nil
}

// HasBaseDataOffset - interpreted flags value
func (t *TfhdBox) HasBaseDataOffset() bool { _ = "STUB: not implemented"; return false }

// HasSampleDescriptionIndex - interpreted flags value
func (t *TfhdBox) HasSampleDescriptionIndex() bool { _ = "STUB: not implemented"; return false }

// HasDefaultSampleDuration - interpreted flags value
func (t *TfhdBox) HasDefaultSampleDuration() bool { _ = "STUB: not implemented"; return false }

// HasDefaultSampleSize - interpreted flags value
func (t *TfhdBox) HasDefaultSampleSize() bool { _ = "STUB: not implemented"; return false }

// HasDefaultSampleFlags - interpreted flags value
func (t *TfhdBox) HasDefaultSampleFlags() bool { _ = "STUB: not implemented"; return false }

// DurationIsEmpty - interpreted flags value
func (t *TfhdBox) DurationIsEmpty() bool { _ = "STUB: not implemented"; return false }

// DefaultBaseIfMoof - interpreted flags value
func (t *TfhdBox) DefaultBaseIfMoof() bool { _ = "STUB: not implemented"; return false }

// Type - returns box type
func (t *TfhdBox) Type() string {
	_ = "STUB: not implemented"

	// Size - returns calculated size
	return ""
}

func (t *TfhdBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (t *TfhdBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (t *TfhdBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write specific box information
func (t *TfhdBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
