package mp4

// ffmpeg boxes according to https://kdenlive.org/en/project/adding-meta-data-to-mp4-video
import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// CTooBox - ©too box defines the ffmpeg encoding tool information
type CTooBox struct {
	Children []Box
}

// DataBox - data box used by ffmpeg for providing information.
type DataBox struct {
	Data []byte
}

// DecodeData - decode Data (from mov_write_string_data_tag in movenc.c in ffmpeg)
func DecodeData(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeDataSR - decode Data (from mov_write_string_data_tag in movenc.c in ffmpeg)
func DecodeDataSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	// Should be 1
	return *new(Box), nil
}

// Should be 0

// Type - box type
func (b *DataBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *DataBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *DataBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *DataBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - box-specific Info
func (b *DataBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
