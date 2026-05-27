package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/avc"
	"github.com/Eyevinn/mp4ff/bits"
)

// AvcCBox - AVCConfigurationBox (ISO/IEC 14496-15 5.4.2.1.2 and 5.3.3.1.2)
// Contains one AVCDecoderConfigurationRecord
type AvcCBox struct {
	avc.DecConfRec
}

// CreateAvcC - Create an avcC box based on SPS and PPS
func CreateAvcC(spsNALUs [][]byte, ppsNALUs [][]byte, includePS bool) (*AvcCBox, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DecodeAvcC - box-specific decode
func DecodeAvcC(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeAvcCSR - box-specific decode
func DecodeAvcCSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (a *AvcCBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (a *AvcCBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (a *AvcCBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (a *AvcCBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (a *AvcCBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
