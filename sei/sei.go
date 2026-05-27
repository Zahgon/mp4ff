package sei

import (
	"errors"
	"io"
)

var ErrRbspTrailingBitsMissing = errors.New("rbsp_trailing_bits byte 0x80 is missing")

const (
	// Combined definition of AVC and HEVC messages. When names collide for the same number, AVC or HEVC is added in the name.
	// AVC defined in ISO/IEC 14496-10:2020 Annex D and HEVC in ISO/IEC 23008-2:2020 Annex D.
	// SEIBufferingPeriodType is defined in AVC D.1.2 and HEVC D.2.2. Definitions differ.
	SEIBufferingPeriodType = 0
	// SEIPicTimingType is defined in AVC D.1.3 and HEVC D.2.3. Definitions differ.
	SEIPicTimingType = 1
	// SEIPanScanRectType is defined in AVC D.1.4 and HEVC D.2.4. Definitions differ.
	SEIPanScanRectType = 2
	// SEIFillerPayloadType is defined in AVC D.1.5 and HEVC D.2.5. Definitions agree.
	SEIFillerPayloadType = 3
	// SEIUserDataRegisteredITUtT35Type is defined in AVC D.1.6 and HEVC D.2.6. Definitions agree.
	SEIUserDataRegisteredITUtT35Type = 4
	// SEIUserDataUnregisteredType is defined in AVC D.1.7 and HEVC D.2.7. Definitions agree.
	SEIUserDataUnregisteredType = 5
	// SEIRecoveryPointType is defined in AVC D.1.8 and HEVC D.2.8. Definitions differ.
	SEIRecoveryPointType = 6
	// SEIDecRefPicMarkingRepetitionType is defined in AVC D.1.9. Absent in HEVC.
	SEIDecRefPicMarkingRepetitionType = 7
	// SEISparePicType is defined in AVC D.1.10. Absent in HEVC.
	SEISparePicType = 8
	// SEISceneInfoType is defined in AVC D.1.11 and HEVC 2.9. Definitions differ.
	SEISceneInfoType = 9
	// SEISubSeqInfoType is defined in AVC D.1.12. Absent in HEVC.
	SEISubSeqInfoType = 10
	// SEISubSeqLayerCharacteristicsType is defined in AVC D.1.13. Absent in HEVC.
	SEISubSeqLayerCharacteristicsType = 11
	// SEISubSeqCharacteristicsType is defined in AVC D.1.14. Absent in HEVC.
	SEISubSeqCharacteristicsType = 12
	// SEIFullFrameFreezeType is defined in AVC D.1.15. Absent in HEVC.
	SEIFullFrameFreezeType = 13
	// SEIFullFrameFreezeReleaseType is defined in AVC D.1.16. Absent in HEVC.
	SEIFullFrameFreezeReleaseType = 14
	// SEIPictureSnapShotType is defined in AVC D.1.17 and HEVC D.2.10. Definitions agree, but called FullFrameSnapshot in AVC.
	SEIPictureSnapShotType = 15
	// SEIProgressiveRefinementSegmentStartType is defined in AVC D.1.18 and HEVC D.2.11. Definitions differ.
	SEIProgressiveRefinementSegmentStartType = 16
	// SEIProgressiveRefinementSegmentStartEnd is defined in AVC D.1.19 and HEVC D.2.12. Definitions agree.
	SEIProgressiveRefinementSegmentStartEnd = 17 // AVC and HEVC. Same definition
	// SEIMotionConstrainedSliceGroupSetType is defined in AVC D.1.20. Absent in HEVC.
	SEIMotionConstrainedSliceGroupSetType = 18
	// SEIFilmGrainCharacteristicsType is defined in AVC D.1.21 and HEVC D.2.13. Definitions differ.
	SEIFilmGrainCharacteristicsType = 19
	// SEIDeblockingFilterDisplayPreferenceType is defined in AVC D.1.22. Absent in HEVC.
	SEIDeblockingFilterDisplayPreferenceType = 20
	// SEIStereoVideoInfoType is defined in AVC D.1.23. Absent in HEVC.
	SEIStereoVideoInfoType = 21
	// SEIPostFilterHintType is defined in AVC D.1.24 and HEVC D.2.14. Definitions differ.
	SEIPostFilterHintType = 22
	// SEIToneMappingInfoType is defined in AVC D.1.25 and HEVC D.15. Definitions agree.
	SEIToneMappingInfoType = 23 // AVC and HEVC. Same definition
	// SEIScalabilityInfoType is defined in AVC Annex F. Absent in HEVC.
	SEIScalabilityInfoType = 24
	// SEISubPicScalableLayerType is defined in AVC Annex F. Absent in HEVC.
	SEISubPicScalableLayerType = 25
	// SEINonRequiredLayerRepType is defined in AVC Annex F. Absent in HEVC.
	SEINonRequiredLayerRepType = 26
	// SEIPriorityLayerInfoType is defined in AVC Annex F. Absent in HEVC.
	SEIPriorityLayerInfoType = 27
	// SEILayersNotPresentAVCType is defined in AVC Annex F. Absent in HEVC.
	SEILayersNotPresentAVCType = 28
	// SEILayerDependencyChangeType is defined in AVC Annex F. Absent in HEVC.
	SEILayerDependencyChangeType = 29
	// SEIScalableNestingAVCType is defined in AVC Annex F. Absent in HEVC.
	SEIScalableNestingAVCType = 30
	// SEIBaseLayerTemporalHrdType is defined in AVC Annex F. Absent in HEVC.
	SEIBaseLayerTemporalHrdType = 31
	// SEIQualityLayerIntegrityCheckTpe is defined in AVC Annex F. Absent in HEVC.
	SEIQualityLayerIntegrityCheckTpe = 32
	// SEIRedundantPicPropertyType is defined in AVC Annex F. Absent in HEVC.
	SEIRedundantPicPropertyType = 33
	// SEITl0DepRepIndexType is defined in AVC Annex F. Absent in HEVC.
	SEITl0DepRepIndexType = 34
	// SEITlSwitchingPointType is defined in AVC Annex F. Absent in HEVC.
	SEITlSwitchingPointType = 35
	// SEIParallelDecodingInfoType is defined in AVC Annex G. Absent in HEVC.
	SEIParallelDecodingInfoType = 36
	// SEIMVCScalableNestingType is defined in AVC Annex G. Absent in HEVC.
	SEIMVCScalableNestingType = 37
	// SEIViewScalabilityInfoType is defined in AVC Annex G. Absent in HEVC.
	SEIViewScalabilityInfoType = 38
	// SEIMultiviewSceneInfoAVCType is defined in AVC Annex G. Absent in HEVC.
	SEIMultiviewSceneInfoAVCType = 39
	// SEIMultiviewAcquisitionInfoAVCType is defined in AVC Annex G. Absent in HEVC.
	SEIMultiviewAcquisitionInfoAVCType = 40
	// SEINonRequiredViewComponentType is defined in AVC Annex G. Absent in HEVC.
	SEINonRequiredViewComponentType = 41
	// SEIViewDependencyChangeType is defined in AVC Annex G. Absent in HEVC.C.
	SEIViewDependencyChangeType = 42
	// SEIOperationPointsNotPresentType is defined in AVC Annex G. Absent in HEVC.
	SEIOperationPointsNotPresentType = 43
	// SEIBaseViewTemporalHrdType is defined in AVC Annex G. Absent in HEVC.
	SEIBaseViewTemporalHrdType = 44
	// SEIFramePackingArrangementType is defined in AVC D.1.26 and HEVC D.2.16. Definitions differ.
	SEIFramePackingArrangementType = 45
	// SEIMultiviewViewPositionAVCType is defined in AVC Annex G. Absent in HEVC.
	SEIMultiviewViewPositionAVCType = 46
	// SEIDisplayOrientationType is defined in AVC D.1.27 and HEVC D.2.17. Definitions differ.
	SEIDisplayOrientationType = 47
	// SEIMvcdScalableNestingType is defined in AVC Annex H. Absent in HEVC.
	SEIMvcdScalableNestingType = 48
	// SEIMvcdViewScalabilityInfoType is defined in AVC Annex H. Absent in HEVC.
	SEIMvcdViewScalabilityInfoType = 49
	// SEIDepthRepresentationInfoAVCType is defined in AVC Annex H. Absent in HEVC.
	SEIDepthRepresentationInfoAVCType = 50
	// SEIThreeDimensionalReferenceDisplaysInfoAVCType is defined in AVC Annex H. Absent in HEVC.
	SEIThreeDimensionalReferenceDisplaysInfoAVCType = 51
	// SEIDepthTimingType is defined in AVC Annex H. Absent in HEVC.
	SEIDepthTimingType = 52
	// SEIDepthSamplingInfoType is defined in AVC Annex H. Absent in HEVC.
	SEIDepthSamplingInfoType = 53
	// SEIConstrainedDepthParameterSetIdentifierType is defined in AVC Annex H. Absent in HEVC.
	SEIConstrainedDepthParameterSetIdentifierType = 54
	// SEIGreenMetaDataType is defined in AVC D.1.28 and HEVC D.2.17. Definitions agree and point to ISO/IEC 23001-11.
	SEIGreenMetaDataType = 56
	// SEIStructureOfPicturesInfoType is defined in HEVC D.2.19. Absent in AVC.
	SEIStructureOfPicturesInfoType = 128
	// SEIActiveParameterSetsType is defined in HEVC D.2.21. Absent in AVC.
	SEIActiveParameterSetsType = 129
	// SEIDecodingUnitInfoType is defined in HEVC D.2.22. Absent in AVC.
	SEIDecodingUnitInfoType = 130
	// SEITemporalSubLayerZeroIndexType is defined in HEVC D.2.23. Absent in AVC.
	SEITemporalSubLayerZeroIndexType = 131
	// SEIDecodedPictureHashType is defined in HEVC D.2.20. Only used as suffix SEI. Absent in AVC.
	SEIDecodedPictureHashType = 132
	// SEIScalableNestingHEVCType is defined in HEVC D.2.24. Absent in AVC.
	SEIScalableNestingHEVCType = 133
	// SEIRegionRefreshInfoType is defined in HEVC D.2.25. Absent in AVC.
	SEIRegionRefreshInfoType = 134
	// SEINoDisplayType is defined in HEVC D.2.26. Absent in AVC.
	SEINoDisplayType = 135
	// SEITimeCodeType is defined in HEVC D.2.27. Absent in AVC.
	SEITimeCodeType = 136
	// SEIMasteringDisplayColourVolumeType is defined in AVC D.1.29 and HEVC D.2.28. Definitions agree.
	SEIMasteringDisplayColourVolumeType = 137
	// SEISegmentedRectFramePackingArrangementType is defined in HEVC D.2.29. Absent in AVC.
	SEISegmentedRectFramePackingArrangementType = 138
	// SEITemporalMotionConstrainedTileSetsType is defined in HEVC D.2.30. Absent in AVC.
	SEITemporalMotionConstrainedTileSetsType = 139
	// SEIChromaResamplingFilterHintType is defined in HEVC D.2.31. Absent in AVC.
	SEIChromaResamplingFilterHintType = 140
	// SEIKneeFunctionInfoType is defined in HEVC D.2.32. Absent in AVC.
	SEIKneeFunctionInfoType = 141
	// SEIColourRemappingInfoType is defined in AVC D.1.30 and HEVC D.2.33. Definitions differ.
	SEIColourRemappingInfoType = 142
	// SEIDeinterlacedFieldIdentificationType is defined in HEVC D.2.34. Absent in AVC.
	SEIDeinterlacedFieldIdentificationType = 143
	// SEIContentLightLevelInformationType is defined in AVC D.1.31 and HEVC D.2.35. Definitions agree.
	SEIContentLightLevelInformationType = 144
	// SEIDependentRapIndicationType is defined in HEVC D.2.36. Absent in AVC.
	SEIDependentRapIndicationType = 145
	// SEICodedRegionCompletionType is defined in HEVC D.2.37. Absent in AVC.
	SEICodedRegionCompletionType = 146
	// SEIAlternativeTransferCharacteristicsType is defined in AVC D.1.32 and HEVC D.2.38. Definitions agree.
	SEIAlternativeTransferCharacteristicsType = 147
	// SEIAmbientViewingEnvironmentType is defined in HEVC D.2.39. Absent in AVC.
	SEIAmbientViewingEnvironmentType = 148
	// SEIContentColourVolumeType is defined in AVC D.1.33 and HEVC D.2.40. Definitions agree.
	SEIContentColourVolumeType = 149
	// SEIEquirectangularProjectionType is defined in AVC D.1.35.1 and HEVC D.2.41.1. Definitions agree.
	SEIEquirectangularProjectionType = 150
	// SEICubemapProjectionType is defined in AVC D.1.35.2 and HEVC D.2.41.2. Definitions agree.
	SEICubemapProjectionType = 151
	// SEIFisheyeVideoInfoType is defined in HEVC D.41.3. Absent in AVC.
	SEIFisheyeVideoInfoType = 152
	// SEISphereRotationType is defined in AVC D.1.35.3 and HEVC D.2.41.4. Definitions agree.
	SEISphereRotationType = 154
	// SEIRegionwisePackingType is defined in AVC D.1.35.4 and HEVC D.2.41.5. Definitions agree.
	SEIRegionwisePackingType = 155
	// SEIOmniViewportType is defined in AVC D.1.35.5 and HEVC D.2.41.6. Definitions agree.
	SEIOmniViewportType = 156
	// SEIRegionalNestingType is defined in HEVC D.2.42. Absent in AVC.
	SEIRegionalNestingType = 157
	// SEIMctsExtractionInfoSetsType is defined in HEVC D.2.43. Absent in AVC.
	SEIMctsExtractionInfoSetsType = 158
	// SEIMctsExtractionInfoNesting is defined in HEVC D.2.44. Absent in AVC.
	SEIMctsExtractionInfoNesting = 159
	// SEILayersNotPresentHEVCType is defined in HEVC Annex F. Absent in AVC.
	SEILayersNotPresentHEVCType = 160
	// SEIInterLayerConstrainedTileSetsType is defined in HEVC Annex F. Absent in AVC.
	SEIInterLayerConstrainedTileSetsType = 161
	// SEIBspNestingType is defined in HEVC Annex F. Absent in AVC.
	SEIBspNestingType = 162
	// SEIBspInitialArrivalTimeType is defined in HEVC Annex F. Absent in AVC.
	SEIBspInitialArrivalTimeType = 163
	// SEISubBitstreamPropertyType is defined in HEVC Annex F. Absent in AVC.
	SEISubBitstreamPropertyType = 164
	// SEIAlphaChannelInfoType is defined in HEVC Annex F. Absent in AVC.
	SEIAlphaChannelInfoType = 165
	// SEIOverlayInfoType is defined in HEVC Annex F. Absent in AVC.
	SEIOverlayInfoType = 166
	// SEITemporalMvPredictionConstraintsType is defined in HEVC Annex F. Absent in AVC.
	SEITemporalMvPredictionConstraintsType = 167
	// SEIFrameFieldInfoType is defined in HEVC Annex F. Absent in AVC.
	SEIFrameFieldInfoType = 168
	// SEIThreeDimensionalReferenceDisplaysInfoHEVCType is defined in HEVC Annex G. Absent in AVC.
	SEIThreeDimensionalReferenceDisplaysInfoHEVCType = 176
	// SEIDepthRepresentationInfoHEVCType is defined in HEVC Annex G. Absent in AVC.
	SEIDepthRepresentationInfoHEVCType = 177
	// SEIMultiviewSceneInfoHEVCType is defined in HEVC Annex G. Absent in AVC.
	SEIMultiviewSceneInfoHEVCType = 178
	// SEIMultiviewAcquisitionInfoHEVCType is defined in HEVC Annex G. Absent in AVC.
	SEIMultiviewAcquisitionInfoHEVCType = 179
	// SEIMultiviewViewPositionHEVCType is defined in HEVC Annex G. Absent in AVC.
	SEIMultiviewViewPositionHEVCType = 180
	// SEIAlternativeDepthInfoType is defined in HEVC Annex I. Absent in AVC.
	SEIAlternativeDepthInfoType = 181
	// SEISeiManifestType is defined in AVC D.1.36 and HEVC D.2.45. Definitions agree.
	SEISeiManifestType = 200
	// SEISeiPrefixIndicationType is defined in HEVC D.2.46. Absent in AVC.
	SEISeiPrefixIndicationType = 201
	// SEIAnnotatedRegionsType is defined in HEVC D.2.47. Absent in AVC.
	SEIAnnotatedRegionsType = 202
)

// SEIType is SEI payload type in AVC or HEVC.
type SEIType uint

// String provides the camel-case name for the SEIType.
func (h SEIType) String() string { _ = "STUB: not implemented"; return "" }

type Codec uint

const (
	AVC Codec = iota
	HEVC
)

// SEI is Supplementary Enhancement Information.
// High level syntax in ISO/IEC 14496-10 Section 7.3.2.3.
// The actual types are listed in Annex D.
type SEI struct {
	SEIMessages []SEIMessage
}

// SEIMessage is common part of any SEI message.
type SEIMessage interface {
	Type() uint
	Size() uint
	String() string
	Payload() []byte
}

// DecodeSEIMessage decodes or at least provides some information about an SEIMessage.
func DecodeSEIMessage(sd *SEIData, codec Codec) (SEIMessage, error) {
	_ = "STUB: not implemented"
	return *new(SEIMessage), nil
}

// DecodeGeneralSEI is a fallback decoder for non-implemented SEI message types.
func DecodeGeneralSEI(sd *SEIData) SEIMessage { _ = "STUB: not implemented"; return *new(SEIMessage) }

// SEIData is raw parsed SEI message including payload rbsp data.
type SEIData struct {
	payloadType uint
	payload     []byte
}

// NewSEIData returns SEIData struct.
func NewSEIData(msgType uint, payload []byte) *SEIData { _ = "STUB: not implemented"; return nil }

// Type returns the SEI payload type.
func (s *SEIData) Type() uint { _ = "STUB: not implemented"; return 0 }

// Payload returns the SEI raw rbsp payload.
func (s *SEIData) Payload() []byte {
	_ = "STUB: not implemented"

	// String provides a description of the SEI message.
	return nil
}

func (s *SEIData) String() string { _ = "STUB: not implemented"; return "" }

// Size is the size in bytes of the raw SEI message rbsp payload.
func (s *SEIData) Size() uint { _ = "STUB: not implemented"; return 0 }

// ExtractSEIData parses ebsp (after NALU header) and returns a slice of SEIData in rbsp format.
// In case the rbsp_trailing_bits 0x80 byte is missing at end, []seiData and
// an ErrMissingRbspTrailingBits error are both returned.
func ExtractSEIData(r io.ReadSeeker) (seiData []SEIData, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Break loop if no more rbsp data (end of sei messages)

// WriteSEIMessages writes the messages in EBSP format with RBSPTrailing bits.
// The output corresponds to an SEI NAL unit payload.
func WriteSEIMessages(w io.Writer, msgs []SEIMessage) error { _ = "STUB: not implemented"; return nil }
