package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// Boxes needed for wvtt according to ISO/IEC 14496-30

////////////////////////////// wvtt //////////////////////////////

// WvttBox - WVTTSampleEntry (wvtt)
// Extends PlainTextSampleEntry which extends SampleEntry
type WvttBox struct {
	VttC               *VttCBox
	Vlab               *VlabBox
	Btrt               *BtrtBox
	Children           []Box
	DataReferenceIndex uint16
}

// NewWvttBox - Create new empty wvtt box
func NewWvttBox() *WvttBox { _ = "STUB: not implemented"; return nil }

// AddChild - add a child box
func (b *WvttBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// Other box

const nrWvttBytesBeforeChildren = 16

// DecodeWvtt - Decoder wvtt Sample Entry (wvtt)
func DecodeWvtt(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeWvttSR - Decoder wvtt Sample Entry (wvtt)
func DecodeWvttSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"

	// 14496-12 8.5.2.2 Sample entry (8 bytes)
	return *new(Box), nil
}

// Skip 6 reserved bytes

// Type - return box type
func (b *WvttBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *WvttBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *WvttBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Only write written bytes

// Next output child boxes in order

// EncodeSW - write box to w
func (b *WvttBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Next output child boxes in order

// Info - write box-specific information
func (b *WvttBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

////////////////////////////// vttC //////////////////////////////

// VttCBox - WebVTTConfigurationBox (vttC)
type VttCBox struct {
	Config string
}

// DecodeVttC - box-specific decode
func DecodeVttC(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeVttCSR - box-specific decode
func DecodeVttCSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *VttCBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *VttCBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *VttCBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *VttCBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *VttCBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

////////////////////////////// vlab //////////////////////////////

// VlabBox - WebVTTSourceLabelBox (vlab)
type VlabBox struct {
	SourceLabel string
}

// DecodeVlab - box-specific decode
func DecodeVlab(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeVlabSR - box-specific decode
func DecodeVlabSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *VlabBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *VlabBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *VlabBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *VlabBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *VlabBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// wvtt Sample boxes
// A sample is either one vtte box or one or more vttc or vta boxes

////////////////////////////// vtte //////////////////////////////

// VtteBox - VTTEmptyBox (vtte)
type VtteBox struct {
}

// Type - box-specific type
func (b *VtteBox) Type() string {
	_ = "STUB: not implemented"

	// DecodeVtte - box-specific decode
	return ""
}

func DecodeVtte(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *

	// DecodeVtteSR - box-specific decode
	new(Box), nil
}

func DecodeVtteSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *

	// Size - calculated size of box
	new(Box), nil
}

func (b *VtteBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *VtteBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *VtteBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *VtteBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

////////////////////////////// vttc //////////////////////////////

// VttcBox - VTTCueBox (vttc)
type VttcBox struct {
	Vsid     *VsidBox
	Iden     *IdenBox
	Ctim     *CtimBox
	Sttg     *SttgBox
	Payl     *PaylBox
	Children []Box
}

// AddChild - Add a child box
func (b *VttcBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

// Type outside ISO/IEC 14496-30 spec

// DecodeVttc - box-specific decode
func DecodeVttc(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeVttcSR - box-specific decode
func DecodeVttcSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *VttcBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *VttcBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren - list of child boxes
func (b *VttcBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Encode - write mvex container to w
	return nil
}

func (b *VttcBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write vttc container to sw
func (b *VttcBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *VttcBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

////////////////////////////// vsid //////////////////////////////

// VsidBox - CueSourceIDBox (vsid)
type VsidBox struct {
	SourceID uint32
}

// DecodeVsid - box-specific decode
func DecodeVsid(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeVsidSR - box-specific decode
func DecodeVsidSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *VsidBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *VsidBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// len of uint32

// Encode - write box to w
func (b *VsidBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *VsidBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *VsidBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

////////////////////////////// ctim //////////////////////////////

// CtimBox - CueTimeBox (ctim)
// CueCurrentTime is current time indication (for split cues)
type CtimBox struct {
	CueCurrentTime string
}

// DecodeCtim - box-specific decode
func DecodeCtim(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeCtimSR - box-specific decode
func DecodeCtimSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *CtimBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *CtimBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *CtimBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *CtimBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *CtimBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

////////////////////////////// iden //////////////////////////////

// IdenBox - CueIDBox (iden)
type IdenBox struct {
	CueID string
}

// DecodeIden - box-specific decode
func DecodeIden(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeIdenSR - box-specific decode
func DecodeIdenSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *IdenBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *IdenBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *IdenBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *IdenBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *IdenBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

////////////////////////////// sttg //////////////////////////////

// SttgBox - CueSettingsBox (sttg)
type SttgBox struct {
	Settings string
}

// DecodeSttg - box-specific decode
func DecodeSttg(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeSttgSR - box-specific decode
func DecodeSttgSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *SttgBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *SttgBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *SttgBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *SttgBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *SttgBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

////////////////////////////// payl //////////////////////////////

// PaylBox - CuePayloadBox (payl)
type PaylBox struct {
	CueText string
}

// DecodePayl - box-specific decode
func DecodePayl(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodePaylSR - box-specific decode
func DecodePaylSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *PaylBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *PaylBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *PaylBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *PaylBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *PaylBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

////////////////////////////// vtta //////////////////////////////

// VttaBox - VTTAdditionalTextBox (vtta) (corresponds to NOTE in WebVTT)
type VttaBox struct {
	CueAdditionalText string
}

// DecodeVtta - box-specific decode
func DecodeVtta(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeVttaSR - box-specific decode
func DecodeVttaSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box-specific type
func (b *VttaBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *VttaBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *VttaBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *VttaBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *VttaBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
