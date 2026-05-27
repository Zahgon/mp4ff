package mp4

import (
	"github.com/Eyevinn/mp4ff/bits"
)

var decodersSR map[string]BoxDecoderSR

func init() {
	decodersSR = map[string]BoxDecoderSR{
		"\xa9ART": DecodeGenericContainerBoxSR,
		"\xa9cpy": DecodeGenericContainerBoxSR,
		"\xa9nam": DecodeGenericContainerBoxSR,
		"\xa9too": DecodeGenericContainerBoxSR,
		"ac-3":    DecodeAudioSampleEntrySR,
		"ac-4":    DecodeAudioSampleEntrySR,
		"alou":    DecodeLoudnessBaseBoxSR,
		"av01":    DecodeVisualSampleEntrySR,
		"av1C":    DecodeAv1CSR,
		"avc1":    DecodeVisualSampleEntrySR,
		"avc3":    DecodeVisualSampleEntrySR,
		"av3c":    DecodeAv3cSR,
		"avcC":    DecodeAvcCSR,
		"avs3":    DecodeVisualSampleEntrySR,
		"btrt":    DecodeBtrtSR,
		"cdat":    DecodeCdatSR,
		"cdsc":    DecodeTrefTypeSR,
		"clap":    DecodeClapSR,
		"co64":    DecodeCo64SR,
		"CoLL":    DecodeCoLLSR,
		"colr":    DecodeColrSR,
		"cslg":    DecodeCslgSR,
		"ctim":    DecodeCtimSR,
		"ctts":    DecodeCttsSR,
		"dac3":    DecodeDac3SR,
		"dac4":    DecodeDac4SR,
		"mhaC":    DecodeMhaCSR,
		"data":    DecodeDataSR,
		"dec3":    DecodeDec3SR,
		"dfLa":    DecodeDfLaSR,
		"dOps":    DecodeDopsSR,
		"desc":    DecodeGenericContainerBoxSR,
		"dinf":    DecodeDinfSR,
		"dpnd":    DecodeTrefTypeSR,
		"dref":    DecodeDrefSR,
		"ec-3":    DecodeAudioSampleEntrySR,
		"edts":    DecodeEdtsSR,
		"elng":    DecodeElngSR,
		"elst":    DecodeElstSR,
		"emeb":    DecodeEmebSR,
		"emib":    DecodeEmibSR,
		"emsg":    DecodeEmsgSR,
		"enca":    DecodeAudioSampleEntrySR,
		"encv":    DecodeVisualSampleEntrySR,
		"esds":    DecodeEsdsSR,
		"evte":    DecodeEvteSR,
		"fLaC":    DecodeAudioSampleEntrySR,
		"font":    DecodeTrefTypeSR,
		"free":    DecodeFreeSR,
		"frma":    DecodeFrmaSR,
		"ftyp":    DecodeFtypSR,
		"hdlr":    DecodeHdlrSR,
		"hev1":    DecodeVisualSampleEntrySR,
		"hind":    DecodeTrefTypeSR,
		"hint":    DecodeTrefTypeSR,
		"hvc1":    DecodeVisualSampleEntrySR,
		"hvcC":    DecodeHvcCSR,
		"iacb":    DecodeIacbSR,
		"iamf":    DecodeAudioSampleEntrySR,
		"iden":    DecodeIdenSR,
		"ID32":    DecodeID32SR,
		"ilst":    DecodeIlstSR,
		"iods":    DecodeUnknownSR,
		"ipir":    DecodeTrefTypeSR,
		"kind":    DecodeKindSR,
		"leva":    DecodeLevaSR,
		"ludt":    DecodeLudtSR,
		"mdat":    DecodeMdatSR,
		"mehd":    DecodeMehdSR,
		"mdhd":    DecodeMdhdSR,
		"mdia":    DecodeMdiaSR,
		"meta":    DecodeMetaSR,
		"mfhd":    DecodeMfhdSR,
		"mfra":    DecodeMfraSR,
		"mfro":    DecodeMfroSR,
		"mha1":    DecodeAudioSampleEntrySR,
		"mha2":    DecodeAudioSampleEntrySR,
		"mhm1":    DecodeAudioSampleEntrySR,
		"mhm2":    DecodeAudioSampleEntrySR,
		"mime":    DecodeMimeSR,
		"minf":    DecodeMinfSR,
		"moof":    DecodeMoofSR,
		"moov":    DecodeMoovSR,
		"mp4a":    DecodeAudioSampleEntrySR,
		"mpod":    DecodeTrefTypeSR,
		"mvex":    DecodeMvexSR,
		"mvhd":    DecodeMvhdSR,
		"nmhd":    DecodeNmhdSR,
		"Opus":    DecodeAudioSampleEntrySR,
		"pasp":    DecodePaspSR,
		"payl":    DecodePaylSR,
		"prft":    DecodePrftSR,
		"pssh":    DecodePsshSR,
		"saio":    DecodeSaioSR,
		"saiz":    DecodeSaizSR,
		"sbgp":    DecodeSbgpSR,
		"schi":    DecodeSchiSR,
		"schm":    DecodeSchmSR,
		"sdtp":    DecodeSdtpSR,
		"senc":    DecodeSencSR,
		"sgpd":    DecodeSgpdSR,
		"sidx":    DecodeSidxSR,
		"silb":    DecodeSilbSR,
		"sinf":    DecodeSinfSR,
		"skip":    DecodeFreeSR,
		"SmDm":    DecodeSmDmSR,
		"smhd":    DecodeSmhdSR,
		"ssix":    DecodeSsixSR,
		"stbl":    DecodeStblSR,
		"stco":    DecodeStcoSR,
		"sthd":    DecodeSthdSR,
		"stpp":    DecodeStppSR,
		"stsc":    DecodeStscSR,
		"stsd":    DecodeStsdSR,
		"stss":    DecodeStssSR,
		"stsz":    DecodeStszSR,
		"sttg":    DecodeSttgSR,
		"stts":    DecodeSttsSR,
		"styp":    DecodeStypSR,
		"subs":    DecodeSubsSR,
		"subt":    DecodeTrefTypeSR,
		"sync":    DecodeTrefTypeSR,
		"tenc":    DecodeTencSR,
		"tfdt":    DecodeTfdtSR,
		"tfhd":    DecodeTfhdSR,
		"tfra":    DecodeTfraSR,
		"tkhd":    DecodeTkhdSR,
		"tlou":    DecodeLoudnessBaseBoxSR,
		"traf":    DecodeTrafSR,
		"trak":    DecodeTrakSR,
		"tref":    DecodeTrefSR,
		"trep":    DecodeTrepSR,
		"trex":    DecodeTrexSR,
		"trun":    DecodeTrunSR,
		"udta":    DecodeUdtaSR,
		"url ":    DecodeURLBoxSR,
		"uuid":    DecodeUUIDBoxSR,
		"vdep":    DecodeTrefTypeSR,
		"vlab":    DecodeVlabSR,
		"vmhd":    DecodeVmhdSR,
		"vp08":    DecodeVisualSampleEntrySR,
		"vp09":    DecodeVisualSampleEntrySR,
		"vpcC":    DecodeVppCSR,
		"vplx":    DecodeTrefTypeSR,
		"vsid":    DecodeVsidSR,
		"vtta":    DecodeVttaSR,
		"vvc1":    DecodeVisualSampleEntrySR,
		"vvcC":    DecodeVvcCSR,
		"vvi1":    DecodeVisualSampleEntrySR,
		"vttc":    DecodeVttcSR,
		"vttC":    DecodeVttCSR,
		"vtte":    DecodeVtteSR,
		"wvtt":    DecodeWvttSR,
	}
}

// BoxDecoderSR is function signature of the Box DecodeSR method
type BoxDecoderSR func(hdr BoxHeader, startPos uint64, sw bits.SliceReader) (Box, error)

// DecodeBoxSR - decode a box from SliceReader
func DecodeBoxSR(startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeHeaderSR - decode a box header (size + box type + possible largeSize) from sr
func DecodeHeaderSR(sr bits.SliceReader) (BoxHeader, error) {
	_ = "STUB: not implemented"
	return *new(BoxHeader), nil
}

// size 1 means large size in next 8 bytes

// size 0 means to end of file

// DecodeBoxBodySR - decode box body from SliceReader given BoxHeader
func DecodeBoxBodySR(startPos uint64, hdr BoxHeader, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// In the following, we do not block mdat to allow for the case
// that the first kiloBytes of a file are fetched and parsed to
// get the init part of a file. In the future, a new decode option that
// stops before the mdat starts is a better alternative.

// DecodeFile - parse and decode a file from reader r with optional file options.
// For example, the file options overwrite the default decode or encode mode.
func DecodeFileSR(sr bits.SliceReader, options ...Option) (*File, error) {
	_ = "STUB: not implemented"

	// apply options to change the default decode or encode mode
	return nil, nil
}

// No moov and heuristic failed.
// Leave senc deferred for caller to parse later with init info.
