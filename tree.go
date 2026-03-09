package merkletree

func NewMerkleTree(data [][]byte, hash HashFunc) *MerkleTree {

	var leaves [][]byte

	for _, d := range data {
		leaves = append(leaves, hash(d))
	}

	tree := &MerkleTree{
		Leaves: leaves,
		Hash:   hash,
	}

	tree.buildTree()

	return tree
}

func (m *MerkleTree) buildTree() {

	current := m.Leaves
	m.Levels = append(m.Levels, current)

	for len(current) > 1 {
		var next [][]byte
		for i := 0; i < len(current); i += 2 {

			if i+1 == len(current) {
				next = append(next, current[i])
				continue
			}

			combined := append(current[i], current[i+1]...)
			parent := m.Hash(combined)

			next = append(next, parent)
		}

		m.Levels = append(m.Levels, next)
		current = next
	}
}

func (m *MerkleTree) Root() []byte {
	last := m.Levels[len(m.Levels)-1]
	return last[0]
}
