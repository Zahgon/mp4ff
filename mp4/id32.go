package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// ID32Box - ID3v2Box (ID32)
// Defined in https://mp4ra.org/references#id3v2
//
//	aligned(8) class ID3v2Box extends FullBox('ID32', version=0, 0) {
//	    const bit(1) pad = 0;
//	    unsigned int(5)[3] language; // ISO-639-2/T language code
//	    unsigned int(8) ID3v2data [];
//	}
type ID32Box struct {
	Version   byte
	Flags     uint32
	Language  string // 3-letter ISO-639-2/T language code
	ID3v2Data []byte
}

// DecodeID32 - box-specific decode
func DecodeID32(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeID32SR - box-specific decode
func DecodeID32SR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Read language code (pad bit + 3 x 5-bit characters)
// The language is packed into 16 bits (1 bit pad + 15 bits for 3x5 chars)

// Extract 3 characters, each 5 bits, after skipping the pad bit
// Bits: [pad:1][char1:5][char2:5][char3:5] = 16 bits

// Convert to ASCII (add 0x60 to get lowercase letters)

// Read remaining ID3v2 data
// subtract version/flags and language

// Type - return box type
func (b *ID32Box) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *ID32Box) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *ID32Box) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *ID32Box) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Encode language code
// Convert 3-letter code to 5-bit packed format

// default to "und" (undetermined)

// Pack into 16 bits: [pad:1][char1:5][char2:5][char3:5]

// Write ID3v2 data

// Info - write box-specific information
func (b *ID32Box) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
