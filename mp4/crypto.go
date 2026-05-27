package mp4

import (
	"github.com/Eyevinn/mp4ff/avc"
	"github.com/Eyevinn/mp4ff/hevc"
)

type cryptoDir int

const (
	minClearSize           = 96 // to generate same output as Bento4
	naluHdrLen             = 4
	dirEnc       cryptoDir = iota
	dirDec
)

// GetAVCProtectRanges for common encryption from a sample with 4-byte NALU lengths.
// THe spsMap and ppsMap are only needed for CBCS mode.
// For scheme cenc, protection ranges must be a multiple of 16 bytes leaving header and some more in the clear
// For scheme cbcs, protection range must start after the slice header.
func GetAVCProtectRanges(spsMap map[uint32]*avc.SPS, ppsMap map[uint32]*avc.PPS, sample []byte,
	scheme string) ([]SubSamplePattern, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Calculate a multiple of 16 bytes to protect

func GetHEVCProtectRanges(spsMap map[uint32]*hevc.SPS, ppsMap map[uint32]*hevc.PPS,
	sample []byte, scheme string) ([]SubSamplePattern, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Calculate a multiple of 16 bytes to protect

// AppendProtectRange appends a SubSamplePattern to a slice of SubSamplePattern, splitting into multiple if needed.
func AppendProtectRange(ssps []SubSamplePattern, nrClear, nrProtected uint32) []SubSamplePattern {
	_ = "STUB: not implemented"
	return nil
}

func getAudioProtectRanges(sample []byte, scheme string) ([]SubSamplePattern, error) {
	_ = "STUB: not implemented"

	// CryptSampleCenc encrypts/decrypts cenc-schema sample in place provided key, iv, and subSamplePatterns.
	return nil, nil
}

func CryptSampleCenc(sample []byte, key []byte, iv []byte, subSamplePatterns []SubSamplePattern) error {
	_ = "STUB: not implemented"
	return nil
}

// DecryptSampleCenc does in-place decryption of cbcs-schema encrypted sample.
// Each protected byte range is striped with with pattern defined by pattern in tenc.
func DecryptSampleCbcs(sample []byte, key []byte, iv []byte, subSamplePatterns []SubSamplePattern, tenc *TencBox) error {
	_ = "STUB: not implemented"
	return nil
}

// EncryptSampleCenc does in-place encryption using cbcs schema.
// Each protected byte range is striped with with pattern defined by pattern in tenc.
func EncryptSampleCbcs(sample []byte, key []byte, iv []byte, subSamplePatterns []SubSamplePattern, tenc *TencBox) error {
	_ = "STUB: not implemented"
	return nil
}

// cryptSampleCbcs does either encryption of decryption of a sample using cbcs scheme.
func cryptSampleCbcs(dir cryptoDir, sample []byte, key []byte, iv []byte, subSamplePatterns []SubSamplePattern, tenc *TencBox) error {
	_ = "STUB: not implemented"
	return nil
}

// Full encryption as used for audio

// cbcsCrypt does one in-place CBC encryption/decryption. Full if nrInSkipBlock == 0.
// The normal case is that nrInCryptBlock == 16 and nrInSkipBlock == 144.
func cbcsCrypt(dir cryptoDir, data []byte, key []byte, iv []byte, nrInCryptBlock, nrInSkipBlock int) error {
	_ = "STUB: not implemented"
	return nil

	// This is the bytes that we should stripe decrypt
}

// Drops 4 last bits -> multiple of 16

// incrementIV increments the IV by the number of encrypted blocks and return a new IV.
func incrementIV(inIV []byte, subsamplePatterns []SubSamplePattern, sampleLen int) []byte {
	_ = "STUB: not implemented"
	return nil
}

func incrementIVInPlace(iv []byte, nrSteps int) { _ = "STUB: not implemented"; return }

type ProtectionRangeFunc func(sample []byte, scheme string) ([]SubSamplePattern, error)
type InitProtectData struct {
	Tenc     *TencBox
	ProtFunc ProtectionRangeFunc
	Trex     *TrexBox
	Scheme   string
}

// InitProtect modifies the init segment to add protection information and return what is needed to encrypt fragments.
func InitProtect(init *InitSegment, key, iv []byte, scheme string, kid UUID, psshBoxes []*PsshBox) (*InitProtectData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert to 16 bytes

func getAVCPSMaps(spss [][]byte, ppss [][]byte) (map[uint32]*avc.SPS, map[uint32]*avc.PPS, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func getAVCProtFunc(avcC *AvcCBox) (ProtectionRangeFunc, error) {
	_ = "STUB: not implemented"
	return *new(ProtectionRangeFunc), nil
}

func getHEVCPSMaps(arrays []hevc.NaluArray) (map[uint32]*hevc.SPS, map[uint32]*hevc.PPS, error) {
	_ = "STUB: not implemented"
	// First check SPS
	return nil, nil, nil
}

// Then check PPS

func getHEVCProtFunc(hvcC *HvcCBox) (ProtectionRangeFunc, error) {
	_ = "STUB: not implemented"
	return *new(ProtectionRangeFunc), nil
}

// EncryptFragment encrypts a fragment in place and returns the next IV to use
// for the following fragment. For cenc, the returned IV is the input IV
// incremented by the number of encrypted AES blocks in the fragment, so that
// callers processing multiple fragments with the same key can chain calls and
// avoid IV reuse. For cbcs the IV is constant and the returned slice is a
// copy of the input IV.
func EncryptFragment(f *Fragment, key, iv []byte, ipd *InitProtectData) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert to 16 bytes

// Store IVs in the senc box and update depending on blocks of encrypted data

// iv is constant and not sent t senc

// Offset to the senc box data to be set in saio

// 12 for full box and 4 for sample count

type DecryptInfo struct {
	Psshs      []*PsshBox
	TrackInfos []DecryptTrackInfo
}

type DecryptTrackInfo struct {
	TrackID uint32
	Sinf    *SinfBox
	Trex    *TrexBox
	Psshs   []*PsshBox
}

func (d DecryptInfo) findTrackInfo(trackID uint32) DecryptTrackInfo {
	_ = "STUB: not implemented"
	return *new(DecryptTrackInfo)
}

// normalizePiffScheme rewrites a PIFF-style sinf so it looks like cenc/cbcs to
// the rest of the decryption pipeline. Per PIFF 1.1 §5.3.3, the PIFF
// TrackEncryptionBox (UUID 8974dbce-7be7-4c51-84f9-7148f9882554) carries
// default_AlgorithmID, default_IV_size and default_KID. AlgorithmID values are
// listed in PIFF 1.1 §5.3.2: 1=AES 128-bit CTR (equivalent to cenc),
// 2=AES 128-bit CBC (equivalent to cbcs).
func normalizePiffScheme(sinf *SinfBox) error { _ = "STUB: not implemented"; return nil }

// DecryptInit modifies init segment in place and returns decryption info and a clean init segment.
func DecryptInit(init *InitSegment) (DecryptInfo, error) {
	_ = "STUB: not implemented"
	return *new(DecryptInfo), nil
}

// Should be track in the clear

// DecryptSegment decrypts a media segment in place
func DecryptSegment(seg *MediaSegment, di DecryptInfo, key []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// DecryptSegmentWithKeys decrypts a media segment in place using either a legacy key
// or keys selected by KID. KID values are expected as 32-char lowercase hex without dashes.
// If strictKIDMode is true, encrypted tracks must have a matching key in keysByKID.
func DecryptSegmentWithKeys(seg *MediaSegment, di DecryptInfo, key []byte, keysByKID map[string][]byte, strictKIDMode bool) error {
	_ = "STUB: not implemented"
	return nil
}

// drop sidx inside segment, since not modified properly

// DecryptFragment decrypts a fragment in place
func DecryptFragment(frag *Fragment, di DecryptInfo, key []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func getTrackKIDHex(ti DecryptTrackInfo) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getTrackKey(ti DecryptTrackInfo, key []byte, keysByKID map[string][]byte, strictKIDMode bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DecryptFragmentWithKeys decrypts a fragment in place using either a legacy key
// or keys selected by KID. KID values are expected as 32-char lowercase hex without dashes.
// If strictKIDMode is true, encrypted tracks must have a matching key in keysByKID.
func DecryptFragmentWithKeys(frag *Fragment, di DecryptInfo, key []byte, keysByKID map[string][]byte, strictKIDMode bool) error {
	_ = "STUB: not implemented"
	return nil
}

// decryptSample - decrypt samples inplace
func decryptSamplesInPlace(schemeType string, samples []FullSample, key []byte, tenc *TencBox, senc *SencBox) error {
	_ = "STUB: not implemented"

	// TODO. Interpret saio and saiz to get to the right place
	// Saio tells where the IV starts relative to moof start
	// It typically ends up inside senc (16 bytes after start)
	return nil
}

// ExtractInitProtectData extracts protection data from init segment
func ExtractInitProtectData(inSeg *InitSegment) (*InitProtectData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
