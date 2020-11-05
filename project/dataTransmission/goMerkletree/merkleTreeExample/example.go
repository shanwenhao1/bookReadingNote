/*
	struct说明
	// merkle tree结构
	type MerkleTree struct {
		Root         *Node				// 根节点
		merkleRoot   []byte				// 根节点的hash值
		Leafs        []*Node			// 叶子节点的信息
		hashStrategy func() hash.Hash
	}
	// 节点结构
	type Node struct {
		Tree   *MerkleTree
		Parent *Node					// 父节点信息
		Left   *Node					// 左节点
		Right  *Node					// 右节点
		leaf   bool						// 是否是叶子节点
		dup    bool						// 该叶子节点是否重复(实际传输过程中, 带该标签的数据可以剔除)
		Hash   []byte
		C      Content
	}
*/
package merkleTreeExample

import (
	"crypto/sha256"
	"github.com/cbergoon/merkletree"
	"log"
)

// TestContent 实现了github.com/cbergoon/merkletree中的content interface, 用来存放merkle tree
type TestContent struct {
	x string
}

// Hash加密
func (t TestContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(t.x)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// 校验数据与hash值是否对应上s
func (t TestContent) Equals(other merkletree.Content) (bool, error) {
	return t.x == other.(TestContent).x, nil
}

func Example() {
	// 模拟切片后的数据, 用于生成merkle tree.
	var listCon []merkletree.Content
	for _, content := range []string{"Hello", "Hi", "Hey", "Hola"} {
		listCon = append(listCon, TestContent{content})
	}

	//Create a new Merkle Tree from the list of Content
	// merkle tree就生成好了, 这里生成的是完全二叉树, 叶子节点为奇数的最后一个叶子会重复一次
	t, err := merkletree.NewTree(listCon)
	if err != nil {
		log.Fatal(err)
	}

	// ------------------- 接下来就是验证example创建的数据是否正确----------------
	//Get the Merkle Root(hash) of the tree
	mr := t.MerkleRoot()
	log.Println(mr)

	//Verify the entire tree (hashes for each node) is valid(对比整个merkle tree的hash是否一致, 即生成的根节点的hash值
	//与接收的整个tree的数据接收生成的根hash是否一致)
	vt, err := t.VerifyTree()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Verify Tree: ", vt)

	//Verify a specific content in in the tree(验证某个分支节点的hash是否一致)
	vc, err := t.VerifyContent(listCon[0])
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Verify Content: ", vc)

	//String representation
	log.Println("================")
	log.Println(t)
}
