package avc

//Functions to handle AnnexB Byte stream format"

import (
	"math/bits"
)

// ExtractNalusFromByteStream extracts NALUs without startcode from ByteStream.
// This function is codec agnostic.
func ExtractNalusFromByteStream(data []byte) [][]byte { _ = "STUB: not implemented"; return nil }

// Remove zeros from end of NAL unit

func extractSlice(data []byte, start, stop int) []byte { _ = "STUB: not implemented"; return nil }

type scNalu struct {
	startCodeLength int
	startPos        int
}

// ConvertByteStreamToNaluSample changes start codes to 4-byte length fields.
// This function is codec agnostic.
func ConvertByteStreamToNaluSample(stream []byte) []byte { _ = "STUB: not implemented"; return nil }

// In-place replacement of startcodes for length fields

// Build new output slice with one extra byte per NALU

// Cut overflow bits at compile time to use it safely on < 64-bit systems
const (
	magicLeft  uint = 0x0101010101010101 >> (64 - bits.UintSize)
	magicRight uint = 0x8080808080808080 >> (64 - bits.UintSize)
)

// This function implement bit-trick to search zero byte in numbered type.
// You can find detail explanation here https://graphics.stanford.edu/~seander/bithacks.html#ZeroInWord .
func hasZeroByte(x uint) bool { _ = "STUB: not implemented"; return false }

func getStartCodePositions(stream []byte) (scNalus []scNalu, minStartCodeLength int) {
	_ = "STUB: not implemented"
	return nil,

		// Platform limitation must be known to iterate over slice effectively and safely.
		0
}

// Faster approach for searching the NALU start codes is applicable only for slices which length is multiple of uintSize.
// Max length of slice should be limited accordingly.

// Iterator value is declared outside to continue iteration in second loop.

// Iterating over slice by uintSize as all intermediate bytes will be checked as needed.

// This code is inspired by ffmpeg https://ffmpeg.org/doxygen/trunk/avc_8c_source.html#l00030 .

// Reference to the current byte in slice converted to untyped Pointer than it's cast to uint reference
// and finally dereference to value. hasZeroByte() func check every byte of uint for zero.
// hasZeroByte() func is endianness agnostic.

// Minimal start code size is 3, so checking every odd byte of uint is enough.

// Next branch will check every neighbor byte to find start code pattern.
// Be aware! It will check two first bytes from next uint.

// We should check remain bytes with old approach.

// ConvertSampleToByteStream replaces 4-byte NALU lengths with start codes.
// This function is codec agnostic.
func ConvertSampleToByteStream(sample []byte) []byte { _ = "STUB: not implemented"; return nil }

// GetParameterSetsFromByteStream copies AVC SPS and PPS nalus from bytestream (Annex B)
func GetParameterSetsFromByteStream(data []byte) (spss, ppss [][]byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Remove zeros from end of NAL unit

// Video NALU types are below 6

// ExtractNalusOfTypeFromByteStream returns all AVC nalus of wanted type from bytestream.
// If stopAtVideo, the stream is not scanned beyond the first video NAL unit.
func ExtractNalusOfTypeFromByteStream(nType NaluType, data []byte, stopAtVideo bool) [][]byte {
	_ = "STUB: not implemented"
	return nil
}

// Remove zeros from end of NAL unit

// Video nal unit type

// GetFirstAVCVideoNALUFromByteStream returns a slice with the first video nal unit.
// No new memory is allocated, but a subslice of data is returned.
func GetFirstAVCVideoNALUFromByteStream(data []byte) []byte { _ = "STUB: not implemented"; return nil }

// Remove zeros from end of NAL unit
