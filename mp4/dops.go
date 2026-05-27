package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// DopsBox - Opus Specific Box (dOps)
// Following https://opus-codec.org/docs/opus_in_isobmff.html
type DopsBox struct {
	Version              byte
	OutputChannelCount   byte
	PreSkip              uint16
	InputSampleRate      uint32
	OutputGain           int16
	ChannelMappingFamily byte
	StreamCount          byte
	CoupledCount         byte
	ChannelMapping       []byte
}

// DecodeDops - box-specific decode
func DecodeDops(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeDopsSR - box-specific decode
func DecodeDopsSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (d *DopsBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (d *DopsBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Version + OutputChannelCount + PreSkip + InputSampleRate + OutputGain + ChannelMappingFamily

// StreamCount + CoupledCount
// ChannelMapping

// Encode - write box to w
func (d *DopsBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - write box to sw
func (d *DopsBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box info to w
func (d *DopsBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
