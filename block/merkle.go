package block

import (
	"crypto/sha256"
)

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  []byte
}

//创建结构体
type MerkleTree struct {
	RootNode *MerkleNode
	Hashs  [][]byte
}

func NewMerkleTree() *MerkleTree {
	return &MerkleTree{nil, [][]byte{}}
}

func (self *MerkleTree) RootHash()(hash []byte) {
	return self.RootNode.Hash
}

func (self *MerkleTree) Append(hash [32]byte) {
	self.Hashs = append(self.Hashs, hash[:])
}

func (self *MerkleTree) CalTree() {
	//循环一层一层的生成节点，知道到最上面的根节点为止
	nodes := self.AddLeafs()
	for i := 0; i < (len(self.Hashs) + 1) / 2; i++ {
		var levels []MerkleNode
		for j := 0; j < len(nodes); j += 2 {
			left, right := j, j+1
			if right == len(nodes) {
				right = left
			}
			node := self.AddNodes(&nodes[left], &nodes[right], nil)
			levels = append(levels, *node)
		}
		nodes = levels
	}
	self.RootNode = &nodes[0]
}

func (self *MerkleTree) AddLeafs() (leafs []MerkleNode) {
	//通过数据生成叶子节点
	len := len(self.Hashs)

	for _, hash := range self.Hashs {
		//输入的交易个数如果是单数的话，就复制最后一个，成为复数
		node := self.AddNodes(nil, nil, hash)
		leafs = append(leafs, *node)
	}
	if len % 2!= 0{
		node := self.AddNodes(nil, nil, self.Hashs[len -1])
		leafs = append(leafs, *node)
	}
	return
}

func (self *MerkleTree) AddNodes(left, right *MerkleNode, hash []byte) *MerkleNode {
	merkleNode := MerkleNode{}
	if left == nil && right == nil {
		//叶子节点
		merkleNode.Hash = hash[:]
	} else {
		prevHashes := append(left.Hash, right.Hash...)
		hash := sha256.Sum256(prevHashes)
		merkleNode.Hash = hash[:]
	}
	merkleNode.Left = left
	merkleNode.Right = right
	return &merkleNode
}


//生成一颗新树
func MerkleTreeFromData(data [][]byte) *MerkleTree {
	var nodes []MerkleNode
	//输入的交易个数如果是单数的话，就复制最后一个，成为复数
	if len(data) % 2 != 0 {
		data = append(data, data[len(data)-1])
	}
	//通过数据生成叶子节点
	for _, dataNode := range data {
		node := NewMerkleNode(nil, nil, dataNode)
		nodes = append(nodes, *node)
	}
	//循环一层一层的生成节点，知道到最上面的根节点为止
	for i := 0; i < len(data) / 2; i++ {
		var newLevel []MerkleNode
		for j := 0; j < len(nodes); j += 2 {
			left, right := j, j+1
			if right == len(nodes) {
				right = left
			}
			node := NewMerkleNode(&nodes[left], &nodes[right], nil)
			newLevel = append(newLevel, *node)
		}
		nodes = newLevel
	}
	mTree := MerkleTree{&nodes[0], nil}
	return &mTree
}

//创建一个新的节点
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	merkleNode := MerkleNode{}
	if left == nil && right == nil {
		//叶子节点
		hash := sha256.Sum256(data)
		merkleNode.Hash = hash[:]
	} else {
		prevHashes := append(left.Hash, right.Hash...)
		hash := sha256.Sum256(prevHashes)
		merkleNode.Hash = hash[:]
	}
	merkleNode.Left = left
	merkleNode.Right = right
	return &merkleNode
}
