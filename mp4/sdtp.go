package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// SdtpBox - Sample Dependency Box (sdtp - optional)
//
// ISO/IEC 14496-12 Ed. 6 2020 Section 8.6.4
// Contained in Sample Table Box (stbl)
//
// Table to determine whether a sample depends or is depended on by other samples
type SdtpBox struct {
	Version byte
	Flags   uint32
	Entries []SdtpEntry
}

// SdtpEntry (uint8)
//
// ISO/IEC 14496-12 Ed. 6 2020 Section 8.6.4.2
type SdtpEntry uint8

// NewSdtpEntry - make new SdtpEntry from 2-bit parameters
func NewSdtpEntry(isLeading, sampleDependsOn, sampleDependedOn, hasRedundancy uint8) SdtpEntry {
	_ = "STUB: not implemented"
	return *new(SdtpEntry)
}

// IsLeading (bits 0-1)
// 0: Leading unknown
// 1: Has dependency before referenced I-picture (not decodable)
// 2: Not a leading sample
// 3: Has no dependency before referenced I-picture (decodable)
func (entry SdtpEntry) IsLeading() uint8 { _ = "STUB: not implemented"; return 0 }

// SampleDependsOn (bits 2-3)
// 0: Dependency is unknown
// 1: Depends on others (not an I-picture)
// 2: Does not depend on others (I-picture)
// 3: Reservced
func (entry SdtpEntry) SampleDependsOn() uint8 { _ = "STUB: not implemented"; return 0 }

// SampleIsDependedOn (bits 4-5)
// 0: Dependency unknown
// 1: Other samples may depend on this (not disposable)
// 2: No other samples depend on this (disposable)
// 3: Reserved
func (entry SdtpEntry) SampleIsDependedOn() uint8 { _ = "STUB: not implemented"; return 0 }

// SampleHasRedundancy (bits 6-7)
// 0: Redundant coding unknown
// 1: Redundant coding in this sample
// 2: No redundant coding in this sample
// 3: Reserved
func (entry SdtpEntry) SampleHasRedundancy() uint8 { _ = "STUB: not implemented"; return 0 }

// CreateSdtpBox - create a new SdtpBox
func CreateSdtpBox(entries []SdtpEntry) *SdtpBox { _ = "STUB: not implemented"; return nil }

// DecodeSdtp - box-specific decode
func DecodeSdtp(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSdtpSR - box-specific decode
func DecodeSdtpSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Supposed to get count from stsz. Use rest of payload

// Type - return box type
func (b *SdtpBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *SdtpBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *SdtpBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *SdtpBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *SdtpBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
