package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// File - an MPEG-4 file asset
//
// A progressive MPEG-4 file contains three main boxes:
//
//	ftyp : the file type box
//	moov : the movie box (meta-data)
//	mdat : the media data (chunks and samples). Only used for pror
//
// where mdat may come before moov.
// If fragmented, there are many more boxes and they are collected
// in the InitSegment, Segment and Segments structures.
// The sample metadata in the fragments in the Segments will be
// optimized unless EncModeBoxTree is set.
// To Encode the same data as Decoded, this flag must therefore be set.
// In all cases, Children contain all top-level boxes
type File struct {
	Ftyp         *FtypBox
	Moov         *MoovBox
	Mdat         *MdatBox        // mdat box for non-fragmented files. Extra empty boxes allowed.
	Init         *InitSegment    // Init data (ftyp + moov for fragmented file)
	Sidx         *SidxBox        // The first sidx box for a DASH OnDemand file
	Sidxs        []*SidxBox      // All sidx boxes for a DASH OnDemand file
	tfra         *TfraBox        // Single tfra box read first if DecISMFlag set
	Mfra         *MfraBox        // MfraBox for ISM files
	Segments     []*MediaSegment // Media segments
	Children     []Box           // All top-level boxes in order
	FragEncMode  EncFragFileMode // Determine how fragmented files are encoded
	EncOptimize  EncOptimize     // Bit field with optimizations being done at encoding
	fileDecFlags DecFileFlags    // Bit field with flags for decoding
	isFragmented bool
	fileDecMode  DecFileMode
}

// EncFragFileMode - mode for writing file
type EncFragFileMode byte

const (
	// EncModeSegment - only encode boxes that are part of Init and MediaSegments
	EncModeSegment = EncFragFileMode(0)
	// EncModeBoxTree - encode all boxes in file tree
	EncModeBoxTree = EncFragFileMode(1)
)

// DecFileMode - mode for decoding file
type DecFileMode byte

const (
	// DecModeNormal - read Mdat data into memory during file decoding.
	DecModeNormal DecFileMode = iota
	// DecModeLazyMdat - do not read mdat data into memory.
	// Thus, decode process requires less memory and faster.
	DecModeLazyMdat
)

// DecFileFlags can be combined for special decoding options
type DecFileFlags uint32

const (
	DecNoFlags DecFileFlags = 0
	// DecISMFlag tries to read mfra box at end to find segment boundaries (for ISM files)
	DecISMFlag DecFileFlags = (1 << 0)
	// DecStartOnMoof starts a segment at each moof boundary
	// This is provided no styp, or sidx/mfra box gives other information
	DecStartOnMoof = (1 << 1)
	// if no styp box, or sidx/mfra strudture
)

// EncOptimize - encoder optimization mode
type EncOptimize uint32

const (
	// OptimizeNone - no optimization
	OptimizeNone = EncOptimize(0)
	// OptimizeTrun - optimize trun box by moving default values to tfhd
	OptimizeTrun = EncOptimize(1 << 0)
)

func (eo EncOptimize) String() string { _ = "STUB: not implemented"; return "" }

// NewFile - create MP4 file
func NewFile() *File { _ = "STUB: not implemented"; return nil }

// Reasonable number of children

// ReadMP4File - read an mp4 file from path
func ReadMP4File(path string) (*File, error) { _ = "STUB: not implemented"; return nil, nil }

// BoxStructure represent a box or similar entity such as a Segment
type BoxStructure interface {
	Encode(w io.Writer) error
}

// WriteToFile - write a box structure to a file at filePath
func WriteToFile(boxStructure BoxStructure, filePath string) error {
	_ = "STUB: not implemented"
	return nil
}

// AddMediaSegment - add a mediasegment to file f
func (f *File) AddMediaSegment(m *MediaSegment) { _ = "STUB: not implemented"; return }

// DecodeFile - parse and decode a file from reader r with optional file options.
// For example, the file options overwrite the default decode or encode mode.
// On decode problems, the returned File may contain some top-level boxes, but not all.
func DecodeFile(r io.Reader, options ...Option) (*File, error) {
	_ = "STUB: not implemented"

	// apply options to change the default decode or encode mode
	return nil, nil
}

// No moov and heuristic failed.
// Leave senc deferred for caller to parse later with init info.

// Not needed anymore

// Size - total size of all boxes
func (f *File) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// AddChild - add child with start position
func (f *File) AddChild(child Box, boxStartPos uint64) { _ = "STUB: not implemented"; return }

// sidx boxes are either added to the File or to a later media segment.
// Since sidx boxes for a segment come before the moof, it is important that a new
// segment is started with a styp box for the sidx to be associated with the
// right segment. An alternative is that the sidx box is coming after
// an mdat box. This is the end of a fragment or segment, but since the sidx box
// should not be inside a segment, a new segment is started.
//
// A more general solution could possibly be implemented by looking at the
// sidx details like reference_ID to understand the sidx chain structure,
// and/or by waiting with associating the sidx box until more boxes are read.
// Given the rareness of multiple sidx boxes and the complexity of implementing
// and testing such a solution, that track is not deemed worth the effort for now.

// Start a new segment since we cannot have sidx later in a segment

// Starts a new segment

// emsg box is only added at the start of a fragment (inside a segment).
// The case that a segment starts without an emsg is also handled.

// Only add if previous mdat is nil or empty

// startSegmentIfNeeded starts a new segment if there is none or if position match with sidx of tfra.
func (f *File) startSegmentIfNeeded(_ Box, boxStartPos uint64) { _ = "STUB: not implemented"; return }

// findAndReadMfra tries to find a tfra box inside an mfra box at the end of the file
// If no mfro box is found, no error is reported.
func (f *File) findAndReadMfra(r io.Reader) error { _ = "STUB: not implemented"; return nil }

// This is the fixed size of the mfro box

// mfro

// Not an mfro box here, just reset and return

// mfra

// AddSidx adds a sidx box to the File and not a MediaSegment.
func (f *File) AddSidx(sidx *SidxBox) { _ = "STUB: not implemented"; return }

// Encode - encode a file to a Writer
// Fragmented files are encoded based on InitSegment and MediaSegments, unless EncModeBoxTree is set.
func (f *File) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Progressive file

// EncodeSW - encode a file to a SliceWriter
// Fragmented files are encoded based on InitSegment and MediaSegments, unless EncModeBoxTree is set.
func (f *File) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Progressive file

// Info - write box tree with indent for each level
func (f *File) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// LastSegment - Currently last segment
func (f *File) LastSegment() *MediaSegment { _ = "STUB: not implemented"; return nil }

// IsFragmented - is file made of multiple segments (Mp4 fragments)
func (f *File) IsFragmented() bool { _ = "STUB: not implemented"; return false }

// ApplyOptions - applies options for decoding or encoding a file
func (f *File) ApplyOptions(opts ...Option) { _ = "STUB: not implemented"; return }

// Option is function signature of file options.
// The design follows functional options pattern.
type Option func(f *File)

// WithEncodeMode sets up EncFragFileMode
func WithEncodeMode(mode EncFragFileMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDecodeMode sets up DecFileMode
func WithDecodeMode(mode DecFileMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDecodeFlags sets up DecodeFlags
func WithDecodeFlags(flags DecFileFlags) Option { _ = "STUB: not implemented"; return *new(Option) }

// CopySampleData copies sample data from a track in a progressive mp4 file to w.
// Use rs for lazy read and workSpace as an intermediate storage to avoid memory allocations.
func (f *File) CopySampleData(w io.Writer, rs io.ReadSeeker, trak *TrakBox,
	startSampleNr, endSampleNr uint32, workSpace []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) UpdateSidx(addIfNotExists, nonZeroEPT bool) error {
	_ = "STUB: not implemented"
	return nil
}

func findReferenceTrak(initSeg *InitSegment) (*TrakBox, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type segData struct {
	startPos         uint64
	presentationTime uint64
	baseDecodeTime   uint64
	dur              uint32
	size             uint32
}

// findSegmentData returns a slice of segment media data using a reference track.
func findSegmentData(segs []*MediaSegment, refTrak *TrakBox, trex *TrexBox) ([]segData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Find track that gives sidx time values

func fillSidx(sidx *SidxBox, refTrak *TrakBox, segDatas []segData, nonZeroEPT bool) {
	_ = "STUB: not implemented"
	return
}

func insertSidx(inFile *File, segDatas []segData, sidx *SidxBox) error {
	_ = "STUB: not implemented"
	// insert sidx box before first media segment
	// TODO. Handle case where startPos is not reliable. Maybe first box of first segment
	return nil
}

func min(a, b int) int { _ = "STUB: not implemented"; return 0 }
