package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// SaioBox - Sample Auxiliary Information Offsets Box (saiz) (in stbl or traf box)
type SaioBox struct {
	Version              byte
	Flags                uint32
	AuxInfoType          string // Used for Common Encryption Scheme (4-bytes uint32 according to spec)
	AuxInfoTypeParameter uint32
	Offset               []int64
}

// Return a new SaioBox with one offset to be updated later
func NewSaioBox() *SaioBox { _ = "STUB: not implemented"; return nil }

// SetOffset sets offset for first (and only) entry
func (b *SaioBox) SetOffset(offset int64) { _ = "STUB: not implemented"; return }

// DecodeSaio - box-specific decode
func DecodeSaio(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSaioSR - box-specific decode
func DecodeSaioSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *SaioBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *SaioBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// expectedSize - calculate size for a given entry count
func (b *SaioBox) expectedSize(entryCount uint32) uint64 { _ = "STUB: not implemented"; return 0 }

// 8 = version + flags + entryCount

// 4 for AuxInfoType + 4 for AuxInfoTypeParameter

// 4 bytes per offset for version 0

// 8 bytes per offset for version 1

// Encode - write box to w
func (b *SaioBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *SaioBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write SaioBox details. Get offset list with level >= 1
func (b *SaioBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
