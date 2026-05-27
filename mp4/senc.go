package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// UseSubSampleEncryption - flag for subsample encryption
const UseSubSampleEncryption = 0x2

// SubSamplePattern - pattern of subsample encryption
type SubSamplePattern struct {
	BytesOfClearData     uint16
	BytesOfProtectedData uint32
}

// InitializationVector (8 or 16 bytes)
type InitializationVector []byte

// SencBox - Sample Encryption Box (senc) (in trak or traf box)
// Should only be decoded after saio and saiz provide relevant offset and sizes
// Here we make a two-step decode, with first step reading, and other parsing.
// See ISO/IEC 23001-7 Section 7.2 and CMAF specification
// Full Box + SampleCount
type SencBox struct {
	Version          byte
	readButNotParsed bool
	isParsedByGuess  bool // true if parsed by heuristic, can be re-parsed with authoritative value
	perSampleIVSize  byte
	Flags            uint32
	SampleCount      uint32
	StartPos         uint64
	rawData          []byte                 // intermediate storage when reading
	IVs              []InitializationVector // 8 or 16 bytes if present
	SubSamples       [][]SubSamplePattern
	readBoxSize      uint64 // As read from box header
}

// CreateSencBox - create an empty SencBox
func CreateSencBox() *SencBox {
	_ = "STUB: not implemented"

	// NewSencBox returns a SencBox with capacity for IVs and SubSamples.
	return nil
}

func NewSencBox(ivCapacity, subSampleCapacity int) *SencBox { _ = "STUB: not implemented"; return nil }

// SencSample - sample in SencBox
type SencSample struct {
	IV         InitializationVector // 0,8,16 byte length
	SubSamples []SubSamplePattern
}

// AddSample - add a senc sample with possible IV and subsamples
func (s *SencBox) AddSample(sample SencSample) error { _ = "STUB: not implemented"; return nil }

// SetPerSampleIVSize sets the per-sample IV size. Should be 0, 8 or 16.
func (s *SencBox) SetPerSampleIVSize(size byte) { _ = "STUB: not implemented"; return }

// PerSampleIVSize returns the per-sample IV size, 0 if not known yet.
// This will be automatically determined when parsing the box, or when
// adding samples. It can also be set explicitly.
func (s *SencBox) PerSampleIVSize() byte { _ = "STUB: not implemented"; return 0 }

// ReadButNotParsed returns true if box has been read but not parsed.
// The parsing happens as a second step after perSampleIVSize is known.
// ParseReadBox should be called to parse the box.
func (s *SencBox) ReadButNotParsed() bool { _ = "STUB: not implemented"; return false }

// IsParsedByGuess returns true if the box was parsed using a heuristic
// rather than an authoritative perSampleIVSize from tenc or seig.
// Such a result can be replaced by calling ParseReadBox with a known value.
func (s *SencBox) IsParsedByGuess() bool { _ = "STUB: not implemented"; return false }

// DecodeSenc - box-specific decode
func DecodeSenc(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSencSR - box-specific decode
func DecodeSencSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// After the first 8 bytes of box content

// ParseReadBox parses a previously read senc box.
// perSampleIVSize should be known from seig sample group or tenc box.
// If perSampleIVSize is 0, a heuristic using saiz sample sizes is attempted.
// A heuristic result can be replaced later by calling this method again with
// an authoritative perSampleIVSize from a tenc box.
func (s *SencBox) ParseReadBox(perSampleIVSize byte, saiz *SaizBox) error {
	_ = "STUB: not implemented"
	return nil
}

// Already parsed by heuristic, no better info available.

// Reset parsed state for re-parsing

// No subsamples
// Infer the size

// Nothing to do

// With subsamples and known perSampleIVSize

// With subsamples and unknown perSampleIVSize, try to find a valid size.
// Use saiz sample sizes to quickly reject invalid candidates.

// saizMatchesIVSize checks whether a candidate perSampleIVSize is consistent
// with the sample auxiliary information sizes in the saiz box.
// Each saiz entry should equal ivSize + 2 (subsample count) + N*6 (subsample entries)
// for some non-negative integer N.
func saizMatchesIVSize(saiz *SaizBox, sampleCount uint32, ivSize byte) bool {
	_ = "STUB: not implemented"
	return false
}

// unprotected sample

// parseAndFillSamples - parse and fill senc samples given perSampleIVSize
func (s *SencBox) parseAndFillSamples(sr bits.SliceReader, perSampleIVSize byte) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

// Cleanup the IVs and SubSamples which may have been partially set

// Type - box-specific type
func (s *SencBox) Type() string {
	_ = "STUB: not implemented"

	// setSubSamplesUsedFlag - set flag if subsamples are used
	return ""
}

func (s *SencBox) setSubSamplesUsedFlag() { _ = "STUB: not implemented"; return }

// Size - box-specific type
func (s *SencBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *SencBox) calcSize() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (s *SencBox) Encode(w io.Writer) error {
	_ = "STUB: not implemented"
	// First check if subsamplencryption is to be used since it influences the box size
	return nil
}

// EncodeSW - box-specific encode to slicewriter
func (s *SencBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// EncodeSWNoHdr encodes without header (useful for PIFF box)
func (s *SencBox) EncodeSWNoHdr(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (s *SencBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPerSampleIVSize - return perSampleIVSize
func (s *SencBox) GetPerSampleIVSize() int { _ = "STUB: not implemented"; return 0 }
