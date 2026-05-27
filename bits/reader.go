package bits

import (
	"io"
)

// Reader is a bit reader that stops reading at first error and stores it.
// First error can be fetched usiin AccError().
type Reader struct {
	rd    io.Reader
	err   error
	n     int  // current number of bits
	value uint // current accumulated value
	pos   int  // current position in reader (in bytes)
}

// AccError - accumulated error is first error that occurred
func (r *Reader) AccError() error {
	_ = "STUB: not implemented"

	// NewReader return a new Reader that accumulates errors.
	return nil
}

func NewReader(rd io.Reader) *Reader { _ = "STUB: not implemented"; return nil }

// Read - read n bits. Return 0, if error now or previously
func (r *Reader) Read(n int) uint { _ = "STUB: not implemented"; return 0 }

// ReadSigned reads a 2-complemented signed int with n bits.
func (r *Reader) ReadSigned(n int) int { _ = "STUB: not implemented"; return 0 }

// ReadFlag reads 1 bit and interprets as a boolean flag. Returns false if error now or previously.
func (r *Reader) ReadFlag() bool { _ = "STUB: not implemented"; return false }

// ReadRemainingBytes reads remaining bytes if byte-aligned. Returns nil if error now or previously.
func (r *Reader) ReadRemainingBytes() []byte { _ = "STUB: not implemented"; return nil }

// NrBytesRead returns how many bytes read into parser.
func (r *Reader) NrBytesRead() int {
	_ = "STUB: not implemented"
	// Starts at -1
	return 0
}

// NrBitsRead returns total number of bits read into parser.
func (r *Reader) NrBitsRead() int { _ = "STUB: not implemented"; return 0 }

// NrBitsReadInCurrentByte returns number of bits read in current byte.
func (r *Reader) NrBitsReadInCurrentByte() int {
	_ = "STUB: not implemented"

	// ByteAlign aligns the reader to the next byte boundary by discarding
	// any remaining bits in the current byte. This is commonly used in
	// multimedia formats where data structures need to be byte-aligned.
	return 0
}

func (r *Reader) ByteAlign() { _ = "STUB: not implemented"; return }

// Discard remaining bits in current byte to align to byte boundary
