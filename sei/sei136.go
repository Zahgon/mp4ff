package sei

import (
	"github.com/Eyevinn/mp4ff/bits"
)

// TimeCodeSEI carries the data of an SEI 136 TimeCode message.
type TimeCodeSEI struct {
	Clocks []ClockTS
}

// ClockTS carries a clock time stamp.
type ClockTS struct {
	TimeOffsetValue     uint32
	NFrames             uint16
	Hours               byte
	Minutes             byte
	Seconds             byte
	ClockTimeStampFlag  bool
	UnitsFieldBasedFlag bool
	FullTimeStampFlag   bool
	SecondsFlag         bool
	MinutesFlag         bool
	HoursFlag           bool
	DiscontinuityFlag   bool
	CntDroppedFlag      bool
	CountingType        byte
	TimeOffsetLength    byte
}

// String returns time stamp
func (c ClockTS) String() string { _ = "STUB: not implemented"; return "" }

// CreateClockTS creates a clock timestamp with time parts set to -1.
func CreateClockTS() ClockTS { _ = "STUB: not implemented"; return *new(ClockTS) }

func DecodeClockTS(br *bits.Reader) ClockTS { _ = "STUB: not implemented"; return *new(ClockTS) }

// DecodeTimeCodeSEI decodes SEI message 136 TimeCode.
func DecodeTimeCodeSEI(sd *SEIData) (SEIMessage, error) {
	_ = "STUB: not implemented"
	return *new(SEIMessage), nil
}

// Type returns the SEI payload type.
func (s *TimeCodeSEI) Type() uint { _ = "STUB: not implemented"; return 0 }

// Payload returns the SEI raw rbsp payload.
func (s *TimeCodeSEI) Payload() []byte { _ = "STUB: not implemented"; return nil }

// Final 1 and then byte align

// String returns string representation of TimeCodeSEI.
func (s *TimeCodeSEI) String() string { _ = "STUB: not implemented"; return "" }

// Size is size in bytes of raw SEI message rbsp payload.
func (s *TimeCodeSEI) Size() uint { _ = "STUB: not implemented"; return 0 }
