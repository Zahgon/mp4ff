package bits

import (
	"io"
)

// Writer writes bits into underlying io.Writer. Stops writing at first error.
// That first error is stored and can later be checked with AccError().
type Writer struct {
	wr  io.Writer
	err error  // The first error caused by any write operation
	out []byte // Slice of length 1 to avoid allocation at output
	n   int    // current number of bits
	v   uint   // current accumulated value
}

// NewWriter returns a new Writer
func NewWriter(w io.Writer) *Writer { _ = "STUB: not implemented"; return nil }

// Write writes n bits from bits and saves error state.
func (w *Writer) Write(bits uint, n int) { _ = "STUB: not implemented"; return }

// Flush writes remaining bits to the underlying io.Writer by adding zeros to the right.
func (w *Writer) Flush() { _ = "STUB: not implemented"; return }

// AccError returns the first error that occurred and stopped writing.
func (w *Writer) AccError() error {
	_ = "STUB: not implemented"

	// Mask returns a binary mask for the n least significant bits.
	return nil
}

func Mask(n int) uint { _ = "STUB: not implemented"; return 0 }
