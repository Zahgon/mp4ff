package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// FLACMetadataBlock - FLAC metadata block
type FLACMetadataBlock struct {
	LastMetadataBlockFlag bool
	BlockType             byte
	Length                uint32
	BlockData             []byte
}

// DfLaBox - FLACSpecificBox (dfLa)
// Defined in https://github.com/xiph/flac/blob/master/doc/isoflac.txt
//
// aligned(8) class FLACSpecificBox
//
//	extends FullBox('dfLa', version=0, 0){
//	  for (i=0; ; i++) { // to end of box
//	    FLACMetadataBlock();
//	  }
//	}
type DfLaBox struct {
	Version        byte
	Flags          uint32
	MetadataBlocks []FLACMetadataBlock
}

// DecodeDfLa - box-specific decode
func DecodeDfLa(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeDfLaSR - box-specific decode
func DecodeDfLaSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Read metadata blocks until end of box
// subtract version and flags

// Read first byte containing last flag and block type

// Read 24-bit length

// Read block data

// 1 byte header + 3 bytes length + data

// If this was the last block, stop reading

// Type - return box type
func (b *DfLaBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *DfLaBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// header + version/flags

// 1 byte header + 3 bytes length + data

// Encode - write box to w
func (b *DfLaBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *DfLaBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Write first byte with last flag and block type

// If this is the last block in the array, set the last flag

// Write 24-bit length

// Write block data

// Info - write box-specific information
func (b *DfLaBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
