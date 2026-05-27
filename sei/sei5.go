package sei

// UnregisteredSEI is SEI message of type 5.
type UnregisteredSEI struct {
	UUID    []byte
	payload []byte // raw rbsp payload including UUID
}

// Type returns SEI payload type.
func (s *UnregisteredSEI) Type() uint { _ = "STUB: not implemented"; return 0 }

// Size returns size in bytes of raw SEI message rbsp payload.
func (s *UnregisteredSEI) Size() uint { _ = "STUB: not implemented"; return 0 }

// String provides a short description of the SEI message.
func (s *UnregisteredSEI) String() string { _ = "STUB: not implemented"; return "" }

// Payload returns the SEI raw rbsp payload.
func (s *UnregisteredSEI) Payload() []byte {
	_ = "STUB: not implemented"

	// DecodeUserDataUnregisteredSEI decodes an unregistered SEI message (type 5).
	return nil
}

func DecodeUserDataUnregisteredSEI(sd *SEIData) (SEIMessage, error) {
	_ = "STUB: not implemented"
	return *new(SEIMessage), nil
}

// NewUnregisteredSEI creates an unregistered SEI message (type 5).
func NewUnregisteredSEI(sd *SEIData, uuid []byte) *UnregisteredSEI {
	_ = "STUB: not implemented"
	return nil
}
