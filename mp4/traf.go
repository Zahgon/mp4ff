package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// TrafBox - Track Fragment Box (traf)
//
// Contained in : Movie Fragment Box (moof)
type TrafBox struct {
	Tfhd     *TfhdBox
	Tfdt     *TfdtBox
	Saiz     *SaizBox
	Saio     *SaioBox
	Sbgp     *SbgpBox
	Sgpd     *SgpdBox
	Senc     *SencBox
	UUIDSenc *UUIDBox // A PIFF box of subtype senc
	Trun     *TrunBox // The first TrunBox
	Truns    []*TrunBox
	Children []Box
}

// DecodeTraf - box-specific decode
func DecodeTraf(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeTrafSR - box-specific decode
func DecodeTrafSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// ContainsSencBox - is there a senc box in traf and is it parsed
// If not parsed, call ParseReadSenc to parse it
func (t *TrafBox) ContainsSencBox() (ok, parsed bool) {
	_ = "STUB: not implemented"
	return false, false
}

// PIFF

// needsSencParsing returns true if the traf contains a senc box that is either
// not yet parsed or was parsed by heuristic (and can benefit from re-parsing
// with authoritative tenc info).
func (t *TrafBox) needsSencParsing() bool { _ = "STUB: not implemented"; return false }

// ParseReadSenc makes a second round to parse a senc box previously read
func (t *TrafBox) ParseReadSenc(defaultIVSize byte, moofStartPos uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// saio should be present, but we try without it, if it doesn't exist

//TODO. Re-enable

//fmt.Printf("offset from saio (%d) and moof differs from senc data start %d\n", posFromSaio, senc.StartPos+16)

// AddChild - add child box
func (t *TrafBox) AddChild(child Box) error { _ = "STUB: not implemented"; return nil }

// Type - return box type
func (t *TrafBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (t *TrafBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// GetChildren - list of child boxes
func (t *TrafBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// Encode - write box to w
	return nil
}

func (t *TrafBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write minf container to sw
func (b *TrafBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box-specific information
func (t *TrafBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// OptimizeTfhdTrun - optimize trun by default values in tfhd box
// Only look at first trun, even if there is more than one
// Don't optimize again, if already done so that no data is present
func (t *TrafBox) OptimizeTfhdTrun() error { _ = "STUB: not implemented"; return nil }

// No need to optimize

// Set defaultSampleDuration in tfhd and remove from trun

// Set defaultSampleSize in tfhd and remove from trun

// RemoveEncryptionBoxes - remove encryption boxes and return number of bytes removed
func (t *TrafBox) RemoveEncryptionBoxes() uint64 { _ = "STUB: not implemented"; return 0 }
