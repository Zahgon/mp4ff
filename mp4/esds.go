package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// EsdsBox as used for MPEG-audio, see ISO 14496-1 Section 7.2.6.6  for DecoderConfigDescriptor
type EsdsBox struct {
	Version byte
	Flags   uint32
	ESDescriptor
}

// CreateEsdsBox - Create an EsdsBox geiven decConfig
func CreateEsdsBox(decConfig []byte) *EsdsBox { _ = "STUB: not implemented"; return nil }

// DecodeEsds - box-specific decode
func DecodeEsds(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeEsdsSR - box-specific decode
func DecodeEsdsSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box type
func (e *EsdsBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (e *EsdsBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (e *EsdsBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (e *EsdsBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (e *EsdsBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
