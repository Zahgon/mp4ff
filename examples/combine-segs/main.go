package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Eyevinn/mp4ff/mp4"
)

func main() {
	if err := run("."); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(outDir string) error { _ = "STUB: not implemented"; return nil }

func combineInitSegments(files []string, newTrackIDs []uint32) (*mp4.InitSegment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func combineMediaSegments(files []string, newTrackIDs []uint32) (*mp4.MediaSegment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Here we should have the trex from the corresponding init segment

func writeSeg(seg encoder, filename string) error { _ = "STUB: not implemented"; return nil }

type encoder interface {
	Encode(w io.Writer) error
}
