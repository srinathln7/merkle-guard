package merkle

import (
	"crypto/sha256"
	"fmt"
	"log"

	mterr "github.com/srinathln7/merkle_gaurd/lib/err"
)

// TreeNode represents a node in the Merkle tree.
type TreeNode struct {
	Hash     string    // Hash value of the node
	LeftIdx  int       // Left index of the node
	RightIdx int       // Right index of the node
	Left     *TreeNode // Left child node
	Right    *TreeNode // Right child node
}

// MerkleTree represents a Merkle tree.
type MerkleTree struct {
	root *TreeNode // Root node of the Merkle tree
}

// BuildMerkleTree builds a Merkle tree from the given file data.
func BuildMerkleTree(file [][]byte) (*MerkleTree, error) {
	log.Println("[merkle-tree] starting to build merkle trees")
	n := len(file)
	if n == 0 {
		return nil, mterr.ErrEmptyFile
	}

	l, r := 0, n-1
	root := buildTree(file, l, r)
	return &MerkleTree{root: root}, nil
}

// GenerateMerkleProof generates a Merkle proof for the given leaf index.
func (mt *MerkleTree) GenerateMerkleProof(leafIdx int) ([]*TreeNode, error) {
	log.Printf("[merkle-tree] starting to generate merkle proof for file index %d with root %T \n", leafIdx, mt.root)
	return genProof(mt.root, leafIdx)
}

// VerifyMerkleProof verifies the Merkle proof for the given file data and leaf index.
func (mt *MerkleTree) VerifyMerkleProof(rootHash, fileHash string, fileIdx int, proofs []*TreeNode) (bool, error) {
	log.Printf("[merkle-tree] verifying merkle proof for file index %d with merkle root hash %s \n", fileIdx, mt.root.Hash)
	if mt.root == nil {
		return false, mterr.ErrEmptyRoot
	}

	if mt.root.Hash != rootHash {
		return false, mterr.ErrMerkleRootHashMisMatch
	}

	merkleHash := fileHash
	leaf, err := findLeaf(mt.root, fileIdx)
	if err != nil {
		return false, err
	}

	if leaf == nil {
		return false, mterr.ErrIndexOutOfBound
	}

	if leaf.Hash != merkleHash {
		return false, nil
	}

	// If the root has either a left or right child
	if mt.root.Left != nil || mt.root.Right != nil {
		curr := &TreeNode{}
		*curr = *leaf
		for _, proof := range proofs {
			if curr.LeftIdx < proof.LeftIdx && curr.RightIdx < proof.RightIdx {
				merkleHash = CalcHash(append([]byte(merkleHash), []byte(proof.Hash)...))
			} else {
				merkleHash = CalcHash(append([]byte(proof.Hash), []byte(merkleHash)...))
			}
			curr.LeftIdx = min(curr.LeftIdx, proof.LeftIdx)
			curr.RightIdx = max(curr.RightIdx, proof.RightIdx)
		}
	}

	return mt.root.Hash == merkleHash && rootHash == merkleHash, nil
}

// GetMerkleRoot returns the root node of the Merkle tree.
// If the MerkleTree instance is nil, it returns nil.
func (mt *MerkleTree) GetMerkleRoot() *TreeNode {
	if mt == nil {
		return nil
	}
	return mt.root
}

// PrintTreeInfo prints information about the Merkle tree.
// It displays the total number of nodes in the tree and its height.
// Additionally, it prints the Merkle tree structure.
func (mt *MerkleTree) PrintTreeInfo() {
	fmt.Println(" ******************************** Merkle Tree Metadata ***************************************************************")

	fmt.Printf("Total number of nodes: %d \n", countNodes(mt.root))
	fmt.Printf("Height of the merkle tree: %d \n", maxDepth(mt.root))

	fmt.Println(" ******************************** Merkle Tree  ***********************************************************************")
	printTree(mt.root, "", true)
}

// buildTree recursively builds the Merkle tree.
func buildTree(file [][]byte, l, r int) *TreeNode {
	if l == r {
		return &TreeNode{Hash: CalcHash(file[l]), LeftIdx: l, RightIdx: r}
	}
	mid := l + (r-l)/2
	left := buildTree(file, l, mid)
	right := buildTree(file, mid+1, r)
	return &TreeNode{
		Hash:     CalcHash(append([]byte(left.Hash), []byte(right.Hash)...)),
		LeftIdx:  l,
		RightIdx: r,
		Left:     left,
		Right:    right,
	}
}

// genProof generates a Merkle proof for the given leaf index.
func genProof(root *TreeNode, leafIdx int) ([]*TreeNode, error) {
	switch {
	case root == nil:
		return nil, mterr.ErrEmptyFile
	case leafIdx < root.LeftIdx || leafIdx > root.RightIdx:
		return nil, mterr.ErrIndexOutOfBound
	}

	if root.Left == nil && root.Right == nil {
		return []*TreeNode{root}, nil
	}

	var result []*TreeNode
	sibling, err := findSiblingByLeafIndex(root, leafIdx)
	if err != nil {
		return nil, err
	}
	result = append(result, sibling)

	parent, err := findParentByLeafIndex(root, leafIdx)
	if err != nil {
		return nil, err
	}

	for parent != root {
		sibling, _ = findSibling(root, parent)
		result = append(result, sibling)
		parent, _ = findParent(root, parent)
	}

	return result, nil
}

// genProofIdx generates proof indices for the leaf node corresponding to the given leaf index.
// It traverses the Merkle tree from the root to the leaf node, collecting the left and right indices
// of each node in the proof path and appends them to the result.
func genProofIdx(root *TreeNode, leafIdx int) ([][]int, error) {
	var result [][]int
	nodes, err := genProof(root, leafIdx)
	if err != nil {
		return nil, err
	}

	for _, node := range nodes {
		result = append(result, []int{node.LeftIdx, node.RightIdx})
	}
	return result, nil
}

// printTree prints the Merkle tree in a tree-like format.
func printTree(node *TreeNode, prefix string, isLeft bool) {
	if node != nil {
		fmt.Printf("%s", prefix)
		if isLeft {
			fmt.Printf("├── L ")
		} else {
			fmt.Printf("└── R ")
		}
		fmt.Printf("(%d, %d) ==> %s \n", node.LeftIdx, node.RightIdx, node.Hash)
		printTree(node.Left, prefix+"│   ", true)
		printTree(node.Right, prefix+"    ", false)
	}
}

// findLeaf finds the leaf node corresponding to the given leaf index.
func findLeaf(root *TreeNode, leafIdx int) (*TreeNode, error) {
	switch {
	case root == nil:
		return nil, mterr.ErrEmptyFile
	case leafIdx < root.LeftIdx || leafIdx > root.RightIdx:
		return nil, mterr.ErrIndexOutOfBound
	}

	if root.Left == nil && root.Right == nil && root.LeftIdx == leafIdx && root.RightIdx == leafIdx {
		return root, nil
	}
	midIdx := root.LeftIdx + (root.RightIdx-root.LeftIdx)/2
	if leafIdx <= midIdx {
		return findLeaf(root.Left, leafIdx)
	}
	return findLeaf(root.Right, leafIdx)
}

// findParent finds the parent node of the given node.
func findParent(root, node *TreeNode) (*TreeNode, error) {

	switch {
	case root == nil:
		return nil, mterr.ErrEmptyRoot
	case node == nil:
		return nil, mterr.ErrEmptyNode
	}

	if root == node {
		return nil, nil
	}

	if root.Left == node || root.Right == node {
		return root, nil
	}
	if parent, err := findParent(root.Left, node); parent != nil && err == nil {
		return parent, nil
	}
	return findParent(root.Right, node)
}

// findSibling finds the sibling node of the given node.
func findSibling(root, node *TreeNode) (*TreeNode, error) {
	parent, err := findParent(root, node)
	if err != nil {
		return nil, err
	}

	if parent == nil {
		return nil, nil
	}
	if parent.Left == node {
		return parent.Right, nil
	}
	return parent.Left, nil
}

// findParentByLeafIndex finds the parent node of the leaf node corresponding to the given leaf index.
func findParentByLeafIndex(root *TreeNode, leafIdx int) (*TreeNode, error) {
	leaf, err := findLeaf(root, leafIdx)
	if err != nil {
		return nil, err
	}

	if leaf == root {
		return nil, nil
	}

	return findParent(root, leaf)
}

// findSiblingByLeafIndex finds the sibling node of the leaf node corresponding to the given leaf index.
func findSiblingByLeafIndex(root *TreeNode, leafIdx int) (*TreeNode, error) {
	leaf, err := findLeaf(root, leafIdx)
	if err != nil {
		return nil, err
	}

	if leaf == nil || leaf == root {
		return nil, nil
	}
	return findSibling(root, leaf)
}

// calcHash calculates the SHA-256 hash of the given byte slice and returns it as a hexadecimal string.
func CalcHash(file []byte) string {
	hash := sha256.Sum256(file)
	return fmt.Sprintf("%x", hash)
}

// countNodes counts the total number of nodes in the Merkle tree.
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// maxDepth calculates the maximum depth of the Merkle tree.
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := 1 + maxDepth(root.Left)
	rightDepth := 1 + maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}
