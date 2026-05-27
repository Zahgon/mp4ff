package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// SchiBox -  Schema Information Box
type SchiBox struct {
	Tenc     *TencBox
	Children []Box
}

// AddChild - Add a child box
func (b *SchiBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// PIFF TrackEncryptionBox carries the same per-track info as tenc

// piffTencToTenc synthesizes a TencBox from PIFF TrackEncryption data
// (PIFF 1.1 §5.3.3) so the rest of the decryption pipeline can treat piff
// like cenc/cbcs.
func piffTencToTenc(p *PiffTencData) *TencBox { _ = "STUB: not implemented"; return nil }

// DecodeSchi - box-specific decode
func DecodeSchi(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSchiSR - box-specific decode
func DecodeSchiSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box type
func (b *SchiBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *SchiBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren - list of child boxes
func (b *SchiBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Encode - write minf container to w
	return nil
}

func (b *SchiBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write minf container to sw
func (b *SchiBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *SchiBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
