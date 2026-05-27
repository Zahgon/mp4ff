package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/av1"
	"github.com/Eyevinn/mp4ff/bits"
)

type Av1CBox struct {
	av1.CodecConfRec
}

// DecodeAv1C - box-specific decode
func DecodeAv1C(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeAv1CSR - box-specific decode
func DecodeAv1CSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *Av1CBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *Av1CBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *Av1CBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write box to sw
func (b *Av1CBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - box-specific Info
func (b *Av1CBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
