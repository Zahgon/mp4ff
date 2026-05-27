package mp4

import (
	"io"
)

// TopBoxInfo - information about a top-level box
type TopBoxInfo struct {
	// Type - box type
	Type string
	// Size - box size
	Size uint64
	// StartPos - where in file does box start
	StartPos uint64
}

// GetTopBoxInfoList - get top boxes until stopBoxType or end of file
func GetTopBoxInfoList(rs io.ReadSeeker, stopBoxType string) ([]TopBoxInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
