package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// KindBox - Track Kind Box
type KindBox struct {
	Version   byte
	Flags     uint32
	SchemeURI string
	Value     string
}

// DecodeKind - box-specific decode
func DecodeKind(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeKindSR - box-specific decode
func DecodeKindSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box type
func (b *KindBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *KindBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *KindBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write box to w
func (b *KindBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *KindBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
