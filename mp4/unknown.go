package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// UnknownBox - box that we don't know how to parse
type UnknownBox struct {
	name       string
	size       uint64
	notDecoded []byte
}

// DecodeUnknown - decode an unknown box
func DecodeUnknown(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// CreateUnknownBox creates an unknown box. Set the size to match
// the payload size + header size to get a well-formed box.
func CreateUnknownBox(name string, size uint64, payload []byte) *UnknownBox {
	_ = "STUB: not implemented"
	return nil
}

// DecodeUnknownSR - decode an unknown box
func DecodeUnknownSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *UnknownBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *UnknownBox) Size() uint64 {
	_ = "STUB: not implemented"

	// Payload returns the (non-decoded) payload.
	return 0
}

func (b *UnknownBox) Payload() []byte { _ = "STUB: not implemented"; return nil }

// Encode - write box to w
func (b *UnknownBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *UnknownBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *UnknownBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
