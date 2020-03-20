package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestNewMerkleTree(t *testing.T) {
	data := [][]byte{
		[]byte("abc"),
		[]byte("123"),
		[]byte("+-/"),
		[]byte("abc"),
		[]byte("123"),

	}
	mk := MerkleTreeFromData(data)
	fmt.Println(hex.Dump(mk.RootNode.Hash))

	//mTree := MerkleTree{nil, [][]byte{}}
	mTree := NewMerkleTree()
	mTree.Append(sha256.Sum256(data[0]))
	mTree.Append(sha256.Sum256(data[1]))
	mTree.Append(sha256.Sum256(data[2]))
	mTree.Append(sha256.Sum256(data[3]))
	mTree.Append(sha256.Sum256(data[4]))

	mTree.CalTree()
	fmt.Println(hex.Dump(mTree.RootHash()))

}

