package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// MhaCBox - MPEG-H MHACConfigurationBox
// According to ISO/IEC 23008-3: 2018, Section 20.5.2
type MhaCBox struct {
	MHADecoderConfigRecord MHADecoderConfigurationRecord
}

// MHADecoderConfigurationRecord - MPEG-H MHADecoderConfigurationRecord
// According to ISO/IEC 23008-3: 2018, Section 20.4.2
type MHADecoderConfigurationRecord struct {
	ConfigVersion                  uint8
	MpegH3DAProfileLevelIndication uint8
	ReferenceChannelLayout         uint8
	MpegH3DAConfigLength           uint16
	MpegH3DAConfig                 []byte
}

// DecodeMhaC - box-specific decode
func DecodeMhaC(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeMhaCSR - box-specific decode
func DecodeMhaCSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

func decodeMhacFromData(data []byte) (Box, error) { _ = "STUB: not implemented"; return *new(Box), nil }

// Type - box type
func (b *MhaCBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *MhaCBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *MhaCBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - write box to sw
func (b *MhaCBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *MhaCBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
