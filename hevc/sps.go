package hevc

import (
	"github.com/Eyevinn/mp4ff/bits"
)

// SPS - HEVC SPS parameters
// ISO/IEC 23008-2 Sec. 7.3.2.2
type SPS struct {
	VpsID                                byte
	MaxSubLayersMinus1                   byte
	TemporalIDNestingFlag                bool
	ProfileTierLevel                     ProfileTierLevel
	SpsID                                byte
	ChromaFormatIDC                      byte
	SeparateColourPlaneFlag              bool
	ConformanceWindowFlag                bool
	PicWidthInLumaSamples                uint32
	PicHeightInLumaSamples               uint32
	ConformanceWindow                    ConformanceWindow
	BitDepthLumaMinus8                   byte
	BitDepthChromaMinus8                 byte
	Log2MaxPicOrderCntLsbMinus4          byte
	SubLayerOrderingInfoPresentFlag      bool
	SubLayeringOrderingInfos             []SubLayerOrderingInfo
	Log2MinLumaCodingBlockSizeMinus3     byte
	Log2DiffMaxMinLumaCodingBlockSize    byte
	Log2MinLumaTransformBlockSizeMinus2  byte
	Log2DiffMaxMinLumaTransformBlockSize byte
	MaxTransformHierarchyDepthInter      byte
	MaxTransformHierarchyDepthIntra      byte
	ScalingListEnabledFlag               bool
	ScalingListDataPresentFlag           bool
	AmpEnabledFlag                       bool
	SampleAdaptiveOffsetEnabledFlag      bool
	PCMEnabledFlag                       bool
	PcmSampleBitDepthLumaMinus1          byte
	PcmSampleBitDepthChromaMinus1        byte
	Log2MinPcmLumaCodingBlockSize        uint16
	Log2DiffMaxMinPcmLumaCodingBlockSize uint16
	PcmLoopFilterDisabledFlag            bool
	NumShortTermRefPicSets               byte
	ShortTermRefPicSets                  []ShortTermRPS
	LongTermRefPicsPresentFlag           bool
	NumLongTermRefPics                   uint8
	LongTermRefPicSets                   []LongTermRPS
	SpsTemporalMvpEnabledFlag            bool
	StrongIntraSmoothingEnabledFlag      bool
	VUIParametersPresentFlag             bool
	VUI                                  *VUIParameters
	ExtensionPresentFlag                 bool
	Extension4bits                       uint8
	RangeExtensionFlag                   bool
	RangeExtension                       *SPSRangeExtension
	MultilayerExtensionFlag              bool
	MultilayerExtension                  *SPSMultilayerExtension
	// SPS 3D extension
	D3ExtensionFlag   bool
	D3Extension       *SPS3dExtension
	SccExtensionFlag  bool
	SccExtension      *SPSSccExtension
	ExtensionDataFlag []bool
}

// ProfileTierLevel according to ISO/IEC 23008-2 Section 7.3.3
type ProfileTierLevel struct {
	GeneralProfileSpace              byte
	GeneralTierFlag                  bool
	GeneralProfileIDC                byte
	GeneralProfileCompatibilityFlags uint32
	GeneralProgressiveSourceFlag     bool
	GeneralInterlacedSourceFlag      bool
	GeneralNonPackedConstraintFlag   bool
	GeneralFrameOnlyConstraintFlag   bool
	// GeneralConstraintIndicatorFlags is 4 + 43+1 bits including the 4 flags above from GeneralProgressiveSourceFlag
	GeneralConstraintIndicatorFlags uint64
	GeneralLevelIDC                 byte
	SubLayers                       []SubLayer
}

type SubLayer struct {
	ProfilePresentFlag        bool
	LevelPresentFlag          bool
	ProfileSpace              byte
	TierFlag                  bool
	ProfileIDC                byte
	ProfileCompatibilityFlags uint32
	ProgressiveSourceFlag     bool
	InterlacedSourceFlag      bool
	NonPackedConstraintFlag   bool
	FrameOnlyConstraintFlag   bool
	ConstraintFlags           uint64 // 43+1 bits
	LayerIDC                  byte
}

func flagFrom(flags uint64, bitNr uint) bool { _ = "STUB: not implemented"; return false }

// parseProfileTierLevel follows ISO/IEC 23008-2 Section 7.3.3
func parseProfileTierLevel(r *bits.EBSPReader, profilePresentFlag bool, maxNumSubLayersMinus1 byte) ProfileTierLevel {
	_ = "STUB: not implemented"
	return *new(ProfileTierLevel)
}

// Including 4 flags from ProgressiveSourceFlag and forward

// ConformanceWindow according to ISO/IEC 23008-2
type ConformanceWindow struct {
	LeftOffset   uint32
	RightOffset  uint32
	TopOffset    uint32
	BottomOffset uint32
}

// SubLayerOrderingInfo according to ISO/IEC 23008-2
type SubLayerOrderingInfo struct {
	MaxDecPicBufferingMinus1 byte
	MaxNumReorderPics        byte
	MaxLatencyIncreasePlus1  byte
}

// VUIParameters - Visual Usability Information as defined in Section E.2
type VUIParameters struct {
	SampleAspectRatioWidth         uint
	SampleAspectRatioHeight        uint
	OverscanInfoPresentFlag        bool
	OverscanAppropriateFlag        bool
	VideoSignalTypePresentFlag     bool
	VideoFormat                    byte
	VideoFullRangeFlag             bool
	ColourDescriptionFlag          bool
	ColourPrimaries                byte
	TransferCharacteristics        byte
	MatrixCoefficients             byte
	ChromaLocInfoPresentFlag       bool
	ChromaSampleLocTypeTopField    uint
	ChromaSampleLocTypeBottomField uint
	NeutralChromaIndicationFlag    bool
	FieldSeqFlag                   bool
	FrameFieldInfoPresentFlag      bool
	DefaultDisplayWindowFlag       bool
	DefDispWinLeftOffset           uint
	DefDispWinRightOffset          uint
	DefDispWinTopOffset            uint
	DefDispWinBottomOffset         uint
	TimingInfoPresentFlag          bool
	NumUnitsInTick                 uint
	TimeScale                      uint
	PocProportionalToTimingFlag    bool
	NumTicksPocDiffOneMinus1       uint
	HrdParametersPresentFlag       bool
	HrdParameters                  *HrdParameters
	BitstreamRestrictionFlag       bool
	BitstreamResctrictions         *BitstreamRestrictions
}

type HrdParameters struct {
	NalHrdParametersPresentFlag            bool
	VclHrdParametersPresentFlag            bool
	SubPicHrdParamsPresentFlag             bool
	TickDivisorMinus2                      uint8
	DuCpbRemovalDelayIncrementLengthMinus1 uint8
	SubPicCpbParamsInPicTimingSeiFlag      bool
	DpbOutputDelayDuLengthMinus1           uint8
	BitRateScale                           uint8
	CpbSizeScale                           uint8
	CpbSizeDuScale                         uint8
	InitialCpbRemovalDelayLengthMinus1     uint8
	AuCpbRemovalDelayLengthMinus1          uint8
	DpbOutputDelayLengthMinus1             uint8
	SubLayerHrd                            []SubLayerHrd
}

// CpbDpbDelaysPresentFlag is defined in ISO/IEC 23008-2 Section E.3.2.
func (h *HrdParameters) CpbDpbDelaysPresentFlag() bool { _ = "STUB: not implemented"; return false }

type SubLayerHrd struct {
	FixedPicRateGeneralFlag     bool
	FixedPicRateWithinCvsFlag   bool
	ElementalDurationInTcMinus1 uint16
	LowDelayHrdFlag             bool
	CpbCntMinus1                uint8
	NalHrdParameters            []SubLayerHrdParameters
	VclHrdParameters            []SubLayerHrdParameters
}

type SubLayerHrdParameters struct {
	BitRateValueMinus1   uint32
	CpbSizeValueMinus1   uint32
	CpbSizeDuValueMinus1 uint32
	BitRateDuValueMinus1 uint32
	CbrFlag              bool
}

// BitstreamRestrictrictions - optional information
type BitstreamRestrictions struct {
	TilesFixedStructureFlag     bool
	MVOverPicBoundariesFlag     bool
	RestrictedRefsPicsListsFlag bool
	MinSpatialSegmentationIDC   uint
	MaxBytesPerPicDenom         uint
	MaxBitsPerMinCuDenom        uint
	Log2MaxMvLengthHorizontal   uint
	Log2MaxMvLengthVertical     uint
}

type LongTermRPS struct {
	PocLsbLt               uint16
	UsedByCurrPicLtFlag    bool
	DeltaPocMsbPresentFlag bool
	DeltaPocMsbCycleLt     uint
}

type SPSRangeExtension struct {
	TransformSkipRotationEnabledFlag    bool
	TransformSkipContextEnabledFlag     bool
	ImplicitRdpcmEnabledFlag            bool
	ExplicitRdpcmEnabledFlag            bool
	ExtendedPrecisionProcessingFlag     bool
	IntraSmoothingDisabledFlag          bool
	HighPrecisionOffsetsEnabledFlag     bool
	PersistentRiceAdaptationEnabledFlag bool
	CabacBypassAlignmentEnabledFlag     bool
}

type SPSMultilayerExtension struct {
	InterViewMvVertConstraintFlag bool
}

type SPS3dExtension struct {
	IvDiMcEnabledFlag0     bool
	IvMvScalEnabledFlag0   bool
	Og2IvmcSubPbSizeMinus3 uint
	IvResPredEnabledFlag   bool
	DepthRefEnabledFlag    bool
	VspMcEnabledFlag       bool
	DbbpEnabledFlag        bool

	IvDiMcEnabledFlag1          bool
	IvMvScalEnabledFlag1        bool
	TexMcEnabledFlag            bool
	Log2TexmcSubPbSizeMinus3    uint
	IntraContourEnabledFlag     bool
	IntraDcOnlyWedgeEnabledFlag bool
	CqtCuPartPredEnabledFlag    bool
	InterDcOnlyEnabledFlag      bool
	SkipIntraEnabledFlag        bool
}

type SPSSccExtension struct {
	CurrPicRefEnabledFlag                   bool
	PaletteModeEnabledFlag                  bool
	PaletteMaxSize                          uint
	DeltaPaletteMaxPredictorSize            uint
	PalettePredictorInitializersPresentFlag bool
	NumPalettePredictorInitializersMinus1   uint
	PalettePredictorInitializer             [][]uint
	MotionVectorResolutionControlIdc        uint8
	IntraBoundaryFilteringDisabledFlag      bool
}

// ParseSPSNALUnit parses SPS NAL unit starting with NAL unit header
func ParseSPSNALUnit(data []byte) (*SPS, error) { _ = "STUB: not implemented"; return nil, nil }

// Note! First two bytes are NALU Header

// Don't continue if we have an issue

// value shall be in the range of 0 to 32, inclusive

// Reserved for future use. Shall be empty

// ImageSize - calculated width and height using ConformanceWindow
func (s *SPS) ImageSize() (width, height uint32) { _ = "STUB: not implemented"; return 0, 0 }

// 4:2:0

// 4:2:2

// parseVUI - parse VUI (Visual Usability Information)
// if parseVUIBeyondAspectRatio is false, stop after AspectRatio has been parsed
func parseVUI(r *bits.EBSPReader, MaxSubLayersMinus1 byte) *VUIParameters {
	_ = "STUB: not implemented"
	return nil
}

func parseHrdParameters(r *bits.EBSPReader,
	commonInfPresentFlag bool, maxNumSubLayersMinus1 byte) *HrdParameters {
	_ = "STUB: not implemented"
	return nil
}

// when fixed_pic_rate_general_flag[ i ] is equal to 1, the value of
// fixed_pic_rate_within_cvs_flag[ i ] is inferred to be equal to 1.

// value shall be in the range of 0 to 2 047, inclusive

// value shall be in the range of 0 to 31, inclusive

func parseSubLayerHrdParameters(r *bits.EBSPReader,
	cpbCntMinus1 uint8, subPicHrdParamsPresentFlag bool) []SubLayerHrdParameters {
	_ = "STUB: not implemented"
	return nil
}

// values shall be in the range of 0 to 2^32 − 2, inclusive

func parseBitstreamRestrictions(r *bits.EBSPReader) *BitstreamRestrictions {
	_ = "STUB: not implemented"
	return nil
}

// ShortTermRPS - Short term Reference Picture Set
type ShortTermRPS struct {
	// Delta Picture Order Count
	DeltaPocS0      []uint32
	DeltaPocS1      []uint32
	UsedByCurrPicS0 []bool
	UsedByCurrPicS1 []bool
	NumNegativePics byte
	NumPositivePics byte
	NumDeltaPocs    byte
}

func (st ShortTermRPS) countInUsePics() uint8 { _ = "STUB: not implemented"; return 0 }

const maxSTRefPics = 16

// parseShortTermRPS - short-term reference pictures with syntax from 7.3.7.
// Focus is on reading/parsing beyond this structure in SPS (and possibly in slice header)
func parseShortTermRPS(r *bits.EBSPReader, idx, numSTRefPicSets byte, sps *SPS) ShortTermRPS {
	_ = "STUB: not implemented"
	return *new(ShortTermRPS)
}

// Slice header

// parse delta_idx_minus1

/* deltaRpsSign */
/* absDeltaRpsMinus1*/
//deltaRps := (1 - (deltaRpsSign << 1)) * (absDeltaRpsMinus1 + 1)

// readPastScalingListData - read and parse all bits of scaling list, without storing values
func readPastScalingListData(r *bits.EBSPReader) { _ = "STUB: not implemented"; return }

// scaling_list_pred_mode_flag[sizeId][matrixId]

// scaling_list_pred_matrix_id_delta[sizeId][matrixId]

// nextCoef = 8;

// scaling_list_dc_coef_minus8[sizeId − 2][matrixId]
// nextCoef = scaling_list_dc_coef_minus8[sizeId − 2][matrixId] + 8

// scaling_list_delta_coef
// nextCoef = ( nextCoef + scaling_list_delta_coef + 256 ) % 256
// ScalingList[sizeId][matrixId][i] = nextCoef

func parseSPS3dExtension(r *bits.EBSPReader) *SPS3dExtension { _ = "STUB: not implemented"; return nil }

func parseSPSSccExtension(r *bits.EBSPReader, ChromaFormatIDC,
	BitDepthLumaMinus8, BitDepthChromaMinus8 byte) *SPSSccExtension {
	_ = "STUB: not implemented"
	return nil
}

// Fill luma

// Fill chroma if any
