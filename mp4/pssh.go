package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// UUIDs for different DRM systems
const (
	UUIDPlayReady   = "9a04f079-9840-4286-ab92-e65be0885f95"
	UUIDWidevine    = "edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"
	UUIDFairPlay    = "94ce86fb-07ff-4f43-adb8-93d2fa968ca2"
	UUID_VCAS       = "9a27dd82-fde2-4725-8cbc-4234aa06ec09"
	UUID_W3C_COMMON = "1077efec-c0b2-4d02-ace3-3c1e52e2fb4b"
)

// ProtectionSystemName returns name of protection system if known.
func ProtectionSystemName(systemID UUID) string { _ = "STUB: not implemented"; return "" }

// PsshBox - Protection System Specific Header Box
// Defined in ISO/IEC 23001-7 Section 8.1
type PsshBox struct {
	Version  byte
	Flags    uint32
	SystemID UUID
	KIDs     []UUID
	Data     []byte
}

// NewPsshBox makes a PsshBox with the given systemID, KIDs and data.
// If there are KIDs, the version is set to 1, otherwise it is set to 0.
// The systemID and KIDs are expected to be in the form of UUIDs
// (e.g., "9a04f079-9840-4286-ab92-e65be0885f95"), hex without hyphens, or base-64 encoded strings.
func NewPsshBox(systemID string, KIDs []string, data []byte) (*PsshBox, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DecodePssh - box-specific decode
func DecodePssh(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodePsshSR - box-specific decode
func DecodePsshSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type - return box type
func (b *PsshBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size
	return ""
}

func (b *PsshBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *PsshBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *PsshBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box info to w
func (b *PsshBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// PsshBoxesFromDBytesextracts pssh boxes from slice of bytes
func PsshBoxesFromBytes(psshData []byte) ([]*PsshBox, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PsshBoxesFromBase64 extracts pssh boxes from base64-encoded string
func PsshBoxesFromBase64(psshBase64 string) ([]*PsshBox, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
