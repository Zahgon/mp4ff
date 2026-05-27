package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

const (
	// boxHeaderSize - standard size + name header
	boxHeaderSize = 8
	largeSizeLen  = 8          // Length of largesize extension
	flagsMask     = 0x00ffffff // Flags for masks from full header
)

var decoders map[string]BoxDecoder

func init() {
	decoders = map[string]BoxDecoder{
		"\xa9ART": DecodeGenericContainerBox,
		"\xa9nam": DecodeGenericContainerBox,
		"\xa9too": DecodeGenericContainerBox,
		"\xa9cpy": DecodeGenericContainerBox,
		"ac-3":    DecodeAudioSampleEntry,
		"ac-4":    DecodeAudioSampleEntry,
		"alou":    DecodeLoudnessBaseBox,
		"av01":    DecodeVisualSampleEntry,
		"av1C":    DecodeAv1C,
		"avc1":    DecodeVisualSampleEntry,
		"avc3":    DecodeVisualSampleEntry,
		"av3c":    DecodeAv3c,
		"avcC":    DecodeAvcC,
		"avs3":    DecodeVisualSampleEntry,
		"btrt":    DecodeBtrt,
		"cdat":    DecodeCdat,
		"cdsc":    DecodeTrefType,
		"clap":    DecodeClap,
		"co64":    DecodeCo64,
		"CoLL":    DecodeCoLL,
		"colr":    DecodeColr,
		"cslg":    DecodeCslg,
		"ctim":    DecodeCtim,
		"ctts":    DecodeCtts,
		"dac3":    DecodeDac3,
		"dac4":    DecodeDac4,
		"mhaC":    DecodeMhaC,
		"data":    DecodeData,
		"dec3":    DecodeDec3,
		"dfLa":    DecodeDfLa,
		"dOps":    DecodeDops,
		"desc":    DecodeGenericContainerBox,
		"dinf":    DecodeDinf,
		"dpnd":    DecodeTrefType,
		"dref":    DecodeDref,
		"ec-3":    DecodeAudioSampleEntry,
		"edts":    DecodeEdts,
		"elng":    DecodeElng,
		"elst":    DecodeElst,
		"emeb":    DecodeEmeb,
		"emib":    DecodeEmib,
		"emsg":    DecodeEmsg,
		"enca":    DecodeAudioSampleEntry,
		"encv":    DecodeVisualSampleEntry,
		"esds":    DecodeEsds,
		"evte":    DecodeEvte,
		"fLaC":    DecodeAudioSampleEntry,
		"font":    DecodeTrefType,
		"free":    DecodeFree,
		"frma":    DecodeFrma,
		"ftyp":    DecodeFtyp,
		"hdlr":    DecodeHdlr,
		"hev1":    DecodeVisualSampleEntry,
		"hind":    DecodeTrefType,
		"hint":    DecodeTrefType,
		"hvc1":    DecodeVisualSampleEntry,
		"hvcC":    DecodeHvcC,
		"iacb":    DecodeIacb,
		"iamf":    DecodeAudioSampleEntry,
		"iden":    DecodeIden,
		"ID32":    DecodeID32,
		"ilst":    DecodeIlst,
		"iods":    DecodeUnknown,
		"ipir":    DecodeTrefType,
		"kind":    DecodeKind,
		"leva":    DecodeLeva,
		"ludt":    DecodeLudt,
		"mdat":    DecodeMdat,
		"mehd":    DecodeMehd,
		"mdhd":    DecodeMdhd,
		"mdia":    DecodeMdia,
		"meta":    DecodeMeta,
		"mfhd":    DecodeMfhd,
		"mfra":    DecodeMfra,
		"mfro":    DecodeMfro,
		"mha1":    DecodeAudioSampleEntry,
		"mha2":    DecodeAudioSampleEntry,
		"mhm1":    DecodeAudioSampleEntry,
		"mhm2":    DecodeAudioSampleEntry,
		"mime":    DecodeMime,
		"minf":    DecodeMinf,
		"moof":    DecodeMoof,
		"moov":    DecodeMoov,
		"mp4a":    DecodeAudioSampleEntry,
		"mpod":    DecodeTrefType,
		"mvex":    DecodeMvex,
		"mvhd":    DecodeMvhd,
		"nmhd":    DecodeNmhd,
		"Opus":    DecodeAudioSampleEntry,
		"pasp":    DecodePasp,
		"payl":    DecodePayl,
		"prft":    DecodePrft,
		"pssh":    DecodePssh,
		"saio":    DecodeSaio,
		"saiz":    DecodeSaiz,
		"sbgp":    DecodeSbgp,
		"schi":    DecodeSchi,
		"schm":    DecodeSchm,
		"sdtp":    DecodeSdtp,
		"senc":    DecodeSenc,
		"sgpd":    DecodeSgpd,
		"sidx":    DecodeSidx,
		"silb":    DecodeSilb,
		"sinf":    DecodeSinf,
		"skip":    DecodeFree,
		"SmDm":    DecodeSmDm,
		"smhd":    DecodeSmhd,
		"ssix":    DecodeSsix,
		"stbl":    DecodeStbl,
		"stco":    DecodeStco,
		"sthd":    DecodeSthd,
		"stpp":    DecodeStpp,
		"stsc":    DecodeStsc,
		"stsd":    DecodeStsd,
		"stss":    DecodeStss,
		"stsz":    DecodeStsz,
		"sttg":    DecodeSttg,
		"stts":    DecodeStts,
		"styp":    DecodeStyp,
		"subs":    DecodeSubs,
		"subt":    DecodeTrefType,
		"sync":    DecodeTrefType,
		"tenc":    DecodeTenc,
		"tfdt":    DecodeTfdt,
		"tfhd":    DecodeTfhd,
		"tfra":    DecodeTfra,
		"tkhd":    DecodeTkhd,
		"tlou":    DecodeLoudnessBaseBox,
		"traf":    DecodeTraf,
		"trak":    DecodeTrak,
		"tref":    DecodeTref,
		"trep":    DecodeTrep,
		"trex":    DecodeTrex,
		"trun":    DecodeTrun,
		"udta":    DecodeUdta,
		"url ":    DecodeURLBox,
		"uuid":    DecodeUUIDBox,
		"vdep":    DecodeTrefType,
		"vlab":    DecodeVlab,
		"vmhd":    DecodeVmhd,
		"vp08":    DecodeVisualSampleEntry,
		"vp09":    DecodeVisualSampleEntry,
		"vpcC":    DecodeVppC,
		"vplx":    DecodeTrefType,
		"vsid":    DecodeVsid,
		"vtta":    DecodeVtta,
		"vvc1":    DecodeVisualSampleEntry,
		"vvcC":    DecodeVvcC,
		"vvi1":    DecodeVisualSampleEntry,
		"vttc":    DecodeVttc,
		"vttC":    DecodeVttC,
		"vtte":    DecodeVtte,
		"wvtt":    DecodeWvtt,
	}
}

// RemoveBoxDecoder removes the decode of boxType. It will be treated as unknown instead.
//
// This is a global change, so use with care.
func RemoveBoxDecoder(boxType string) { _ = "STUB: not implemented"; return }

// SetBoxDecoder sets decoder functions for a specific boxType.
//
// This is a global change, so use with care.
func SetBoxDecoder(boxType string, dec BoxDecoder, decSR BoxDecoderSR) {
	_ = "STUB: not implemented"
	return
}

// BoxHeader - 8 or 16 bytes depending on size
type BoxHeader struct {
	Name   string
	Size   uint64
	Hdrlen int
}

func (b BoxHeader) payloadLen() int { _ = "STUB: not implemented"; return 0 }

// DecodeHeader decodes a box header (size + box type + possible largeSize)
func DecodeHeader(r io.Reader) (BoxHeader, error) {
	_ = "STUB: not implemented"
	return *new(BoxHeader), nil
}

// size 1 means large size in next 8 bytes

// size 0 means to end of file

// EncodeHeader - encode a box header to a writer
func EncodeHeader(b Box, w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeHeaderWithSize - encode a box header to a writer and allow for largeSize
func EncodeHeaderWithSize(boxType string, boxSize uint64, largeSize bool, w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// signals large size

// EncodeHeaderSW - encode a box header to a SliceWriter
func EncodeHeaderSW(b Box, sw bits.SliceWriter) error { _ = "STUB: not implemented"; return nil }

// EncodeHeaderWithSize - encode a box header to a writer and allow for largeSize
func EncodeHeaderWithSizeSW(boxType string, boxSize uint64, largeSize bool, sw bits.SliceWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// signals large size

// Box is the general interface to any ISOBMFF box or similar
type Box interface {
	// Type of box, normally 4 asccii characters, but is uint32 according to spec
	Type() string
	// Size of box including header and all children if any
	Size() uint64
	// Encode box to writer
	Encode(w io.Writer) error
	// Encode box to SliceWriter
	EncodeSW(sw bits.SliceWriter) error
	// Info - write box details
	//   spedificBoxLevels is a comma-separated list box:level or all:level where level >= 0.
	//   Higher levels give more details. 0 is default
	//   indent is indent at this box level.
	//   indentStep is how much to indent at each level
	Info(w io.Writer, specificBoxLevels, indent, indentStep string) error
}

// Informer - write box, segment or file details
type Informer interface {
	// Info - write details via Info method
	//   spedificBoxLevels is a comma-separated list box:level or all:level where level >= 0.
	//   Higher levels give more details. 0 is default
	//   indent is indent at this box level.
	//   indentStep is how much to indent at each level
	Info(w io.Writer, specificBoxLevels, indent, indentStep string) error
}

// BoxDecoder is function signature of the Box Decode method
type BoxDecoder func(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error)

// DecodeBox decodes a box
func DecodeBox(startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeBoxBody decodes a box body from an io.Reader given BoxHeader
func DecodeBoxBody(startPos uint64, hdr BoxHeader, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeBoxLazyMdat decodes a box but doesn't read mdat into memory
func DecodeBoxLazyMdat(startPos uint64, r io.ReadSeeker) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

func DecodeBoxBodyLazily(startPos uint64, h BoxHeader, r io.ReadSeeker) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// Fixed16 - An 8.8 fixed point number
type Fixed16 uint16

func (f Fixed16) String() string { _ = "STUB: not implemented"; return "" }

// Fixed32 -  A 16.16 fixed point number
type Fixed32 uint32

func (f Fixed32) String() string { _ = "STUB: not implemented"; return "" }

func strtobuf(out []byte, in string, l int) { _ = "STUB: not implemented"; return }

func makebuf(b Box) []byte { _ = "STUB: not implemented"; return nil }

// readBoxBody reads complete box body. Returns error if not possible
func readBoxBody(r io.Reader, h BoxHeader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
