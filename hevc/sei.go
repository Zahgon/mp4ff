package hevc

import (
	"errors"

	"github.com/Eyevinn/mp4ff/sei"
)

var (
	ErrNotSEINalu = errors.New("not an SEI NAL unit")
)

// ParseSEINalu - parse SEI NAL unit (incl header) and return messages given SPS.
// Returns sei.ErrRbspTrailingBitsMissing if the NALU is missing the trailing bits.
func ParseSEINalu(nalu []byte, sps *SPS) ([]sei.SEIMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip NALU header

func fillHEVCPicTimingParams(sps *SPS) sei.HEVCPicTimingParams {
	_ = "STUB: not implemented"
	return *new(sei.HEVCPicTimingParams)
}
