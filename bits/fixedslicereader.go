package bits

// FixedSliceReader - read integers and other data from a fixed slice.
// Accumulates error, and the first error can be retrieved.
// If err != nil, 0 or empty string is returned
type FixedSliceReader struct {
	err   error
	slice []byte
	pos   int
	len   int
}

// bits.NewFixedSliceReader creates a new slice reader reading from data
func NewFixedSliceReader(data []byte) *FixedSliceReader { _ = "STUB: not implemented"; return nil }

// AccError - get accumulated error after read operations
func (s *FixedSliceReader) AccError() error {
	_ = "STUB: not implemented"

	// ReadUint8 - read uint8 from slice
	return nil
}

func (s *FixedSliceReader) ReadUint8() byte { _ = "STUB: not implemented"; return 0 }

// ReadUint16 - read uint16 from slice
func (s *FixedSliceReader) ReadUint16() uint16 { _ = "STUB: not implemented"; return 0 }

// ReadInt16 - read int16 from slice
func (s *FixedSliceReader) ReadInt16() int16 { _ = "STUB: not implemented"; return 0 }

// ReadUint24 - read uint24 from slice
func (s *FixedSliceReader) ReadUint24() uint32 { _ = "STUB: not implemented"; return 0 }

// ReadUint32 - read uint32 from slice
func (s *FixedSliceReader) ReadUint32() uint32 { _ = "STUB: not implemented"; return 0 }

// ReadInt32 - read int32 from slice
func (s *FixedSliceReader) ReadInt32() int32 { _ = "STUB: not implemented"; return 0 }

// ReadUint64 - read uint64 from slice
func (s *FixedSliceReader) ReadUint64() uint64 { _ = "STUB: not implemented"; return 0 }

// ReadInt64 - read int64 from slice
func (s *FixedSliceReader) ReadInt64() int64 { _ = "STUB: not implemented"; return 0 }

// ReadFixedLengthString - read string of specified length n.
// Sets err and returns empty string if full length not available
func (s *FixedSliceReader) ReadFixedLengthString(n int) string {
	_ = "STUB: not implemented"
	return ""
}

// ReadZeroTerminatedString - read string until zero byte but at most maxLen
// Set err and return empty string if no zero byte found.
func (s *FixedSliceReader) ReadZeroTerminatedString(maxLen int) string {
	_ = "STUB: not implemented"
	return ""
}

// Next position to read

// ReadPossiblyZeroTerminatedString - read string until zero byte but at most maxLen.
// If maxLen is reached and no zero-byte, return string and ok = false
func (s *FixedSliceReader) ReadPossiblyZeroTerminatedString(maxLen int) (str string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Next position to read

// ReadBytes - read a slice of n bytes
// Return empty slice if n bytes not available
func (s *FixedSliceReader) ReadBytes(n int) []byte { _ = "STUB: not implemented"; return nil }

// RemainingBytes - return remaining bytes of this slice
func (s *FixedSliceReader) RemainingBytes() []byte { _ = "STUB: not implemented"; return nil }

// NrRemainingBytes - return number of bytes remaining
func (s *FixedSliceReader) NrRemainingBytes() int { _ = "STUB: not implemented"; return 0 }

// SkipBytes - skip passed n bytes
func (s *FixedSliceReader) SkipBytes(n int) { _ = "STUB: not implemented"; return }

// SetPos - set read position is slice
func (s *FixedSliceReader) SetPos(pos int) { _ = "STUB: not implemented"; return }

// GetPos - get read position is slice
func (s *FixedSliceReader) GetPos() int {
	_ = "STUB: not implemented"

	// Length - get length of slice
	return 0
}

func (s *FixedSliceReader) Length() int {
	_ = "STUB: not implemented"

	// LookAhead returns data ahead of current pos if within bounds.
	return 0
}

func (s *FixedSliceReader) LookAhead(offset int, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}
