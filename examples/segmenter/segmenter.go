package main

import (
	"io"

	"github.com/Eyevinn/mp4ff/mp4"
)

// Segmenter - segment the progressive inFIle
type Segmenter struct {
	inFile *mp4.File
	tracks []*Track
	nrSegs int //  target number of segments
}

// Track - media track defined by inTrak
type Track struct {
	trackType string
	inTrak    *mp4.TrakBox
	timeScale uint32
	trackID   uint32 // trackID in segmented output
	lang      string
	segments  []sampleInterval
}

// NewSegmenter - create a Segmenter from inFile and fill in track information
func NewSegmenter(inFile *mp4.File) (*Segmenter, error) { _ = "STUB: not implemented"; return nil, nil }

// SetTargetSegmentation - set segment start points
func (s *Segmenter) SetTargetSegmentation(syncTimescale uint32, segStarts []syncPoint) error {
	_ = "STUB: not implemented"
	return nil
}

// MakeInitSegments - initialized and return init segments for all the tracks
func (s *Segmenter) MakeInitSegments() ([]*mp4.InitSegment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MakeMuxedInitSegment - initialized and return one init segments for all the tracks
func (s *Segmenter) MakeMuxedInitSegment() (*mp4.InitSegment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFullSamplesForInterval - get slice of fullsamples with numbers startSampleNr to endSampleNr (inclusive)
func (s *Segmenter) GetFullSamplesForInterval(mp4f *mp4.File, tr *Track, startSampleNr, endSampleNr uint32,
	rs io.ReadSeeker) ([]mp4.FullSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Next find bytes as slice in mdat

//presTime := uint64(int64(decTime) + int64(cto))
//One can either segment on presentationTime or DecodeTime
//presTimeMs := presTime * 1000 / uint64(tr.timeScale)

//fmt.Printf("Sample %d times %d %d, sync %v, offset %d, size %d\n", sampleNr, decTime, cto, isSync, offset, size)

// GetSamplesForInterval - get slice of samples with numbers startSampleNr to endSampleNr (inclusive)
func (s *Segmenter) GetSamplesForInterval(mp4f *mp4.File, trak *mp4.TrakBox, startSampleNr, endSampleNr uint32) ([]mp4.Sample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//presTime := uint64(int64(decTime) + int64(cto))
//One can either segment on presentationTime or DecodeTime
//presTimeMs := presTime * 1000 / uint64(trak.timeScale)

//fmt.Printf("Sample %d times %d %d, sync %v, offset %d, size %d\n", sampleNr, decTime, cto, isSync, offset, size)

// TranslateSampleFlagsForFragment - translate sample flags from stss and sdtp to what is needed in trun
func TranslateSampleFlagsForFragment(stbl *mp4.StblBox, sampleNr uint32) (flags uint32) {
	_ = "STUB: not implemented"
	return 0
}

//2 == does not depend on others (I-picture). May be overridden by sdtp entry

// table starts at 0, but sampleNr is one-based
