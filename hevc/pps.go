package hevc

import (
	"errors"

	"github.com/Eyevinn/mp4ff/bits"
)

// HEVC PPS errors
var (
	ErrNotPPS = errors.New("not an PPS NAL unit")
)

// This parser based on Rec. ITU-T H.265 v5 (02/2018) and ISO/IEC 23008-2 Ed. 5

// PPS - Picture Parameter Set
type PPS struct {
	PicParameterSetID                      uint32
	SeqParameterSetID                      uint32
	DependentSliceSegmentsEnabledFlag      bool
	OutputFlagPresentFlag                  bool
	NumExtraSliceHeaderBits                uint8
	SignDataHidingEnabledFlag              bool
	CabacInitPresentFlag                   bool
	NumRefIdxL0DefaultActiveMinus1         uint8
	NumRefIdxL1DefaultActiveMinus1         uint8
	InitQpMinus26                          int8
	ConstrainedIntraPredFlag               bool
	TransformSkipEnabledFlag               bool
	CuQpDeltaEnabledFlag                   bool
	DiffCuQpDeltaDepth                     uint
	CbQpOffset                             int8
	CrQpOffset                             int8
	SliceChromaQpOffsetsPresentFlag        bool
	WeightedPredFlag                       bool
	WeightedBipredFlag                     bool
	TransquantBypassEnabledFlag            bool
	TilesEnabledFlag                       bool
	EntropyCodingSyncEnabledFlag           bool
	NumTileColumnsMinus1                   uint
	NumTileRowsMinus1                      uint
	UniformSpacingFlag                     bool
	ColumnWidthMinus1                      []uint
	RowHeightMinus1                        []uint
	LoopFilterAcrossTilesEnabledFlag       bool
	LoopFilterAcrossSlicesEnabledFlag      bool
	DeblockingFilterControlPresentFlag     bool
	DeblockingFilterOverrideEnabledFlag    bool
	DeblockingFilterDisabledFlag           bool
	BetaOffsetDiv2                         int8
	TcOffsetDiv2                           int8
	ScalingListDataPresentFlag             bool
	ListsModificationPresentFlag           bool
	Log2ParallelMergeLevelMinus2           uint
	SliceSegmentHeaderExtensionPresentFlag bool
	ExtensionPresentFlag                   bool
	RangeExtensionFlag                     bool
	RangeExtension                         *RangeExtension
	MultilayerExtensionFlag                bool
	MultilayerExtension                    *MultilayerExtension
	// PPS 3D extension
	D3ExtensionFlag   bool
	D3Extension       *D3Extension
	SccExtensionFlag  bool
	SccExtension      *SccExtension
	Extension4bits    uint8
	ExtensionDataFlag []bool
}

type RangeExtension struct {
	Log2MaxTransformSkipBlockSizeMinus2 uint
	CrossComponentPredictionEnabledFlag bool
	ChromaQpOffsetListEnabledFlag       bool
	DiffCuChromaQpOffsetDepth           uint
	ChromaQpOffsetListLenMinus1         uint
	CbQpOffsetList                      []int8
	CrQpOffsetList                      []int8
	Log2SaoOffsetScaleLuma              uint
	Log2SaoOffsetScaleChroma            uint
}

type MultilayerExtension struct {
	PocResetInfoPresentFlag  bool
	InferScalingListFlag     bool
	ScalingListRefLayerId    uint8
	NumRefLocOffsets         uint
	RefLocOffsetLayerIds     []uint8
	RefLocOffsets            map[uint8]RefLocOffset
	ColourMappingEnabledFlag bool
	ColourMappingTable       *ColourMappingTable
}

type RefLocOffset struct {
	ScaledRefLayerOffsetPresentFlag bool
	ScaledRefLayerLeftOffset        int16
	ScaledRefLayerTopOffset         int16
	ScaledRefLayerRightOffset       int16
	ScaledRefLayerBottomOffset      int16
	RefRegionOffsetPresentFlag      bool
	RefRegionLeftOffset             int16
	RefRegionTopOffset              int16
	RefRegionRightOffset            int16
	RefRegionBottomOffset           int16
	ResamplePhaseSetPresentFlag     bool
	PhaseHorLuma                    uint8
	PhaseVerLuma                    uint8
	PhaseHorChromaPlus8             uint8
	PhaseVerChromaPlus8             uint8
}

type ColourMappingTable struct {
	NumCmRefLayersMinus1         uint8
	RefLayerId                   []uint8
	OctantDepth                  uint8
	YPartNumLog2                 uint8
	LumaBitDepthCmInputMinus8    uint
	ChromaBitDepthCmInputMinus8  uint
	LumaBitDepthCmOutputMinus8   uint
	ChromaBitDepthCmOutputMinus8 uint
	ResQuantBits                 uint8
	DeltaFlcBitsMinus1           uint8
	AdaptThresholdUDelta         int
	AdaptThresholdVDelta         int
	Octants                      map[string][4]Octant
}

type Octant struct {
	CodedResFlag bool
	CodedRes     [3]struct {
		ResCoeffQ uint
		ResCoeffR uint
		ResCoeffS bool
	}
}

type SccExtension struct {
	CurrPicRefEnabledFlag                      bool
	ResidualAdaptiveColourTransformEnabledFlag bool
	SliceActQpOffsetsPresentFlag               bool
	ActYQpOffsetPlus5                          int
	ActCbQpOffsetPlus5                         int
	ActCrQpOffsetPlus3                         int
	PalettePredictorInitializersPresentFlag    bool
	NumPalettePredictorInitializers            uint
	MonochromePaletteFlag                      bool
	LumaBitDepthEntryMinus8                    uint
	ChromaBitDepthEntryMinus8                  uint
	PalettePredictorInitializer                [][]uint
}

// D3Extension represent PPS 3D extension
type D3Extension struct {
	DltsPresentFlag              bool
	NumDepthLayersMinus1         uint8
	BitDepthForDepthLayersMinus8 uint8
	DepthLayers                  []DepthLayer
}

type DepthLayer struct {
	DltFlag                bool
	DltPredFlag            bool
	DltValFlagsPresentFlag bool
	DltValueFlag           []bool
	DeltaDlt               *DeltaDlt
}

type DeltaDlt struct {
	NumValDeltaDlt       uint
	MaxDiff              uint
	MinDiffMinus1        uint
	DeltaDltVal0         uint
	DeltaValDiffMinusMin []uint
}

// ParsePPSNALUnit - Parse AVC PPS NAL unit starting with NAL header
func ParsePPSNALUnit(data []byte, spsMap map[uint32]*SPS) (*PPS, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note! First two bytes are NALU Header

// value shall be in the range of 0 to 14, inclusive

// value shall be in the range of −( 26 + QpBdOffsetY ) to +25, inclusive

// values shall be in the range of −12 to +12, inclusive

// values shall be in the range of −6 to 6, inclusive

// Reserved for future use. Shall be empty

func parseRangeExtension(r *bits.EBSPReader, transformSkipEnabled bool) (*RangeExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// values shall be in the range of −12 to +12, inclusive

func parseMultilayerExtension(r *bits.EBSPReader) (*MultilayerExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// value shall be in the range of −2^14 to 2^14 − 1, inclusive

// value shall be in the range of −2^14 to 2^14 − 1, inclusive

// value shall be in the range of 0 to 31, inclusive

// value shall be in the range of 0 to 63, inclusive

func parseColourMappingTable(r *bits.EBSPReader) (*ColourMappingTable, error) {
	_ = "STUB: not implemented"
	return nil,

		// value shall be in the range of 0 to 61, inclusive
		nil
}

//Max( 0, ( 10 + BitDepthCmInputY − BitDepthCmOutputY − cm_res_quant_bits − ( cm_delta_flc_bits_minus1 + 1 ) ) )
//BitDepthCmInputY = 8 + luma_bit_depth_cm_input_minus8
//BitDepthCmOutputY = 8 + luma_bit_depth_cm_output_minus8

func parseColourMappingOctants(r *bits.EBSPReader, octantDepth uint, partNumY uint, resLsBits int,
	inpDepth, idxY, idxCb, idxCr, inpLength uint) (map[string][4]Octant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A map is used instead of the 5-dimensional array in the standard pseudo-code
// Key represent [ idxShiftY ][ idxCb ][ idxCr ] with idxShiftY variable part

func makeKeyOctant(idxShiftY, idxCb, idxCr uint) string { _ = "STUB: not implemented"; return "" }

func parseSccExtension(r *bits.EBSPReader) (*SccExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fill luma

// Fill chroma if any

func parse3dExtension(r *bits.EBSPReader) (*D3Extension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// variable depthMaxValue is set equal to ( 1 << ( pps_bit_depth_for_depth_layers_minus8 + 8 ) ) − 1

func parseDeltaDlt(r *bits.EBSPReader, BitDepthForDepthLayers int) (*DeltaDlt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// variable minDiff is set equal to ( min_diff_minus1 + 1 )
// length of delta_val_diff_minus_min[ k ] syntax element is Ceil( Log2( max_diff − minDiff + 1 ) ) bits
