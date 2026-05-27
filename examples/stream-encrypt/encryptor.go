package main

import (
	"github.com/Eyevinn/mp4ff/mp4"
)

type EncryptConfig struct {
	Key    []byte
	KeyID  []byte
	IV     []byte
	Scheme string
}

type StreamEncryptor struct {
	config        EncryptConfig
	ipd           *mp4.InitProtectData
	fragNum       uint32
	encryptedInit *mp4.InitSegment
}

func NewStreamEncryptor(init *mp4.InitSegment, config EncryptConfig) (*StreamEncryptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (se *StreamEncryptor) GetEncryptedInit() *mp4.InitSegment {
	_ = "STUB: not implemented"
	return nil
}

func (se *StreamEncryptor) EncryptFragment(frag *mp4.Fragment) error {
	_ = "STUB: not implemented"
	return nil
}

func (se *StreamEncryptor) deriveIV(fragNum uint32) []byte { _ = "STUB: not implemented"; return nil }

func ParseHexKey(s string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
