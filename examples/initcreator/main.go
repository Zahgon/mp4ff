package main

import (
	"fmt"
	"os"

	"github.com/Eyevinn/mp4ff/mp4"
)

const (
	avcSPSnalu  = "67640020accac05005bb0169e0000003002000000c9c4c000432380008647c12401cb1c31380"
	avcPPSnalu  = "68b5df20"
	hevcVPSnalu = "40010c01ffff022000000300b0000003000003007b18b024"
	hevcSPSnalu = "420101022000000300b0000003000003007ba0078200887db6718b92448053888892cf24a69272c9124922dc91aa48fca223ff000100016a02020201"
	hevcPPSnalu = "4401c0252f053240"
)

func main() {
	if err := run("."); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}

func run(outDir string) error { _ = "STUB: not implemented"; return nil }

func writeVideoAVCInitSegment(outPath string) error { _ = "STUB: not implemented"; return nil }

func writeVideoHEVCInitSegment(outPath string) error { _ = "STUB: not implemented"; return nil }

func writeAudioAACInitSegment(outPath string) error { _ = "STUB: not implemented"; return nil }

func writeAudioAC3InitSegment(outPath string) error { _ = "STUB: not implemented"; return nil }

func writeAudioEC3InitSegment(outPath string) error { _ = "STUB: not implemented"; return nil }

func writeSubtitlesWvttInitSegment(outPath string) error { _ = "STUB: not implemented"; return nil }

func writeSubtitlesStppInitSegment(outPath string) error { _ = "STUB: not implemented"; return nil }

func writeToFile(init *mp4.InitSegment, filePath string) error {
	_ = "STUB: not implemented"
	// Next write to a file
	return nil
}
