package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Eyevinn/mp4ff/mp4"
)

const (
	appName = "mp4ff-crop"
)

var usg = `%s crops a (progressive) mp4 file to just before a sync frame after specified number of milliseconds.
The goal is to leave the file structure intact except for cropping of samples and
moving mdat to the end of the file, if not already there.

Usage of %s:
`

type options struct {
	durationMS uint
	version    bool
}

func parseOptions(fs *flag.FlagSet, args []string) (*options, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error { _ = "STUB: not implemented"; return nil }

func cropMP4(inMP4 *mp4.File, durationMS int, w io.Writer, ifh io.ReadSeeker) error {
	_ = "STUB: not implemented"
	return nil
}

// findEndTime - find closest video sync frame, or audio frame if no video
func findEndTime(moov *mp4.MoovBox, durationMS int) (endTime, endTimescale uint64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

//trakDur := float64(trak.Tkhd.Duration) / float64(moov.Mvhd.Timescale)
//fmt.Printf("video trak %d duration = %.3fs\n", trak.Tkhd.TrackID, trakDur)

// TimeToSampleBox

func cropToTime(inMP4 *mp4.File, endTime, endTimescale uint64, w io.Writer, ifh io.ReadSeeker) error {
	_ = "STUB: not implemented"
	return nil
}

type trakOut struct {
	lastSampleNr  uint32
	endTime       uint64
	lastChunk     mp4.Chunk
	nextInChunkNr uint32
	chunkOffsets  []uint64
}

// findTrakEnds - find where traks end in form of last chunk, lastSampleNr and endTime
func findTrakEnds(traks []*mp4.TrakBox, endTime, endTimescale uint64) (map[uint32]*trakOut, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fillTrakOutsAndByteRanges(traks []*mp4.TrakBox, tos map[uint32]*trakOut, byteRanges *byteRanges) (firstOffset uint64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//Done

// updateChunkOffsets - calculate new moov size, and update stco/co64 (chunk offsets)
func updateChunkOffsets(inMP4 *mp4.File, firstOffset uint64) { _ = "STUB: not implemented"; return }

func writeUptoMdat(inMP4 *mp4.File, endTime, endTimescale uint64, w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func writeMdat(byteRanges *byteRanges, mdatIn *mp4.MdatBox, w io.Writer, ifh io.ReadSeeker) error {
	_ = "STUB: not implemented"
	// write mdat header
	return nil
}

// write mdat body

func minUint32(a, b uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func cropStblChildren(traks []*mp4.TrakBox, trakOuts map[uint32]*trakOut) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func cropStts(b *mp4.SttsBox, lastSampleNr uint32) { _ = "STUB: not implemented"; return }

func cropStss(b *mp4.StssBox, lastSampleNr uint32) { _ = "STUB: not implemented"; return }

func cropCtts(b *mp4.CttsBox, lastSampleNr uint32) { _ = "STUB: not implemented"; return }

// Finally cut down the endSampleNr for this index

func cropStsc(b *mp4.StscBox, lastSampleNr uint32) error { _ = "STUB: not implemented"; return nil }

func cropStsz(b *mp4.StszBox, lastSampleNr uint32) { _ = "STUB: not implemented"; return }

func cropSdtp(b *mp4.SdtpBox, lastSampleNr uint32) { _ = "STUB: not implemented"; return }

func updateStco(b *mp4.StcoBox, offsets []uint64) { _ = "STUB: not implemented"; return }

func updateCo64(b *mp4.Co64Box, offsets []uint64) { _ = "STUB: not implemented"; return }

type byteRange struct {
	start uint64
	end   uint64 // Included
}

type byteRanges struct {
	ranges []byteRange
}

func createByteRanges() *byteRanges { _ = "STUB: not implemented"; return nil }

func (b *byteRanges) addRange(start, end uint64) { _ = "STUB: not implemented"; return }

func (b *byteRanges) size() uint64 { _ = "STUB: not implemented"; return 0 }
