package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// ElngBox - Extended Language Box
// Defined in ISO/IEC 14496-12 Section 8.4.6
// It should be a full box, but was erroneously implemented
// as a normal box. For backwards compatibility, the
// erroneous box without full header can still be decoded.
// The method MissingFullBoxBytes() returns true if that is the case.
type ElngBox struct {
	missingFullBox bool
	Version        byte
	Flags          uint32
	Language       string
}

// MissingFullBoxBytes indicates that the box is erroneously not including the 4 full box header bytes
func (b *ElngBox) MissingFullBoxBytes() bool { _ = "STUB: not implemented"; return false }

// CreateElng - Create an Extended Language Box
func CreateElng(language string) *ElngBox { _ = "STUB: not implemented"; return nil }

// DecodeElng - box-specific decode
func DecodeElng(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeElngSR - box-specific decode
func DecodeElngSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Less than 4 byte flag and version + 2 letters + 0 termination

// Type - box type
func (b *ElngBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *ElngBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// FixMissingFullBoxBytes adds missing bytes version and flags bytes.
func (b *ElngBox) FixMissingFullBoxBytes() { _ = "STUB: not implemented"; return }

// Encode - write box to w
func (b *ElngBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *ElngBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *ElngBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
