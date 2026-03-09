package merkletree

import (
	"crypto/sha256"
)

func SHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func Keccak256(data []byte) []byte {
	// Placeholder for Keccak256 hash function implementation
	return nil
}
