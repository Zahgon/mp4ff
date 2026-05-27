package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// TkhdBox - Track Header Box (tkhd - mandatory)
//
// This box describes the track. Duration is measured in time units (according to the time scale
// defined in the movie header box). Duration is 0 for fragmented files.
//
// Volume (relevant for audio tracks) is a fixed point number (8 bits + 8 bits). Full volume is 1.0.
// Width and Height (relevant for video tracks) are fixed point numbers (16 bits + 16 bits).
// Video pixels are not necessarily square.
type TkhdBox struct {
	Version          byte
	Flags            uint32
	CreationTime     uint64
	ModificationTime uint64
	TrackID          uint32
	Duration         uint64
	Layer            int16
	AlternateGroup   int16 // should be int16
	Volume           Fixed16
	Width, Height    Fixed32
}

// CreateTkhd - create tkhd box with common settings
func CreateTkhd() *TkhdBox { _ = "STUB: not implemented"; return nil }

// Enabled, inMovie, inPreview set
// Typically just have one track

// DecodeTkhd - box-specific decode
func DecodeTkhd(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeTkhdSR - box-specific decode
func DecodeTkhdSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Reserved = 0

// Reserved = 0

// Reserved 8 x 0

// 3x3 matrixdata

// Type - box type
func (b *TkhdBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *TkhdBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *TkhdBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *TkhdBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Reserved

// Reserved

// Reserved

// Reserved
// unity matrix according to 8.3.2.2

// Info - write box-specific information
func (b *TkhdBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// These are Fixed32 values

// CraetionTimeS returns the creation time in seconds since Jan 1, 1970
func (b *TkhdBox) CreationTimeS() int64 { _ = "STUB: not implemented"; return 0 }

// ModificationTimeS returns the modification time in seconds since Jan 1, 1970
func (b *TkhdBox) ModificationTimeS() int64 { _ = "STUB: not implemented"; return 0 }

// SetCreationTimeS sets the creation time from seconds since Jan 1, 1970
func (b *TkhdBox) SetCreationTimeS(unixTimeS int64) { _ = "STUB: not implemented"; return }

// SetModificationTimeS sets the modification time from seconds since Jan 1, 1970
func (b *TkhdBox) SetModificationTimeS(unixTimeS int64) { _ = "STUB: not implemented"; return }
