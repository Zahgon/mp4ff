package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// MfroBox - Movie Fragment Random Access Offset Box (mfro)
// Contained in : MfraBox (mfra)
type MfroBox struct {
	Version    byte
	Flags      uint32
	ParentSize uint32
}

// TryDecodeMfro only decode an MfroBox and return it.
// If it is not an MfroBox, it returns nil and an error.
func TryDecodeMfro(startPos uint64, r io.Reader) (*MfroBox, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DecodeMfro - box-specific decode
func DecodeMfro(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeMfroSR - box-specific decode
func DecodeMfroSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *MfroBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *MfroBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *MfroBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *MfroBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *MfroBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
