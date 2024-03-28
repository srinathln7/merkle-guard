package test

import (
	"os"
	"testing"
)

// Run the tests
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGRPCServer(t *testing.T) {

	// We need a shared-server instance to mimic persistent storage during the test runtime
	// Hence, we setup the grpc client before running each of the individual test cases
	grpcClient, teardown := SetupGRPCClient(t, nil)

	// gracefully shutdown the server after finishing all the test cases
	defer teardown()

	// To ensure that the test cases are run sequentially, we use subtests.
	// Subtests are a way to group related tests together and run them in a specific order.
	// Each test case is run in its own subtest to provide better isolation, readability, and reporting.

	t.Run("merkle verification success", func(t *testing.T) {
		testClientMerkleVerficationSuccess(t, grpcClient)
	})

}
