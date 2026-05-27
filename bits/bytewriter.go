package bits

import (
	"io"
)

// ByteWriter - writer that wraps an io.Writer and accumulates error.
// Only the first error is saved, but any later calls will not panic.
type ByteWriter struct {
	w   io.Writer
	err error
}

// NewByteWriter creates accumulated error writer around io.Writer.
func NewByteWriter(w io.Writer) *ByteWriter { _ = "STUB: not implemented"; return nil }

// AccError - return accumulated error
func (a *ByteWriter) AccError() error {
	_ = "STUB: not implemented"

	// WriteUint8 - write a byte
	return nil
}

func (a *ByteWriter) WriteUint8(b byte) { _ = "STUB: not implemented"; return }

// WriteUint16 - write uint16
func (a *ByteWriter) WriteUint16(u uint16) { _ = "STUB: not implemented"; return }

// WriteUint32 - write uint32
func (a *ByteWriter) WriteUint32(u uint32) { _ = "STUB: not implemented"; return }

// WriteUint48 - write uint48
func (a *ByteWriter) WriteUint48(u uint64) { _ = "STUB: not implemented"; return }

// WriteUint64 - write uint64
func (a *ByteWriter) WriteUint64(u uint64) { _ = "STUB: not implemented"; return }

// WriteSlice - write a slice
func (a *ByteWriter) WriteSlice(s []byte) { _ = "STUB: not implemented"; return }
