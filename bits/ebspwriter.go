package bits

import (
	"io"
)

// EBSPWriter write bits and insert start-code emulation prevention bytes as necessary.
// Ceases writing at first error.
// The first error can later be checked with AccError().
type EBSPWriter struct {
	wr  io.Writer // underlying writer
	err error     // The first error caused by any write operation
	out []byte    // Slice of length 1 to avoid allocation at output

	n   int  // current number of bits
	v   uint // current accumulated value
	nr0 int  // Number preceding zero bytes

}

// NewEBSPWriter - returns a new Writer
func NewEBSPWriter(w io.Writer) *EBSPWriter { _ = "STUB: not implemented"; return nil }

// Write - write n bits from bits and save error state
func (w *EBSPWriter) Write(bits uint, n int) { _ = "STUB: not implemented"; return }

// start code emulation prevention

// WriteExpGolomb - write an exponential Golomb code
func (w *EBSPWriter) WriteExpGolomb(nr uint) { _ = "STUB: not implemented"; return }

// WriteSEIValue insert 0xFF until value is less than 255. Used in SEI payload type and size.
func (w *EBSPWriter) WriteSEIValue(val uint) { _ = "STUB: not implemented"; return }

// WriteRbspTrailingBits - write rbsp trailing bits (a 1 followed by zeros to a byte boundary)
func (w *EBSPWriter) WriteRbspTrailingBits() { _ = "STUB: not implemented"; return }

// StuffByteWithZeros - write zero bits until byte boundary (0-7bits)
func (w *EBSPWriter) StuffByteWithZeros() { _ = "STUB: not implemented"; return }

// AccError - return accumulated error
func (w *EBSPWriter) AccError() error {
	_ = "STUB: not implemented"

	// NrBitsInBuffer - number bits written in buffer byte
	return nil
}

func (w *EBSPWriter) NrBitsInBuffer() uint {
	_ = "STUB: not implemented"

	// BitsInBuffer - n bits written in buffer byte, not written to underlying writer
	return 0
}

func (w *EBSPWriter) BitsInBuffer() (bits, n uint) { _ = "STUB: not implemented"; return 0, 0 }
