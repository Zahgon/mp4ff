package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// StscBox is Sample To Chunk Box in progressive file.
//
// A chunk contains samples. This table defines to which chunk a sample is associated.
// Each entry is defined by :
//
//   - first chunk : all chunks starting at this index up to the next first chunk have the same sample count/description
//   - samples per chunk : number of samples in the chunk
//   - sample description id : description (see the sample description box - stsd)
//     this value is most often the same for all samples, so it is stored as a single value if possible.
//
// FirstSampleNr is a helper value for fast lookup. Something that is often a bottleneck.
type StscBox struct {
	Version                   byte
	Flags                     uint32
	singleSampleDescriptionID uint32 // Used instead of slice if all values are the same
	Entries                   []StscEntry
	SampleDescriptionID       []uint32
}

type StscEntry struct {
	FirstChunk      uint32
	SamplesPerChunk uint32
	FirstSampleNr   uint32
}

// DecodeStsc - box-specific decode
func DecodeStsc(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeStscSR - box-specific decode
func DecodeStscSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Type box-specific type
func (b *StscBox) Type() string {
	_ = "STUB: not implemented"

	// Size - box-specific size
	return ""
}

func (b *StscBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *StscBox) expectedSize(nrEntries int) uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write box to w
func (b *StscBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - box-specific encode to slicewriter
func (b *StscBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write specific box info to w
func (b *StscBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// AddEntry adds a new entry and calculates helper values.
func (b *StscBox) AddEntry(firstChunk, samplesPerChunk, sampleDescriptionID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSampleDescriptionID returns the sample description ID from common or individual values for chunk.
// chunkNr is 1-based.
func (b *StscBox) GetSampleDescriptionID(chunkNr int) uint32 { _ = "STUB: not implemented"; return 0 }

// SetSingleSampleDescriptionID - use this for efficiency if all samples have same sample description
func (b *StscBox) SetSingleSampleDescriptionID(sampleDescriptionID uint32) {
	_ = "STUB: not implemented"
	return
}

// ChunkNrFromSampleNr - get chunk number from sampleNr (one-based)
func (b *StscBox) ChunkNrFromSampleNr(sampleNr int) (chunkNr, firstSampleInChunk int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// Chunk defines a chunk with number, starting sampleNr and nrSamples.
type Chunk struct {
	ChunkNr       uint32
	StartSampleNr uint32
	NrSamples     uint32
}

// GetContainingChunks returns chunks containing the sample interval including endSampleNr.
// startSampleNr and endSampleNr are 1-based.
func (b *StscBox) GetContainingChunks(startSampleNr, endSampleNr uint32) ([]Chunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetChunk returns chunk for chunkNr (one-based).
func (b *StscBox) GetChunk(chunkNr uint32) Chunk { _ = "STUB: not implemented"; return *new(Chunk) }

// findEntryNrForChunkNr returns the entry where chunkNr belongs.
// The resulting entryNr is 0-based index.
func (b *StscBox) findEntryNrForChunkNr(chunkNr uint32) uint32 {
	_ = "STUB: not implemented"
	// The following is essentially the sort.Search() code specialized to this case
	return 0
}

// avoid overflow when computing h
// low ≤ mid < high

// FindEntryNrForSampleNr returns the entry where sampleNr belongs. lowEntryIdx is entry index (zero-based).
// The resulting entryNr is 0-based index.
func (b *StscBox) FindEntryNrForSampleNr(sampleNr, lowEntryIdx uint32) uint32 {
	_ = "STUB: not implemented"
	// The following is essentially the sort.Search() code specialized to this case
	return 0
}

// low ≤ mid < high
