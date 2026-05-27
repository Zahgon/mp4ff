package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// DrefBox - Data Reference Box (dref - mandatory)
//
// Contained id: Data Information Box (dinf)
//
// Defines the location of the media data. If the data for the track is located in the same file
// it contains nothing useful.
type DrefBox struct {
	Version    byte
	Flags      uint32
	EntryCount uint32
	Children   []Box
}

// CreateDref - Create an DataReferenceBox for selfcontained content
func CreateDref() *DrefBox { _ = "STUB: not implemented"; return nil }

// AddChild - Add a child box and update EntryCount
func (d *DrefBox) AddChild(box Box) { _ = "STUB: not implemented"; return }

// DecodeDref - box-specific decode
func DecodeDref(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Note higher startPos for children since not simple container.

// DecodeDrefSR - box-specific decode
func DecodeDrefSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Note higher startPos for children since not simple container.

// incremented by AddChild

// Type - box type
func (d *DrefBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (d *DrefBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write dref box to w including children
func (d *DrefBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - write dref box to w including children
func (d *DrefBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (d *DrefBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
