package bits

import (
	"errors"
	"io"
)

// ESBPReader errors
var (
	ErrNotReadSeeker = errors.New("reader does not support Seek")
)

const (
	startCodeEmulationPreventionByte = 0x03
)

// EBSPReader reads an EBSP bitstream dropping start-code emulation bytes.
// It also supports checking for more rbsp data and reading rbsp_trailing_bits.
type EBSPReader struct {
	rd        io.Reader
	err       error
	n         int  // current number of bits
	v         uint // current accumulated value
	pos       int
	zeroCount int // Count number of zero bytes read
}

// NewEBSPReader return a new EBSP reader stopping reading at first error.
func NewEBSPReader(rd io.Reader) *EBSPReader { _ = "STUB: not implemented"; return nil }

// AccError returns the accumulated error. If no error, returns nil.
func (r *EBSPReader) AccError() error {
	_ = "STUB: not implemented"

	// NrBytesRead returns how many bytes read into parser.
	return nil
}

func (r *EBSPReader) NrBytesRead() int {
	_ = "STUB: not implemented"
	// Starts at -1
	return 0
}

// NrBitsRead returns total number of bits read into parser.
func (r *EBSPReader) NrBitsRead() int { _ = "STUB: not implemented"; return 0 }

// NrBitsReadInCurrentByte returns number of bits read in current byte.
func (r *EBSPReader) NrBitsReadInCurrentByte() int {
	_ = "STUB: not implemented"

	// Read reads n bits and respects and accumulates errors. If error, returns 0.
	return 0
}

func (r *EBSPReader) Read(n int) uint { _ = "STUB: not implemented"; return 0 }

// ReadBytes read n bytes and return nil if new or accumulated error.
func (r *EBSPReader) ReadBytes(n int) []byte { _ = "STUB: not implemented"; return nil }

// ReadFlag reads 1 bit and translates a bool.
func (r *EBSPReader) ReadFlag() bool { _ = "STUB: not implemented"; return false }

// ReadExpGolomb reads one unsigned exponential Golomb code.
func (r *EBSPReader) ReadExpGolomb() uint { _ = "STUB: not implemented"; return 0 }

// ReadSignedGolomb reads one signed exponential Golomb code.
func (r *EBSPReader) ReadSignedGolomb() int { _ = "STUB: not implemented"; return 0 }

// IsSeeker returns true if underluing reader supports Seek interface.
func (r *EBSPReader) IsSeeker() bool { _ = "STUB: not implemented"; return false }

// MoreRbspData returns false if next bit is 1 and last 1-bit in fullSlice.
// Underlying reader must support ReadSeeker interface to reset after check.
// Return false, nil if underlying error.
func (r *EBSPReader) MoreRbspData() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Find out if next position is the last 1

// If all remaining bits are zero, there is no more rbsp data

// Reset

// reset resets EBSPReader based on copy of previous state.
func (r *EBSPReader) reset(prevState EBSPReader) error { _ = "STUB: not implemented"; return nil }

// ReadRbspTrailingBits reads rbsp_traling_bits. Returns error if wrong pattern.
// If other error, returns nil and let AccError() provide that error.
func (r *EBSPReader) ReadRbspTrailingBits() error { _ = "STUB: not implemented"; return nil }

// Reset

// SetError sets an error if not already set.
func (r *EBSPReader) SetError(err error) { _ = "STUB: not implemented"; return }
