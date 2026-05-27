package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	appName = "mp4ff-info"
)

var usg = `%s prints the box tree of input mp4 (ISOBMFF) file.

Usage of %s:
`

type options struct {
	levels  string
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

func run(args []string, w io.Writer) error { _ = "STUB: not implemented"; return nil }
