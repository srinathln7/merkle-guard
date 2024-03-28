package client

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	api "github.com/srinathln7/merkle_gaurd/api/v1/proto"
	"github.com/srinathln7/merkle_gaurd/lib/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SetupGRPCClient() (*api.MerkleTreeClient, error) {

	// Set up the gRPC client
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
		return nil, err
	}

	grpcServerAddr := os.Getenv("SERVER_ADDRESS")
	log.Printf("grpc client dialing on server address %s", grpcServerAddr)

	grpcClientOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(grpcServerAddr, grpcClientOptions...)
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
		return nil, err
	}

	// Create the gRPC client
	grpcClient := api.NewMerkleTreeClient(conn)
	return &grpcClient, nil
}

type UploadResponse struct {
	Msg      string `json:"msg"`
	RootHash string `json:"merkle_root_hash"`
}

func Upload(grpcClient api.MerkleTreeClient, files [][]byte) (*UploadResponse, error) {
	ctx := context.Background()
	resp, err := grpcClient.Upload(
		ctx,
		&api.UploadRequest{
			Files: files,
		},
	)

	if err != nil {
		util.ErrLog(err.Error())
		return nil, err
	}

	return &UploadResponse{
		Msg:      "All files uploaded successfully",
		RootHash: string(resp.MerkleRootHash),
	}, nil
}

type DownloadResponse struct {
	Msg  string `json:"msg"`
	File []byte `json:"file_content"`
}

func Download(grpcClient api.MerkleTreeClient, fileIdx int) (*DownloadResponse, error) {
	ctx := context.Background()
	resp, err := grpcClient.Download(
		ctx,
		&api.DownloadRequest{
			FileIndex: int64(fileIdx),
		},
	)

	if err != nil {
		util.ErrLog(err.Error())
		return nil, err
	}

	msg := fmt.Sprintf("File%d downloaded successfully \n", fileIdx)
	return &DownloadResponse{
		Msg:  msg,
		File: resp.FileContent,
	}, nil
}
