package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// TfraBox - Track Fragment Random Access Box (tfra)
// Contained it MfraBox (mfra)
type TfraBox struct {
	Version               byte
	Flags                 uint32
	TrackID               uint32
	LengthSizeOfTrafNum   byte
	LengthSizeOfTrunNum   byte
	LengthSizeOfSampleNum byte
	Entries               []TfraEntry
}

// TfraEntry - reference as used inside TfraBox
type TfraEntry struct {
	Time         uint64
	MoofOffset   uint64
	TrafNumber   uint32
	TrunNumber   uint32
	SampleNumber uint32
}

// DecodeTfra - box-specific decode
func DecodeTfra(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeTfraSR - box-specific decode
func DecodeTfraSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *TfraBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *TfraBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// expectedSize - calculate size for a given number of entries
func (b *TfraBox) expectedSize(entryCount uint32) uint64 { _ = "STUB: not implemented"; return 0 }

// 16 = version + flags(4) + trackID(4) + sizesBlock(4) + entryCount(4)

// Size per entry depends on version
// Version 0: time(4) + moof_offset(4)

// Version 1: time(8) + moof_offset(8)

// Add size for traf/trun/sample numbers based on their length settings

// 1-4 bytes
// 1-4 bytes
// 1-4 bytes

// Encode - write box to w
func (b *TfraBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *TfraBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - box-specific info. More for level 1
func (b *TfraBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// FindEntry - find entry for moofStart. Return nil if not found
func (b *TfraBox) FindEntry(moofStart uint64) *TfraEntry { _ = "STUB: not implemented"; return nil }
