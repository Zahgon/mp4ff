package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

const (
	PrftTimeEncoderInput       = 0
	PrftTimeEncoderOutput      = 1
	PrftTimeMoofFinalized      = 2
	PrftTimeMoofWritten        = 4
	PrftTimeArbitraryConsitent = 8
	PrftTimeCaptured           = 24
)

var PrftFlagsInterpretation = map[uint32]string{
	PrftTimeEncoderInput:       "time_encoder_input",
	PrftTimeEncoderOutput:      "time_encoder_output",
	PrftTimeMoofFinalized:      "time_moof_finalized",
	PrftTimeMoofWritten:        "time_moof_written",
	PrftTimeArbitraryConsitent: "time_arbitrary_consistent",
	PrftTimeCaptured:           "time_captured",
}

// PrftBox - Producer Reference Box (prft)
//
// Contained in File before moof box
type PrftBox struct {
	Version          byte
	Flags            uint32
	ReferenceTrackID uint32
	NTPTimestamp     NTP64
	MediaTime        uint64
}

// CreatePrftBox creates a new PrftBox.
func CreatePrftBox(version byte, flags, refTrackID uint32, ntp NTP64, mediatime uint64) *PrftBox {
	_ = "STUB: not implemented"
	return nil
}

// DecodePrft - box-specific decode
func DecodePrft(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodePrftSR - box-specific decode
func DecodePrftSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *PrftBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *PrftBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *PrftBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *PrftBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *PrftBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// InterpretFlags - return string representation of flags.
func (b *PrftBox) InterpretFlags() string { _ = "STUB: not implemented"; return "" }
