package mp4

import (
	"io"
	"time"

	"github.com/Eyevinn/mp4ff/bits"
)

// SttsBox -  Decoding Time to Sample Box (stts - mandatory)
//
// This table contains the duration in time units for each sample.
//
//   - SampleCount : the number of consecutive samples having the same duration
//   - SampleTimeDelta : duration in time units
type SttsBox struct {
	Version         byte
	Flags           uint32
	SampleCount     []uint32
	SampleTimeDelta []uint32
}

// DecodeStts - box-specific decode
func DecodeStts(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSttsSR - box-specific decode
func DecodeSttsSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *SttsBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *SttsBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// expectedSize - calculate size for a given entry count
func (b *SttsBox) expectedSize(entryCount uint32) uint64 {
	_ = "STUB: not implemented"
	// 8 = version + flags(4) + entryCount(4)
	// 8 = sampleCount(4) + sampleTimeDelta(4)
	return 0
}

// GetTimeCode - return the timecode (duration since the beginning of the media)
// of the beginning of a sample
func (b *SttsBox) GetTimeCode(sample, timescale uint32) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// GetDecodeTime - decode time and duration for (one-based) sampleNr in track timescale
func (b *SttsBox) GetDecodeTime(sampleNr uint32) (decTime uint64, dur uint32) {
	_ = "STUB: not implemented"

	// This is bad index input. Should never happen
	return 0, 0
}

// GetDur - get dur for a specific sample
func (b *SttsBox) GetDur(sampleNr uint32) (dur uint32) {
	_ = "STUB: not implemented"

	// This is bad index input. Should never happen
	return 0
}

// one-based -> zero-based

// Encode - write box to w
func (b *SttsBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *SttsBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *SttsBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSampleNrAtTime returns the 1-based sample number at or as soon as possible after time.
// Match a final single zero duration if present.
// If time is too big to reach, an error is returned.
// Time is calculated by summing up durations of previous samples
func (b *SttsBox) GetSampleNrAtTime(sampleStartTime uint64) (sampleNr uint32, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If not exact, increase number to next sample

// Check if there is a final single zero duration and time matches.
