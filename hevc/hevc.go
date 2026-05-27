package hevc

// NaluType - HEVC nal type according to ISO/IEC 23008-2 Table 7.1
type NaluType uint16

// HEVC NALU types
const (
	NALU_TRAIL_N = NaluType(0)
	NALU_TRAIL_R = NaluType(1)
	NALU_TSA_N   = NaluType(2)
	NALU_TSA_R   = NaluType(3)
	NALU_STSA_N  = NaluType(4)
	NALU_STSA_R  = NaluType(5)
	NALU_RADL_N  = NaluType(6)
	NALU_RADL_R  = NaluType(7)
	NALU_RASL_N  = NaluType(8)
	NALU_RASL_R  = NaluType(9)
	// BLA_W_LP and the following types are Random Access
	NALU_BLA_W_LP   = NaluType(16)
	NALU_BLA_W_RADL = NaluType(17)
	NALU_BLA_N_LP   = NaluType(18)
	NALU_IDR_W_RADL = NaluType(19)
	NALU_IDR_N_LP   = NaluType(20)
	NALU_CRA        = NaluType(21)
	// Reserved IRAP VCL NAL unit types
	NALU_IRAP_VCL22 = NaluType(22)
	NALU_IRAP_VCL23 = NaluType(23)
	// NALU_VPS - VideoParameterSet NAL Unit
	NALU_VPS = NaluType(32)
	// NALU_SPS - SequenceParameterSet NAL Unit
	NALU_SPS = NaluType(33)
	// NALU_PPS - PictureParameterSet NAL Unit
	NALU_PPS = NaluType(34)
	// NALU_AUD - AccessUnitDelimiter NAL Unit
	NALU_AUD = NaluType(35)
	//NALU_EOS - End of Sequence NAL Unit
	NALU_EOS = NaluType(36)
	//NALU_EOB - End of Bitstream NAL Unit
	NALU_EOB = NaluType(37)
	//NALU_FD - Filler data NAL Unit
	NALU_FD = NaluType(38)
	//NALU_SEI_PREFIX - Prefix SEI NAL Unit
	NALU_SEI_PREFIX = NaluType(39)
	//NALU_SEI_SUFFIX - Suffix SEI NAL Unit
	NALU_SEI_SUFFIX = NaluType(40)

	highestVideoNaluType = 31
)

func (n NaluType) String() string { _ = "STUB: not implemented"; return "" }

// GetNaluType - extract NALU type from first byte of NALU Header
func GetNaluType(naluHeaderStart byte) NaluType { _ = "STUB: not implemented"; return *new(NaluType) }

// GetNaluLayerID extracts nuh_layer_id (6 bits) from the 2-byte HEVC NAL unit header.
// HEVC NAL header: forbidden(1) | nal_unit_type(6) | nuh_layer_id(6) | nuh_temporal_id_plus1(3)
func GetNaluLayerID(naluHeader []byte) byte { _ = "STUB: not implemented"; return 0 }

// GetNaluTemporalID extracts nuh_temporal_id (nuh_temporal_id_plus1 - 1) from the 2-byte HEVC NAL unit header.
func GetNaluTemporalID(naluHeader []byte) byte { _ = "STUB: not implemented"; return 0 }

// NaluInfo holds parsed information from a HEVC NAL unit header.
type NaluInfo struct {
	Type       NaluType
	LayerID    byte
	TemporalID byte
}

// ParseNaluHeader parses a 2-byte HEVC NAL unit header.
func ParseNaluHeader(naluHeader []byte) NaluInfo { _ = "STUB: not implemented"; return *new(NaluInfo) }

// SplitNalusByLayerID splits length-prefixed NALUs in a sample by nuh_layer_id.
// The lengthSize is typically 4 bytes.
func SplitNalusByLayerID(sample []byte, lengthSize int) map[byte][][]byte {
	_ = "STUB: not implemented"
	return nil
}

// FindNaluTypes - find list of nalu types in sample
func FindNaluTypes(sample []byte) []NaluType { _ = "STUB: not implemented"; return nil }

// FindNaluTypesUpToFirstVideoNalu - all nalu types up to first video nalu
func FindNaluTypesUpToFirstVideoNalu(sample []byte) []NaluType {
	_ = "STUB: not implemented"
	return nil
}

// Video has started

// IsVideoNaluType returns true if NaluType is a video type (<= 31)
func IsVideoNaluType(naluType NaluType) bool { _ = "STUB: not implemented"; return false }

// ContainsNaluType - is specific NaluType present in sample
func ContainsNaluType(sample []byte, specificNaluType NaluType) bool {
	_ = "STUB: not implemented"
	return false
}

// IsRAPSample - is Random Access picture (NALU 16-23)
func IsRAPSample(sample []byte) bool { _ = "STUB: not implemented"; return false }

// IsIDRSample - is IDR picture (NALU 19-20)
func IsIDRSample(sample []byte) bool { _ = "STUB: not implemented"; return false }

// HasParameterSets - Check if HEVC VPS, SPS and PPS are present
func HasParameterSets(b []byte) bool { _ = "STUB: not implemented"; return false }

// GetParameterSets - get (multiple) VPS,  SPS, and PPS from a sample
func GetParameterSets(sample []byte) (vps, sps, pps [][]byte) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
