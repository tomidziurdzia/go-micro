package todo

import (
	"crypto/rand"
	"encoding/hex"
)

type ID string

func NewID() (ID, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return ID(hex.EncodeToString(b)), nil
}

func ParseID(raw string) (ID, error) {
	if len(raw) != 32 {
		return "", ErrInvalidID
	}
	// hex decoding validates allowed chars
	_, err := hex.DecodeString(raw)
	if err != nil {
		return "", ErrInvalidID
	}
	return ID(raw), nil
}
