package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	appName = "mp4ff-decrypt"
)

var usg = `%s decrypts a fragmented mp4 file encrypted with Common Encryption scheme cenc or cbcs.
For a media segment, it needs an init segment with encryption information.

Usage of %s:
`

type options struct {
	initFilePath string
	keyStrs      stringSliceFlag
	version      bool
}

type stringSliceFlag []string

func (s *stringSliceFlag) String() string { _ = "STUB: not implemented"; return "" }

func (s *stringSliceFlag) Set(value string) error { _ = "STUB: not implemented"; return nil }

func parseKeys(keyStrs []string) (key []byte, keysByKID map[string][]byte, strictKIDMode bool, err error) {
	_ = "STUB: not implemented"
	return nil, nil, false, nil
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

func decryptFileWithKeyMap(r, initR io.Reader, w io.Writer, key []byte, keysByKID map[string][]byte, strictKIDMode bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Write output to file

// Init segment was separate, so senc boxes were not parsed during decode.
// Parse them now using the init segment's encryption info.

// Write output to file
