package avc

import (
	"errors"
)

// Errors for parsing and handling AVC slices
var (
	ErrNoSliceHeader      = errors.New("no slice header")
	ErrInvalidSliceType   = errors.New("invalid slice type")
	ErrTooFewBytesToParse = errors.New("too few bytes to parse symbol")
)

// SliceType - AVC slice type
type SliceType uint

func (s SliceType) String() string { _ = "STUB: not implemented"; return "" }

// AVC slice types
const (
	SLICE_P  = SliceType(0)
	SLICE_B  = SliceType(1)
	SLICE_I  = SliceType(2)
	SLICE_SP = SliceType(3)
	SLICE_SI = SliceType(4)
)

// GetSliceTypeFromNALU - parse slice header to get slice type in interval 0 to 4
func GetSliceTypeFromNALU(data []byte) (sliceType SliceType, err error) {
	_ = "STUB: not implemented"
	return *new(SliceType), nil
}

// slice_layer_without_partitioning_rbsp
// slice_data_partition_a_layer_rbsp

// first_mb_in_slice

// The same type is repeated twice to tell if all slices in picture are the same

type SliceHeader struct {
	SliceType                     SliceType
	FirstMBInSlice                uint32
	PicParamID                    uint32
	SeqParamID                    uint32
	ColorPlaneID                  uint32
	FrameNum                      uint32
	IDRPicID                      uint32
	PicOrderCntLsb                uint32
	DeltaPicOrderCntBottom        int32
	DeltaPicOrderCnt              [2]int32
	RedundantPicCnt               uint32
	NumRefIdxL0ActiveMinus1       uint32
	NumRefIdxL1ActiveMinus1       uint32
	ModificationOfPicNumsIDC      uint32
	AbsDiffPicNumMinus1           uint32
	LongTermPicNum                uint32
	AbsDiffViewIdxMinus1          uint32
	LumaLog2WeightDenom           uint32
	ChromaLog2WeightDenom         uint32
	DifferenceOfPicNumsMinus1     uint32
	LongTermFramIdx               uint32
	MaxLongTermFrameIdxPlus1      uint32
	CabacInitIDC                  uint32
	SliceQPDelta                  int32
	SliceQSDelta                  int32
	DisableDeblockingFilterIDC    uint32
	SliceAlphaC0OffsetDiv2        int32
	SliceBetaOffsetDiv2           int32
	SliceGroupChangeCycle         uint32
	Size                          uint32
	FieldPicFlag                  bool
	BottomFieldFlag               bool
	DirectSpatialMvPredFlag       bool
	NumRefIdxActiveOverrideFlag   bool
	RefPicListModificationL0Flag  bool
	RefPicListModificationL1Flag  bool
	NoOutputOfPriorPicsFlag       bool
	LongTermReferenceFlag         bool
	SPForSwitchFlag               bool
	AdaptiveRefPicMarkingModeFlag bool
}

// ParseSliceHeader parses AVC slice header following the syntax in ISO/IEC 14496-10 section 7.3.3
func ParseSliceHeader(nalu []byte, spsMap map[uint32]*SPS, ppsMap map[uint32]*PPS) (*SliceHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// slice_layer_without_partitioning_rbsp
// slice_data_partition_a_layer_rbsp

// ref_pic_list_modification (nal unit type != 20 or 21) Section G.3.3.3.1.1

// end ref_pic_list_modification

// pred_weight_table, section 7.3.3.2

// chroma_idc != 0 in Bento4

// Just parse, don't store this
// luma_weight_l0[i] = SignedGolomb()
// luma_offset_l0[i] = SignedGolomb()

// Just parse, don't store this
// chroma_weight_l0[i][j] = SignedGolomb()
// chroma_offset_l0[i][j] = SignedGolomb()

// Just parse, don't store this
// luma_weight_l1[i] = SignedGolomb()
// luma_offset_l1[i] = SignedGolomb()

// Just parse, don't store this

// chroma_weight_l1[i][j] = SignedGolomb()
// chroma_offset_l1[i][j] = SignedGolomb()

// end pred_weight_table

// dec_ref_pic_marking

// end dec_ref_pic_marking

// compute the size in bytes. The last byte may not be fully parsed
