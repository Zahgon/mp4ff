package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// CoLLBox - Content Light Level Box (coll)
// Can be used for VP9 codec in vp09 box (VisualSampleEntryBox).
// Defined in [WebM Project].
//
// [WebM Project]: https://www.webmproject.org/vp9/mp4/
type CoLLBox struct {
	Version byte
	Flags   uint32
	MaxCLL  uint16 // Maximum Content Light Level
	MaxFALL uint16 // Maximum Frame-Average Light Level
}

// CreateCoLLBox - Create a new CoLLBox with specified values
func CreateCoLLBox(maxCLL, maxFALL uint16) *CoLLBox { _ = "STUB: not implemented"; return nil }

const coLLBoxSize = boxHeaderSize + 4 + 2*2 // Header + version/flags + 2 uint16s

// DecodeCoLL - box-specific decode
func DecodeCoLL(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	// Only allow header size of 8 and correct total box size
	return *new(Box), nil
}

// DecodeCoLLSR - decode box from SliceReader
func DecodeCoLLSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	// Only allow header size of 8 and correct total box size
	return *new(Box), nil
}

// Type - box type
func (b *CoLLBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *CoLLBox) Size() uint64 {
	_ = "STUB: not implemented"

	// Encode - write box to w
	return 0
}

func (b *CoLLBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *CoLLBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *CoLLBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
