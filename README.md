# merkletree

A simple and efficient Merkle Tree implementation in Go.

## Features

- Build Merkle Trees from arbitrary data
- Supports custom hash functions (e.g., SHA256, Keccak256)
- Generate and verify Merkle proofs
- Easy integration and usage

## Installation

```bash
go get github.com/danielmercer0503/merkletree
```

## Usage

```go
package main

import (
    "github.com/danielmercer0503/merkletree"
    "fmt"
)

func main() {
    data := [][]byte{
        []byte("a"),
        []byte("b"),
        []byte("c"),
        []byte("d"),
    }

    tree := merkletree.NewMerkleTree(data, merkletree.SHA256)
    root := tree.Root()
    fmt.Printf("Merkle Root: %x\n", root)

    proof, _ := tree.GenerateProof(2) // Proof for "c"
    valid := merkletree.VerifyProof(data[2], proof, root, merkletree.SHA256)
    fmt.Printf("Proof valid: %v\n", valid)
}
```

## API

- `NewMerkleTree(data [][]byte, hash HashFunc) *MerkleTree`
- `(*MerkleTree) Root() []byte`
- `(*MerkleTree) GenerateProof(index int) ([]Proof, error)`
- `VerifyProof(leaf []byte, proof []Proof, root []byte, hash HashFunc) bool`
- `SHA256(data []byte) []byte`
- `Keccak256(data []byte) []byte` (placeholder)

## Types

- `MerkleTree`
- `Proof`
- `HashFunc`

## License

MIT
