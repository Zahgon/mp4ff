package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// MvhdBox - Movie Header Box (mvhd - mandatory)
//
// Contained in : Movie Box (‘moov’)
//
// Contains all media information (duration, ...).
//
// Duration is measured in "time units", and timescale defines the number of time units per second.
type MvhdBox struct {
	Version          byte
	Flags            uint32
	CreationTime     uint64 // Seconds since 1904-01-01
	ModificationTime uint64 // Seconds since 1904-01-01
	Timescale        uint32
	Duration         uint64
	NextTrackID      uint32
	Rate             Fixed32
	Volume           Fixed16
}

// EpochDiffS is the difference in seconds between Jan 1, 1904 and Jan 1, 1970
const EpochDiffS = int64((66*365 + 17) * 24 * 3600)

// CreateMvhd - create mvhd box with reasonable values
func CreateMvhd() *MvhdBox { _ = "STUB: not implemented"; return nil }

// Irrelevant since mdhd timescale is used
// There will typically only be one track
// This is 1.0
// Full volume

// DecodeMvhd - box-specific decode
func DecodeMvhd(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeMvhdSR - box-specific decode
func DecodeMvhdSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Reserved bytes
// Matrix patterndata
// Predefined 0

// Type - return box type
func (b *MvhdBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *MvhdBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Full header + variable part + fixed part

// Full header + variable part + fixed part

// Encode - write box to w
func (b *MvhdBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *MvhdBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Reserved bytes
// unity matrix according to 8.2.2.2
// Predefined 0

// Info - write box-specific information
func (b *MvhdBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreationTimeS returns the creation time in seconds since Jan 1, 1970
func (b *MvhdBox) CreationTimeS() int64 { _ = "STUB: not implemented"; return 0 }

// ModificationTimeS returns the modification time in seconds since Jan 1, 1970
func (b *MvhdBox) ModificationTimeS() int64 { _ = "STUB: not implemented"; return 0 }

// SetCreationTimeS sets the creation time from seconds since Jan 1, 1970
func (b *MvhdBox) SetCreationTimeS(unixTimeS int64) { _ = "STUB: not implemented"; return }

// SetModificationTimeS sets the modification time from seconds since Jan 1, 1970
func (b *MvhdBox) SetModificationTimeS(unixTimeS int64) { _ = "STUB: not implemented"; return }

// Make time string from t which is seconds since Jan. 1 1904
func timeStr(t uint64) string { _ = "STUB: not implemented"; return "" }
