package vvc

// NaluType - VVC NAL unit type according to ISO/IEC 23090-3 Table 5
type NaluType uint8

// VVC NAL unit types (0-31)
const (
	// VCL NAL unit types
	NALU_TRAIL      = NaluType(0)  // Coded slice of a trailing picture or subpicture
	NALU_STSA       = NaluType(1)  // Coded slice of an STSA picture or subpicture
	NALU_RADL       = NaluType(2)  // Coded slice of a RADL picture or subpicture
	NALU_RASL       = NaluType(3)  // Coded slice of a RASL picture or subpicture
	NALU_RSV_VCL_4  = NaluType(4)  // Reserved non-IRAP VCL NAL unit type
	NALU_RSV_VCL_5  = NaluType(5)  // Reserved non-IRAP VCL NAL unit type
	NALU_RSV_VCL_6  = NaluType(6)  // Reserved non-IRAP VCL NAL unit type
	NALU_IDR_W_RADL = NaluType(7)  // Coded slice of an IDR picture or subpicture
	NALU_IDR_N_LP   = NaluType(8)  // Coded slice of an IDR picture or subpicture
	NALU_CRA        = NaluType(9)  // Coded slice of a CRA picture or subpicture
	NALU_GDR        = NaluType(10) // Coded slice of a GDR picture or subpicture
	NALU_RSV_IRAP   = NaluType(11) // Reserved IRAP VCL NAL unit type

	// Non-VCL NAL unit types
	NALU_OPI         = NaluType(12) // Operating point information
	NALU_DCI         = NaluType(13) // Decoding capability information
	NALU_VPS         = NaluType(14) // Video parameter set
	NALU_SPS         = NaluType(15) // Sequence parameter set
	NALU_PPS         = NaluType(16) // Picture parameter set
	NALU_PREFIX_APS  = NaluType(17) // Adaptation parameter set
	NALU_SUFFIX_APS  = NaluType(18) // Adaptation parameter set
	NALU_PH          = NaluType(19) // Picture header
	NALU_AUD         = NaluType(20) // AU delimiter
	NALU_EOS         = NaluType(21) // End of sequence
	NALU_EOB         = NaluType(22) // End of bitstream
	NALU_SEI_PREFIX  = NaluType(23) // Supplemental enhancement information
	NALU_SEI_SUFFIX  = NaluType(24) // Supplemental enhancement information
	NALU_FD          = NaluType(25) // Filler data
	NALU_RSV_NVCL_26 = NaluType(26) // Reserved non-VCL NAL unit type
	NALU_RSV_NVCL_27 = NaluType(27) // Reserved non-VCL NAL unit type
	NALU_UNSPEC_28   = NaluType(28) // Unspecified non-VCL NAL unit type
	NALU_UNSPEC_29   = NaluType(29) // Unspecified non-VCL NAL unit type
	NALU_UNSPEC_30   = NaluType(30) // Unspecified non-VCL NAL unit type
	NALU_UNSPEC_31   = NaluType(31) // Unspecified non-VCL NAL unit type
)

func (n NaluType) String() string { _ = "STUB: not implemented"; return "" }

// NaluTypeName returns the name of the NAL unit type (backward compatibility)
func NaluTypeName(naluType uint8) string { _ = "STUB: not implemented"; return "" }

// NaluArray represents an array of NAL units of the same type
type NaluArray struct {
	NaluType NaluType
	Complete bool
	Nalus    [][]byte
}

// NewNaluArray creates a new NaluArray
func NewNaluArray(complete bool, naluType NaluType, nalus [][]byte) NaluArray {
	_ = "STUB: not implemented"
	return *new(NaluArray)
}

// NaluTypeName returns the NAL unit type name
func (n NaluArray) NaluTypeName() string { _ = "STUB: not implemented"; return "" }

// NaluHeader is VVC NAL unit header
type NaluHeader struct {
	NuhLayerId         uint8 // NAL unit header layer ID
	NaluType           NaluType
	NuhTemporalIdPlus1 uint8 // NAL unit header temporal ID plus 1
}

// ParseNaluHeader parses the NAL unit header from raw bytes
func ParseNaluHeader(rawBytes []byte) (NaluHeader, error) {
	_ = "STUB: not implemented"
	return *new(NaluHeader), nil
}

// 6 bits for layer ID
// 5 bits for NALU type
// 3 bits for temporal ID plus 1
