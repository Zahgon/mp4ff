package sei

// PicTimingHevcSEI carries the data of an SEI 1 PicTiming message for HEVC.
// The corresponding SEI 1 for AVC is very different. Time code is in SEI 136 for HEVC.
// Defined in ISO/IEC 23008-2 Ed 5. Section D.2.3 (page 372) and D.3.2.3 (page 405)
type PicTimingHevcSEI struct {
	ExternalParams                         HEVCPicTimingParams `json:"-"`
	FrameFieldInfo                         *HEVCFrameFieldInfo `json:"FrameFieldInfo,omitempty"`
	AuCpbRemovalDelayMinus1                uint32              `json:"AuCpbRemovalDelayMinus1,omitempty"`
	PicDpbOutputDelay                      uint32              `json:"PicDpbOutputDelay,omitempty"`
	PicDpbOutputDuDelay                    uint32              `json:"PicDpbOutputDuDelay,omitempty"`
	NumDecodingUnitsMinus1                 uint32              `json:"NumDecodingUnitsMinus1,omitempty"`
	DuCommonCpbRemovalDelayFlag            bool                `json:"DuCommonCpbRemovalDelayFlag,omitempty"`
	DuCommonCpbRemovalDelayIncrementMinus1 uint32              `json:"DuCommonCpbRemovalDelayIncrementMinus1,omitempty"`
	NumNalusInDuMinus1                     []uint32            `json:"NumNalusInDuMinus1,omitempty"`
	DuCpbRemovalDelayIncrementMinus1       []uint32            `json:"DuCpbRemovalDelayIncrementMinus1,omitempty"`
	payload                                []byte              `json:"-"`
}

type HEVCPicTimingParams struct {
	FrameFieldInfoPresentFlag              bool
	CpbDpbDelaysPresentFlag                bool
	SubPicHrdParamsPresentFlag             bool
	SubPicCpbParamsInPicTimingSeiFlag      bool
	AuCbpRemovalDelayLengthMinus1          uint8
	DpbOutputDelayLengthMinus1             uint8
	DpbOutputDelayDuLengthMinus1           uint8
	DuCpbRemovalDelayIncrementLengthMinus1 uint8
}

type HEVCFrameFieldInfo struct {
	PicStruct      uint8 // 4 bits
	SourceScanType uint8 // 2 bits
	DuplicateFlag  bool  `json:"DuplicateFlag,omitempty"` // 1bit
}

func DecodePicTimingHevcSEI(sd *SEIData, exPar HEVCPicTimingParams) (SEIMessage, error) {
	_ = "STUB: not implemented"
	return *new(SEIMessage), nil
}

// Type returns the SEI payload type.
func (s *PicTimingHevcSEI) Type() uint { _ = "STUB: not implemented"; return 0 }

// Payload returns the SEI raw rbsp payload.
func (s *PicTimingHevcSEI) Payload() []byte {
	_ = "STUB: not implemented"

	// String returns string representation of PicTiming SEI1.
	return nil
}

func (s *PicTimingHevcSEI) String() string { _ = "STUB: not implemented"; return "" }

// Size is size in bytes of raw SEI message rbsp payload.
func (s *PicTimingHevcSEI) Size() uint { _ = "STUB: not implemented"; return 0 }
