package vvc

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

/*
PTL represents profile-tier-level information (VvcPTLRecord) Section 11.2.4.1.2

	aligned(8) class VvcPTLRecord(num_sublayers) {
		bit(2) reserved = 0;
		unsigned int(6) num_bytes_constraint_info;
		unsigned int(7) general_profile_idc;
		unsigned int(1) general_tier_flag;
		unsigned int(8) general_level_idc;
		unsigned int(1) ptl_frame_only_constraint_flag;
		unsigned int(1) ptl_multi_layer_enabled_flag;
		unsigned int(8*num_bytes_constraint_info - 2) general_constraint_info;
		for (i=num_sublayers - 2; i >= 0; i--)
			unsigned int(1) ptl_sublayer_level_present_flag[i];
		for (j=num_sublayers; j<=8 && num_sublayers > 1; j++)
			bit(1) ptl_reserved_zero_bit = 0;
		for (i=num_sublayers-2; i >= 0; i--)
			if (ptl_sublayer_level_present_flag[i])
				unsigned int(8) sublayer_level_idc[i];
		unsigned int(8) ptl_num_sub_profiles;
		for (j=0; j < ptl_num_sub_profiles; j++)
		unsigned int(32) general_sub_profile_idc[j];
	}
*/
type PTL struct {
	NumBytesConstraintInfo      uint8
	GeneralProfileIDC           uint8
	GeneralTierFlag             bool
	GeneralLevelIDC             uint8
	PtlFrameOnlyConstraintFlag  bool
	PtlMultiLayerEnabledFlag    bool
	GeneralConstraintInfo       []byte
	PtlSublayerLevelPresentFlag []bool
	SublayerLevelIDC            []uint8
	PtlNumSubProfiles           uint8
	GeneralSubProfileIDC        []uint32
}

/*
DecConfRec represents VVC decoder configuration record, 11.2.4.2.2

	aligned(8) class VvcDecoderConfigurationRecord {
		bit(5) reserved = '11111'b;
		unsigned int(2) LengthSizeMinusOne;
		unsigned int(1) ptl_present_flag;
		if (ptl_present_flag) {
			unsigned int(9) ols_idx;
			unsigned int(3) num_sublayers;
			unsigned int(2) constant_frame_rate;
			unsigned int(2) chroma_format_idc;
			unsigned int(3) bit_depth_minus8;
			bit(5) reserved = '11111'b;
			VvcPTLRecord(num_sublayers) native_ptl;
			unsigned_int(16) max_picture_width;
			unsigned_int(16) max_picture_height;
			unsigned int(16) avg_frame_rate;
		}
		unsigned int(8) num_of_arrays;
		for (j=0; j < num_of_arrays; j++) {
			unsigned int(1) array_completeness;
			bit(2) reserved = 0;
			unsigned int(5) NAL_unit_type;
			if (NAL_unit_type != DCI_NUT && NAL_unit_type != OPI_NUT)
				unsigned int(16) num_nalus;
			for (i=0; i< num_nalus; i++) {
				unsigned int(16) nal_unit_length;
				bit(8*nal_unit_length) nal_unit;
			}
		}
	}
*/
type DecConfRec struct {
	LengthSizeMinusOne uint8
	PtlPresentFlag     bool
	OlsIdx             uint16
	NumSublayers       uint8
	ConstantFrameRate  uint8
	ChromaFormatIDC    uint8
	BitDepthMinus8     uint8
	NativePTL          PTL
	MaxPictureWidth    uint16
	MaxPictureHeight   uint16
	AvgFrameRate       uint16
	NaluArrays         []NaluArray
}

// Size returns the size of the decoder configuration record
func (d *DecConfRec) Size() int {
	_ = "STUB: not implemented"
	// reserved + lengthSizeMinusOne + ptlPresentFlag
	return 0
}

// olsIdx + numSublayers + constantFrameRate + chromaFormatIDC
// bitDepthMinus8 + reserved

// PTL fields (VvcPTLRecord)
// num_bytes_constraint_info
// general_profile_idc + general_tier_flag
// general_level_idc
// constraint info (includes frame_only/multi_layer flags)

// Sublayer flags and levels

// The spec requires exactly 8 bits total for flags + reserved bits = 1 byte

// Sublayer level IDCs (only for present ones)

// ptl_num_sub_profiles
// general_sub_profile_idc array

// maxPictureWidth + maxPictureHeight + avgFrameRate

// numOfArrays

// arrayCompleteness + reserved + NALUnitType

// numNalus

// naluLength + nalu

// Encode writes the decoder configuration record to w
func (d *DecConfRec) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW writes the decoder configuration record to sw
func (d *DecConfRec) EncodeSW(sw bits.SliceWriter) error {
	_ = "STUB: not implemented"
	// First byte: reserved (5 bits) + lengthSizeMinusOne (2 bits) + ptlPresentFlag (1 bit)
	return nil
}

// olsIdx (9 bits) + numSublayers (3 bits) + constantFrameRate (2 bits) + chromaFormatIDC (2 bits)

// bitDepthMinus8 (3 bits) + reserved (5 bits)

// PTL fields - VvcPTLRecord structure
// First byte: reserved (2 bits) + num_bytes_constraint_info (6 bits)

// Second byte: general_profile_idc (7 bits) + general_tier_flag (1 bit)

// general_level_idc (8 bits)

// ptl_frame_only_constraint_flag (1 bit) + ptl_multi_layer_enabled_flag (1 bit) + constraint info

// Handle sublayer level present flags

// The spec requires writing exactly 8 bits total:
// - (numSublayers - 1) bits for ptl_sublayer_level_present_flag
// - (9 - numSublayers) bits for ptl_reserved_zero_bit
// This always totals to 8 bits when numSublayers > 1

// Write sublayer level present flags (from MSB to LSB)

// The remaining bits are reserved zero bits (already 0)

// Write sublayer level IDCs

// Default value

// ptl_num_sub_profiles

// general_sub_profile_idc

// maxPictureWidth, maxPictureHeight, avgFrameRate

// numOfArrays

// NAL unit arrays

// arrayCompleteness (1 bit) + reserved (2 bits) + NALUnitType (5 bits)

// NALU_DCI and NALU_OPI do not have numNalus field but default to 1 NALU

// DecodeVVCDecConfRec decodes a VVC decoder configuration record
func DecodeVVCDecConfRec(data []byte) (DecConfRec, error) {
	_ = "STUB: not implemented"
	return *new(DecConfRec), nil
}

// First byte: reserved (5 bits) + lengthSizeMinusOne (2 bits) + ptlPresentFlag (1 bit)

// olsIdx (9 bits) + numSublayers (3 bits) + constantFrameRate (2 bits) + chromaFormatIDC (2 bits)

// bitDepthMinus8 (3 bits) + reserved (5 bits)

// PTL fields - VvcPTLRecord structure
// First byte: reserved (2 bits) + num_bytes_constraint_info (6 bits)

// Validate num_bytes_constraint_info - must be > 0 according to VVC spec

// Second byte: general_profile_idc (7 bits) + general_tier_flag (1 bit)

// general_level_idc (8 bits)

// ptl_frame_only_constraint_flag + ptl_multi_layer_enabled_flag + general_constraint_info
// These flags are always present, followed by (8*num_bytes_constraint_info - 2) bits of constraint info

// Read remaining constraint info bytes

// Mask to get the last 6 bits

// Handle sublayer level present flags

// The spec requires reading exactly 8 bits total:
// - (numSublayers - 1) bits for ptl_sublayer_level_present_flag
// - (9 - numSublayers) bits for ptl_reserved_zero_bit
// This always totals to 8 bits when numSublayers > 1

// Read sublayer level present flags (from MSB to LSB)

// The remaining bits in flagByte are reserved zero bits (already consumed)

// Read sublayer level IDCs

// ptl_num_sub_profiles

// general_sub_profile_idc

// maxPictureWidth, maxPictureHeight, avgFrameRate

// numOfArrays

// NAL unit arrays

// arrayCompleteness (1 bit) + reserved (2 bits) + NALUnitType (5 bits)

// Default to 1 NALU for DCI and OPI

// Helper function to convert bool to uint8
func boolToUint8(b bool) uint8 { _ = "STUB: not implemented"; return 0 }
