package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

const (
	colrType                         = "colr"
	ColorTypeOnScreenColors          = "nclx" // on-screen colours acc. to ISO/IEC 14496-12 Sec. 12.1.5.2
	ColorTypeRestrictedICCProfile    = "rICC" // restricted ICC profile acc. to ISO/IEC 14496-12 Sec. 12.1.5.2
	ColorTypeUnrestrictedICCTProfile = "prof" // unrestricted ICC profile acc. to ISO/IEC 14496-12 Sec. 12.1.5.2
	// QuickTimeColorParameters defined in [nclc]
	//
	// [nclc]: https://developer.apple.com/library/archive/technotes/tn2162/_index.html#//apple_ref/doc/uid/DTS40013070-CH1-TNTAG10
	QuickTimeColorParameters = "nclc"
	fullRangeBit             = 0x80
)

// ColrBox is colr box defined in ISO/IEC 14496-12 2021 Sec. 12.1.5.
type ColrBox struct {
	ColorType               string
	ICCProfile              []byte
	ColorPrimaries          uint16
	TransferCharacteristics uint16
	MatrixCoefficients      uint16
	FullRangeFlag           bool
	UnknownPayload          []byte
}

// DecodeColr decodes a ColrBox
func DecodeColr(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeColrSR decodes a ColrBox from a SliceReader
func DecodeColrSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type returns the box type
func (c *ColrBox) Type() string {
	_ = "STUB: not implemented"

	// Size returns the calculated size of the box
	return ""
}

func (c *ColrBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode writes box to w
func (c *ColrBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW writes box to sw
func (c *ColrBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info writes box information
func (c *ColrBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
