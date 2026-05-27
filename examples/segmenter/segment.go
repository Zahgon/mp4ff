package main

import (
	"io"

	"github.com/Eyevinn/mp4ff/mp4"
)

func makeSingleTrackSegments(segmenter *Segmenter, parsedMp4 *mp4.File, rs io.ReadSeeker, outFilePath string) error {
	_ = "STUB: not implemented"
	return nil
}

func makeSingleTrackSegmentsLazyWrite(segmenter *Segmenter, parsedMp4 *mp4.File, rs io.ReadSeeker, outFilePath string) error {
	_ = "STUB: not implemented"
	return nil
}

// Also write media data

func makeMultiTrackSegments(segmenter *Segmenter, parsedMp4 *mp4.File, rs io.ReadSeeker, outFilePath string) error {
	_ = "STUB: not implemented"
	return nil
}

type syncPoint struct {
	sampleNr   uint32
	decodeTime uint64
	presTime   uint64
}

func getSegmentStartsFromVideo(parsedMp4 *mp4.File, segDurMS uint32) (timeScale uint32, syncPoints []syncPoint) {
	_ = "STUB: not implemented"
	return 0, nil
}

type sampleInterval struct {
	startNr uint32
	endNr   uint32 // included in interval
}

func getSegmentIntervals(syncTimescale uint32, syncPoints []syncPoint, trak *mp4.TrakBox) ([]sampleInterval, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func copyMediaData(trak *mp4.TrakBox, startSampleNr, endSampleNr uint32, rs io.ReadSeeker, w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}
