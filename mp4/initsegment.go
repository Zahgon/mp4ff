package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// InitSegment - MP4/CMAF init segment
type InitSegment struct {
	MediaType string
	Ftyp      *FtypBox
	Moov      *MoovBox
	Children  []Box // All top-level boxes in order
}

// NewMP4Init - Create MP4Init
func NewMP4Init() *InitSegment { _ = "STUB: not implemented"; return nil }

// AddChild - Add a top-level box to InitSegment
func (s *InitSegment) AddChild(b Box) { _ = "STUB: not implemented"; return }

// Size - size of init segment
func (s *InitSegment) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - encode an initsegment to a Writer
func (s *InitSegment) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeSW - encode an initsegment to a SliceWriter
func (s *InitSegment) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// Info - write box tree with indent for each level
func (s *InitSegment) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateEmptyInit - create an init segment for fragmented files
func CreateEmptyInit() *InitSegment { _ = "STUB: not implemented"; return nil }

// AddEmptyTrack - add trak + trex box with appropriate trackID value and returns a reference to the new TrakBox
func (s *InitSegment) AddEmptyTrack(timeScale uint32, mediaType, language string) *TrakBox {
	_ = "STUB: not implemented"
	return nil
}

// CreateEmptyTrak - create a full trak-tree for an empty (fragmented) track with no samples or stsd content
func CreateEmptyTrak(trackID, timeScale uint32, mediaType, language string) *TrakBox {
	_ = "STUB: not implemented"
	/*  Built tree like
	- trak
	+ tkhd (trakID, flags, width, height)
	+ mdia
	  - mdhd (track Timescale, language (3letters))
	  - hdlr (hdlr showing mediaType)
	  - elng (only if language is not 3 letters)
	  - minf
		+ vmhd/smhd etc (media header box)
		+ dinf
		  - dref
			+ url
		+ stbl
		  - stsd
			+ empty on purpose
		  - stts
		  - stsc
		  - stsz
		  - stco
	*/return nil
}

// Fixed 16 value 1.0

// SetAVCDescriptor - Set AVC SampleDescriptor based on SPS and PPS
func (t *TrakBox) SetAVCDescriptor(sampleDescriptorType string, spsNALUs, ppsNALUs [][]byte, includePS bool) error {
	_ = "STUB: not implemented"
	return nil
}

// This is display width
// This is display height

// SetHEVCDescriptor sets HEVC SampleDescriptor based on descriptorType, VPS, SPS, PPS and SEI.
func (t *TrakBox) SetHEVCDescriptor(sampleDescriptorType string, vpsNALUs, spsNALUs, ppsNALUs, seiNALUs [][]byte, includePS bool) error {
	_ = "STUB: not implemented"
	return nil
}

// This is display width
// This is display height

// hvc1 must include parameter sets (PS) and they must be complete
// hev1 may include PS and they may not be complete
// here we choose to include PS in both cases

// GetMediaType - should return video or audio (at present)
func (s *InitSegment) GetMediaType() string { _ = "STUB: not implemented"; return "" }

// TweakSingleTrakLive assures that there is only one track and removes any mehd box.
func (s *InitSegment) TweakSingleTrakLive() error { _ = "STUB: not implemented"; return nil }

// SetAACDescriptor - Modify a TrakBox by adding AAC SampleDescriptor
// objType is one of AAClc, HEAACv1, HEAACv2
// For HEAAC, the samplingFrequency is the base frequency (normally 24000)
func (t *TrakBox) SetAACDescriptor(objType byte, samplingFrequency int) error {
	_ = "STUB: not implemented"
	return nil
}

// SetAC3Descriptor  - Modify a TrakBox by adding AC-3 SampleDescriptor
func (t *TrakBox) SetAC3Descriptor(dac3 *Dac3Box) error { _ = "STUB: not implemented"; return nil }

//  Not to be used, but we set it anyway

// SetEC3Descriptor  - Modify a TrakBox by adding EC-3 SampleDescriptor
func (t *TrakBox) SetEC3Descriptor(dec3 *Dec3Box) error { _ = "STUB: not implemented"; return nil }

//  Not to be used, but we set it anyway

// SetWvttDescriptor - Set wvtt descriptor with a vttC box. config should start with WEBVTT or be empty.
func (t *TrakBox) SetWvttDescriptor(config string) error { _ = "STUB: not implemented"; return nil }

// SetStppDescriptor - add stpp box with utf8-lists namespace, schemaLocation and auxiliaryMimeType
// The utf8-lists have space-separated items, but no zero-termination
func (t *TrakBox) SetStppDescriptor(namespace, schemaLocation, auxiliaryMimeTypes string) error {
	_ = "STUB: not implemented"
	return nil
}
