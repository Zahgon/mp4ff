package sei

// DecodeUserDataRegisteredSEI decodes a SEI message of type 4.
func DecodeUserDataRegisteredSEI(sd *SEIData) (SEIMessage, error) {
	_ = "STUB: not implemented"
	return *new(SEIMessage), nil
}

// ITUData identifies registered payload in SEI of type 4 (User data registered by ITU-T Rec T 35).
type ITUData struct {
	CountryCode      byte
	UserDataTypeCode byte
	ProviderCode     uint16
	UserIdentifier   uint32
}

// IsCEA608 checks if ITU-T data corresponds to CEA-608.
func (i ITUData) IsCEA608() bool { _ = "STUB: not implemented"; return false }

// RegisteredSEI is user_data_registered_itu_t_t35 (type 4) SEI message.
type RegisteredSEI struct {
	payload  []byte
	ITUTData ITUData
}

// NewRegisteredSEI creates an ITU-T registered SEI message (type 4).
func NewRegisteredSEI(sd *SEIData, ituData ITUData) *RegisteredSEI {
	_ = "STUB: not implemented"
	return nil
}

// Type returns the SEI payload type.
func (s *RegisteredSEI) Type() uint { _ = "STUB: not implemented"; return 0 }

// Size returns size in bytes of raw SEI message rbsp payload.
func (s *RegisteredSEI) Size() uint { _ = "STUB: not implemented"; return 0 }

// String provides a short description of the SEI message.
func (s *RegisteredSEI) String() string { _ = "STUB: not implemented"; return "" }

// Payload returns the SEI raw rbsp payload.
func (s *RegisteredSEI) Payload() []byte {
	_ = "STUB: not implemented"

	// ExtractCEA608sei returns payload and parsed field for CEA 608 SEI message.
	// CEA-608 encapsulation in SEI nal unit is defined in ATSC-120 and further
	// in CTA-708 specification (previously CEA-708).
	return nil
}

func ExtractCEA608sei(sd *SEIData) (*CEA608sei, error) { _ = "STUB: not implemented"; return nil, nil }

// CEA608sei data structure.
type CEA608sei struct {
	payload []byte // full raw payload
	Field1  []byte
	Field2  []byte
}

// Type returns the SEI payload type.
func (s *CEA608sei) Type() uint { _ = "STUB: not implemented"; return 0 }

// Size is size in bytes of raw SEI message rbsp payload.
func (s *CEA608sei) Size() uint { _ = "STUB: not implemented"; return 0 }

// String provides a simple representation of the CEA608 data.
func (s *CEA608sei) String() string { _ = "STUB: not implemented"; return "" }

// Payload returns the SEI raw rbsp payload.
func (s *CEA608sei) Payload() []byte {
	_ = "STUB: not implemented"

	// ParseCEA608 parsers the the fields of data from CEA-708 encapsulation.
	// This is specified in Section 4.3 of ANSI/CTA-708-E R-2018.
	return nil
}

func ParseCEA608(payload []byte) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Advance 1 and skip reserved byte

// Keep parity bit

// Keep parity bit

//Check validity and non-empty data

// There should also be a 0xff marker bits byte before the end of the NALU
