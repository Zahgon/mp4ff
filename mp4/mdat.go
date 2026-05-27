package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// MdatBox - Media Data Box (mdat)
// The mdat box contains media chunks/samples.
// DataParts is to be able to gather output data without
// new allocations
type MdatBox struct {
	StartPos     uint64
	Data         []byte
	DataParts    [][]byte
	lazyDataSize uint64
	LargeSize    bool
}

const maxNormalPayloadSize = (1 << 32) - 1 - 8

// DecodeMdat - box-specific decode
func DecodeMdat(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeMdatSR decodes an mdat box
//
// Currently no content and no error is returned if not full length available.
// If not enough content, an accumulated error is stored in sr, though
func DecodeMdatSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// IsLazy - is the mdat data handled lazily (with separate writer/reader).
func (m *MdatBox) IsLazy() bool { _ = "STUB: not implemented"; return false }

// DecodeMdatLazily - box-specific decode but Data is not in memory
func DecodeMdatLazily(hdr BoxHeader, startPos uint64) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// SetLazyDataSize - set size of mdat lazy data so that the data can be written separately
// Don't put any data in m.Data in this mode.
func (m *MdatBox) SetLazyDataSize(newSize uint64) { _ = "STUB: not implemented"; return }

// GetLazyDataSize - size of the box if filled with data
func (m *MdatBox) GetLazyDataSize() uint64 { _ = "STUB: not implemented"; return 0 }

// Type - return box type
func (m *MdatBox) Type() string {
	_ = "STUB: not implemented"

	// Size - return calculated size, depending on largeSize set or not
	return ""
}

func (m *MdatBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// AddSampleData -  a sample data to an mdat box
func (m *MdatBox) AddSampleData(s []byte) { _ = "STUB: not implemented"; return }

// SetData - set the mdat data to given slice. No copying is done
func (m *MdatBox) SetData(data []byte) { _ = "STUB: not implemented"; return }

// AddSampleDataPart - add a data part (for output)
func (m *MdatBox) AddSampleDataPart(s []byte) { _ = "STUB: not implemented"; return }

// Reasonable size

// Encode - write box to w. If m.lazyDataSize > 0, the mdat data needs to be written separately
func (m *MdatBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write box to sw. If m.lazyDataSize > 0, the mdat data needs to be written separately
func (m *MdatBox) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// DataLength - length of data stored in box either as one or multiple parts
func (m *MdatBox) DataLength() uint64 { _ = "STUB: not implemented"; return 0 }

// Info - write box-specific information
func (m *MdatBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// HeaderSize - 8 or 16 (bytes) depending o whether largeSize is used
func (m *MdatBox) HeaderSize() uint64 { _ = "STUB: not implemented"; return 0 }

// PayloadAbsoluteOffset - position of mdat payload start (works after header)
func (m *MdatBox) PayloadAbsoluteOffset() uint64 { _ = "STUB: not implemented"; return 0 }

// ReadData reads Mdat data specified by the start and size.
// Input argument start is the position relative to the start of a file.
// The ReadSeeker is used for lazily loaded mdat case.
func (m *MdatBox) ReadData(start, size int64, rs io.ReadSeeker) ([]byte, error) {
	_ = "STUB: not implemented"
	// The Mdat box was decoded lazily
	return nil, nil
}

// Otherwise, all Mdat data is in memory, either as parts or as one big slice

// validate if indexes are valid to avoid panics

// CopyData - copy data range from mdat to w.
// The ReadSeeker is used for lazily loaded mdat case.
func (m *MdatBox) CopyData(start, size int64, rs io.ReadSeeker, w io.Writer) (nrWritten int64, err error) {
	_ = "STUB: not implemented"
	// The Mdat box was decoded lazily
	return 0, nil
}

// Otherwise, all Mdat data is in memory

// validate if indexes are valid to avoid panics
