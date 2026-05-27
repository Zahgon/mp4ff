package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// TfdtBox - Track Fragment Decode Time (tfdt)
//
// Contained in : Track Fragment box (traf)
type TfdtBox struct {
	Version             byte
	Flags               uint32
	baseMediaDecodeTime uint64
}

// DecodeTfdt - box-specific decode
func DecodeTfdt(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeTfdtSR - box-specific decode
func DecodeTfdtSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// CreateTfdt - Create a new TfdtBox with baseMediaDecodeTime
func CreateTfdt(baseMediaDecodeTime uint64) *TfdtBox { _ = "STUB: not implemented"; return nil }

// BaseMediaDecodeTime is the base media decode time.
func (t *TfdtBox) BaseMediaDecodeTime() uint64 { _ = "STUB: not implemented"; return 0 }

// SetBaseMediaDecodeTime sets base media decode time of TfdtBox.
func (t *TfdtBox) SetBaseMediaDecodeTime(bTime uint64) { _ = "STUB: not implemented"; return }

// Type - return box type
func (t *TfdtBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (t *TfdtBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (t *TfdtBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (t *TfdtBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box info to w
func (t *TfdtBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
