# Package merkle 

We use segment trees to build the actual merkle trees for the underlying file system. The rationale behind using segment trees to build Merkle trees lies in their efficiency and effectiveness in organizing and representing data. Here's a breakdown of the approach and its rationale:

1. **File Representation:** 
   - Each file is treated as a slice of bytes. This representation is intuitive and aligns well with how files are stored and processed in computer systems.

2. **Group of Files Representation:**
   - A group of files is represented as a slice of slices of bytes. This hierarchical structure allows for organizing multiple files efficiently while maintaining individual file integrity.

3. **Segment Trees:**
   - Segment trees are utilized to organize the slices of bytes into a tree-like structure. Each leaf node in the segment tree represents the cryptographic hash
   of a single file in the set of files from `f0` to `fn`. In the segment tree, each leaf node has the same left and right index value. For example `l=r=0` represents
   the cryptographc hash (SHA256 in this case) of `f0`. 

4. **Efficient Construction:**
   - The use of segment trees allows for the efficient construction of Merkle trees from the given file data. By recursively dividing the data into segments and constructing nodes for each segment, the Merkle tree can be built in a logarithmic time complexity.

5. **Effective Proof Generation:**
   - Segment trees enable the generation of Merkle proofs for individual file indices or segments. These proofs provide a succinct and verifiable representation of the data's integrity, making it easy to verify the authenticity of specific file segments.

6. **Space Optimization:**
   - Segment trees optimize space utilization by representing file data and Merkle tree nodes in a structured and hierarchical manner. This helps in minimizing memory overhead while ensuring efficient access and manipulation of data.

In summary, the approach of using segment trees to build Merkle trees offers a balanced trade-off between efficiency, effectiveness, and space optimization. It leverages the inherent structure of file data and utilizes segment trees to organize and represent the data in a way that facilitates efficient construction, querying, and verification of Merkle trees and proofs.

## merkle.go

This file implements the Merkle tree data structure and related functionalities for building the tree, generating Merkle proofs, verifying proofs, and retrieving tree metadata.

- **TreeNode:** Represents a node in the Merkle tree, containing its hash value, left and right indices, and child nodes.

- **MerkleTree:** Represents the Merkle tree itself, consisting of a root node.

- **BuildMerkleTree:** Builds a Merkle tree recursively from the given file data.

- **GenerateMerkleProof:** Generates a Merkle proof for a specified leaf index. It uses the `genProof` function which is responsible for generating Merkle proofs for a given leaf node in a Merkle tree. Merkle proofs are cryptographic constructs that provide evidence of the inclusion or absence of a specific data item (represented by a leaf node) in the Merkle tree. This function takes as input the root node of the Merkle tree and the index of the leaf node for which the proof is to be generated. It traverses the tree from the leaf node to the root, collecting **sibling** nodes along the way. These sibling nodes, along with intermediate parent nodes, form the proof.
The function performs input validation to ensure the integrity of the Merkle tree structure and returns an error if the root node is nil or if the leaf index is out of bounds. Once the traversal is complete, the function constructs and returns an array of sibling nodes, which collectively form the Merkle proof for the specified leaf node. Overall, `genProof` facilitates the generation of Merkle proofs, a crucial aspect of ensuring data integrity and security.

- **VerifyMerkleProof:** Verifies a Merkle proof for a given file data and leaf index.

- **GetMerkleRoot:** Returns the root node of the Merkle tree.

- **PrintTreeInfo:** Prints information about the Merkle tree, including the number of nodes and its height.

- **CalcHash:** Calculates the SHA-256 hash of a byte slice and returns it as a hexadecimal string.

- **countNodes:** Counts the total number of nodes in the Merkle tree.

- **maxDepth:** Calculates the maximum depth of the Merkle tree.

## merkle_test.go

This file contains unit tests for the functionalities implemented in `merkle.go`. It covers scenarios for building Merkle trees, generating Merkle proofs, and verifying proofs.

- **TestMerkleTree:** Tests various scenarios of building Merkle trees from different file data.
- **TestMain:** Runs the tests defined in the file.

