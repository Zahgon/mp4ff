package hevc

import (
	"errors"
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// HEVC errors
var (
	ErrLengthSize = errors.New("can only handle 4byte NALU length size")
)

// DecConfRec - HEVCDecoderConfigurationRecord
// Specified in ISO/IEC 14496-15 4't ed 2017 Sec. 8.3.3
type DecConfRec struct {
	ConfigurationVersion             byte
	GeneralProfileSpace              byte
	GeneralTierFlag                  bool
	GeneralProfileIDC                byte
	GeneralProfileCompatibilityFlags uint32
	GeneralConstraintIndicatorFlags  uint64
	GeneralLevelIDC                  byte
	MinSpatialSegmentationIDC        uint16
	ParallellismType                 byte
	ChromaFormatIDC                  byte
	BitDepthLumaMinus8               byte
	BitDepthChromaMinus8             byte
	AvgFrameRate                     uint16
	ConstantFrameRate                byte
	NumTemporalLayers                byte
	TemporalIDNested                 byte
	LengthSizeMinusOne               byte
	NaluArrays                       []NaluArray
}

// NaluArray - HEVC NALU array including complete bit and type
type NaluArray struct {
	completeAndType byte
	Nalus           [][]byte
}

// NewNaluArray - create an HEVC NaluArray
func NewNaluArray(complete bool, naluType NaluType, nalus [][]byte) NaluArray {
	_ = "STUB: not implemented"
	return *new(NaluArray)
}

// NaluType - return NaluType for NaluArray
func (n *NaluArray) NaluType() NaluType { _ = "STUB: not implemented"; return *new(NaluType) }

// Complete - return 0x1 if complete
func (n *NaluArray) Complete() byte { _ = "STUB: not implemented"; return 0 }

// CreateHEVCDecConfRec - extract information from sps and insert vps, sps, pps if includePS set
func CreateHEVCDecConfRec(vpsNalus, spsNalus, ppsNalus [][]byte,
	vpsComplete, spsComplete, ppsComplete, includePS bool) (DecConfRec, error) {
	_ = "STUB: not implemented"
	return *new(DecConfRec), nil
}

// Set as default value
// Set as default value

// Set as default value
// Set as default value
// Set as default value
// Set as default value
// only support 4-byte length
// VPS, SPS, PPS nalus with complete flag

// DecodeHEVCDecConfRec - decode an HEVCDecConfRec
func DecodeHEVCDecConfRec(data []byte) (DecConfRec, error) {
	_ = "STUB: not implemented"
	return *new(DecConfRec), nil
}

// Size - total size in bytes
func (h *DecConfRec) Size() uint64 {
	_ = "STUB: not implemented"
	// Up to and including numArrays
	return 0
}

// complete + nalu type + num nalus

// nal unit length

func (h *DecConfRec) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW- write an HEVCDecConfRec to sw
func (h *DecConfRec) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// GetNalusForType - get all nalus for a specific naluType
func (h *DecConfRec) GetNalusForType(naluType NaluType) [][]byte {
	_ = "STUB: not implemented"
	return nil
}

// AddNaluArrays appends new nalus to HEVCDecConfRec.
func (h *DecConfRec) AddNaluArrays(na []NaluArray) { _ = "STUB: not implemented"; return }
