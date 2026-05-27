package sei

import (
	"github.com/Eyevinn/mp4ff/bits"
)

// PicTimingAvcSEI carries the data of an SEI 1 PicTiming message for AVC.
// The corresponding SEI 1 for HEVC is very different.
type PicTimingAvcSEI struct {
	// CbpDbpDelay is optional and triggered by VUI HRD data
	CbpDbpDelay *CbpDbpDelay `json:"-"`
	// TimeOffsetLength is 5 bits and comes from SPS HRD if present
	TimeOffsetLength uint8        `json:"-"`
	PictStruct       uint8        `json:"pict_struct"`
	Clocks           []ClockTSAvc `json:"clocks"`
}

// CbpDbpDelay carries the optional data on CpbDpbDelay.
// This being set corresponds CpbDpbDelaysPresentFlag = true,
// which in turn is a calculated value if HRD info is present,
// i.e. NalHrdBpPresentFlag or VclHrdBpPresentFlag is set
type CbpDbpDelay struct {
	CpbRemovalDelay uint
	DpbOutputDelay  uint
	// InitialCpbRemovalDelayLengthMinus1 comes from SPS HRD and is 5 bits
	InitialCpbRemovalDelayLengthMinus1 byte
	// CpbRemovalDelayLengthMinus1 comes from SPS HRD and is 5 bits
	CpbRemovalDelayLengthMinus1 byte
	// DpbOutputDelayLengthMinus1 comes from SPS HRD and is 5 bits
	DpbOutputDelayLengthMinus1 byte
}

// DecodePicTimingAvcSEI decodes SEI message 1 TimeCode without HRD parameters.
func DecodePicTimingAvcSEI(sd *SEIData) (SEIMessage, error) {
	_ = "STUB: not implemented"
	return *new(SEIMessage), nil
}

// DecodePicTimingAvcSEIHRD decodes AVC SEI message 1 PicTiming with HRD parameters.
// cbpDbpDelay length fields must be properly set if cbpDbpDelay is not nil.
// The delay values in cbpDbpDelay will then be set by the decoder by reading the bits.
// It is assumed that pict_struct_present_flag is true, so that a 4-bit pict_struct value is present.
func DecodePicTimingAvcSEIHRD(sd *SEIData, cbpDbpDelay *CbpDbpDelay, timeOffsetLen byte) (SEIMessage, error) {
	_ = "STUB: not implemented"
	return *new(SEIMessage), nil
}

// Type returns the SEI payload type.
func (s *PicTimingAvcSEI) Type() uint { _ = "STUB: not implemented"; return 0 }

// Payload returns the SEI raw rbsp payload.
func (s *PicTimingAvcSEI) Payload() []byte { _ = "STUB: not implemented"; return nil }

// String returns string representation of PicTiming SEI1.
func (s *PicTimingAvcSEI) String() string { _ = "STUB: not implemented"; return "" }

// Size is size in bytes of raw SEI message rbsp payload.
func (s *PicTimingAvcSEI) Size() uint { _ = "STUB: not implemented"; return 0 }

// pict_struct

// ClockTSAvc carries a clock time stamp for SEI type 1.
type ClockTSAvc struct {
	CtType             byte // scan type
	NuitFieldBasedFlag bool // misspelled in AVC spec, changed in HEVC to UnitsFieldBasedFlag
	CountingType       byte
	NFrames            byte
	Hours              byte
	Minutes            byte
	Seconds            byte
	ClockTimeStampFlag bool
	FullTimeStampFlag  bool
	SecondsFlag        bool
	MinutesFlag        bool
	HoursFlag          bool
	DiscontinuityFlag  bool
	CntDroppedFlag     bool
	TimeOffsetLength   byte
	TimeOffsetValue    int
}

// String returns time stamp
func (c ClockTSAvc) String() string { _ = "STUB: not implemented"; return "" }

func (c *ClockTSAvc) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// CreatePTClockTS creates a clock timestamp.
func CreateClockTSAvc(timeOffsetLen byte) ClockTSAvc {
	_ = "STUB: not implemented"
	return *new(ClockTSAvc)
}

func DecodeClockTSAvc(br *bits.Reader, timeOffsetLen byte) ClockTSAvc {
	_ = "STUB: not implemented"
	return *new(ClockTSAvc)
}

// 0 progressive, 1 interlaced, 2 unknown, 3 reserved

// NrBits returns size of PTClockTS in bits.
func (c ClockTSAvc) NrBits() int { _ = "STUB: not implemented"; return 0 }

// WriteToSliceWriter writes PTClockTS to slice writer.
func (c ClockTSAvc) WriteToSliceWriter(sw bits.SliceWriter) { _ = "STUB: not implemented"; return }
