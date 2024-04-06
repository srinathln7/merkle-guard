# Package test

The test files in this package are used to validate the functionality of the server and client implementations respectively. Let's break down each one:

## `server_test.go`

1. **TestMain Function**:
   - It sets up the test environment by initializing any necessary resources. In this case, it just runs all the tests.

2. **TestGRPCServer Function**:
   - Sets up the gRPC client and runs a series of subtests to validate different aspects of the server functionality.

   - Subtests:
     - **merkle verification for four files success**: Tests the successful verification of Merkle trees for four files.
     - **merkle verification for empty file**: Tests the behavior when an empty file is uploaded.
     - **merkle root mis-match**: Tests the behavior when the calculated Merkle root hash mismatches the expected hash.

## `client_test.go`

1. **SetupGRPCClient Function**:
   - Sets up the gRPC client for testing purposes.
   - Binds the gRPC client to a random port and initializes the server.

2. **Individual Test Functions**:
   - **testClientMerkleVerficationSuccess**: Tests the successful verification of Merkle trees for a set of files.
   - **testClientMerkleVerficationEmptyFile**: Tests the behavior when attempting to upload an empty file.
   - **testClientMerkleRootMisMatch**: Tests the behavior when the calculated Merkle root hash mismatches the expected hash.

These tests ensure that the server and client implementations behave correctly under various scenarios, including successful cases, error handling, and edge cases. They provide a safety net for developers to refactor and extend the codebase while maintaining the expected functionality.