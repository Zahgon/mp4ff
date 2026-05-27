package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Eyevinn/mp4ff/mp4"
)

const (
	appName = "add-sidx"
)

var usg = `%s shows how to add a top-level sidx box to a fragmented file provided it does not exist.
Segments are identified by styp boxes if they exist, otherwise by
the start of moof or emsg boxes. It is possible to interpret
every moof box as the start of a new segment, by specifying the "-startSegOnMoof" option.
One can further remove unused encryption boxes with the "-removeEnc" option.

Usage of %s:
`

type options struct {
	removeEncBoxes bool
	nonZeroEPT     bool
	segOnMoof      bool
	version        bool
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

func removeEncryptionBoxes(inFile *mp4.File) { _ = "STUB: not implemented"; return }
