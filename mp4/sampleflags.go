package mp4

// SampleFlags according to 14496-12 Sec. 8.8.3.1
type SampleFlags struct {
	IsLeading                 byte
	SampleDependsOn           byte
	SampleIsDependedOn        byte
	SampleHasRedundancy       byte
	SamplePaddingValue        byte
	SampleIsNonSync           bool
	SampleDegradationPriority uint16
}

func (sf SampleFlags) String() string { _ = "STUB: not implemented"; return "" }

// Encode - convert sampleflags to uint32 bit pattern
func (sf SampleFlags) Encode() uint32 { _ = "STUB: not implemented"; return 0 }

// SyncSampleFlags - flags for I-frame or other sync sample
const SyncSampleFlags uint32 = 0x02000000

// NonSyncSampleFlags - flags for non-sync sample
const NonSyncSampleFlags uint32 = 0x00010000

// SampleDependsOn1 - this sample depends on others (not an I picture)
const SampleDependsOn1 uint32 = 0x01000000

// IsSyncSampleFlags - flags is set correctly for sync sample
func IsSyncSampleFlags(flags uint32) bool { _ = "STUB: not implemented"; return false }

// SetSyncSampleFlags - return flags with syncsample pattern
func SetSyncSampleFlags(flags uint32) uint32 { _ = "STUB: not implemented"; return 0 }

// SetNonSyncSampleFlags - return flags with nonsyncsample pattern
func SetNonSyncSampleFlags(flags uint32) uint32 { _ = "STUB: not implemented"; return 0 }

// DecodeSampleFlags - decode a uint32 flags field
func DecodeSampleFlags(u uint32) SampleFlags { _ = "STUB: not implemented"; return *new(SampleFlags) }
