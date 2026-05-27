package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Eyevinn/mp4ff/mp4"
)

const (
	appName = "mp4ff-encrypt"
)

var usg = `%s encrypts a fragmented mp4 file using Common Encryption with cenc or cbcs scheme.
A combined fragmented file with init segment and media segment(s) will be encrypted.
For a pure media segment, an init segment with encryption information is needed.
For video, only AVC with avc1 and HEVC with hvc1 sample entries are currently supported.
For audio, all supported audio codecs should work.

Usage of %s:
`

type options struct {
	initFile string
	kidStr   string
	keyStr   string
	ivHex    string
	scheme   string
	psshFile string
	version  bool
}

func parseOptions(fs *flag.FlagSet, args []string) (*options, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error { _ = "STUB: not implemented"; return nil }

func encryptFile(ifh io.Reader, ofh io.Writer, initSeg *mp4.InitSegment,
	scheme, kidStr, keyStr, ivHex string, psshData []byte) error {
	_ = "STUB: not implemented"
	return nil
}
