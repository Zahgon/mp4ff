package av1

import (
	"errors"
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// AV1 parsing errors
var (
	ErrInvalidMarker       = errors.New("invalid marker value found in AV1CodecConfigurationRecord")
	ErrInvalidVersion      = errors.New("unsupported AV1CodecConfigurationRecord version")
	ErrNonZeroReservedBits = errors.New("non-zero reserved bits found in AV1CodecConfigurationRecord")
)

// CodecConfRec - AV1CodecConfigurationRecord
// Specified in https://github.com/AOMediaCodec/av1-isobmff/releases/tag/v1.2.0
type CodecConfRec struct {
	Version                          byte
	SeqProfile                       byte
	SeqLevelIdx0                     byte
	SeqTier0                         byte
	HighBitdepth                     byte
	TwelveBit                        byte
	MonoChrome                       byte
	ChromaSubsamplingX               byte
	ChromaSubsamplingY               byte
	ChromaSamplePosition             byte
	InitialPresentationDelayPresent  byte
	InitialPresentationDelayMinusOne byte
	ConfigOBUs                       []byte
}

// DecodeAVCDecConfRec - decode an AV1CodecConfRec
func DecodeAV1CodecConfRec(data []byte) (CodecConfRec, error) {
	_ = "STUB: not implemented"
	// Minimum size is 4 bytes for the fixed header fields
	return *new(CodecConfRec), nil
}

// Size - total size in bytes
func (a *CodecConfRec) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// EncodeSW- write an AV1CodecConfRec to w
func (a *CodecConfRec) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW- write an AV1CodecConfRec to sw
func (a *CodecConfRec) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }
