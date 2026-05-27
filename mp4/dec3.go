package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// dec3

// ETSI TS 102 366 V1.4.1 (2017) Table E.1.4
// chanmap - Custom channel map - 16 bits
var CustomChannelMapLocations = map[string]uint16{
	"L":       1 << 15, // Left (MSB)
	"C":       1 << 14, // Center
	"R":       1 << 13, // Right
	"Ls":      1 << 12, // Left Surround
	"Rs":      1 << 11, // Right Surround
	"Lc/Rc":   1 << 10, // Front Left/Right of Center
	"Lrs/Rrs": 1 << 9,  // Left/Right Rear Surround
	"Cs":      1 << 8,  // Back Center
	"Ts":      1 << 7,  // Top Center
	"Lsd/Rsd": 1 << 6,  // Left/Right Surround Direct
	"Lw/Rw":   1 << 5,  // Left/Right Wide
	"Vhl/Vhr": 1 << 4,  // Top Front Left/Right
	"Vhc":     1 << 3,  // Top Front Center
	"Lts/Rts": 1 << 2,  // Left/Right Top Surround
	"LFE2":    1 << 1,  // Low Frequency 2
	"LFE":     1 << 0,  // Low Frequency
}

// EC3ChannelLocationBits - channel location signal in 9bits Table F.6.1
var EC3ChannelLocationBits = []string{
	"Lc/Rc",
	"Lrs/Rrs",
	"Cs",
	"Ts",
	"Lsd/Rsd",
	"Lw/Rw",
	"Lvh/Rvh",
	"Cvh",
	"LFE2", //MSB
}

// Dec3Box - AC3SpecificBox from ETSI TS 102 366 V1.4.1 F.4 (2017)
type Dec3Box struct {
	DataRate  uint16
	NumIndSub uint16
	EC3Subs   []EC3Sub
	Reserved  []byte
}

// EC3Sub - Enhanced AC-3 substream information
type EC3Sub struct {
	FSCod     byte
	BSID      byte
	ASVC      byte
	BSMod     byte
	ACMod     byte
	LFEOn     byte
	NumDepSub byte
	ChanLoc   uint16
}

// DecodeDec3 - box-specific decode
func DecodeDec3(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeDec3SR - box-specific decode
func DecodeDec3SR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

func decodeDec3FromData(data []byte) (Box, error) { _ = "STUB: not implemented"; return *new(Box), nil }

// There must be one base stream

// Reserved 0

// Reserved 000

// Reserved 0

// Type - box type
func (b *Dec3Box) Type() string {
	_ = "STUB: not implemented"

	// Size - calculated size of box
	return ""
}

func (b *Dec3Box) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *Dec3Box) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - write box to sw
func (b *Dec3Box) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// reserved 0

// reserved 000

// Reserved 0d

func (b *Dec3Box) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Dec3Box) ChannelInfo() (nrChannels int, chanmap uint16) {
	_ = "STUB: not implemented"

	// All Enhanced AC-3 bit streams shall contain an independent substream
	// assigned substream ID 0 (E.1.3.1.2)
	return 0, 0
}

// Get base channel configuration according to acmod

// Dependent substreams associated with this independent substream

// Check if a channel pair (contains /)
