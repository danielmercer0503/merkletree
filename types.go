package merkletree

type HashFunc func(data []byte) []byte

// HashFunc defines the function signature for hashing data in the Merkle tree.
type MerkleTree struct {
	// Leaves contains the hashed leaves of the Merkle tree.
	Leaves [][]byte
	// Levels contains all levels of the Merkle tree, from leaves to root.
	Levels [][][]byte
	// Hash is the hash function used by the Merkle tree.
	Hash   HashFunc
}

type Proof struct {
	// Hash is the sibling hash at this proof step.
	Hash []byte
	// Left indicates if the sibling is a left node.
	Left bool
}
