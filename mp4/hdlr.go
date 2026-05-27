package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// HdlrBox - Handler Reference Box (hdlr - mandatory)
//
// Contained in: Media Box (mdia) or Meta Box (meta)
//
// This box describes the type of data contained in the trak.
// Most common hnadler types are: "vide" (video track), "soun" (audio track), "subt" (subtitle track),
// "text" (text track). "meta" (timed Metadata track), clcp (Closed Captions (QuickTime))
type HdlrBox struct {
	Version              byte
	Flags                uint32
	PreDefined           uint32
	HandlerType          string
	Name                 string // Null-terminated UTF-8 string according to ISO/IEC 14496-12 Sec. 8.4.3.3
	LacksNullTermination bool   // This should be false, but we allow true as well
}

// CreateHdlr - create mediaType-specific hdlr box
func CreateHdlr(mediaOrHdlrType string) (*HdlrBox, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DecodeHdlr - box-specific decode
func DecodeHdlr(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeHdlrSR - box-specific decode
func DecodeHdlrSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// 12 bytes of zero

// Type - box type
func (b *HdlrBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *HdlrBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *HdlrBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *HdlrBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *HdlrBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
