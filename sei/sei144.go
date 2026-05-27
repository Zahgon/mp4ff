package sei

// ContentLightLevelInformationSEI is HEVC SEI Message 144.
// Defined in ISO/IEC 23008-2 D.2.35
type ContentLightLevelInformationSEI struct {
	MaxContentLightLevel    uint16
	MaxPicAverageLightLevel uint16
}

func (c ContentLightLevelInformationSEI) Type() uint { _ = "STUB: not implemented"; return 0 }

func (c ContentLightLevelInformationSEI) Size() uint { _ = "STUB: not implemented"; return 0 }

func (c ContentLightLevelInformationSEI) Payload() []byte { _ = "STUB: not implemented"; return nil }

func (c ContentLightLevelInformationSEI) String() string { _ = "STUB: not implemented"; return "" }

// DecodeContentLightLevelInformationSEI decodes HEVC SEI 144.
func DecodeContentLightLevelInformationSEI(sd *SEIData) (SEIMessage, error) {
	_ = "STUB: not implemented"
	return *new(SEIMessage), nil
}
