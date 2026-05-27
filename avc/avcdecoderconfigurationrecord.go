package avc

import (
	"errors"
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// AVC parsing errors
var (
	ErrCannotParseAVCExtension = errors.New("cannot parse SPS extensions")
	ErrLengthSize              = errors.New("can only handle 4byte NAL length size")
)

// DecConfRec - AVCDecoderConfigurationRecord
type DecConfRec struct {
	AVCProfileIndication byte
	ProfileCompatibility byte
	AVCLevelIndication   byte
	SPSnalus             [][]byte
	PPSnalus             [][]byte
	ChromaFormat         byte
	BitDepthLumaMinus1   byte
	BitDepthChromaMinus1 byte
	NumSPSExt            byte
	NoTrailingInfo       bool // To handle strange cases where trailing info is missing
	SkipBytes            int
}

// CreateAVCDecConfRec - extract information from sps and insert sps, pps if includePS set
func CreateAVCDecConfRec(spsNalus [][]byte, ppsNalus [][]byte, includePS bool) (*DecConfRec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// false -> parse only start of VUI

// DecodeAVCDecConfRec - decode an AVCDecConfRec
func DecodeAVCDecConfRec(data []byte) (DecConfRec, error) {
	_ = "STUB: not implemented"
	// Check minimum length for fixed header (6 bytes)
	return *new(DecConfRec), nil
}

// Should be 1

// The first 5 bits are 1

// 5 bits following 3 reserved bits

// Check if we have enough bytes to read NALU length

// Check if we have enough bytes to read NALU

// Check if we have enough bytes to read numPPS

// Check if we have enough bytes to read NALU length

// Check if we have enough bytes to read NALU

// The rest of this structure may vary
// ISO/IEC 14496-15 2017 says that
// Compatible extensions to this record will extend it and
// will not change the configuration version code.
// Readers should be prepared to ignore unrecognized
// data beyond the definition of the data they understand
// (e.g. after the parameter sets in this specification).

// From ISO/IEC 14496-15 2017 Section 5.3.3.1.2
// No extra bytes

// Check if we have enough bytes for the trailing info
// Not according to standard, but have been seen

// Size - total size in bytes
func (a *DecConfRec) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// From ISO/IEC 14496-15 2017 Section 5.3.1.1.2
// No extra bytes

// Encode - write box to w
func (a *DecConfRec) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write an AVCDecConfRec to w
func (a *DecConfRec) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Set length to 4

// Added reserved 3 bits

// From ISO/IEC 14496-15 2017 Section 5.3.3.1.2
// Strange content, but consistent with Size()

//Nothing more to write
