package main

import (
	"github.com/Eyevinn/mp4ff/mp4"
)

type RefragmentConfig struct {
	SamplesPerFrag uint32
}

func processFragment(
	inputFrag *mp4.Fragment,
	sa mp4.SampleAccessor,
	config RefragmentConfig,
	writeFunc func(*mp4.Fragment) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func getTotalSampleCount(trun *mp4.TrunBox) uint32 { _ = "STUB: not implemented"; return 0 }

func createFragmentFromSamples(
	inputFrag *mp4.Fragment,
	samples []mp4.FullSample,
	_startNr, _endNr uint32,
	isFirstSubFrag bool,
) (*mp4.Fragment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func min(a, b uint32) uint32 { _ = "STUB: not implemented"; return 0 }
