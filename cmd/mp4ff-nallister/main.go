package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Eyevinn/mp4ff/avc"
	"github.com/Eyevinn/mp4ff/mp4"
)

const (
	appName = "mp4ff-nallister"
)

var usg = `%s lists NAL units and slice types of AVC, HEVC, or VVC tracks of an mp4 (ISOBMFF) file
or a file containing a byte stream in Annex B format.

Takes first video track in a progressive file and the first track in a fragmented file.
It can also output information about SEI NAL units.

The parameter-sets can be further analyzed using mp4ff-pslister.

Usage of %s:
`

type options struct {
	maxNrSamples int
	codec        string
	seiLevel     int
	printRaw     int
	annexB       bool
	printPsHex   bool
	version      bool
}

func parseOptions(fs *flag.FlagSet, args []string) (*options, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error { _ = "STUB: not implemented"; return nil }

// First try to handle Annex B file

// Need to handle progressive files as well as fragmented files

func parseProgressiveMp4(w io.Writer, f *mp4.File, maxNrSamples int, codec string, seiLevel int, parameterSets bool, nrRaw int) error {
	_ = "STUB: not implemented"
	return nil
}

// Next find sample bytes as slice in mdat

func findFirstVideoTrak(moov *mp4.MoovBox) (*mp4.TrakBox, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func getChunkOffset(stbl *mp4.StblBox, chunkNr int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func parseFragmentedMp4(w io.Writer, f *mp4.File, maxNrSamples int, codec string, seiLevel int, parameterSets bool, nrRaw int) error {
	_ = "STUB: not implemented"
	return nil
}

// Auto-detect codec if moov box is there

func getVideoListOffset(moov *mp4.MoovBox, videoTrak *mp4.TrakBox) int64 {
	_ = "STUB: not implemented"
	return 0
}

func printAVCNalus(w io.Writer, avcSPS *avc.SPS, nalus [][]byte, nr int, pts int64, seiLevel int, parameterSets bool, nrRaw int) error {
	_ = "STUB: not implemented"
	return nil
}

func printHEVCNalus(w io.Writer, nalus [][]byte, nr int, pts int64, seiLevel int, parameterSets bool, nrRaw int) error {
	_ = "STUB: not implemented"
	return nil
}

func printVVCNalus(w io.Writer, nalus [][]byte, nr int, pts int64, seiLevel int, parameterSets bool, nrRaw int) error {
	_ = "STUB: not implemented"
	return nil
}

// printSEINALus - print interpreted information if seiLevel is >= 1. Add hex dump if seiLevel >= 2
func printSEINALus(w io.Writer, seiNALUs [][]byte, codec string, seiLevel int, avcSPS *avc.SPS) {
	_ = "STUB: not implemented"
	return
}

// VVC uses same SEI format as HEVC

func bytesToStringN(data []byte, maxNrBytes int) string { _ = "STUB: not implemented"; return "" }

func findAnnexBFrames(nalus [][]byte, codec string) ([][][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isAvcAudNalu(nalu []byte) bool { _ = "STUB: not implemented"; return false }

func isHEVCAudNalu(nalu []byte) bool { _ = "STUB: not implemented"; return false }

func isVVCAudNalu(nalu []byte) bool { _ = "STUB: not implemented"; return false }
