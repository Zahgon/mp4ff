package aac

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

const (
	// AAClc - AAC-LC Low Complexity
	AAClc = 2
	// HEAACv1 - HE-AAC version 1 with SBR
	HEAACv1 = 5
	// HEAACv2 - HE-AAC version 2 with SBR and PS
	HEAACv2 = 29
)

// AudioSpecificConfig according to ISO/IEC 14496-3
// Syntax specified in Table 1.15
type AudioSpecificConfig struct {
	ObjectType           byte
	ChannelConfiguration byte // Defined in Table 1.19
	SamplingFrequency    int
	ExtensionFrequency   int
	SBRPresentFlag       bool
	PSPresentFlag        bool
}

// FrequencyTable maps frequency index to sample rate in Hz
var FrequencyTable = map[byte]int{
	0:  96000,
	1:  88200,
	2:  64000,
	3:  48000,
	4:  44100,
	5:  32000,
	6:  24000,
	7:  22050,
	8:  16000,
	9:  12000,
	10: 11025,
	11: 8000,
	12: 7350,
}

// ReverseFrequencies converts sample frequency to index
var ReverseFrequencies = map[int]byte{
	96000: 0,
	88200: 1,
	64000: 2,
	48000: 3,
	44100: 4,
	32000: 5,
	24000: 6,
	22050: 7,
	16000: 8,
	12000: 9,
	11025: 10,
	8000:  11,
	7350:  12,
}

/* Channel configurations according to table 1.19 in ISO/IEC 14496-3
0: Defined in AOT Specific Config
1: 1 channel: front-center
2: 2 channels: front-left, front-right
3: 3 channels: front-center, front-left, front-right
4: 4 channels: front-center, front-left, front-right, back-center
5: 5 channels: front-center, front-left, front-right, back-left, back-right
6: 6 channels: front-center, front-left, front-right, back-left, back-right, LFE-channel
7: 8 channels: front-center, front-left, front-right, side-left, side-right, back-left, back-right, LFE-channel
8-15: Reserved
*/

// DecodeAudioSpecificConfig -
func DecodeAudioSpecificConfig(r io.Reader) (*AudioSpecificConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// do nothing

// Shall be set to AAC-LC here again

//GASpecificConfig()
//GASpecificConfig
// Done (there may be trailing bits)

// Encode - write AudioSpecificConfig to w for AAC-LC and HE-AAC
func (a *AudioSpecificConfig) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// fine

// base audioObjectType

// GASpecificConfig

// getFrequency - either from 4-bit index or 24-bit value
func getFrequency(br *bits.Reader) (frequency int, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}
