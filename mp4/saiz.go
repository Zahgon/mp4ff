package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// SaizBox - Sample Auxiliary Information Sizes Box (saiz)  (in stbl or traf box)
type SaizBox struct {
	Version               byte
	Flags                 uint32
	AuxInfoType           string // Used for Common Encryption Scheme (4-bytes uint32 according to spec)
	AuxInfoTypeParameter  uint32
	SampleCount           uint32
	SampleInfo            []byte
	DefaultSampleInfoSize byte
}

// DecodeSaiz - box-specific decode
func DecodeSaiz(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSaizSR - box-specific decode
func DecodeSaizSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// NewSaizBox creates a SaizBox with appropriate size allocated.
func NewSaizBox(capacity int) *SaizBox { _ = "STUB: not implemented"; return nil }

// AddSampleInfo adds a sampleinfo info based on parameters provided.
// If no length field, don't update the sample field (typically audio cbcs)
func (b *SaizBox) AddSampleInfo(iv []byte, subsamplePatterns []SubSamplePattern) {
	_ = "STUB: not implemented"
	return
}

// Type - return box type
func (b *SaizBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *SaizBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// expectedSize - calculate size based on flags and sample count
func (b *SaizBox) expectedSize() uint64 { _ = "STUB: not implemented"; return 0 }

// 9 = version + flags(4) + defaultSampleInfoSize(1) + sampleCount(4)

// auxInfoType(4) + auxInfoTypeParameter(4)

// 1 byte per sample info when default size is 0

// Encode - write box to w
func (b *SaizBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *SaizBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write SaizBox details. Get sampleInfo list with level >= 1
func (b *SaizBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
