package main

import (
	"crypto/sha256"
	"fmt"
)

type TreeNode struct {
	Hash     string
	LeftIdx  int
	RightIdx int
	Left     *TreeNode
	Right    *TreeNode
}

type MerkleTree struct {
	root *TreeNode
}

func BuildMerkleTree(file [][]byte) *MerkleTree {
	n := len(file)
	l, r := 0, n-1
	root := buildTree(file, l, r)
	return &MerkleTree{root: root}
}

func (mt *MerkleTree) GenerateMerkleProof(leafIdx int) []string {
	var result []string
	for _, node := range genProof(mt.root, leafIdx) {
		result = append(result, node.Hash)
	}
	return result
}

func genProofIdx(root *TreeNode, leafIdx int) [][]int {
	var result [][]int
	for _, node := range genProof(root, leafIdx) {
		result = append(result, []int{node.LeftIdx, node.RightIdx})
	}
	return result
}

func genProof(root *TreeNode, leafIdx int) []*TreeNode {
	var result []*TreeNode

	// Base cases - Empty tree
	if root == nil {
		return nil
	}

	// Only one node
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

func calculateHash(file []byte) string {
	hash := sha256.Sum256(file)
	return fmt.Sprintf("%x", hash)
}

func buildTree(file [][]byte, l, r int) *TreeNode {
	if l == r {
		return &TreeNode{Hash: calculateHash(file[l]), LeftIdx: l, RightIdx: r}
	}
	mid := l + (r-l)/2
	left := buildTree(file, l, mid)
	right := buildTree(file, mid+1, r)
	return &TreeNode{
		Hash:     calculateHash(append([]byte(left.Hash), []byte(right.Hash)...)),
		LeftIdx:  l,
		RightIdx: r,
		Left:     left,
		Right:    right,
	}
}

func printTree(node *TreeNode, prefix string, isLeft bool) {
	if node != nil {
		fmt.Printf("%s", prefix)
		if isLeft {
			fmt.Printf("├── L ")
		} else {
			fmt.Printf("└── R ")
		}
		fmt.Printf("(%d, %d) ==> %s \n", node.LeftIdx, node.RightIdx, node.Hash)

		// Recursive call for children TreeNodes
		printTree(node.Left, prefix+"│   ", true)
		printTree(node.Right, prefix+"    ", false)
	}
}

func countNodes(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + countNodes(node.Left) + countNodes(node.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDepth := 1 + maxDepth(root.Left)
	rightDepth := 1 + maxDepth(root.Right)
	return max(leftDepth, rightDepth)
}

func findLeaf(root *TreeNode, leafIndex int) *TreeNode {
	// Edge cases  - Empty tree or out-of bound index
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

func findParent(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil || root == node {
		return nil
	}
	if root.Left == node || root.Right == node {
		return root
	}
	// Recursively search in the left and right subtrees
	if parent := findParent(root.Left, node); parent != nil {
		return parent
	}
	return findParent(root.Right, node)
}

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

func findParentByLeafIndex(root *TreeNode, leafIdx int) *TreeNode {
	leaf := findLeaf(root, leafIdx)
	if leaf == nil || leaf == root {
		return nil
	}
	return findParent(root, leaf)
}

func findSiblingByLeafIndex(root *TreeNode, leafIdx int) *TreeNode {
	leaf := findLeaf(root, leafIdx)
	if leaf == nil || leaf == root {
		return nil
	}
	return findSibling(root, leaf)
}

func main() {
	files := [][]byte{[]byte("A"), []byte("B"), []byte("C"), []byte("D"), []byte("E"), []byte("F")}
	merkleTree := BuildMerkleTree(files)

	fmt.Println("************************************************  FILE **************************************************************************************")
	for i, file := range files {
		fmt.Printf("Hash of file %d  -> %s \n", i, calculateHash(file))
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
		fmt.Printf(" merkle proof for index:%d  is %d \n", idx, genProofIdx(merkleTree.root, idx))
	}
}
