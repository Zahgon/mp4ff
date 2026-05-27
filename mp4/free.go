package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// FreeBox - Free Space Box (free or skip)
type FreeBox struct {
	Name       string
	notDecoded []byte
}

// NewFreeBox creates a new FreeBox with arbitrary data payload.
func NewFreeBox(data []byte) *FreeBox { _ = "STUB: not implemented"; return nil }

// NewSkipBox creates a new SkipBox with arbitrary data payload.
func NewSkipBox(data []byte) *FreeBox { _ = "STUB: not implemented"; return nil }

// Payload returns the payload of the box (everything after the box header)
func (b *FreeBox) Payload() []byte { _ = "STUB: not implemented"; return nil }

// DecodeFree - box-specific decode
func DecodeFree(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeFreeSR - box-specific decode
func DecodeFreeSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box type
func (b *FreeBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *FreeBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *FreeBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *FreeBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *FreeBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
