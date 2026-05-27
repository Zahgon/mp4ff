package avc

import (
	"errors"
)

// PPS - Picture Parameter Set
type PPS struct {
	PicParameterSetID                     uint32
	SeqParameterSetID                     uint32
	EntropyCodingModeFlag                 bool
	BottomFieldPicOrderInFramePresentFlag bool
	NumSliceGroupsMinus1                  uint
	SliceGroupMapType                     uint
	RunLengthMinus1                       []uint
	TopLeft                               []uint
	BottomRight                           []uint
	SliceGroupChangeDirectionFlag         bool
	SliceGroupChangeRateMinus1            uint
	PicSizeInMapUnitsMinus1               uint
	SliceGroupID                          []uint
	NumRefIdxI0DefaultActiveMinus1        uint
	NumRefIdxI1DefaultActiveMinus1        uint
	WeightedPredFlag                      bool
	WeightedBipredIDC                     uint
	PicInitQpMinus26                      int
	PicInitQsMinus26                      int
	ChromaQpIndexOffset                   int
	DeblockingFilterControlPresentFlag    bool
	ConstrainedIntraPredFlag              bool
	RedundantPicCntPresentFlag            bool
	Transform8x8ModeFlag                  bool
	PicScalingMatrixPresentFlag           bool
	PicScalingLists                       []ScalingList
	SecondChromaQpIndexOffset             int
}

// AVC PPS errors
var (
	ErrNotPPS = errors.New("not an PPS NAL unit")
)

// ParsePPSNALUnit - Parse AVC PPS NAL unit starting with NAL header
func ParsePPSNALUnit(data []byte, spsMap map[uint32]*SPS) (*PPS, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note! First byte is NAL Header

// slice_group_id[i] has Ceil(Log2(num_slice_groups_minus1 +1) bits)

// Cannot call MoreRbspData, so cannot parse further

// 4x4 for i < 6

// 8x8 for i >= 6
