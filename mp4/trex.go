package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// TrexBox - Track Extends Box
//
// Contained in : Mvex Box (mvex)
type TrexBox struct {
	Version                       byte
	Flags                         uint32
	TrackID                       uint32
	DefaultSampleDescriptionIndex uint32
	DefaultSampleDuration         uint32
	DefaultSampleSize             uint32
	DefaultSampleFlags            uint32
}

// CreateTrex - create trex box with trackID
func CreateTrex(trackID uint32) *TrexBox { _ = "STUB: not implemented"; return nil }

// DecodeTrex - box-specific decode
func DecodeTrex(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeTrexSR - box-specific decode
func DecodeTrexSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// interpreted as SampleFlags

// Type - return box type
func (b *TrexBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *TrexBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *TrexBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *TrexBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *TrexBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
