package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// StppBox - XMLSubtitleSampleEntry Box (stpp)
// Defined in ISO/IEC 14496-12 Sec. 12.6.3.2 and ISO/IEC 14496-30.
//
// Contained in : Media Information Box (minf)
type StppBox struct {
	Namespace                 string   // Mandatory
	SchemaLocation            string   // Optional
	AuxiliaryMimeTypes        string   // Optional, but required if auxiliary types present
	Btrt                      *BtrtBox // Optional
	Children                  []Box
	DataReferenceIndex        uint16
	nrMissingOptionalEndBytes byte // 0, 1, or 2 depending on whether SchemaLocation and AuxiliaryMimeTypes have a zero end byte
}

// NewStppBox - Create new stpp box
// namespace, schemaLocation and auxiliaryMimeType are space-separated utf8-lists with zero-termination
// schemaLocation and auxiliaryMimeTypes are optional but must at least have a zero byte.
func NewStppBox(namespace, schemaLocation, auxiliaryMimeTypes string) *StppBox {
	_ = "STUB: not implemented"
	return nil
}

// AddChild - add a child box (avcC normally, but clap and pasp could be part of visual entry)
func (b *StppBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// Other box

// DecodeStpp - Decode XMLSubtitleSampleEntry (stpp)
func DecodeStpp(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeStppSR - Decode XMLSubtitleSampleEntry (stpp)
func DecodeStppSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// 14496-12 8.5.2.2 Sample entry (8 bytes)

// Skip 6 reserved bytes

// Type - return box type
func (b *StppBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *StppBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w via a SliceWriter
func (b *StppBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - write box to sw
func (b *StppBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Next output child boxes in order

// Info - write specific box info to w
func (b *StppBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
