package sei

// MasteringDisplayColourVolumeSEI is HEVC SEI Message 137.
// Defined in ISO/IEC 23008-2 D.2.28
type MasteringDisplayColourVolumeSEI struct {
	DisplayPrimariesX            [3]uint16
	DisplayPrimariesY            [3]uint16
	WhitePointX                  uint16
	WhitePointY                  uint16
	MaxDisplayMasteringLuminance uint32
	MinDisplayMasteringLuminance uint32
}

func (m MasteringDisplayColourVolumeSEI) Type() uint { _ = "STUB: not implemented"; return 0 }

func (m MasteringDisplayColourVolumeSEI) Size() uint { _ = "STUB: not implemented"; return 0 }

func (m MasteringDisplayColourVolumeSEI) Payload() []byte { _ = "STUB: not implemented"; return nil }

func (m MasteringDisplayColourVolumeSEI) String() string { _ = "STUB: not implemented"; return "" }

// DecodeUserDataUnregisteredSEI - Decode an unregistered SEI message (type 5)
func DecodeMasteringDisplayColourVolumeSEI(sd *SEIData) (SEIMessage, error) {
	_ = "STUB: not implemented"
	return *new(SEIMessage), nil
}
