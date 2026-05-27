package aac

import (
	"io"
)

// ADTSHeader - data for an unencrypted ADTS Header with one AAC frame.
// Not used in mp4 files, but in MPEG-2 TS.
// Defined in ISO/IEC 13818-7
type ADTSHeader struct {
	ID                     byte // 0 is MPEG-4, 1 is MPEG-2
	ObjectType             byte
	SamplingFrequencyIndex byte
	ChannelConfig          byte
	HeaderLength           byte // Should be 7 or 9
	PayloadLength          uint16
	BufferFullness         uint16
}

// NewADTSHeader - create a new ADTS header
func NewADTSHeader(samplingFrequency int, channelConfig byte, objectType byte, plLen uint16) (*ADTSHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// variable bitrate

// Encode - encode ADTSHeader into byte slice
func (a ADTSHeader) Encode() []byte { _ = "STUB: not implemented"; return nil }

// sync word
//ID=0 for MPEG-4 + layer + protection absent
// profile
// sampling frequency index (3 = 48KHz)
// private
// Channel configuration
// Copyright etc
// The length should include this 7-byte header
// Buffer fullness value
// Nr AAC frames in ADTS frame minus 1

// DecodeADTSHeader by first looking for sync word
func DecodeADTSHeader(r io.Reader) (header *ADTSHeader, offset int, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Find sync 0xfff in first 188 bytes (MPEG-TS related)

// 16-bit CRC

// ignore private

// ignore original/copy, home, copyright

// CRC

// Frequency looks up the sampling frequency for index in ADTSHeader
func (a ADTSHeader) Frequency() uint16 { _ = "STUB: not implemented"; return 0 }
