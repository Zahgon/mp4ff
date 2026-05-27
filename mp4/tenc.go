package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// TencBox - Track Encryption Box
// Defined in ISO/IEC 23001-7 Section 8.2
type TencBox struct {
	Version                byte
	Flags                  uint32
	DefaultCryptByteBlock  byte
	DefaultSkipByteBlock   byte
	DefaultIsProtected     byte
	DefaultPerSampleIVSize byte
	DefaultKID             UUID
	// DefaultConstantIVSize  byte given by len(DefaultConstantIV)
	DefaultConstantIV []byte
}

// DecodeTenc - box-specific decode
func DecodeTenc(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeTencSR - box-specific decode
func DecodeTencSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Skip reserved == 0

// Skip reserved == 0

// Type - return box type
func (b *TencBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *TencBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *TencBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *TencBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// reserved

// reserved

// Info - write box info to w
func (b *TencBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
