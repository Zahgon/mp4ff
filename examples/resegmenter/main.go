package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	appName = "resegmenter"
)

var usg = `%s is an example on how to resegment a fragmented file to a new target segment duration.
The duration is given in ticks (in the track timescale).

If no init segment in the input, the trex defaults will not be known which may cause an issue.
The  input must be a fragmented file.

Usage of %s:
`

type options struct {
	chunkDur uint64
	verbose  bool
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

func run(args []string, w io.Writer) error { _ = "STUB: not implemented"; return nil }
