package mp4

import (
	"io"

	"github.com/Eyevinn/mp4ff/bits"
)

// ContainerBox is interface for ContainerBoxes
type ContainerBox interface {
	Type() string
	Size() uint64
	Encode(w io.Writer) error
	EncodeSW(w bits.SliceWriter) error
	GetChildren() []Box
	Info(w io.Writer, specificBoxLevels, indent, indentStep string) error
}

// GenericContainerBox is a generic container box with no special child pointers
type GenericContainerBox struct {
	name     string
	Children []Box
}

func NewGenericContainerBox(name string) *GenericContainerBox {
	_ = "STUB: not implemented"
	return nil
}

func (b *GenericContainerBox) Type() string { _ = "STUB: not implemented"; return "" }

func (b *GenericContainerBox) Size() uint64 { _ = "STUB: not implemented"; return 0 }

// Encode - write GenericContainerBox to w
func (b *GenericContainerBox) Encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Encode - write minf container to sw
func (b *GenericContainerBox) EncodeSW(sw bits.SliceWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// Info - write box-specific information
func (b *GenericContainerBox) Info(w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetChildren - list of child boxes
func (b *GenericContainerBox) GetChildren() []Box {
	_ = "STUB: not implemented"

	// DecodeGenericContainerBox - box-specific decode
	return nil
}

func DecodeGenericContainerBox(hdr BoxHeader, startPos uint64, r io.Reader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// DecodeGenericContainerBoxSR - box-specific decode
func DecodeGenericContainerBoxSR(hdr BoxHeader, startPos uint64, sr bits.SliceReader) (Box, error) {
	_ = "STUB: not implemented"
	return *new(Box), nil
}

// AddChild - Add a child box
func (b *GenericContainerBox) AddChild(child Box) { _ = "STUB: not implemented"; return }

func containerSize(children []Box) uint64 { _ = "STUB: not implemented"; return 0 }

// DecodeContainerChildren decodes a container box
func DecodeContainerChildren(hdr BoxHeader, startPos, endPos uint64, r io.Reader) ([]Box, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DecodeContainerChildren decodes a container box
func DecodeContainerChildrenSR(hdr BoxHeader, startPos, endPos uint64, sr bits.SliceReader) ([]Box, error) {
	_ = "STUB: not implemented"
	return nil,
		// Good initial size
		nil
}

// EncodeContainer - marshal container c to w
func EncodeContainer(c ContainerBox, w io.Writer) error { _ = "STUB: not implemented"; return nil }

// EncodeContainerSW - marshal container c to sw
func EncodeContainerSW(c ContainerBox, sw bits.SliceWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// ContainerInfo - write container-box information
func ContainerInfo(c ContainerBox, w io.Writer, specificBoxLevels, indent, indentStep string) error {
	_ = "STUB: not implemented"
	return nil
}
