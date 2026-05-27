package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

/* Elementary Stream Descriptors are defined in ISO/IEC 14496-1.
The full spec looks like, below, but we don't parse all that stuff.
*/

const (
	// Following Table 1 of Class Tags for descriptors in ISO/IEC 14496-1. There are more types
	ObjectDescrTag        = 1
	InitialObjectDescrTag = 2
	ES_DescrTag           = 3
	DecoderConfigDescrTag = 4
	DecSpecificInfoTag    = 5
	SLConfigDescrTag      = 6
)

func TagType(tag byte) string { _ = "STUB: not implemented"; return "" }

type Descriptor interface {
	// Tag - descriptor tag. Fixed for each descriptor type
	Tag() byte
	// Type is string describing Tag, making descriptor fulfill boxLike interface
	Type() string
	// Size - size of descriptor, excluding tag byte and size field
	Size() uint64
	// SizeSize - size of descriptor including tag byte and size field
	SizeSize() uint64
	// EncodeSW - Write descriptor to slice writer
	EncodeSW(sw bits.SliceWriter) error
	// Info - write information about descriptor
	//   Higher levels give more details. 0 is default
	//   indent is indent at this box level.
	//   indentStep is how much to indent at each level
	Info(w io.Writer, specificLevels, indent, indentStep string) error
}

/*
ESDescriptor is defined in ISO/IEC 14496-1 7.2.6.5

	class ES_Descriptor extends BaseDescriptor : bit(8) tag=ES_DescrTag {
	  bit(16) ES_ID;
	  bit(1) streamDependenceFlag;
	  bit(1) URL_Flag;
	  bit(1) OCRstreamFlag;
	  bit(5) streamPriority;
	  if (streamDependenceFlag)
	    bit(16) dependsOn_ES_ID;
	  if (URL_Flag) {
	    bit(8) URLlength;
	    bit(8) URLstring[URLlength];
	  }
	  if (OCRstreamFlag)
	    bit(16) OCR_ES_Id;
	  DecoderConfigDescriptor decConfigDescr;
	  if (ODProfileLevelIndication==0x01) //no SL extension.
	  {
	    SLConfigDescriptor slConfigDescr;
	  } else  { // SL extension is possible.
	    SLConfigDescriptor slConfigDescr;
	  }
	  IPI_DescrPointer ipiPtr[0 .. 1];
	  IP_IdentificationDataSet ipIDS[0 .. 255];
	  IPMP_DescriptorPointer ipmpDescrPtr[0 .. 255];
	  LanguageDescriptor langDescr[0 .. 255];
	  QoS_Descriptor qosDescr[0 .. 1];
	  RegistrationDescriptor regDescr[0 .. 1];
	  ExtensionDescriptor extDescr[0 .. 255];
	}
*/
type ESDescriptor struct {
	EsID                uint16
	DependsOnEsID       uint16
	OCResID             uint16
	FlagsAndPriority    byte
	sizeFieldSizeMinus1 byte
	URLString           string
	DecConfigDescriptor *DecoderConfigDescriptor
	SLConfigDescriptor  *SLConfigDescriptor
	OtherDescriptors    []Descriptor
	UnknownData         []byte // Data, probably erroneous, that we don't understand
}

func DecodeDescriptor(sr bits.SliceReader, maxNrBytes int) (Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(Descriptor), nil
}

func DecodeESDescriptor(sr bits.SliceReader, descSize uint32) (ESDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(ESDescriptor), nil
}

// streamPriority := ed.FlagsAndPriority & 0x1f

func (e *ESDescriptor) Tag() byte { _ = "STUB: not implemented"; return 0 }

func (e *ESDescriptor) Type() string { _ = "STUB: not implemented"; return "" }

// Size is size of payload after tag and size field
func (e *ESDescriptor) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// SizeSize is size of size field.
func (e *ESDescriptor) SizeSize() uint64 { _ = "STUB: not implemented"; return 0 }

func (e *ESDescriptor) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// streamPriority := ed.FlagsAndPriority & 0x1f

/* no zero-termination */

func (e *ESDescriptor) Info(w io.Writer, specificLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// DecoderConfigDescriptor is defined in ISO/IEC 14496-1 Section 7.2.6.6.1
//
//	class DecoderConfigDescriptor extends BaseDescriptor : bit(8) tag=DecoderConfigDescrTag {
//	  bit(8) objectTypeIndication;
//	  bit(6) streamType;
//	  bit(1) upStream;
//	  const bit(1) reserved=1;
//	  bit(24) bufferSizeDB;
//	  bit(32) maxBitrate;
//	  bit(32) avgBitrate;
//	  DecoderSpecificInfo decSpecificInfo[0 .. 1];
//	  profileLevelIndicationIndexDescriptor profileLevelIndicationIndexDescr [0..255];
//	}
type DecoderConfigDescriptor struct {
	ObjectType          byte
	StreamType          byte
	sizeFieldSizeMinus1 byte
	BufferSizeDB        uint32
	MaxBitrate          uint32
	AvgBitrate          uint32
	DecSpecificInfo     *DecSpecificInfoDescriptor
	OtherDescriptors    []Descriptor
	UnknownData         []byte // Data, probably erroneous, that we don't understand
}

func exceedsMaxNrBytes(sizeFieldSizeMinus1 byte, size uint64, maxNrBytes int) bool {
	_ = "STUB: not implemented"
	return false
}

func DecodeDecoderConfigDescriptor(tag byte, sr bits.SliceReader, maxNrBytes int) (Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(Descriptor), nil
}

// The optional decoderSpeicificInfo is not present

func (d *DecoderConfigDescriptor) Tag() byte { _ = "STUB: not implemented"; return 0 }

func (d *DecoderConfigDescriptor) Type() string { _ = "STUB: not implemented"; return "" }

func (d *DecoderConfigDescriptor) Size() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *DecoderConfigDescriptor) SizeSize() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *DecoderConfigDescriptor) EncodeSW(sw bits.SliceWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DecoderConfigDescriptor) Info(w io.Writer, specificLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// DecSpecificInfoDescriptor is a generic DecoderSpecificInfoDescriptor.
//
// The meaning of the MPEG-4 audio descriptor is defined in  ISO/IEC 14496-3 Section 1.6.2.1.

type DecSpecificInfoDescriptor struct {
	sizeFieldSizeMinus1 byte
	DecConfig           []byte
}

func DecodeDecSpecificInfoDescriptor(tag byte, sr bits.SliceReader, maxNrBytes int) (Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(Descriptor), nil
}

func (d *DecSpecificInfoDescriptor) Tag() byte { _ = "STUB: not implemented"; return 0 }

func (d *DecSpecificInfoDescriptor) Type() string { _ = "STUB: not implemented"; return "" }

func (d *DecSpecificInfoDescriptor) Size() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *DecSpecificInfoDescriptor) SizeSize() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *DecSpecificInfoDescriptor) EncodeSW(sw bits.SliceWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DecSpecificInfoDescriptor) Info(w io.Writer, specificLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

type SLConfigDescriptor struct {
	sizeFieldSizeMinus1 byte
	ConfigValue         byte
	MoreData            []byte
}

func DecodeSLConfigDescriptor(tag byte, sr bits.SliceReader, maxNrBytes int) (Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(Descriptor), nil
}

func (d *SLConfigDescriptor) Tag() byte { _ = "STUB: not implemented"; return 0 }

func (d *SLConfigDescriptor) Type() string { _ = "STUB: not implemented"; return "" }

func (d *SLConfigDescriptor) Size() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *SLConfigDescriptor) SizeSize() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *SLConfigDescriptor) EncodeSW(sw bits.SliceWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *SLConfigDescriptor) Info(w io.Writer, specificLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// RawDescriptor - raw representation of any descriptor
type RawDescriptor struct {
	tag                 byte
	sizeFieldSizeMinus1 byte
	data                []byte
}

func DecodeRawDescriptor(tag byte, sr bits.SliceReader, maxNrBytes int) (Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(Descriptor), nil
}

func CreateRawDescriptor(tag, sizeFieldSizeMinus1 byte, data []byte) (RawDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(RawDescriptor), nil
}

func (s *RawDescriptor) Tag() byte { _ = "STUB: not implemented"; return 0 }

func (d *RawDescriptor) Type() string { _ = "STUB: not implemented"; return "" }

func (d *RawDescriptor) Size() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *RawDescriptor) SizeSize() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *RawDescriptor) EncodeSW(sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

func (d *RawDescriptor) Info(w io.Writer, specificLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateESDescriptor creates an ESDescriptor with a DecoderConfigDescriptor for audio.
func CreateESDescriptor(decConfig []byte) ESDescriptor {
	_ = "STUB: not implemented"
	return *new(ESDescriptor)
}

// Audio ISO/IEC 14496-3,
// 0x5 << 2 + 0x01 (audioType + upstreamFlag + reserved)

// readTagAndSize - get size by accumulate 7 bits from each byte. MSB = 1 indicates more bytes.
// Defined in ISO 14496-1 Section 8.3.3
func readSizeSize(sr bits.SliceReader) (sizeFieldSizeMinus1 byte, size uint64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// writeDescriptorSize - write descriptor size 7-bit at a time in as many bytes as prescribed
func writeDescriptorSize(sw bits.SliceWriter, size uint64, sizeFieldSizeMinus1 byte) {
	_ = "STUB: not implemented"
	return
}
