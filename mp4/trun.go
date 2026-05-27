package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// TrunBox - Track Fragment Run Box (trun)
//
// Contained in :  Track Fragment Box (traf)
type TrunBox struct {
	Version          byte
	Flags            uint32
	DataOffset       int32
	firstSampleFlags uint32 // interpreted same way as SampleFlags
	Samples          []Sample
	writeOrderNr     uint32 // Used for multi trun offsets
}

const TrunDataOffsetPresentFlag uint32 = 0x01
const TrunFirstSampleFlagsPresentFlag uint32 = 0x04
const TrunSampleDurationPresentFlag uint32 = 0x100
const TrunSampleSizePresentFlag uint32 = 0x200
const TrunSampleFlagsPresentFlag uint32 = 0x400
const TrunSampleCompositionTimeOffsetPresentFlag uint32 = 0x800

// DecodeTrun - box-specific decode
func DecodeTrun(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeTrun - box-specific decode
func DecodeTrunSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// CreateTrun - create a TrunBox for filling up with samples.
// writeOrderNr is only used for multi-trun offsets.
func CreateTrun(writeOrderNr uint32) *TrunBox { _ = "STUB: not implemented"; return nil }

// Signed composition_time_offset
// Data offset and all sample data present

// AddSampleDefaultValues - add values from tfhd and trex boxes if needed
// Return total duration
func (t *TrunBox) AddSampleDefaultValues(tfhd *TfhdBox, trex *TrexBox) (totalDur uint64) {
	_ = "STUB: not implemented"
	return 0
}

// FirstSampleFlags - return firstSampleFlags and indicator if present
func (t *TrunBox) FirstSampleFlags() (flags uint32, present bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetFirstSampleFlags - set firstSampleFlags and bit indicating its presence
func (t *TrunBox) SetFirstSampleFlags(flags uint32) { _ = "STUB: not implemented"; return }

// RemoveFirstSampleFlags - remove firstSampleFlags and its indicator
func (t *TrunBox) RemoveFirstSampleFlags() { _ = "STUB: not implemented"; return }

// SampleCount - return how many samples are defined
func (t *TrunBox) SampleCount() uint32 { _ = "STUB: not implemented"; return 0 }

// HasDataOffset - interpreted dataOffsetPresent flag
func (t *TrunBox) HasDataOffset() bool { _ = "STUB: not implemented"; return false }

// HasFirstSampleFlags - interpreted firstSampleFlagsPresent flag
func (t *TrunBox) HasFirstSampleFlags() bool { _ = "STUB: not implemented"; return false }

// HasSampleDuration - interpreted sampleDurationPresent flag
func (t *TrunBox) HasSampleDuration() bool { _ = "STUB: not implemented"; return false }

// HasSampleFlags - interpreted sampleFlagsPresent flag
func (t *TrunBox) HasSampleFlags() bool { _ = "STUB: not implemented"; return false }

// HasSampleSize - interpreted sampleSizePresent flag
func (t *TrunBox) HasSampleSize() bool { _ = "STUB: not implemented"; return false }

// HasSampleCompositionTimeOffset - interpreted sampleCompositionTimeOffset flag
func (t *TrunBox) HasSampleCompositionTimeOffset() bool { _ = "STUB: not implemented"; return false }

// Type - return box type
func (t *TrunBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (t *TrunBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// expectedSize - calculate size for a given sample count
func (t *TrunBox) expectedSize(sampleCount uint32) uint64 { _ = "STUB: not implemented"; return 0 }

// flags + entryCount

// Encode - write box to w
func (t *TrunBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (t *TrunBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - specificBoxLevels trun:1 gives details
func (t *TrunBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetFullSamples - get all sample data including accumulated time and binary media data
// offsetInMdat is offset in mdat data (data normally starts 8 or 16 bytes after start of mdat box)
// baseDecodeTime is decodeTime in tfdt in track timescale (timescale in mfhd)
// To fill missing individual values from tfhd and trex defaults, call trun.AddSampleDefaultValues() before this call
func (t *TrunBox) GetFullSamples(offsetInMdat uint32, baseDecodeTime uint64, mdat *MdatBox) []FullSample {
	_ = "STUB: not implemented"
	return nil
}

// GetSamples - get all trun sample data
// To fill missing individual values from tfhd and trex defaults, call AddSampleDefaultValues() before this call
func (t *TrunBox) GetSamples() []Sample {
	_ = "STUB: not implemented"

	// GetSampleRange - get a one-based range of samples
	// To fill missing individual values from tfhd and trex defaults, call AddSampleDefaultValues() before this call
	return nil
}

func (t *TrunBox) GetSampleRange(startSampleNr, endSampleNr uint32) []Sample {
	_ = "STUB: not implemented"
	return nil
}

// GetSampleInterval - get sample interval [startSampleNr, endSampleNr] (1-based and inclusive)
// This includes mdat data (if not lazy), in which case only offsetInMdat is given.
// baseDecodeTime is decodeTime in tfdt in track timescale (timescale from mfhd).
// To fill missing individual values from tfhd and trex defaults, call AddSampleDefaultValues() before this call.
func (t *TrunBox) GetSampleInterval(startSampleNr, endSampleNr uint32, baseDecodeTime uint64,
	mdat *MdatBox, offsetInMdat uint32) (SampleInterval, error) {
	_ = "STUB: not implemented"
	return *new(SampleInterval), nil
}

// AddFullSample - add Sample part of FullSample
func (t *TrunBox) AddFullSample(s *FullSample) { _ = "STUB: not implemented"; return }

// AddSample - add a Sample
func (t *TrunBox) AddSample(s Sample) { _ = "STUB: not implemented"; return }

// AddSamples - add a a slice of Sample
func (t *TrunBox) AddSamples(s []Sample) { _ = "STUB: not implemented"; return }

// Duration returns the total duration of all samples given defaultSampleDuration
func (t *TrunBox) Duration(defaultSampleDuration uint32) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// CommonSampleDuration returns the common duration if all samples have the same duration, otherwise 0.
func (t *TrunBox) CommonSampleDuration(defaultSampleDuration uint32) uint32 {
	_ = "STUB: not implemented"
	return 0
}

// GetSampleNrForRelativeTime - get sample number for exact relative time (calculated from summing durations)
func (t *TrunBox) GetSampleNrForRelativeTime(deltaTime uint64, defaultSampleDuration uint32) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// SizeOfData - size of mediasamples in bytes
func (t *TrunBox) SizeOfData() (totalSize uint64) { _ = "STUB: not implemented"; return 0 }
