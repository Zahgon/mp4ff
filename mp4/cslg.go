package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// CslgBox - CompositionToDecodeBox -ISO/IEC 14496-12 2015 Sec. 8.6.1.4
//
// Contained in: Sample Table Box (stbl) or Track Extension Properties Box (trep)
type CslgBox struct {
	Version                      byte
	Flags                        uint32
	CompositionToDTSShift        int64
	LeastDecodeToDisplayDelta    int64
	GreatestDecodeToDisplayDelta int64
	CompositionStartTime         int64
	CompositionEndTime           int64
}

// DecodeCslg - box-specific decode
func DecodeCslg(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeCslgSR - box-specific decode
func DecodeCslgSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box type
func (b *CslgBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *CslgBox) Size() uint64 {
	_ = "STUB: not implemented"
	// full Box + 5 * 4 + version * 5*4
	return 0
}

// Encode - write box to w
func (b *CslgBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *CslgBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - get details with specificBoxLevels cslg:1 or higher
func (b *CslgBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
