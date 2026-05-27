package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// SampleGroupEntry - like a box, but size and type are not in a header
type SampleGroupEntry interface {
	// Type - GroupingType SampleGroupEntry (uint32 according to spec)
	Type() string // actually
	// Size of SampleGroup Entry
	Size() uint64
	// Encode SampleGroupEntry to SliceWriter
	Encode(sw bits.SliceWriter)
	// Info - description of content.
	Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error)
}

// SampleGroupEntryDecoder is function signature of the SampleGroupEntry Decode method
type SampleGroupEntryDecoder func(name string, length uint32, sr bits.SliceReader) (SampleGroupEntry, error)

var sgeDecoders map[string]SampleGroupEntryDecoder

func init() {
	sgeDecoders = map[string]SampleGroupEntryDecoder{
		"seig": DecodeSeigSampleGroupEntry,
		"roll": DecodeRollSampleGroupEntry,
		"rap ": DecodeRapSampleGroupEntry,
		"alst": DecodeAlstSampleGroupEntry,
	}
}

func decodeSampleGroupEntry(name string, length uint32, sr bits.SliceReader) (SampleGroupEntry, error) {
	_ = "STUB: not implemented"
	return *new(SampleGroupEntry), nil
}

// SeigSampleGroupEntry - CencSampleEncryptionInformationGroupEntry as defined in
// CEF ISO/IEC 23001-7 3rd edition 2016
type SeigSampleGroupEntry struct {
	CryptByteBlock  byte
	SkipByteBlock   byte
	IsProtected     byte
	PerSampleIVSize byte
	KID             UUID
	// ConstantIVSize byte given by len(ConstantIV)
	ConstantIV []byte
}

// DecodeSeigSampleGroupEntry - decode Common Encryption Sample Group Entry
func DecodeSeigSampleGroupEntry(name string, length uint32, sr bits.SliceReader) (SampleGroupEntry, error) {
	_ = "STUB: not implemented"
	return *new(SampleGroupEntry), nil
}

// Reserved

// ConstantIVSize - non-zero if protected and perSampleIVSize == 0
func (s *SeigSampleGroupEntry) ConstantIVSize() byte { _ = "STUB: not implemented"; return 0 }

// Type - GroupingType SampleGroupEntry (uint32 according to spec)
func (s *SeigSampleGroupEntry) Type() string {
	_ = "STUB: not implemented"

	// Size of SampleGroup Entry
	return ""
}

func (s *SeigSampleGroupEntry) Size() uint64 {
	_ = "STUB: not implemented"
	// reserved: 1
	// cryptByteBlock + SkipByteBlock : 1
	// isProtected: 1
	// perSampleIVSize: 1
	// KID: 16
	return 0
}

// Encode SampleGroupEntry to SliceWriter
func (s *SeigSampleGroupEntry) Encode(sw bits.SliceWriter) {
	_ = "STUB: not implemented"
	// Reserved
	return
}

// Info - write box info to w
func (s *SeigSampleGroupEntry) Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// UnknownSampleGroupEntry - unknown or not implemented SampleGroupEntry
type UnknownSampleGroupEntry struct {
	Name   string
	Length uint32
	Data   []byte
}

// DecodeUnknownSampleGroupEntry - decode an unknown sample group entry
func DecodeUnknownSampleGroupEntry(name string, length uint32, sr bits.SliceReader) (SampleGroupEntry, error) {
	_ = "STUB: not implemented"
	return *new(SampleGroupEntry), nil
}

// Type - GroupingType SampleGroupEntry (uint32 according to spec)
func (s *UnknownSampleGroupEntry) Type() string {
	_ = "STUB: not implemented"

	// Size of SampleGroup Entry
	return ""
}

func (s *UnknownSampleGroupEntry) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode SampleGroupEntry to SliceWriter
func (s *UnknownSampleGroupEntry) Encode(sw bits.SliceWriter) { _ = "STUB: not implemented"; return }

// Info - write box info to w
func (s *UnknownSampleGroupEntry) Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// RollSampleGroupEntry - Gradual Decoding Refresh "roll"
//
// ISO/IEC 14496-12 Ed. 6 2020 Section 10.1
//
// VisualRollRecoveryEntry / AudioRollRecoveryEntry / AudioPreRollEntry
type RollSampleGroupEntry struct {
	RollDistance int16
}

// DecodeRollSampleGroupEntry - decode Roll Sample Group Entry
func DecodeRollSampleGroupEntry(name string, length uint32, sr bits.SliceReader) (SampleGroupEntry, error) {
	_ = "STUB: not implemented"
	return *new(SampleGroupEntry), nil
}

// Type - GroupingType SampleGroupEntry (uint32 according to spec)
func (s *RollSampleGroupEntry) Type() string {
	_ = "STUB: not implemented"

	// Size of sample group entry
	return ""
}

func (s *RollSampleGroupEntry) Size() uint64 {
	_ = "STUB: not implemented"

	// Encode SampleGroupEntry to SliceWriter
	return 0
}

func (s *RollSampleGroupEntry) Encode(sw bits.SliceWriter) { _ = "STUB: not implemented"; return }

// Info - write box info to w
func (s *RollSampleGroupEntry) Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// RapSampleGroupEntry - Random Access Point "rap "
//
// ISO/IEC 14496-12 Ed. 6 2020 Section 10.4 - VisualRandomAccessEntry
type RapSampleGroupEntry struct {
	NumLeadingSamplesKnown uint8
	NumLeadingSamples      uint8
}

// DecodeRapSampleGroupEntry - decode Rap Sample Sample Group Entry
func DecodeRapSampleGroupEntry(name string, length uint32, sr bits.SliceReader) (SampleGroupEntry, error) {
	_ = "STUB: not implemented"
	return *new(SampleGroupEntry), nil
}

// Type - GroupingType SampleGroupEntry (uint32 according to spec)
func (s *RapSampleGroupEntry) Type() string {
	_ = "STUB: not implemented"

	// Size of sample group entry
	return ""
}

func (s *RapSampleGroupEntry) Size() uint64 {
	_ = "STUB: not implemented"

	// Encode SampleGroupEntry to SliceWriter
	return 0
}

func (s *RapSampleGroupEntry) Encode(sw bits.SliceWriter) { _ = "STUB: not implemented"; return }

// Info - write box info to w
func (s *RapSampleGroupEntry) Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// AlstSampleGroupEntry - Alternative Startup Entry "alst"
//
// ISO/IEC 14496-12 Ed. 6 2020 Section 10.3 - AlternativeStartupEntry
type AlstSampleGroupEntry struct {
	RollCount         uint16
	FirstOutputSample uint16
	SampleOffset      []uint32
	NumOutputSamples  []uint16
	NumTotalSamples   []uint16
}

// Type - GroupingType SampleGroupEntry (uint32 according to spec)
func (s *AlstSampleGroupEntry) Type() string {
	_ = "STUB: not implemented"

	// Size of sample group entry
	return ""
}

func (s *AlstSampleGroupEntry) Size() uint64 {
	_ = "STUB: not implemented"
	// RollCount: 2
	// FirstOutputSample: 2
	// SampleOffset: 4 * count
	// NumOutputSamples: 2 * count
	// NumTotalSamples: 2 * count
	return 0
}

// DecodeAlstSampleGroupEntry - decode ALST Sample Group Entry
func DecodeAlstSampleGroupEntry(name string, length uint32, sr bits.SliceReader) (SampleGroupEntry, error) {
	_ = "STUB: not implemented"
	return *new(SampleGroupEntry), nil
}

// Optional

// Encode SampleGroupEntry to SliceWriter
func (s *AlstSampleGroupEntry) Encode(sw bits.SliceWriter) { _ = "STUB: not implemented"; return }

// Info - write box info to w
func (s *AlstSampleGroupEntry) Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
