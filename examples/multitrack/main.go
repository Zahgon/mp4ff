// multitrack - decode example multitrack fragmented file with video and closed caption tracks
package main

import (
	"io"
	"log"
	"os"

	"github.com/Eyevinn/mp4ff/mp4"
)

const (
	// filePath is start of a track of track v1 from
	// https://devstreaming-cdn.apple.com/videos/streaming/examples/bipbop_adv_example_hevc/master.m3u8
	filePath = "testdata/main_1.mp4"
)

// Track - information for an mp4 track
type Track struct {
	trackID   uint32
	hdlrType  string
	timeScale uint64
	trak      *mp4.TrakBox
	trex      *mp4.TrexBox
	samples   []mp4.FullSample
}

func main() {
	ifd, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer ifd.Close()

	tracks, err := getTracksAndSamplesFromMultiTrackFragmentedFile(ifd)
	if err != nil {
		log.Fatalln(err)
	}

	err = writeTrackInfo(os.Stdout, tracks)
	if err != nil {
		log.Fatalln(err)
	}

	for _, track := range tracks {
		if track.hdlrType == "clcp" {
			err = writeScenaristFile(os.Stdout, track)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func getTracksAndSamplesFromMultiTrackFragmentedFile(ifd io.Reader) (tracks []*Track, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func writeTrackInfo(w io.Writer, tracks []*Track) error { _ = "STUB: not implemented"; return nil }

// Should contain cdat boxes with CEA-608 byte pairs

// writeScenaristFile - write file from clcp track with cdat samples
func writeScenaristFile(w io.Writer, clcpTrack *Track) error { _ = "STUB: not implemented"; return nil }

// timeFromMs - return time string hh:mm:ss:fr where fr is frame (~29.97Hz)
func timeFromMs(tMs uint64) string { _ = "STUB: not implemented"; return "" }
