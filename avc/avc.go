package avc

// NaluType - AVC NAL unit type
type NaluType uint16

const (
	// NALU_NON_IDR - Non-IDR Slice NAL unit
	NALU_NON_IDR = NaluType(1)
	// NALU_IDR - IDR Random Access Slice NAL Unit
	NALU_IDR = NaluType(5)
	// NALU_SEI - Supplementary Enhancement Information NAL Unit
	NALU_SEI = NaluType(6)
	// NALU_SPS - SequenceParameterSet NAL Unit
	NALU_SPS = NaluType(7)
	// NALU_PPS - PictureParameterSet NAL Unit
	NALU_PPS = NaluType(8)
	// NALU_AUD - AccessUnitDelimiter NAL Unit
	NALU_AUD = NaluType(9)
	// NALU_EO_SEQ - End of Sequence NAL Unit
	NALU_EO_SEQ = NaluType(10)
	// NALU_EO_STREAM - End of Stream NAL Unit
	NALU_EO_STREAM = NaluType(11)
	// NALU_FILL - Filler NAL Unit
	NALU_FILL = NaluType(12)
)

func (a NaluType) String() string { _ = "STUB: not implemented"; return "" }

// GetNaluType - get NALU type from  NALU Header byte
func GetNaluType(naluHeader byte) NaluType { _ = "STUB: not implemented"; return *new(NaluType) }

// FindNaluTypes - find list of NAL unit types in sample
func FindNaluTypes(sample []byte) []NaluType { _ = "STUB: not implemented"; return nil }

// FindNaluTypesUpToFirstVideoNALU - find list of NAL unit types in sample
func FindNaluTypesUpToFirstVideoNALU(sample []byte) []NaluType {
	_ = "STUB: not implemented"
	return nil
}

// first video nalu

// IsIDRSample - does sample contain IDR NALU
func IsIDRSample(sample []byte) bool { _ = "STUB: not implemented"; return false }

// ContainsNaluType - is specific NaluType present in sample
func ContainsNaluType(sample []byte, specificNalType NaluType) bool {
	_ = "STUB: not implemented"
	return false
}

// HasParameterSets - Check if H.264 SPS and PPS are present
func HasParameterSets(b []byte) bool { _ = "STUB: not implemented"; return false }

// GetParameterSets - get (multiple) SPS and PPS from a sample
func GetParameterSets(sample []byte) (sps [][]byte, pps [][]byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

//SPS and PPS must come before video

// IsVideoNaluType returns true if nalu type is a VCL nalu.
func IsVideoNaluType(naluType NaluType) bool { _ = "STUB: not implemented"; return false }
