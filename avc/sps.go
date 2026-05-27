package avc

import (
	"errors"

	"github.com/Eyevinn/mp4ff/bits"
)

// ExtendedSAR - Extended Sample Aspect Ratio Code
const ExtendedSAR = 255

// SPS errors
var (
	ErrNotSPS = errors.New("not an SPS NAL unit")
)

// SPS - AVC SPS parameters
type SPS struct {
	Profile                         uint32
	ProfileCompatibility            uint32
	Level                           uint32
	ParameterID                     uint32
	ChromaFormatIDC                 byte
	SeparateColourPlaneFlag         bool
	BitDepthLumaMinus8              uint
	BitDepthChromaMinus8            uint
	QPPrimeYZeroTransformBypassFlag bool
	SeqScalingMatrixPresentFlag     bool
	SeqScalingLists                 []ScalingList
	Log2MaxFrameNumMinus4           uint
	PicOrderCntType                 uint
	Log2MaxPicOrderCntLsbMinus4     uint
	DeltaPicOrderAlwaysZeroFlag     bool
	OffsetForNonRefPic              uint
	OffsetForTopToBottomField       uint
	RefFramesInPicOrderCntCycle     []uint
	NumRefFrames                    uint
	GapsInFrameNumValueAllowedFlag  bool
	FrameMbsOnlyFlag                bool
	MbAdaptiveFrameFieldFlag        bool
	Direct8x8InferenceFlag          bool
	FrameCroppingFlag               bool
	FrameCropLeftOffset             uint
	FrameCropRightOffset            uint
	FrameCropTopOffset              uint
	FrameCropBottomOffset           uint
	Width                           uint
	Height                          uint
	NrBytesBeforeVUI                int
	NrBytesRead                     int
	VUI                             *VUIParameters
}

// ScalingList - 4x4 or 8x8 Scaling lists. Nil if not present
type ScalingList []int

// VUIParameters - extra parameters according to 14496-10, E.1
type VUIParameters struct {
	SampleAspectRatioWidth             uint
	SampleAspectRatioHeight            uint
	OverscanInfoPresentFlag            bool
	OverscanAppropriateFlag            bool
	VideoSignalTypePresentFlag         bool
	VideoFormat                        uint
	VideoFullRangeFlag                 bool
	ColourDescriptionFlag              bool
	ColourPrimaries                    uint
	TransferCharacteristics            uint
	MatrixCoefficients                 uint
	ChromaLocInfoPresentFlag           bool
	ChromaSampleLocTypeTopField        uint
	ChromaSampleLocTypeBottomField     uint
	TimingInfoPresentFlag              bool
	NumUnitsInTick                     uint
	TimeScale                          uint
	FixedFrameRateFlag                 bool
	NalHrdParametersPresentFlag        bool
	NalHrdParameters                   *HrdParameters
	VclHrdParametersPresentFlag        bool
	VclHrdParameters                   *HrdParameters
	LowDelayHrdFlag                    bool // Only present with HrdParameters
	PicStructPresentFlag               bool
	BitstreamRestrictionFlag           bool
	MotionVectorsOverPicBoundariesFlag bool
	MaxBytesPerPicDenom                uint
	MaxBitsPerMbDenom                  uint
	Log2MaxMvLengthHorizontal          uint
	Log2MaxMvLengthVertical            uint
	MaxNumReorderFrames                uint
	MaxDecFrameBuffering               uint
}

// HrdParameters inside VUI
type HrdParameters struct {
	CpbCountMinus1                     uint
	BitRateScale                       uint
	CpbSizeScale                       uint
	CpbEntries                         []CpbEntry
	InitialCpbRemovalDelayLengthMinus1 uint
	CpbRemovalDelayLengthMinus1        uint
	DpbOutputDelayLengthMinus1         uint
	TimeOffsetLength                   uint
}

// CpbEntry inside HrdParameters
type CpbEntry struct {
	BitRateValueMinus1 uint
	CpbSizeValueMinus1 uint
	CbrFlag            bool
}

// ParseSPSNALUnit - Parse AVC SPS NAL unit starting with NAL header
func ParseSPSNALUnit(data []byte, parseVUIBeyondAspectRatio bool) (*SPS, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note! First byte is NAL Header

// Default value if no explicit value present

// The following table is from 14496-10:2020 Section 7.3.2.1.1

// 4x4 for i < 6

// 8x8 for i >= 6

// Empty

// Interlaced so the height should be doubled

//This lacks one extra check?

// CpbDbpDelaysPresent signals if Cpb and Dbp can be found in Picture Timing SEI
func (s *SPS) CpbDpbDelaysPresent() bool { _ = "STUB: not implemented"; return false }

// PicStructPresent signals if pic struct can be found in Picture Timing SEI
func (s *SPS) PicStructPresent() bool { _ = "STUB: not implemented"; return false }

// ChromaArrayType as defined in Section 7.4.2.1.1 under separate_colour_plane_flag
func (s *SPS) ChromaArrayType() byte { _ = "STUB: not implemented"; return 0 }

// parseVUI - parse VUI (Visual Usability Information)
// if parseVUIBeyondAspectRatio is false, stop after AspectRatio has been parsed
func parseVUI(reader *bits.EBSPReader, parseVUIBeyondAspectRatio bool) *VUIParameters {
	_ = "STUB: not implemented"
	return nil
}

func parseHrdParameters(r *bits.EBSPReader) *HrdParameters { _ = "STUB: not implemented"; return nil }

// ConstraintFlags - return the four ConstraintFlag bits
func (a *SPS) ConstraintFlags() byte { _ = "STUB: not implemented"; return 0 }

// GetSARfromIDC - get Sample Aspect Ratio from IDC index
func GetSARfromIDC(index uint) (uint, uint, error) {
	_ = "STUB: not implemented"
	// index 0 is unspecified in the standard. Return 1:1 as a reasonable default
	return 0, 0, nil
}

// 255 is custom and should be handled outside

func readScalingList(reader *bits.EBSPReader, sizeOfScalingList int) ScalingList {
	_ = "STUB: not implemented"
	return *new(ScalingList)
}
