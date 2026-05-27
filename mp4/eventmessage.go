package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// EvteBox - EventMessageSampleEntry box as defined in ISO/IEC 23001-18 Section 7.2
type EvteBox struct {
	Btrt               *BtrtBox
	Silb               *SilbBox
	Children           []Box
	DataReferenceIndex uint16
}

// DecodeEvte - Decode EventMessageSampleEntry (evte)
func DecodeEvte(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeEvteSR - Decode EventMessageSampleEntry (evte)
func DecodeEvteSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// 14496-12 8.5.2.2 Sample entry (8 bytes)

// Skip 6 reserved bytes

// AddChild - add a child box (should only be btrt and silb)
func (b *EvteBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// Other box

func (b *EvteBox) Type() string { _ = "STUB: not implemented"; return "" }

func (b *EvteBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w via a SliceWriter
func (b *EvteBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - write box to sw
func (b *EvteBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Next output child boxes in order

// Info - write specific box info to w
func (b *EvteBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// SilbBox - Scheme Identifier Box as defined in ISO/IEC 23001-18 Section 7.3
type SilbBox struct {
	Version          uint8
	Flags            uint32
	Schemes          []SilbEntry
	OtherSchemesFlag bool
}

// SilbEntry - Scheme Identifier Box entry
type SilbEntry struct {
	SchemeIdURI    string
	Value          string
	AtLeastOneFlag bool
}

// DecodeSilb - Decode Scheme Identifier Box (silb)
func DecodeSilb(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSilbSR - Decode Scheme Identifier Box (silb)
func DecodeSilbSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

func (b *SilbBox) Type() string { _ = "STUB: not implemented"; return "" }

func (b *SilbBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// OtherSchemesFlag

// Encode - write box to w via a SliceWriter
func (b *SilbBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - write box to sw
func (b *SilbBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write specific box info to w
func (b *SilbBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// EmibBox - EventMessageInstanceBox as defined in ISO/IEC 23001-18 Section 6.1
type EmibBox struct {
	Version               uint8
	Flags                 uint32
	PresentationTimeDelta int64
	EventDuration         uint32
	Id                    uint32
	SchemeIdURI           string
	Value                 string
	MessageData           []byte
}

// DecodeEmib - box-specific decode
func DecodeEmib(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeEmibSR - box-specific decode
func DecodeEmibSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// reserved

func (b *EmibBox) Type() string { _ = "STUB: not implemented"; return "" }

func (b *EmibBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *EmibBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (b *EmibBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// reserved

func (b *EmibBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// EmebBox - EventMessageBox as defined in ISO/IEC 23001-18 Section 6.2
type EmebBox struct {
}

// DecodeEmeb - box-specific decode
func DecodeEmeb(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeEmebSR - box-specific decode
func DecodeEmebSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *EmebBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *EmebBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *EmebBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *EmebBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *EmebBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
