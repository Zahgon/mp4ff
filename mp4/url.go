package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// URLBox - DataEntryUrlBox ('url ')
//
// Contained in : DrefBox (dref
type URLBox struct {
	Version           byte
	Flags             uint32
	Location          string // Zero-terminated string
	NoLocation        bool
	NoZeroTermination bool
}

const dataIsSelfContainedFlag = 0x000001

// DecodeURLBox - box-specific decode
func DecodeURLBox(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeURLBoxSR - box-specific decode
func DecodeURLBoxSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// CreateURLBox - Create a self-referencing URL box
func CreateURLBox() *URLBox { _ = "STUB: not implemented"; return nil }

// Type - return box type
func (b *URLBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *URLBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *URLBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *URLBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write specific box information
func (b *URLBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
