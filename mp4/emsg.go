package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// EmsgBox - DASHEventMessageBox as defined in ISO/IEC 23009-1
type EmsgBox struct {
	Version               byte
	Flags                 uint32
	TimeScale             uint32
	PresentationTimeDelta uint32
	PresentationTime      uint64
	EventDuration         uint32
	ID                    uint32
	SchemeIDURI           string
	Value                 string
	MessageData           []byte
}

// DecodeEmsg - box-specific decode
func DecodeEmsg(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeEmsgSR - box-specific decode
func DecodeEmsgSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - box type
func (b *EmsgBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *EmsgBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// m.Version == 0

// Encode - write box to w
func (b *EmsgBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *EmsgBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (b *EmsgBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
