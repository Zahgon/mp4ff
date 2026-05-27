package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
	"github.com/Eyevinn/mp4ff/hevc"
)

// HvcCBox - HEVCConfigurationBox (ISO/IEC 14496-15 8.4.1.1.2)
// Contains one HEVCDecoderConfigurationRecord
type HvcCBox struct {
	hevc.DecConfRec
}

// CreateHvcC - create an hvcC box based on VPS, SPS and PPS and signal completeness
// If includePS is false, the nalus are not included, but information from sps is extracted.
func CreateHvcC(vpsNalus, spsNalus, ppsNalus [][]byte, vpsComplete, spsComplete, ppsComplete, includePS bool) (*HvcCBox, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DecodeHvcC - box-specific decode
func DecodeHvcC(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeHvcCSR - box-specific decode
func DecodeHvcCSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *HvcCBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *HvcCBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *HvcCBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write box to w
func (b *HvcCBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - box-specific Info
func (b *HvcCBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
