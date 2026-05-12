package merkletree

func (m *MerkleTree) GenerateProof(index int) ([]Proof, error) {

	var proof []Proof

	for level := 0; level < len(m.Levels)-1; level++ {
		layer := m.Levels[level]

		var siblingIndex int
		var left bool
		if index%2 == 0 {
			siblingIndex = index + 1
			left = false
		} else {
			siblingIndex = index - 1
			left = true
		}

		if siblingIndex < len(layer) {
			proof = append(proof, Proof{
				Hash: layer[siblingIndex],
				Left: left,
			})
		}

		index /= 2
	}

	return proof, nil
}

// GenerateProof generates a Merkle proof for the leaf at the given index.
// It returns a slice of Proof objects and an error if the index is invalid.
func VerifyProof(
	leaf []byte,
	proof []Proof,
	root []byte,
	hash HashFunc,
) bool {
	computed := hash(leaf)

	for _, p := range proof {
		if p.Left {
			computed = hash(append(p.Hash, computed...))
		} else {
			computed = hash(append(computed, p.Hash...))
		}
	}

	return string(computed) == string(root)
}

// VerifyProof verifies a Merkle proof for a given leaf and root using the provided hash function.
// It returns true if the proof is valid, false otherwise.
