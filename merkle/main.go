package main

import (
	"crypto/sha256"
	"fmt"
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
func BuildMerkleTree(file [][]byte) *MerkleTree {
	n := len(file)
	l, r := 0, n-1
	root := buildTree(file, l, r)
	return &MerkleTree{root: root}
}

// GenerateMerkleProof generates a Merkle proof for the given leaf index.
func (mt *MerkleTree) GenerateMerkleProof(leafIdx int) []*TreeNode {
	return genProof(mt.root, leafIdx)
}

// VerifyMerkleProof verifies the Merkle proof for the given file data and leaf index.
func (mt *MerkleTree) VerifyMerkleProof(file []byte, leafIdx int, proofs []*TreeNode) bool {

	if mt.root == nil {
		return false
	}

	merkleHash := calcHash(file)
	leaf := findLeaf(mt.root, leafIdx)
	if leaf == nil || leaf.Hash != merkleHash {
		return false
	}

	// If the root has either a left or right child
	if mt.root.Left != nil || mt.root.Right != nil {
		curr := &TreeNode{}
		*curr = *leaf
		for _, proof := range proofs {
			if curr.LeftIdx < proof.LeftIdx && curr.RightIdx < proof.RightIdx {
				merkleHash = calcHash(append([]byte(merkleHash), []byte(proof.Hash)...))
			} else {
				merkleHash = calcHash(append([]byte(proof.Hash), []byte(merkleHash)...))
			}
			curr.LeftIdx = min(curr.LeftIdx, proof.LeftIdx)
			curr.RightIdx = max(curr.RightIdx, proof.RightIdx)
		}
	}

	return mt.root.Hash == merkleHash
}

// buildTree recursively builds the Merkle tree.
func buildTree(file [][]byte, l, r int) *TreeNode {
	if l == r {
		return &TreeNode{Hash: calcHash(file[l]), LeftIdx: l, RightIdx: r}
	}
	mid := l + (r-l)/2
	left := buildTree(file, l, mid)
	right := buildTree(file, mid+1, r)
	return &TreeNode{
		Hash:     calcHash(append([]byte(left.Hash), []byte(right.Hash)...)),
		LeftIdx:  l,
		RightIdx: r,
		Left:     left,
		Right:    right,
	}
}

// genProof generates a Merkle proof for the given leaf index.
func genProof(root *TreeNode, leafIdx int) []*TreeNode {
	var result []*TreeNode
	if root == nil || (leafIdx < root.LeftIdx || leafIdx > root.RightIdx) {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []*TreeNode{root}
	}
	sibling := findSiblingByLeafIndex(root, leafIdx)
	result = append(result, sibling)
	parent := findParentByLeafIndex(root, leafIdx)
	for parent != root {
		sibling = findSibling(root, parent)
		result = append(result, sibling)
		parent = findParent(root, parent)
	}
	return result
}

// genProofIdx generates proof indices for the leaf node corresponding to the given leaf index.
// It traverses the Merkle tree from the root to the leaf node, collecting the left and right indices
// of each node in the proof path and appends them to the result.
func genProofIdx(root *TreeNode, leafIdx int) [][]int {
	var result [][]int
	for _, node := range genProof(root, leafIdx) {
		result = append(result, []int{node.LeftIdx, node.RightIdx})
	}
	return result
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
func findLeaf(root *TreeNode, leafIndex int) *TreeNode {
	if root == nil || (leafIndex < root.LeftIdx || leafIndex > root.RightIdx) {
		return nil
	}
	if root.Left == nil && root.Right == nil && root.LeftIdx == leafIndex && root.RightIdx == leafIndex {
		return root
	}
	midIdx := root.LeftIdx + (root.RightIdx-root.LeftIdx)/2
	if leafIndex <= midIdx {
		return findLeaf(root.Left, leafIndex)
	}
	return findLeaf(root.Right, leafIndex)
}

// findParent finds the parent node of the given node.
func findParent(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil || root == node {
		return nil
	}
	if root.Left == node || root.Right == node {
		return root
	}
	if parent := findParent(root.Left, node); parent != nil {
		return parent
	}
	return findParent(root.Right, node)
}

// findSibling finds the sibling node of the given node.
func findSibling(root, node *TreeNode) *TreeNode {
	parent := findParent(root, node)
	if parent == nil {
		return nil
	}
	if parent.Left == node {
		return parent.Right
	}
	return parent.Left
}

// findParentByLeafIndex finds the parent node of the leaf node corresponding to the given leaf index.
func findParentByLeafIndex(root *TreeNode, leafIdx int) *TreeNode {
	leaf := findLeaf(root, leafIdx)
	if leaf == nil || leaf == root {
		return nil
	}
	return findParent(root, leaf)
}

// findSiblingByLeafIndex finds the sibling node of the leaf node corresponding to the given leaf index.
func findSiblingByLeafIndex(root *TreeNode, leafIdx int) *TreeNode {
	leaf := findLeaf(root, leafIdx)
	if leaf == nil || leaf == root {
		return nil
	}
	return findSibling(root, leaf)
}

// calcHash calculates the SHA-256 hash of the given byte slice and returns it as a hexadecimal string.
func calcHash(file []byte) string {
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

func main() {
	files := [][]byte{
		[]byte("A"), []byte("B"), []byte("C"), []byte("D"),
		[]byte("E"), []byte("F"), []byte("G"), []byte("H"),
		[]byte("I"), []byte("J"), []byte("K"), []byte("L"),
		[]byte("M"), []byte("N"), []byte("O"), []byte("P"),
		[]byte("Q"), []byte("R"), []byte("S"), []byte("T"),
		[]byte("U"), []byte("V"), []byte("W"), []byte("X"),
		[]byte("Y"), []byte("Z"),
	}

	merkleTree := BuildMerkleTree(files)

	fmt.Println("************************************************  FILE **************************************************************************************")
	for i, file := range files {
		fmt.Printf("Hash of file %d  -> %s \n", i, calcHash(file))
	}

	fmt.Println("************************************************  METADATA **************************************************************************************")
	fmt.Println("Root Hash:", merkleTree.root.Hash)
	fmt.Println("Total nodes:", countNodes(merkleTree.root))
	fmt.Println("Height:", maxDepth(merkleTree.root))

	fmt.Println("************************************************  MERKLE TREE  **************************************************************************************")
	fmt.Println(merkleTree.root)
	printTree(merkleTree.root, "", true)
	fmt.Println("************************************************  MERKLE TREE  **************************************************************************************")

	fmt.Println("************************************************  GENERATE PROOF  **************************************************************************************")

	for idx := range files {
		fmt.Printf("merkle proof for index:%d  is %d \n", idx, genProofIdx(merkleTree.root, idx))
	}

	fmt.Println("************************************************  VERIFY PROOF FORWARD **************************************************************************************")

	n := len(files)
	for idx := 0; idx <= n-1; idx++ {
		proofs := merkleTree.GenerateMerkleProof(idx)
		fmt.Printf("merkle proof for index:%d  is %d \n", idx, genProofIdx(merkleTree.root, idx))
		fmt.Printf("merkle verification for index:%d is: %t \n", idx, merkleTree.VerifyMerkleProof(files[idx], idx, proofs))
	}

	fmt.Println("************************************************  VERIFY PROOF REVERSE **************************************************************************************")

	for idx := n - 1; idx >= 0; idx-- {
		proofs := merkleTree.GenerateMerkleProof(idx)
		fmt.Printf("merkle proof for index:%d  is %d \n", idx, genProofIdx(merkleTree.root, idx))
		fmt.Printf("merkle verification for index:%d is: %t \n", idx, merkleTree.VerifyMerkleProof(files[idx], idx, proofs))
	}

}
