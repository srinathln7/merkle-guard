# Package client 

This package represents a client implementation for interacting with a gRPC server providing Merkle tree-related functionalities. Here's a high-level overview of what the code does:

1. **Initialization**:
   - The client initializes its gRPC connection to the server by dialing the server address obtained from the environment variables.

2. **Handling Uploads**:
   - The client can upload files to the server by calling the `Upload` function, which sends a gRPC request containing the files to the server.
   - Upon successful upload, the client stores the resulting Merkle root hash on its disk.

3. **Handling Downloads**:
   - Clients can request to download specific files from the server by calling the `Download` function, which sends a gRPC request for the file content based on the file index.
   - The downloaded file content is returned to the client.

4. **Generating Merkle Proofs**:
   - Clients can request Merkle proofs for specific files from the server by calling the `GetMerkleProof` function, which sends a gRPC request for the Merkle proof based on the file index.
   - The generated Merkle proofs are returned to the client.

5. **Verifying Merkle Proofs**:
   - Clients can verify Merkle proofs for specific files by calling the `VerifyMerkleProof` function, which sends a gRPC request containing the Merkle proofs, file index, file content, and root hash to the server.
   - The server verifies the provided Merkle proofs against its stored Merkle tree and returns the verification result to the client.

Overall, this client provides a convenient interface for interacting with the Merkle tree server, allowing users to upload, download, generate proofs, and verify file integrity using Merkle trees over gRPC.