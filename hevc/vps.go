package hevc

// VPS is HEVC VPS parameters
// ISO/IEC 23008-2 (Ed. 5) Sec. 7.3.2.1 page 47 and 7.4.3.1 page 92
type VPS struct {
	VpsID                           byte
	BaseLayerInternalFlag           bool
	BaseLayerAvailableFlag          bool
	MaxLayersMinus1                 byte
	MaxSubLayersMinus1              byte
	TemporalIdNestingFlag           bool
	ProfileTierLevel                ProfileTierLevel
	SubLayerOrderingInfoPresentFlag bool
	MaxDecPicBufferingMinus1        []uint
	MaxNumReorderPics               []uint
	MaxLatencyIncreasePlus1         []uint
	MaxLayerID                      byte
	NumLayerSetsMinus1              uint
	TimingInfoPresentFlag           bool
	TimingInfo                      *VPSTimingInfo
	ExtensionFlag                   bool
}

// VPSTimingInfo contains VPS timing info parameters.
type VPSTimingInfo struct {
	NumUnitsInTick              uint32
	TimeScale                   uint32
	PocProportionalToTimingFlag bool
	NumTicksPocDiffOneMinus1    uint
	NumHrdParameters            uint
	HrdParameters               []*HrdParameters
}

// ParseVPSNALUnit parses HEVC VPS NAL unit starting with NAL unit header.
func ParseVPSNALUnit(data []byte) (*VPS, error) { _ = "STUB: not implemented"; return nil, nil }

// Note! First two bytes are NALU Header

// vps_reserved_0xffff_16bits

// layer_id_included_flag[i][j]

// hrd_layer_set_idx[i]
