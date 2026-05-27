package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Eyevinn/mp4ff/mp4"
)

const (
	appName = "mp4ff-pslister"
)

var usg = `%s lists parameter sets for AVC/H.264 or HEVC/H.265 from mp4 sample description, bytestream, or hex input.

It prints them as hex and in verbose mode it also prints details in JSON format.
Usage of %s:
`

type options struct {
	inFile  string
	vpsHex  string
	spsHex  string
	ppsHex  string
	codec   string
	verbose bool
	version bool
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

// Don't just print hex again

// Assume bytestream,AnnexB

// SPS coming back again

// hevc

// VPS coming back again

// Ignore other NALUs

// Now we have hex case left

func getNalusFromBytestream(f io.Reader) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseMp4File(w io.Writer, r io.Reader, codec string, verbose bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Non-fragmented mp4 file with PS in samples

// Next find bytes as slice in mdat

func parseMp4Init(w io.Writer, parsedMp4 *mp4.File, verbose bool) (trackID uint32, codec string, foundPS bool, err error) {
	_ = "STUB: not implemented"
	return 0, "", false, nil
}

func parseMp4Fragment(w io.Writer, parsedMp4 *mp4.File, trackID uint32, codec string, verbose bool) error {
	_ = "STUB: not implemented"
	return nil
}

func printAvcPS(w io.Writer, spsNalus, ppsNalus [][]byte, verbose bool) error {
	_ = "STUB: not implemented"
	return nil
}

/*fullVui*/

/*fullVui*/

func printHevcPS(w io.Writer, vpsNalus, spsNalus, ppsNalus [][]byte, verbose bool) error {
	_ = "STUB: not implemented"
	return nil
}

func printPS(w io.Writer, name string, nr int, ps []byte, psInfo interface{}, verbose bool) {
	_ = "STUB: not implemented"
	return
}
