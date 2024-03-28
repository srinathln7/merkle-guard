package test

import (
	"context"
	"net"
	"testing"

	"github.com/srinathln7/merkle_gaurd/internal/server"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	api "github.com/srinathln7/merkle_gaurd/api/v1/proto"
	mt "github.com/srinathln7/merkle_gaurd/internal/merkle"
)

// SetupGRPCClient: sets up the grpc client
func SetupGRPCClient(t *testing.T, fn func()) (
	grpcClient api.MerkleTreeClient,
	teardown func(),
) {
	// Helper marks the calling function as a test helper function.
	// When printing file and line information, that function will be skipped
	t.Helper()

	// Get a free available port for testing purpose
	// Alt. we can also use the `.env` variable port
	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	grpcClientOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	cc, err := grpc.Dial(listener.Addr().String(), grpcClientOptions...)
	require.NoError(t, err)

	grpcServer, err := server.NewgrpcServer()
	require.NoError(t, err)

	go func() {
		grpcServer.Serve(listener)
	}()

	grpcClient = api.NewMerkleTreeClient(cc)

	return grpcClient, func() {
		grpcServer.Stop()
		cc.Close()
		listener.Close()
	}
}

func testClientMerkleVerficationSuccess(t *testing.T, grpcClient api.MerkleTreeClient) {
	ctx := context.Background()
	files := [][]byte{
		[]byte("A"), []byte("B"), []byte("C"), []byte("D"),
	}

	expectedResp := &api.UploadResponse{
		MerkleRootHash: []byte("50a504831bd50fee3581d287168a85a8dcdd6aa777ffd0fe35e37290268a0153"),
	}

	// First upload the file
	uploadResp, err := grpcClient.Upload(
		ctx,
		&api.UploadRequest{
			Files: files,
		},
	)

	require.NoError(t, err)
	require.Equal(t, expectedResp.MerkleRootHash, uploadResp.MerkleRootHash)

	// Start downloading files in a psuedo-random order
	fileIdxs := []int{2, 3, 0, 1}
	for _, fileIdx := range fileIdxs {
		expectedResp := &api.DownloadResponse{FileContent: files[fileIdx]}
		downloadResp, err := grpcClient.Download(
			ctx,
			&api.DownloadRequest{
				FileIndex: int64(fileIdx),
			},
		)
		require.NoError(t, err)
		require.Equal(t, expectedResp.FileContent, downloadResp.FileContent)
	}

	for _, fileIdx := range fileIdxs {
		recvResp, err := grpcClient.GetMerkleProof(
			ctx,
			&api.MerkleProofRequest{
				FileIndex: int64(fileIdx),
			},
		)
		require.NoError(t, err)

		// Starting verifying the files
		verifyResp, err := grpcClient.VerifyMerkleProof(
			ctx,
			&api.VerifyProofRequest{
				RootHash:  uploadResp.MerkleRootHash,
				FileHash:  []byte(mt.CalcHash(files[fileIdx])),
				FileIndex: int64(fileIdx),
				Proofs:    recvResp.Proofs,
			},
		)

		require.NoError(t, err)
		require.True(t, verifyResp.IsVerified)
	}
}
