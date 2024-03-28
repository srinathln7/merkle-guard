package merkle

import (
	"log"
	"testing"

	mterr "github.com/srinathln7/merkle_gaurd/lib/err"
	"github.com/stretchr/testify/require"
)

func TestMerkleTree(t *testing.T) {
	tests := []struct {
		name  string
		files [][]byte
	}{
		{
			name:  "EmptyFile",
			files: [][]byte{},
		},
		{
			name: "SingleFile",
			files: [][]byte{
				[]byte("A"),
			},
		},
		{
			name: "FourFiles",
			files: [][]byte{
				[]byte("A"), []byte("B"), []byte("C"), []byte("D"),
			},
		},
		{
			name: "FiveFiles",
			files: [][]byte{
				[]byte("A"), []byte("B"), []byte("C"), []byte("D"),
				[]byte("E"),
			},
		},
		{
			name: "TwentySixFiles",
			files: [][]byte{
				[]byte("A"), []byte("B"), []byte("C"), []byte("D"),
				[]byte("E"), []byte("F"), []byte("G"), []byte("H"),
				[]byte("I"), []byte("J"), []byte("K"), []byte("L"),
				[]byte("M"), []byte("N"), []byte("O"), []byte("P"),
				[]byte("Q"), []byte("R"), []byte("S"), []byte("T"),
				[]byte("U"), []byte("V"), []byte("W"), []byte("X"),
				[]byte("Y"), []byte("Z"),
			},
		},
	}

	// Test for Empty file at idx = 0
	t.Run(tests[0].name, func(t *testing.T) {
		_, err := BuildMerkleTree(tests[0].files)
		require.Error(t, err, mterr.ErrEmptyFile)
	})

	// Merkle proof test for "FiveFiles"
	t.Run(tests[3].name, func(t *testing.T) {
		merkleTree, err := BuildMerkleTree(tests[3].files)
		require.NoError(t, err)
		var merkleProof [][]int
		for fileIdx := range tests[3].files {
			merkleProofIdx, err := genProofIdx(merkleTree.root, fileIdx)
			require.NoError(t, err)
			merkleProof = append(merkleProof, merkleProofIdx...)
		}
		require.Equal(t, merkleProof, [][]int{
			{1, 1}, {2, 2}, {3, 4},
			{0, 0}, {2, 2}, {3, 4},
			{0, 1}, {3, 4},
			{4, 4}, {0, 2},
			{3, 3}, {0, 2},
		},
		)
	})

	// Verification test for non-empty files
	for i := 1; i < len(tests); i++ {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			merkleTree, err := BuildMerkleTree(test.files)
			require.NoError(t, err)

			log.Println("test case - merkle verification forward")
			for idx, file := range test.files {
				merkleProofs, err := merkleTree.GenerateMerkleProof(idx)
				require.NoError(t, err)
				isVerified, err := merkleTree.VerifyMerkleProof(merkleTree.root.Hash, CalcHash(file), idx, merkleProofs)
				require.NoError(t, err)
				require.True(t, isVerified, "merkle proof verification failed for test %s at file index %d \n", test.name, idx)
			}

			log.Println("test case - merkle verification reverse")
			for idx := len(test.files) - 1; idx >= 0; idx-- {
				merkleProofs, err := merkleTree.GenerateMerkleProof(idx)
				require.NoError(t, err)
				isVerified, err := merkleTree.VerifyMerkleProof(merkleTree.root.Hash, CalcHash(test.files[idx]), idx, merkleProofs)
				require.NoError(t, err)
				require.True(t, isVerified, "merkle proof verification failed for test %s at file index %d \n", test.name, idx)
			}

		})
	}
}

// Run the tests
func TestMain(m *testing.M) {
	m.Run()
}
