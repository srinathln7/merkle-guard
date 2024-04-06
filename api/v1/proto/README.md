# GRPC Proto File Specification

The `merkle.proto` protocol buffer file defines the message types and service methods for uploading, downloading, and verifying files using a Merkle tree over gRPC.

1. `syntax = "proto3";`: This line specifies that the protocol buffer file is written in version 3 of the protocol buffers syntax.

2. `package merkle_gaurd;`: This line specifies the package name for the protocol buffer file. It's similar to a namespace in Go, used to avoid naming conflicts.

3. `option go_package = "github.com/srinathln7/api/merkle_gaurd";`: This line specifies the Go package name for the generated Go code. It indicates the directory structure where the generated Go files will be placed.

4. `message UploadRequest { ... }`: This block defines the `UploadRequest` message, which is used to send a request to upload files to the server. It contains a repeated field `files`, which is a list of bytes representing the files to be uploaded.

5. `message UploadResponse { ... }`: This block defines the `UploadResponse` message, which is the response to an upload request. It contains a single field `merkle_root_hash`, which is a byte array representing the Merkle root hash of the uploaded files.

6. `message DownloadRequest { ... }`: This block defines the `DownloadRequest` message, which is used to request downloading a file from the server. It contains a single field `file_index`, an integer representing the index of the file to download.

7. `message DownloadResponse { ... }`: This block defines the `DownloadResponse` message, which is the response to a download request. It contains a single field `file_content`, a byte array representing the content of the downloaded file.

8. `message MerkleProofRequest { ... }`: This block defines the `MerkleProofRequest` message, which is used to request a Merkle proof for a specific file from the server. It contains a single field `file_index`, an integer representing the index of the file for which the proof is requested.

9. `message TreeNode { ... }`: This block defines the `TreeNode` message, which represents a node in a Merkle tree. It contains fields `hash`, `left_idx`, `right_idx`, `left`, and `right`, representing the hash value of the node, indices of its left and right children, and references to its left and right child nodes.

10. `message MerkleProofResponse { ... }`: This block defines the `MerkleProofResponse` message, which is the response to a Merkle proof request. It contains a repeated field `proofs`, which is a list of `TreeNode` messages representing the Merkle proof.

11. `message VerifyProofRequest { ... }`: This block defines the `VerifyProofRequest` message, which is used to request verification of a Merkle proof. It contains fields `root_hash`, `file_hash`, `file_index`, and `proofs`, representing the root hash of the Merkle tree, hash of the file, index of the file, and the Merkle proof.

12. `message VerifyProofResponse { ... }`: This block defines the `VerifyProofResponse` message, which is the response to a Merkle proof verification request. It contains a single field `is_verified`, a boolean indicating whether the proof is verified.

13. `service MerkleTree { ... }`: This block defines the `MerkleTree` service, which contains RPC methods for interacting with the Merkle tree. It specifies four RPC methods: `Upload`, `Download`, `GetMerkleProof`, and `VerifyMerkleProof`, each with its request and response message types.

