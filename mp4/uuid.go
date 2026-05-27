package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// UUID - 16-byte KeyID or SystemID
type UUID []byte

func (u UUID) String() string { _ = "STUB: not implemented"; return "" }

// Equal compares with other UUID
func (u UUID) Equal(a UUID) bool { _ = "STUB: not implemented"; return false }

// NewUUIDFromString creates a UUID from a hexadecimal, uuid-string or base64 string
func NewUUIDFromString(h string) (UUID, error) { _ = "STUB: not implemented"; return *new(UUID), nil }

const (
	// The following UUIDs belong to Microsoft Smooth Streaming Protocol (MSS)

	// UUIDMssSm - MSS StreamManifest UUID [MS-SSTR 2.2.7.2]
	UUIDMssSm = "3c2fe51b-efee-40a3-ae815300199dc348"

	// UUIDMssLs - MSS LiveServerManifest UUID [MS-SSTR 2.2.7.3]
	UUIDMssLsm = "a5d40b30-e814-11dd-ba2f-0800200c9a66"

	// UUIDTfxd - MSS tfxd UUID [MS-SSTR 2.2.4.4]
	UUIDTfxd = "6d1d9b05-42d5-44e6-80e2-141daff757b2"

	// UUIDTfrf - MSS tfrf UUID [MS-SSTR 2.2.4.5]
	UUIDTfrf = "d4807ef2-ca39-4695-8e54-26cb9e46a79f"

	// UUIDPiffSenc - PIFF UUID for Sample Encryption Box (PIFF 1.1 §5.3.2)
	UUIDPiffSenc = "a2394f52-5a9b-4f14-a244-6c427c648df4"

	// UUIDPiffTenc - PIFF UUID for Track Encryption Box (PIFF 1.1 §5.3.3)
	UUIDPiffTenc = "8974dbce-7be7-4c51-84f9-7148f9882554"

	// UUIDSphericalVideoV1 - Spherical Video V1 UUID
	UUIDSphericalVideoV1 = "ffcc8263-f855-4a93-8814-587a02521fdd"
)

// NewTrfrfBox creates a new TfrfBox with values.
// fragmentCount is the number of fragments, andb both
// fragmentAbsoluteTimes and fragmentAbsoluteDurations must be slices of that length.
func NewTfrfBox(fragmentCount byte, fragmentAbsoluteTimes, fragmentAbsoluteDurations []uint64) *UUIDBox {
	_ = "STUB: not implemented"
	return nil
}

// NewTfxdBox creates a new TfxdBox with values.
func NewTfxdBox(fragmentAbsoluteTime, fragmentAbsoluteDuration uint64) *UUIDBox {
	_ = "STUB: not implemented"
	return nil
}

// createUUID - create uuid from hex, uuid-formatted hex, or base64 string
func createUUID(u string) (UUID, error) { _ = "STUB: not implemented"; return *new(UUID), nil }

// mustCreateUUID - create uuid from string. Panic for bad string
func mustCreateUUID(u string) UUID { _ = "STUB: not implemented"; return *new(UUID) }

var (
	uuidTfxd             UUID = mustCreateUUID(UUIDTfxd)
	uuidTfrf             UUID = mustCreateUUID(UUIDTfrf)
	uuidPiffSenc         UUID = mustCreateUUID(UUIDPiffSenc)
	uuidPiffTenc         UUID = mustCreateUUID(UUIDPiffTenc)
	uuidSphericalVideoV1 UUID = mustCreateUUID(UUIDSphericalVideoV1)
)

// UUIDBox - Used as container for MSS boxes tfxd and tfrf
// For unknown UUID, the data after the UUID is stored as UnknownPayload
type UUIDBox struct {
	uuid           UUID
	Tfxd           *TfxdData
	Tfrf           *TfrfData
	Senc           *SencBox
	PiffTenc       *PiffTencData
	SphericalV1    *SphericalVideoV1Data
	StartPos       uint64
	UnknownPayload []byte
}

// UUID - Return UUID as formatted string
func (u *UUIDBox) UUID() string { _ = "STUB: not implemented"; return "" }

// UUID - Set UUID from string corresponding to 16 bytes.
// The input should be a UUID-formatted hex string, plain hex or baset64 encoded.
func (u *UUIDBox) SetUUID(uuid string) (err error) { _ = "STUB: not implemented"; return nil }

// TfxdData - MSS TfxdBox data after UUID part
// Defined in MSS-SSTR v20180912 section 2.2.4.4
type TfxdData struct {
	Version                  byte
	Flags                    uint32
	FragmentAbsoluteTime     uint64
	FragmentAbsoluteDuration uint64
}

// TfrfData - MSS TfrfBox data after UUID part
// Defined in MSS-SSTR v20180912 section 2.2.4.5
type TfrfData struct {
	Version                   byte
	Flags                     uint32
	FragmentCount             byte
	FragmentAbsoluteTimes     []uint64
	FragmentAbsoluteDurations []uint64
}

// SphericalVideoV1Data - Spherical Video V1 metadata
// Defined in Google's Spherical Video V1 RFC
// https://github.com/google/spatial-media/blob/master/docs/spherical-video-rfc.md
type SphericalVideoV1Data struct {
	XMLData string
}

// PiffTencData - PIFF TrackEncryptionBox payload (after the FullBox header
// and UUID). Defined in PIFF 1.1 §5.3.3 as:
//
//	unsigned int(24) default_AlgorithmID;
//	unsigned int(8)  default_IV_size;
//	unsigned int(8)[16] default_KID;
//
// AlgorithmID values (PIFF 1.1 §5.3.2): 0=Not Encrypted, 1=AES 128-bit CTR
// (equivalent to cenc), 2=AES 128-bit CBC (equivalent to cbcs).
type PiffTencData struct {
	Version     byte
	Flags       uint32
	AlgorithmID uint32 // 24 bits on the wire
	IVSize      byte
	KID         UUID
}

func (p *PiffTencData) size() uint64 { _ = "STUB: not implemented"; return 0 }

func decodePiffTenc(s bits.SliceReader) (*PiffTencData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *PiffTencData) encode(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// DecodeUUIDBox - decode a UUID box including tfxd or tfrf
func DecodeUUIDBox(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeUUIDBoxSR - decode a UUID box including tfxd or tfrf
func DecodeUUIDBoxSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// PIFF 1.1 §5.3.2: SampleEncryptionBox flag 0x1 ("Override TrackEncryptionBox
// parameters") prepends 24 bytes (AlgorithmID(24)+IV_size(8)+KID(128)) before
// sample_count, which the generic SencBox decoder doesn't expect. Reject
// explicitly to avoid silent misparse.

// This is like a SencBox except that there is no size and type. Offset and sizes must be slightly adjusted.

// Type - return box type
func (b *UUIDBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size including tfxd/tfrf
	return ""
}

func (b *UUIDBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// -8 because no header

// Encode - write box to w
func (b *UUIDBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *UUIDBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// SubType - interpret the UUID as a known sub type or unknown
func (b *UUIDBox) SubType() string { _ = "STUB: not implemented"; return "" }

func decodeTfxd(s bits.SliceReader) (*TfxdData, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *TfxdData) size() uint64 { _ = "STUB: not implemented"; return 0 }

func (t *TfxdData) encode(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

func decodeTfrf(s bits.SliceReader) (*TfrfData, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *TfrfData) size() uint64 { _ = "STUB: not implemented"; return 0 }

func (t *TfrfData) encode(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - box-specific info
func (b *UUIDBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackKey unpacks a hex or base64 encoded 16-byte key.
// The key can be in uuid formats with hyphens at positions 8, 13, 18, 23.
func UnpackKey(inKey string) (key []byte, err error) { _ = "STUB: not implemented"; return nil, nil }
