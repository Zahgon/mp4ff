package mp4

import (
	"io"
)

const (
	infoVersionNone         = -1
	infoVersionGroupingType = -2
	infoVersionDescriptor   = -3
)

type boxLike interface {
	Type() string
	Size() uint64
	Info(w io.Writer, specificBoxLevels, indent, indentStep string) error
}

// infoDumper - dump box name and size. Allow for more with write.
type infoDumper struct {
	w      io.Writer
	indent string
	box    boxLike
	err    error
}

// fixStartingCopyrightChar - replace starting one byte © with two-bytes UTF-8
func fixStartingCopyrightChar(boxType string) string {
	_ = "STUB: not implemented"
	// © is 0xa9 in latin1 (and in Apple boxes/atoms)
	// In UTF-8 it is two bytes: 0xc2 0xa9
	return ""
}

// newInfoDumper - make an infoDumper with indent
// write version if >= 0
// set Version to -1 if not present for box
// set Version to -2 for sample group entries
// set Version to -3 for descriptors
func newInfoDumper(w io.Writer, indent string, b boxLike, version int, flags uint32) *infoDumper {
	_ = "STUB: not implemented"
	return nil
}

// write - write formatted objecds if level <= bd.level
func (b *infoDumper) write(format string, p ...interface{}) { _ = "STUB: not implemented"; return }

// getInfoLevel - get info level for specific boxLike, or from all
func getInfoLevel(b boxLike, specificBoxLevels string) (level int) {
	_ = "STUB: not implemented"
	return 0
}
