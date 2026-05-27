package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// IacbBox - IAMF Configuration Box (iacb)
// Defined in IAMF v1.0 Section 6.2.4
type IacbBox struct {
	ConfigurationVersion byte
	IASequenceData       []byte // Contains the IAMF descriptors (OBUs)
}

// DecodeIacb - box-specific decode
func DecodeIacb(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeIacbSR - box-specific decode
func DecodeIacbSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"

	// Read configurationVersion (1 byte)
	return *new(Box), nil
}

// Read descriptors_size as LEB128

// Read the IA Sequence descriptors data

// Type - return box type
func (b *IacbBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *IacbBox) Size() uint64 {
	_ = "STUB: not implemented"
	// 1 byte for version + LEB128 size + descriptor data
	return 0
}

// Encode - write box to w
func (b *IacbBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *IacbBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box info to w
func (b *IacbBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// Parse and display OBUs
