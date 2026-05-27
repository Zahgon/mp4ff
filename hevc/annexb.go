package hevc

// GetParameterSetsFromByteStream gets SPS and PPS nalus from bytestream
func GetParameterSetsFromByteStream(data []byte) (vpss [][]byte, spss [][]byte, ppss [][]byte) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Remove zeros from end of NAL unit

// Video NALU types are below 32

// ExtractNalusOfTypeFromByteStream returns all HEVC nalus of wanted type from bytestream.
// If stopAtVideo, the stream is not scanned beyond the first video NAL unit.
func ExtractNalusOfTypeFromByteStream(nType NaluType, data []byte, stopAtVideo bool) [][]byte {
	_ = "STUB: not implemented"
	return nil
}

// Remove zeros from end of NAL unit

// Video nal unit type

func extractSlice(data []byte, start, stop int) []byte { _ = "STUB: not implemented"; return nil }
