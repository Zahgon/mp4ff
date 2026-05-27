package bits

// FixedSliceWriter - write numbers to a fixed []byte slice
type FixedSliceWriter struct {
	accError error
	buf      []byte
	off      int
	n        int  // current number of bits
	v        uint // current accumulated value for bits
}

// NewFixedSliceWriter - create writer around slice.
// The slice will not grow, but stay the same size.
// If too much data is written, there will be
// an accumuluated error. Can be retrieved with AccError()
func NewFixedSliceWriterFromSlice(data []byte) *FixedSliceWriter {
	_ = "STUB: not implemented"
	return nil
}

// NewSliceWriter - create slice writer with fixed size.
func NewFixedSliceWriter(size int) *FixedSliceWriter { _ = "STUB: not implemented"; return nil }

// Len - length of FixedSliceWriter buffer written. Same as Offset()
func (sw *FixedSliceWriter) Len() int {
	_ = "STUB: not implemented"

	// Capacity - max length of FixedSliceWriter buffer
	return 0
}

func (sw *FixedSliceWriter) Capacity() int {
	_ = "STUB: not implemented"

	// Offset - offset for writing in FixedSliceWriter buffer
	return 0
}

func (sw *FixedSliceWriter) Offset() int {
	_ = "STUB: not implemented"

	// Bytes - return buf up to what's written
	return 0
}

func (sw *FixedSliceWriter) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// AccError - return accumulated error
func (sw *FixedSliceWriter) AccError() error {
	_ = "STUB: not implemented"

	// WriteUint8 - write byte to slice
	return nil
}

func (sw *FixedSliceWriter) WriteUint8(n byte) { _ = "STUB: not implemented"; return }

// WriteUint16 - write uint16 to slice
func (sw *FixedSliceWriter) WriteUint16(n uint16) { _ = "STUB: not implemented"; return }

// WriteInt16 - write int16 to slice
func (sw *FixedSliceWriter) WriteInt16(n int16) { _ = "STUB: not implemented"; return }

// WriteUint24 - write uint24 to slice
func (sw *FixedSliceWriter) WriteUint24(n uint32) { _ = "STUB: not implemented"; return }

// WriteUint32 - write uint32 to slice
func (sw *FixedSliceWriter) WriteUint32(n uint32) { _ = "STUB: not implemented"; return }

// WriteInt32 - write int32 to slice
func (sw *FixedSliceWriter) WriteInt32(n int32) { _ = "STUB: not implemented"; return }

// WriteUint48 - write uint48
func (sw *FixedSliceWriter) WriteUint48(u uint64) { _ = "STUB: not implemented"; return }

// WriteUint64 - write uint64 to slice
func (sw *FixedSliceWriter) WriteUint64(n uint64) { _ = "STUB: not implemented"; return }

// WriteInt64 - write int64 to slice
func (sw *FixedSliceWriter) WriteInt64(n int64) { _ = "STUB: not implemented"; return }

// WriteString - write string to slice with or without zero end
func (sw *FixedSliceWriter) WriteString(s string, addZeroEnd bool) {
	_ = "STUB: not implemented"
	return
}

// WriteZeroBytes - write n byte of zeroes
func (sw *FixedSliceWriter) WriteZeroBytes(n int) { _ = "STUB: not implemented"; return }

// WriteBytes - write []byte
func (sw *FixedSliceWriter) WriteBytes(byteSlice []byte) { _ = "STUB: not implemented"; return }

// WriteUnityMatrix - write a unity matrix for mvhd or tkhd
func (sw *FixedSliceWriter) WriteUnityMatrix() { _ = "STUB: not implemented"; return }

// = 1 fixed 16.16

// = 1 fixed 16.16

// = 1 fixed 2.30

func (sw *FixedSliceWriter) WriteBits(bits uint, n int) { _ = "STUB: not implemented"; return }

// WriteFlag writes a flag as 1 bit.
func (sw *FixedSliceWriter) WriteFlag(f bool) { _ = "STUB: not implemented"; return }

// FlushBits - write remaining bits to the underlying .Writer.
// bits will be left-shifted and zeros appended to fill up a byte.
func (sw *FixedSliceWriter) FlushBits() { _ = "STUB: not implemented"; return }
