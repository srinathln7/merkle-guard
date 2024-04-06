# Package server

This package implements a gRPC server for handling Merkle tree-related operations, such as file uploads, downloads, generating Merkle proofs, and verifying Merkle proofs. Here's a high-level overview of its functionality:

1. **Initialization**:
   - The server is initialized with the required dependencies and configurations, such as loading environment variables and setting up a TCP listener.

2. **Handling Uploads**:
   - When a client uploads files, the server constructs a Merkle tree based on the uploaded files.
   - The resulting Merkle tree is stored along with the uploaded files.

3. **Handling Downloads**:
   - Clients can request to download a specific file by providing its index.
   - The server retrieves the requested file content from the stored files and sends it back to the client.

4. **Generating Merkle Proofs**:
   - Clients can request Merkle proofs for specific files.
   - The server generates the Merkle proof for the requested file using the stored Merkle tree and sends it back to the client.

5. **Verifying Merkle Proofs**:
   - Clients can request to verify Merkle proofs for specific files.
   - The server verifies the provided Merkle proofs against the stored Merkle tree and returns the verification result to the client.

Overall, this server facilitates secure file operations using Merkle trees over a gRPC interface, providing functionalities for file uploads, downloads, and integrity verification.