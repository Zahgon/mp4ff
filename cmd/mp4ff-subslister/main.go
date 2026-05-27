package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Eyevinn/mp4ff/mp4"
)

const (
	appName = "mp4ff-subslister"
)

var usg = `%s lists and displays content of wvtt or stpp samples.
These corresponds to WebVTT or TTML subtitles in ISOBMFF files.
Uses track with given non-zero track ID or first subtitle track found in an asset.

Usage of %s:
`

type options struct {
	maxNrSamples int
	trackID      int
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

// Progressive file

// Fragmented file

func findTrack(moov *mp4.MoovBox, hdlrType string, trackID uint32) (*mp4.TrakBox, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type subtitleTrack struct {
	variant string
	trak    *mp4.TrakBox
}

func parseProgressiveMp4(f *mp4.File, w io.Writer, trackID uint32, maxNrSamples int) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip checking compositionTimeOffset since not uset for subtitles
// Next find sample bytes as slice in mdat

func findWvttTrack(moov *mp4.MoovBox, w io.Writer, trackID uint32) (*subtitleTrack, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findStppTrack(moov *mp4.MoovBox, w io.Writer, trackID uint32) (*subtitleTrack, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFragmentedMp4(f *mp4.File, w io.Writer, trackID uint32, maxNrSamples int) error {
	_ = "STUB: not implemented"
	return nil
}

// Print vttC header and timescale if moov-box is present

// Only wvtt start with a length field.

func printWvttSample(w io.Writer, sample []byte, nr int, pts int64, dur uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func printStppSample(w io.Writer, sample []byte, nr int, pts int64, dur uint32) error {
	_ = "STUB: not implemented"
	return nil
}
