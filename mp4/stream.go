package mp4

import (
	"io"
)

// TrailingBoxesErrror indicates that there are unexpected boxes after the last fragment.
type TrailingBoxesErrror struct {
	BoxNames []string
}

func (e *TrailingBoxesErrror) Error() string { _ = "STUB: not implemented"; return "" }

// InitDecodeStream reads and parses only the init segment.
// Stops as soon as it peeks a box that belongs to a fragment (styp, sidx, moof, emsg, prft).
// Returns a StreamFile ready for ProcessFragments to consume fragments.
func InitDecodeStream(r io.Reader, options ...StreamOption) (*StreamFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start with 64KB, will grow as needed

// Peek at next box header to see what's coming

// Reached EOF before any fragments - file may be init-only

// Check if this box belongs to fragments - if so, stop here
// The header is in the buffer, leave it there for ProcessFragments

// These boxes indicate start of fragments
// Header bytes are in buffer, currentPos points after header
// Reset currentPos to boxStartPos so ProcessFragments can re-peek

// This box is part of the init segment - read and parse it

// Parse box from buffer using DecodeBoxSR

// Clear buffer for next box now that we're done parsing

// Unknown boxes in init segment - keep them

// Update stream position

// StreamFile wraps File with streaming capabilities for processing fragments incrementally.
type StreamFile struct {
	*File
	reader          io.Reader
	boxSeekReader   *BoxSeekReader
	onFragmentReady FragmentCallback
	onFragmentDone  FragmentDoneCallback
	maxFragments    int
	streamPos       uint64
}

// FragmentCallback is called when a fragment's moof box has been parsed and mdat is ready to be accessed.
// The SampleAccessor provides lazy access to sample data.
type FragmentCallback func(f *Fragment, sa SampleAccessor) error

// FragmentDoneCallback is called after a fragment has been fully processed.
type FragmentDoneCallback func(f *Fragment) error

// SampleAccessor provides access to samples within a fragment.
type SampleAccessor interface {
	GetSample(trackID uint32, sampleNr uint32) (*FullSample, error)
	GetSampleRange(trackID uint32, startSampleNr, endSampleNr uint32) ([]FullSample, error)
	GetSamples(trackID uint32) ([]FullSample, error)
}

// StreamOption configures streaming behavior.
type StreamOption func(*StreamFile)

// WithFragmentCallback sets the callback invoked when a fragment is ready for processing.
// This corresponds to the point after the moof box has been parsed and mdat is ready to be accessed.
func WithFragmentCallback(cb FragmentCallback) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

// WithFragmentDone sets the callback invoked after fragment processing completes.
func WithFragmentDone(cb FragmentDoneCallback) StreamOption {
	_ = "STUB: not implemented"
	return *new(StreamOption)
}

// WithMaxFragments sets the maximum number of fragments to retain in memory (sliding window).
// Default is 3. Set to 0 to keep all fragments.
func WithMaxFragments(max int) StreamOption { _ = "STUB: not implemented"; return *new(StreamOption) }

// fragmentSampleAccessor implements SampleAccessor for a fragment using the boxSeekReader.
type fragmentSampleAccessor struct {
	fragment      *Fragment
	boxSeekReader io.ReadSeeker
	trex          *TrexBox
}

// GetSample retrieves a specific sample by track ID and sample number (1-based).
func (fsa *fragmentSampleAccessor) GetSample(trackID uint32, sampleNr uint32) (*FullSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Find which trun contains this sample and the sample's position

// This sample is in a later trun

// Sample is in this trun

// Accumulate decode time for samples before this one in the trun

// Calculate file offset for this sample

// Add size of samples before this one in the trun

// Read just this sample's data

func (fsa *fragmentSampleAccessor) GetSampleRange(trackID uint32, startSampleNr, endSampleNr uint32) ([]FullSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Calculate base offset for this trun

// If we're past the end of the range, we're done

// If we haven't reached the start yet, skip this sample

// Read this sample's data

// GetSamples retrieves all samples for a given track ID in the fragment.
// Will not return until the full mdat box has been read.
func (fsa *fragmentSampleAccessor) GetSamples(trackID uint32) ([]FullSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProcessFragments reads and processes fragments from the stream until EOF.
// Returns a TrailingBoxesErrror if there are unexpected boxes after the last fragment.
func (sf *StreamFile) ProcessFragments() error {
	_ = "STUB: not implemented"
	// Collect boxes between fragments (styp, sidx, emsg, etc.)
	return nil
}

// Peek at next box header to get type and size

// Check if this might be trailing data or end of stream

// We successfully read some fragments, this might just be EOF

// For non-moof boxes, collect them to include with the next fragment

// Read entire box into buffer

// Parse box from buffer using DecodeBoxSR

// Copy styp boxes to avoid shared mutable state

// Read entire moof box into buffer

// Parse moof from buffer using DecodeBoxSR

// Process the fragment (moof + mdat)

// Clear pre-fragment boxes for next fragment

// processFragment positions stream at end of mdat, ready for next box

// processFragment handles a complete fragment (moof + mdat).
// moofStartPos is the start position of the moof box.
// preFragmentBoxes are boxes that appeared before the moof (sidx, emsg, styp, etc.)
func (sf *StreamFile) processFragment(moof *MoofBox, moofStartPos uint64, preFragmentBoxes []Box) error {
	_ = "STUB: not implemented"
	return nil
}

// Peek at mdat box header
// Stream should already be positioned at moofEndPos (right after moof box)

// Create lazy mdat box and skip the header in stream

// Skip past mdat header to position at payload start

// Configure boxSeekReader for this mdat's bounds
// This also pre-allocates buffer if mdat is small enough

// Stream is now positioned at start of mdat payload, ready for sample reads
// Verify position is correct

// Create fragment with all boxes (pre-fragment boxes + moof + mdat)

// Invoke callback if set

// Add to file structure

// Invoke done callback and handle cleanup

// Drop old fragments if sliding window is enabled

// Skip to end of mdat box to continue to next box
// mdatBox.Size() includes both header and payload
// So end position is mdatStartPos + mdatBox.Size()

// Reset mdat-specific state after seeking (clears mdatActive flag and buffer)

// Stream is now positioned at mdatEndPos, ready to read next box header

// dropOldestFragment removes the oldest fragment from the file structure.
func (sf *StreamFile) dropOldestFragment() { _ = "STUB: not implemented"; return }

// GetActiveFragments returns the currently retained fragments.
func (sf *StreamFile) GetActiveFragments() []*Fragment { _ = "STUB: not implemented"; return nil }
