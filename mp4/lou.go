package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// LoudnessBaseBox according to ISO/IEC 14496-12 Section 12.2.7.2
// TrackLoudnessInfo (tlou) and AudioLoudnessInfo (alou) boxes
// are extensions of the LoudnessBaseBox.
type LoudnessBaseBox struct {
	Name          string
	Version       byte
	Flags         uint32
	LoudnessBases []*LoudnessBase
}

// LoudnessBase provides a loudness entry in a LoudnessBaseBox
type LoudnessBase struct {
	EQSetID                uint8
	DownmixID              uint8
	DRCSetID               uint8
	BsSamplePeakLevel      int16
	BsTruePeakLevel        int16
	MeasurementSystemForTP uint8
	ReliabilityForTP       uint8
	Measurements           []LoudnessMeasurement
}

// LoudnessMeasurement provides a loudness measurement in a LoudnessBase.
type LoudnessMeasurement struct {
	MethodDefinition  uint8
	MethodValue       uint8
	MeasurementSystem uint8
	Reliability       uint8
}

// DecodeLoudnessBaseBox - box-specific decode
func DecodeLoudnessBaseBox(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeLoudnessBaseBoxSR - box-specific decode
func DecodeLoudnessBaseBoxSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type of LoundessBaseBox, should be tlou or alou
func (b *LoudnessBaseBox) Type() string {
	_ = "STUB: not implemented"

	// Size of LoudnessBaseBox.
	return ""
}

func (b *LoudnessBaseBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *LoudnessBaseBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (b *LoudnessBaseBox) EncodeSW(sw bits.SliceWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *LoudnessBaseBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
