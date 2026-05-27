package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// AC3SampleRates - Sample rates as defined in  ETSI TS 102 366 V1.4.1 (2017) section 4.4.1.3
// Signaled in fscod - Sample rate code - 2 bits
var AC3SampleRates = []int{48000, 44100, 32000}

// AX3acmodChanneTable - channel configurations from ETSI TS 102 366 V1.4.1 (2017) section 4.4.2.3A
// Signaled in acmod - audio coding mode - 3 bits
var AC3acmodChannelTable = []string{
	"L/R", //Ch1 Ch2 dual mono but name them L R
	"C",
	"L/R",
	"L/C/R",
	"L/R/Cs",
	"L/C/R/Cs",
	"L/R/Ls/Rs",
	"L/C/R/Ls/Rs",
}

// AC3BitrateCodesKbps - Bitrates in kbps ETSI TS 102 366 V1.4.1 Table F.4.1 (2017)
var AC3BitrateCodesKbps = []uint16{
	32,
	40,
	48,
	56,
	64,
	80,
	96,
	112,
	128,
	160,
	192,
	224,
	256,
	320,
	384,
	448,
	512,
	576,
	640,
}

// Dac3Box - AC3SpecificBox from ETSI TS 102 366 V1.4.1 F.4 (2017)
// Extra b
type Dac3Box struct {
	FSCod         byte
	BSID          byte
	BSMod         byte
	ACMod         byte
	LFEOn         byte
	BitRateCode   byte
	Reserved      byte
	InitialZeroes byte // Should be zero
}

// DecodeDac3 - box-specific decode
func DecodeDac3(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeDac3SR - box-specific decode
func DecodeDac3SR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

func decodeDac3FromData(data []byte) (Box, error) { _ = "STUB: not implemented"; return *new(Box), nil }

// 5 bits reserved follows

// Type - box type
func (b *Dac3Box) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *Dac3Box) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *Dac3Box) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write box to sw
func (b *Dac3Box) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// 5-bits reserved

// ChannelInfo - number of channels and channelmap according to E.1.3.1.8
func (b *Dac3Box) ChannelInfo() (nrChannels int, chanmap uint16) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (b *Dac3Box) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Dac3Box) BitrateBps() int { _ = "STUB: not implemented"; return 0 }

func (b *Dac3Box) SamplingFrequency() int { _ = "STUB: not implemented"; return 0 }

// GetChannelListFromACMod - get list of channels from acmod byte
func GetChannelListFromACMod(acmod byte) []string { _ = "STUB: not implemented"; return nil }
