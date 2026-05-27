package hevc

import (
	"github.com/Eyevinn/mp4ff/bits"
)

// This parser based on Rec. ITU-T H.265 v5 (02/2018) and ISO/IEC 23008-2 Ed. 5
// It implements specification 7.3.6 . Annex F/I extensions aren't supported yet.

// SliceType - HEVC slice type
type SliceType uint

func (s SliceType) String() string { _ = "STUB: not implemented"; return "" }

// HEVC slice types
const (
	SLICE_B = SliceType(0)
	SLICE_P = SliceType(1)
	SLICE_I = SliceType(2)
)

type SliceHeader struct {
	SliceType                         SliceType
	FirstSliceSegmentInPicFlag        bool
	NoOutputOfPriorPicsFlag           bool
	PicParameterSetId                 uint32
	DependentSliceSegmentFlag         bool
	SegmentAddress                    uint
	PicOutputFlag                     bool
	ColourPlaneId                     uint8
	PicOrderCntLsb                    uint16
	ShortTermRefPicSetSpsFlag         bool
	ShortTermRefPicSet                ShortTermRPS
	ShortTermRefPicSetIdx             byte
	NumLongTermSps                    uint8
	NumLongTermPics                   uint
	LongTermRefPicSets                []LongTermRPS
	TemporalMvpEnabledFlag            bool
	SaoLumaFlag                       bool
	SaoChromaFlag                     bool
	NumRefIdxActiveOverrideFlag       bool
	NumRefIdxL0ActiveMinus1           uint8
	NumRefIdxL1ActiveMinus1           uint8
	RefPicListsModification           *RefPicListsModification
	MvdL1ZeroFlag                     bool
	CabacInitFlag                     bool
	CollocatedFromL0Flag              bool
	CollocatedRefIdx                  uint8
	PredWeightTable                   *PredWeightTable
	FiveMinusMaxNumMergeCand          uint8
	UseIntegerMvFlag                  bool
	QpDelta                           int
	CbQpOffset                        int8
	CrQpOffset                        int8
	ActYQpOffset                      int8
	ActCbQpOffset                     int8
	ActCrQpOffset                     int8
	CuChromaQpOffsetEnabledFlag       bool
	DeblockingFilterOverrideFlag      bool
	DeblockingFilterDisabledFlag      bool
	BetaOffsetDiv2                    int8
	TcOffsetDiv2                      int8
	LoopFilterAcrossSlicesEnabledFlag bool
	NumEntryPointOffsets              uint
	OffsetLenMinus1                   uint8
	EntryPointOffsetMinus1            []uint32
	SegmentHeaderExtensionLength      uint16
	SegmentHeaderExtensionDataByte    []byte
	Size                              uint32
}

type RefPicListsModification struct {
	RefPicListModificationFlagL0 bool
	ListEntryL0                  []uint8
	RefPicListModificationFlagL1 bool
	ListEntryL1                  []uint8
}

type PredWeightTable struct {
	LumaLog2WeightDenom        uint8
	DeltaChromaLog2WeightDenom int8
	WeightsL0                  []WeightingFactors
	WeightsL1                  []WeightingFactors
}

// WeightingFactors fields described in specification 7.4.7.3 (Weighted prediction parameters semantics)
type WeightingFactors struct {
	LumaWeightFlag    bool
	ChromaWeightFlag  bool
	DeltaLumaWeight   int8
	LumaOffset        int
	DeltaChromaWeight [2]int8
	DeltaChromaOffset [2]int
}

func ParseSliceHeader(nalu []byte, spsMap map[uint32]*SPS, ppsMap map[uint32]*PPS) (*SliceHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Default according to Section 7.4.7.1

/*
	Pseudocode from standard:

	MinCbLog2SizeY = log2_min_luma_coding_block_size_minus3 + 3
	CtbLog2SizeY = MinCbLog2SizeY + log2_diff_max_min_luma_coding_block_size
	CtbSizeY = 1 << CtbLog2SizeY
	PicWidthInCtbsY = Ceil( pic_width_in_luma_samples ÷ CtbSizeY )
	PicHeightInCtbsY = Ceil( pic_height_in_luma_samples ÷ CtbSizeY )
	PicSizeInCtbsY = PicWidthInCtbsY * PicHeightInCtbsY
*/

/*
		Pseudocode from standard:

		NumPicTotalCurr = 0
		if( nal_unit_type != IDR_W_RADL && nal_unit_type != IDR_N_LP ) {
			for( i = 0; i < NumNegativePics[ CurrRpsIdx ]; i++ ) if( UsedByCurrPicS0[ CurrRpsIdx ][ i ] )
				NumPicTotalCurr++
			for( i = 0; i < NumPositivePics[ CurrRpsIdx ]; i++ ) if( UsedByCurrPicS1[ CurrRpsIdx ][ i ] )
	    		NumPicTotalCurr++
			for( i = 0; i < num_long_term_sps + num_long_term_pics; i++ ) if( UsedByCurrPicLt[ i ] )
				NumPicTotalCurr++
		}
		if( pps_curr_pic_ref_enabled_flag )
			NumPicTotalCurr++
		NumPicTotalCurr += NumActiveRefLayerPics
*/

// The variable ChromaArrayType is derived as equal to 0 when separate_colour_plane_flag is equal to 1
// and chroma_format_idc is equal to 3.

// Decoders shall ignore the presence and value of slice_reserved_flag[ i ]

// value of log2_max_pic_order_cnt_lsb_minus4 shall be in the range of 0 to 12, inclusive

// value shall be in the range of 0 to num_long_term_ref_pics_sps, inclusive

// When the current slice is a P or B slice and num_ref_idx_l0_active_minus1 is not present,
// num_ref_idx_l0_active_minus1 is inferred to be equal to num_ref_idx_l0_default_active_minus1.

// 0 specifies that the syntax elements num_ref_idx_l0_active_minus1 and num_ref_idx_l1_active_minus1 are not present.

// value shall be in the range of 0 to 14, inclusive

// value shall be in the range of 0 to num_ref_idx_l0_active_minus1, inclusive

// MaxNumMergeCand = 5 − five_minus_max_num_merge_cand
// value of MaxNumMergeCand shall be in the range of 1 to 5, inclusive

// values shall be in the range of −12 to +12, inclusive

// values shall be in the range of −12 to +12, inclusive

// values shall both be in the range of −6 to 6, inclusive

// value shall be in the range of 0 to 31, inclusive

// value shall be in the range of 0 to 256, inclusive

// compute the size in bytes. last byte is always aligned

func parseRefPicListsModification(r *bits.EBSPReader, sliceType SliceType,
	refIdxL0Minus1, refIdxL1Minus1 uint8, numPicTotalCurr uint8) (*RefPicListsModification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parsePredWeightTable(r *bits.EBSPReader, sliceType SliceType,
	refIdxL0Minus1, refIdxL1Minus1 uint8, chromaArrayType byte) (*PredWeightTable, error) {
	_ = "STUB: not implemented"
	return nil,

		// value shall be in the range of 0 to 7, inclusive
		nil
}

// ChromaLog2WeightDenom is derived to be equal to luma_log2_weight_denom + delta_chroma_log2_weight_denom
// and the value shall be in the range of 0 to 7, inclusive

// Not implemented
// if( ( pic_layer_id( RefPicList0[ i ] ) != nuh_layer_id ) | |
//( PicOrderCnt( RefPicList0[ i ] ) != PicOrderCnt( CurrPic ) ) )

// Not implemented
// if( ( pic_layer_id( RefPicList0[ i ] ) != nuh_layer_id ) | |
//( PicOrderCnt( RefPicList0[ i ] ) != PicOrderCnt( CurrPic ) ) )

// value shall be in the range of −128 to 127, inclusive

// value shall be in the range of −128 to 127, inclusive

// Not implemented
// if( ( pic_layer_id( RefPicList0[ i ] ) != nuh_layer_id ) | |
//( PicOrderCnt( RefPicList1[ i ] ) != PicOrderCnt( CurrPic ) ) )

// Not implemented
// if( ( pic_layer_id( RefPicList0[ i ] ) != nuh_layer_id ) | |
//( PicOrderCnt( RefPicList1[ i ] ) != PicOrderCnt( CurrPic ) ) )

// value shall be in the range of −128 to 127, inclusive

// value shall be in the range of −128 to 127, inclusive

func ceilDiv(a, b uint) uint { _ = "STUB: not implemented"; return 0 }
