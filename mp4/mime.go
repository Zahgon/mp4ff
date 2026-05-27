package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// MimeBox - MIME Box as defined in ISO/IEC 14496-12 2020 Section 12.3.3.2
type MimeBox struct {
	Version              byte
	Flags                uint32
	ContentType          string
	LacksZeroTermination bool // Handle non-compliant case as well
}

// DecodeMime - box-specific decode
func DecodeMime(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeMimeSR - box-specific decode
func DecodeMimeSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// zero-termination

// Type - box type
func (b *MimeBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *MimeBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *MimeBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *MimeBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write specific box information
func (b *MimeBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
