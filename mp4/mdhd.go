package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

const charOffset = 0x60 // According to Section 8.4.2.3 of 14496-12

// MdhdBox - Media Header Box (mdhd - mandatory)
//
// Contained in : Media Box (mdia)
//
// Timescale defines the timescale used for this track.
// Language is a ISO-639-2/T language code stored as 1bit padding + [3]int5
type MdhdBox struct {
	Version          byte // Only version 0
	Flags            uint32
	CreationTime     uint64 // Seconds since 1904-01-01
	ModificationTime uint64 // Seconds since 1904-01-01
	Timescale        uint32 // Media timescale for this track
	Duration         uint64 // Trak duration, 0 for fragmented files
	Language         uint16 // Three-letter ISO-639-2/T language code
}

// DecodeMdhd - Decode box
func DecodeMdhd(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeMdhd - Decode box
func DecodeMdhdSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// GetLanguage - Get three-byte language string
func (m *MdhdBox) GetLanguage() string { _ = "STUB: not implemented"; return "" }

// SetLanguage - Set three-byte language string
func (m *MdhdBox) SetLanguage(lang string) { _ = "STUB: not implemented"; return }

// Type - box type
func (m *MdhdBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (m *MdhdBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// m.Version = 0

// Encode - write box to w
func (m *MdhdBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (m *MdhdBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (m *MdhdBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreationTimeS returns the creation time in seconds since Jan 1, 1970
func (b *MdhdBox) CreationTimeS() int64 { _ = "STUB: not implemented"; return 0 }

// ModificationTimeS returns the modification time in seconds since Jan 1, 1970
func (b *MdhdBox) ModificationTimeS() int64 { _ = "STUB: not implemented"; return 0 }

// SetCreationTimeS sets the creation time from seconds since Jan 1, 1970
func (b *MdhdBox) SetCreationTimeS(unixTimeS int64) { _ = "STUB: not implemented"; return }

// SetModificationTimeS sets the modification time from seconds since Jan 1, 1970
func (b *MdhdBox) SetModificationTimeS(unixTimeS int64) { _ = "STUB: not implemented"; return }
