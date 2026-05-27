package iamf

import (
	"github.com/Eyevinn/mp4ff/bits"
)

/**
 * Types were extracted from the ffmpeg implementation.
 * Based on IAMF_C
 */

type channelMask uint64

// define masks for channel layouts
const (
	chmFrontLeft channelMask = 1 << iota
	chmFrontRight
	chmFrontCenter
	chmLowFrequency
	chmBackLeft
	chmBackRight
	chmFrontLeftOfCenter
	chmFrontRightOfCenter
	chmBackCenter
	chmSideLeft
	chmSideRight
	chmTopCenter
	chmTopFrontLeft
	chmTopFrontCenter
	chmTopFrontRight
	chmTopBackLeft
	chmTopBackCenter
	chmTopBackRight
	chmLowFrequency2
	chmTopSideLeft
	chmTopSideRight
	chmBottomFrontCenter
	chmBottomFrontLeft
	chmBottomFrontRight
	chmSideSurroundLeft
	chmSideSurroundRight
	chmBinauralLeft
	chmBinauralRight
	chmBottomBackLeft
	chmBottomBackRight
)

type channelOrder uint8

// define constants for audio channel order
const (
	coScalable channelOrder = iota
	coExpanded
	coAmbisonics
	coCustom
)

// channelLayout defines a struct for scalable audio channel layouts
type channelLayout struct {
	TableIndex int
	Channels   int
	Order      channelOrder
	Mask       channelMask
	Map        *map[int]int
}

func scalable(index int, channels int, mask channelMask) channelLayout {
	_ = "STUB: not implemented"
	return *new(channelLayout)
}

// scalableChannelLayouts defines an array of channel layouts
// Based on IAMF loudspeaker_layout specification
// https://aomediacodec.github.io/iamf/#loudspeaker_layout
var scalableChannelLayouts = []channelLayout{
	// 0: Mono (C)
	// the mono channel
	scalable(0, 1,
		chmFrontCenter),

	// 1: Stereo (L/R)
	// the config of (0+2+0) of [ITU-2051-3] (Sound System A)
	scalable(1, 2,
		chmFrontLeft|chmFrontRight),

	// 2: 5.1ch (L/C/R/Ls/Rs/LFE)
	// the config of (0+5+0) of [ITU-2051-3] (Sound System B)
	scalable(2, 6, 0|
		chmFrontLeft|chmFrontCenter|chmFrontRight|
		chmSideLeft|chmSideRight|chmLowFrequency),

	// 3: 5.1.2ch (L/C/R/Ls/Rs/Ltf/Rtf/LFE)
	// the config of (2+5+0) of [ITU-2051-3] (Sound System C)
	scalable(3, 8, 0|
		chmFrontLeft|chmFrontCenter|chmFrontRight|
		chmSideLeft|chmSideRight|
		chmTopFrontLeft|chmTopFrontRight|
		chmLowFrequency),

	// 4: 5.1.4ch (L/C/R/Ls/Rs/Ltf/Rtf/Ltr/Rtr/LFE)
	// the config of (4+5+0) of [ITU-2051-3] (Sound System D)
	scalable(4, 10, 0|
		chmFrontLeft|chmFrontCenter|chmFrontRight|
		chmSideLeft|chmSideRight|
		chmTopFrontLeft|chmTopFrontRight|
		chmTopBackLeft|chmTopBackRight|
		chmLowFrequency),

	// 5: 7.1ch (L/C/R/Lss/Rss/Lrs/Rrs/LFE)
	// the config of (0+7+0) of [ITU-2051-3] (Sound System I)
	scalable(5, 8, 0|
		chmFrontLeft|chmFrontCenter|chmFrontRight|
		chmSideSurroundLeft|chmSideSurroundRight|
		chmBackLeft|chmBackRight|
		chmLowFrequency),

	// 6: 7.1.2ch (L/C/R/Lss/Rss/Lrs/Rrs/Ltf/Rtf/LFE)
	// The combination of 7.1ch and the Left and Right top front pair of 7.1.4ch
	scalable(6, 10, 0|
		chmFrontLeft|chmFrontCenter|chmFrontRight|
		chmSideSurroundLeft|chmSideSurroundRight|
		chmBackLeft|chmBackRight|
		chmTopFrontLeft|chmTopFrontRight|
		chmLowFrequency),

	// 7: 7.1.4ch (L/C/R/Lss/Rss/Lrs/Rrs/Ltf/Rtf/Ltb/Rtb/LFE)
	// the config of (4+7+0) of [ITU-2051-3] (Sound System J)
	scalable(7, 12, 0|
		chmFrontLeft|chmFrontCenter|chmFrontRight|
		chmSideSurroundLeft|chmSideSurroundRight|
		chmBackLeft|chmBackRight|
		chmTopFrontLeft|chmTopFrontRight|
		chmTopBackLeft|chmTopBackRight|
		chmLowFrequency),

	// 8: 3.1.2ch (L/C/R/Ltf/Rtf/LFE)
	// The front subset of 7.1.4ch (Sound System J)
	scalable(8, 6, 0|
		chmFrontLeft|chmFrontCenter|chmFrontRight|
		chmTopFrontLeft|chmTopFrontRight|
		chmLowFrequency),

	// 9: Binaural (L/R)
	// the binaural channels
	scalable(9, 2,
		chmBinauralLeft|chmBinauralRight),

	// 10-14: Reserved for future use

	// 15: Expanded channel layouts - defined in expanded_loudspeaker_layout field bellow
}

func expanded(index int, channels int, mask channelMask) channelLayout {
	_ = "STUB: not implemented"
	return *new(channelLayout)
}

// expandedScalableChannelLayouts defines an array of expanded scalable channel layouts
// Based on IAMF expanded_loudspeaker_layout specification
// https://aomediacodec.github.io/iamf/#expanded_loudspeaker_layout
// https://www.itu.int/dms_pubrec/itu-r/rec/bs/R-REC-BS.2051-3-202205-I!!PDF-E.pdf
var expandedScalableChannelLayouts = []channelLayout{

	// 0: LFE - The low-frequency effects subset (LFE) of 7.1.4ch (Sound System J)
	expanded(0, 1,
		chmLowFrequency),

	// 1: Stereo-S (Ls/Rs) - The surround subset of 5.1.4ch (Sound System I)
	expanded(1, 2,
		chmSideLeft|chmSideRight),

	// 2: Stereo-SS (Lss/Rss) - The side surround subset of 7.1.4ch (Sound System J)
	// it is wrong in ffmpeg
	expanded(2, 2,
		chmSideSurroundLeft|chmSideSurroundRight),

	// 3: Stereo-RS (Lrs/Rrs) - The rear surround subset of 7.1.4ch (Sound System J)
	expanded(3, 2,
		chmBackLeft|chmBackRight),

	// 4: Stereo-TF (Ltf/Rtf) - The top front subset of 7.1.4ch (Sound System J)
	expanded(4, 2,
		chmTopFrontLeft|chmTopFrontRight),

	// 5: Stereo-TB (Ltb/Rtb) - The top back subset of 7.1.4ch (Sound System J)
	expanded(5, 2,
		chmTopBackLeft|chmTopBackRight),

	// 6: Top-4ch (Ltf/Rtf/Ltb/Rtb) - The top 4 channels of 7.1.4ch (Sound System J)
	expanded(6, 4,
		chmTopFrontLeft|chmTopFrontRight|chmTopBackLeft|chmTopBackRight),

	// 7: 3.0ch (L/C/R) - The front 3 channels of 7.1.4ch (Sound System J)
	expanded(7, 3,
		chmFrontLeft|chmFrontCenter|chmFrontRight),

	// 8: 9.1.6ch - FLc/FC/FRc/FL/FR/SiL/SiR/BL/BR/
	//              TpFL/TpFR/TpSiL/TpSiR/TpBL/TpBR/LFE1
	// The subset of (9+10+3) of [ITU-2051-3] (Sound System H)
	expanded(8, 16, 0|
		chmFrontLeftOfCenter|chmFrontCenter|chmFrontRightOfCenter|
		chmFrontLeft|chmFrontRight|
		chmSideSurroundLeft|chmSideSurroundRight|
		chmBackLeft|chmBackRight|
		chmTopFrontLeft|chmTopFrontRight|
		chmTopSideLeft|chmTopSideRight|
		chmTopBackLeft|chmTopBackRight|
		chmLowFrequency),

	// 9: Stereo-F (FL/FR) - The front subset of 9.1.6ch (Sound System H)
	expanded(9, 2,
		chmFrontLeft|chmFrontRight),

	// 10: Stereo-Si (SiL/SiR) - The side subset of 9.1.6ch (Sound System H)
	expanded(10, 2,
		chmSideSurroundLeft|chmSideSurroundRight),

	// 11: Stereo-TpSi (TpSiL/TpSiR) - The top side subset of 9.1.6ch (Sound System H)
	expanded(11, 2,
		chmTopSideLeft|chmTopSideRight),

	// 12: Top-6ch (TpFL/TpFR/TpSiL/TpSiR/TpBL/TpBR)
	// The top 6 channels of 9.1.6ch (Sound System H)
	expanded(12, 6, 0|
		chmTopFrontLeft|chmTopFrontRight|
		chmTopSideLeft|chmTopSideRight|
		chmTopBackLeft|chmTopBackRight),

	// ffmpeg table is missing the following ones

	// 13: 10.2.9.3ch - FLc/FC/FRc/FL/FR/SiL/SiR/BL/BC/BR/
	//                  TpFL/TpFC/TpFR/TpSiL/TpC/TpSiR/
	//                  TpBL/TpBC/TpBR/BtFL/BtFC/BtFR/LFE1/LFE2
	// The subset of (9+10+3) of [ITU-2051-3] (Sound System H)
	expanded(13, 24, 0|
		chmFrontLeftOfCenter|chmFrontCenter|chmFrontRightOfCenter|
		chmFrontLeft|chmFrontRight|
		chmSideSurroundLeft|chmSideSurroundRight|
		chmBackLeft|chmBackCenter|chmBackRight|
		chmTopFrontLeft|chmTopFrontCenter|chmTopFrontRight|
		chmTopSideLeft|chmTopCenter|chmTopSideRight|
		chmTopBackLeft|chmTopBackCenter|chmTopBackRight|
		chmBottomFrontLeft|chmBottomFrontCenter|chmBottomFrontRight|
		chmLowFrequency|chmLowFrequency2),

	// 14: LFE-Pair (LFE1/LFE2)
	// The low-frequency effects subset of 10.2.9.3ch (Sound System H)
	expanded(14, 2,
		chmLowFrequency|chmLowFrequency2),

	// 15: Bottom-3ch (BtFL/BtFC/BtFR)
	// The bottom 3 channels of 10.2.9.3ch (Sound System H)
	expanded(15, 3,
		chmBottomFrontLeft|chmBottomFrontCenter|chmBottomFrontRight),

	// 16: 7.1.5.4ch - L/C/R/Lss/Rss/Lrs/Rrs/Ltf/Rtf/TpC/
	//                 Ltb/Rtb/BtFL/BtFR/BtBL/BtBR/LFE
	// Top and bottom speakers added to (4+7+0) of [ITU-2051-3] (Sound System J)
	expanded(16, 16, 0|
		chmFrontLeft|chmFrontCenter|chmFrontRight|
		chmSideSurroundLeft|chmSideSurroundRight|
		chmBackLeft|chmBackRight|
		chmTopFrontLeft|chmTopFrontRight|
		chmTopCenter|
		chmTopBackLeft|chmTopBackRight|
		chmBottomFrontLeft|chmBottomFrontRight|
		chmBottomBackLeft|chmBottomBackRight|
		chmLowFrequency),

	// 17: Bottom-4ch (BtFL/BtFR/BtBL/BtBR)
	// The bottom 4 channels of 7.1.5.4ch (Sound System J)
	expanded(17, 4, 0|
		chmBottomFrontLeft|chmBottomFrontRight|
		chmBottomBackLeft|chmBottomBackRight),

	// 18: Top-1ch (TpC) - The top subset of 7.1.5.4ch (Sound System J)
	expanded(18, 1,
		chmTopCenter),

	// 19: Top-5ch (Ltf/Rtf/TpC/Ltb/Rtb)
	// The top 5 channels of 7.1.5.4ch (Sound System J)
	expanded(19, 5, 0|
		chmTopFrontLeft|chmTopFrontRight|
		chmTopCenter|
		chmTopBackLeft|chmTopBackRight),

	// 20-255: Reserved for future use
}

// specific channel layouts
var channelLayoutBinaural = scalableChannelLayouts[9]

// soundSystemMap defines a struct for system maps
type soundSystemMap struct {
	SoundSystem SoundSystem
	Layout      channelLayout
}

// 5.1.4ch + Bottom Front Center (not in standard layouts)
var systemE = expanded(-1, 11, 0|
	chmFrontLeft|chmFrontCenter|chmFrontRight|
	chmSideLeft|chmSideRight|
	chmTopFrontLeft|chmTopFrontRight|
	chmTopBackLeft|chmTopBackRight|
	chmBottomFrontCenter|
	chmLowFrequency)

// 7.1.2ch + TpBC + LFE2 (not in standard layouts)
var systemF = expanded(-2, 12, 0|
	chmFrontLeft|chmFrontCenter|chmFrontRight|
	chmSideLeft|chmSideRight|
	chmBackLeft|chmBackRight|
	chmTopFrontLeft|chmTopFrontRight|
	chmTopBackLeft|chmTopBackCenter|chmLowFrequency2)

// 9.1.4ch (not in standard layouts)
var systemG = expanded(-3, 14, 0|
	chmFrontLeft|chmFrontCenter|chmFrontRight|
	chmFrontLeftOfCenter|chmFrontRightOfCenter|
	chmSideLeft|chmSideRight|
	chmBackLeft|chmBackRight|
	chmTopFrontLeft|chmTopFrontRight|
	chmTopBackLeft|chmTopBackRight|
	chmLowFrequency)

// mapping between IAMF types and the structs above
var iamfSoundSystemMap = []soundSystemMap{
	{SoundSystemA_0_2_0, scalableChannelLayouts[1]},           // Stereo
	{SoundSystemB_0_5_0, scalableChannelLayouts[2]},           // 5.1ch
	{SoundSystemC_2_5_0, scalableChannelLayouts[3]},           // 5.1.2ch
	{SoundSystemD_4_5_0, scalableChannelLayouts[4]},           // 5.1.4ch
	{SoundSystemE_4_5_1, systemE},                             // 5.1.4ch + BFC
	{SoundSystemF_3_7_0, systemF},                             // 7.1.2ch + TpBC + LFE2
	{SoundSystemG_4_9_0, systemG},                             // 9.1.4ch
	{SoundSystemH_9_10_3, expandedScalableChannelLayouts[13]}, // 10.2.9.3ch
	{SoundSystemI_0_7_0, scalableChannelLayouts[5]},           // 7.1ch
	{SoundSystemJ_4_7_0, scalableChannelLayouts[7]},           // 7.1.4ch
	{SoundSystem10_2_7_0, scalableChannelLayouts[6]},          // 7.1.2ch
	{SoundSystem11_2_3_0, scalableChannelLayouts[8]},          // 3.1.2ch
	{SoundSystem12_0_1_0, scalableChannelLayouts[0]},          // Mono
	{SoundSystem13_9_1_6, expandedScalableChannelLayouts[8]},  // 9.1.6ch
}

/**
 * implementation.
 * Based on IAMF_PARSE_C
 */

// Constants for IAMF parsing
const (
	MaxIamfObuHeaderSizeBytes = 1 + 8*3
	MaxIamfLabelSize          = 128
)

// OpusDecoderConfig parses Opus decoder configuration
func OpusDecoderConfig(sr bits.SliceReader, codecConfig *IamfCodecConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func mp4ReadDescr(sr bits.SliceReader) (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

// descLen

// AACDecoderConfig parses AAC decoder configuration
func AacDecoderConfig(sr bits.SliceReader, codecConfig *IamfCodecConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// MP4DecConfigDescrTag

// buffer size db
// rc_max_rate
// avg bitrate

// MP4DecSpecificDescrTag

// FLACDecoderConfig parses FLAC decoder configuration
func FlacDecoderConfig(sr bits.SliceReader, codecConfig *IamfCodecConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip METADATA_BLOCK_HEADER (4 bytes)

// FLAC_STREAMINFO_SIZE

// Extract sample rate from STREAMINFO

// PCMDecoderConfig parses PCM decoder configuration
func PcmDecoderConfig(sr bits.SliceReader, codecConfig *IamfCodecConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// 0 = BE, 1 = LE

// 16, 24, 32 bits

// Map codec IDs based on format and size

// codecConfigObu parses a Codec Config OBU
func codecConfigObu(sr bits.SliceReader, ctx *IamfContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Map codec ID to internal representation

// scalableChannelLayoutConfig parses scalable channel layout configuration
func scalableChannelLayoutConfig(sr bits.SliceReader, audioElement *IamfAudioElement) error {
	_ = "STUB: not implemented"
	return nil
}

// Bits 0-1 are reserved

// ambisonicsConfig parses ambisonics configuration
func ambisonicsConfig(sr bits.SliceReader, audioElement *IamfAudioElement) error {
	_ = "STUB: not implemented"
	return nil
}

/* incomplete order - some harmonics are missing */

// paramParse parses parameter definitions
func paramParse(sr bits.SliceReader, ctx *IamfContext, paramType ParamDefinitionType,
	audioElement *IamfAudioElement) (*ParamDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// audioElementObu parses an Audio Element OBU
func audioElementObu(sr bits.SliceReader, ctx *IamfContext) error {
	_ = "STUB: not implemented"
	return nil
}

// mixPresentationObu parses a Mix Presentation OBU
func mixPresentationObu(sr bits.SliceReader, ctx *IamfContext) error {
	_ = "STUB: not implemented"
	return nil
}

// parseObuSR parses an IAMF OBU
func parseObuSR(sr bits.SliceReader) (ObuInfo, error) {
	_ = "STUB: not implemented"
	return *new(ObuInfo), nil
}

// Read OBU type (5 bits)

/* redundant := */

// num_samples_to_trim_at_end
// num_samples_to_trim_at_start

// headerSize including trimming and extension

type ObuReader struct {
	ctx     *IamfContext
	sr      bits.SliceReader
	maxSize int
}

func NewObuReader(data []byte, maxSize int) ObuReader {
	_ = "STUB: not implemented"
	return *new(ObuReader)
}

func (r ObuReader) ReadObu() (*ObuInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *ObuReader) SkipPayload(obu *ObuInfo) { _ = "STUB: not implemented"; return }

func (r ObuReader) Context() *IamfContext {
	_ = "STUB: not implemented"

	// ReadDescriptors reads and parses IAMF descriptors
	return nil
}

func (obu ObuInfo) ReadDescriptors(r *ObuReader) (*IamfContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// remove already shown

// remove already shown

// Skip unknown OBU types

func (o *ObuInfo) Info(writer func(format string, p ...interface{})) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *IamfContext) Info(f func(level int, format string, p ...interface{})) error {
	_ = "STUB: not implemented"
	return nil
}

func (c channelLayout) toChannelLayout() ChannelLayout {
	_ = "STUB: not implemented"
	return *new(ChannelLayout)
}

// Identify layout based on Order type

// Build detailed channel map showing which speakers are present

// For custom/ambisonics layouts with explicit mapping

// Ambisonics channel naming: ACN (Ambisonic Channel Number)

func signExtend(v uint16) int32 {
	_ = "STUB: not implemented"
	// Sign extend from 16-bit to 32-bit
	return 0
}
