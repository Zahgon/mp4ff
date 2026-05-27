package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

const (
	appName = "stream-encrypt"
)

var usg = `%s is an HTTP streaming server that encrypts and refragments MP4 files on-the-fly.

It serves the specified input MP4 file at /enc.mp4 with optional encryption and refragmentation.

Usage of %s:
`

type options struct {
	port           int
	samplesPerFrag int
	key            string
	keyID          string
	iv             string
	scheme         string
	inputFile      string
}

func parseOptions(fs *flag.FlagSet, args []string) (*options, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func makeStreamHandler(opts options) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error { _ = "STUB: not implemented"; return nil }
