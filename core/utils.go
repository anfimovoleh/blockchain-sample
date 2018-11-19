package core

import (
	"crypto/sha256"
	"encoding/hex"
)

func CalculateHash(block Block) string {
	h := sha256.New()
	h.Write([]byte(block.String()))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
