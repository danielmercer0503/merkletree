package merkletree

type HashFunc func(data []byte) []byte

type MerkleTree struct {
	Leaves [][]byte
	Levels [][][]byte
	Hash   HashFunc
}

type Proof struct {
	Hash []byte
	Left bool
}
