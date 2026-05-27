package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// TrepBox - Track Extension Properties Box (trep)
// Contained in mvex
type TrepBox struct {
	Version  byte
	Flags    uint32
	TrackID  uint32
	Children []Box
}

// AddChild - Add a child box and update SampleCount
func (b *TrepBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// DecodeTrep - box-specific decode
func DecodeTrep(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeTrepSR - box-specific decode
func DecodeTrepSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

//Note higher startPos below since not simple container

// Type - box-specific type
func (b *TrepBox) Type() string {
	_ = "STUB: not implemented"

	// Size - box-specific type
	return ""
}

func (b *TrepBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - box-specific encode of stsd - not a usual container
func (b *TrepBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW- box-specific encode of stsd - not a usual container
func (b *TrepBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *TrepBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
