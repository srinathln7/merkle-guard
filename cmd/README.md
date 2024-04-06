# cmd.go

This file defines the command-line interface (CLI) for interacting with the Merkle-Guard server. It provides commands for uploading files, downloading files, fetching Merkle proofs, and verifying Merkle proofs.

- **SetupFlags:** Sets up command-line flags for the CLI options such as file index, upload directory, merkle root hash directory, download directory, and merkle proofs directory.

- **RootCmd:** Represents the root command of the CLI. It prints welcome messages and instructions for using the CLI, and sets up a signal handler for graceful termination.

- **uploadCmd:** Defines the `upload` command, which uploads a set of files to the server. It sets up a gRPC client, reads files from the specified directory, uploads the files to the server, and writes the merkle root hash to a file.

- **downloadCmd:** Defines the `download` command, which downloads a file from the server. It sets up a gRPC client, downloads the file corresponding to the specified index, and writes it to the specified directory.

- **getMerkleProofsCmd:** Defines the `getMerkleProofs` command, which fetches merkle proofs for a file from the server. It sets up a gRPC client, fetches the merkle proofs, and writes them to a file.

- **verifyMerkleProofsCmd:** Defines the `verifyMerkleProofs` command, which verifies merkle proofs for a file. It sets up a gRPC client, reads the merkle root hash and file, fetches the merkle proofs, verifies them, and prints the verification result.

