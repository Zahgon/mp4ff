package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// SchmBox - Scheme Type Box
type SchmBox struct {
	Version       byte
	Flags         uint32
	SchemeType    string // 4CC represented as uint32
	SchemeVersion uint32
	SchemeURI     string // Absolute null-terminated URL
}

// DecodeSchm - box-specific decode
func DecodeSchm(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSchmSR - box-specific decode
func DecodeSchmSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *SchmBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *SchmBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *SchmBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *SchmBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box info to w
func (b *SchmBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
