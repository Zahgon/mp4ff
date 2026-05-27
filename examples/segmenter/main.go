package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	appName = "segmenter"
)

var usg = `%s segments a progressive mp4 file into init and media segments.

The output is either single-track segments, or muxed multi-track segments.
With the -lazy mode, mdat is read and written lazily. The lazy write
is only for single-track segments, to provide a comparison with the multi-track
implementation.
There should be at most one audio and one video track in the input.
The output files will be named as
init segments: <output>_a.mp4 and <output>_v.mp4
media segments: <output>_a_<n>.m4s and <output>_v_<n>.m4s where n >= 1
or init.mp4 and media_<n>.m4s

Codecs supported are AVC and HEVC for video and AAC and AC-3 for audio.

Usage of %s:
`

type options struct {
	chunkDurMS uint64
	multipex   bool
	lazy       bool
	verbose    bool
}

func parseOptions(fs *flag.FlagSet, args []string) (*options, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	if err := run(os.Args, "."); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, outDir string) error { _ = "STUB: not implemented"; return nil }
