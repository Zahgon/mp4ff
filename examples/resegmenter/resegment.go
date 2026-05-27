package main

import (
	"io"

	"github.com/Eyevinn/mp4ff/mp4"
)

// Resegment file into multiple segments
func Resegment(w io.Writer, in *mp4.File, chunkDur uint64, verbose bool) (*mp4.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// approximative, but good for allocation

//  ftyp + moov + (styp+moof+mdat for each segment)

func addSamplesToFrag(frag *mp4.Fragment, samples []mp4.FullSample, nextSampleNrToWrite, stopNr int, trackID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func addNewSegment(oFile *mp4.File, styp *mp4.StypBox, seqNr, trackID uint32) (*mp4.Fragment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
