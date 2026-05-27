package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// Av3cBox - AVS3 Configuration Box (av3c)
// Defined in AVS3-P6-TAI 109.6-2022-en.pdf Section 5.2.2.3
type Av3cBox struct {
	Avs3Config Avs3DecoderConfigurationRecord
}

// DecodeAv3c - box-specific decode
func DecodeAv3c(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeAv3cSR - box-specific decode
func DecodeAv3cSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	// 8-byte header + minimum 4 bytes for config
	return *new(Box), nil
}

// Check that 6 most significant bits are 111111 (0x3F << 2 = 0xFC)

// Extract 2 least significant bits

// Type - box type
func (b *Av3cBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *Av3cBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *Av3cBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *Av3cBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Set 6 most significant bits to 111111 (0xFC) and OR with 2-bit LibraryDependencyIDC

// Info - write box info to w
func (b *Av3cBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// Avs3DecoderConfigurationRecord - AVS3 Decoder Configuration Record
// Defined in AVS3-P6-TAI 109.6-2022-en.pdf Section 5.2.2.1
type Avs3DecoderConfigurationRecord struct {
	ConfigurationVersion uint8
	SequenceHeaderLength uint16
	SequenceHeader       []byte
	LibraryDependencyIDC uint8 // 2 bits
}
