package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// VppCBox - VP Codec Configuration Box (vpcC)
// The VPCodecConfigurationBox contains decoder configuration information
// formatted according to the VP codec configuration syntax.
//
// [WebM VP Codec Configuration]: https://www.webmproject.org/vp9/mp4/
type VppCBox struct {
	Version                 byte
	Flags                   uint32
	Profile                 byte
	Level                   byte
	BitDepth                byte
	ChromaSubsampling       byte
	VideoFullRangeFlag      byte
	ColourPrimaries         byte
	TransferCharacteristics byte
	MatrixCoefficients      byte
	CodecInitData           []byte
}

// DecodeVppC - box-specific decode
func DecodeVppC(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeVppCSR - box-specific decode
func DecodeVppCSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Read bit depth and chroma subsampling packed in one byte

// top 4 bits
// next 3 bits
// last bit

// Type - box type
func (b *VppCBox) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *VppCBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *VppCBox) expectedSize(codecInitSize uint16) uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *VppCBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *VppCBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Pack bit depth, chroma subsampling and video full range flag into one byte

// Info - write box info to w
func (b *VppCBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
