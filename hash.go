package merkletree

import (
	"crypto/sha256"
)

// SHA256 returns the SHA-256 hash of the input data.
func SHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// Keccak256 is a placeholder for the Keccak256 hash function implementation.
func Keccak256(data []byte) []byte {
	// Placeholder for Keccak256 hash function implementation
	return nil
}
