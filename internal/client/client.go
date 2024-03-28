package client

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/srinathln7/merkle_gaurd/lib/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	api "github.com/srinathln7/merkle_gaurd/api/v1/proto"
	mt "github.com/srinathln7/merkle_gaurd/internal/merkle"
	mterr "github.com/srinathln7/merkle_gaurd/lib/err"
)

var MERKLE_ROOT_FILE string

func SetupGRPCClient() (*api.MerkleTreeClient, error) {

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

	// Store the resulting merkle root hash on client side disk - say a simple `.json` file
	util.ClientLog("storing the merkle tree root hash on client's disk")
	// err = os.WriteFile(os.Getenv("MERKLE_ROOT_FILE"), resp.MerkleRootHash, 0644)
	// if err != nil {
	// 	return nil, err
	// }

	MERKLE_ROOT_FILE = string(resp.MerkleRootHash)
	return &UploadResponse{
		Msg:      "all files uploaded successfully",
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

	msg := fmt.Sprintf("file%d downloaded successfully \n", fileIdx)
	return &DownloadResponse{
		Msg:  msg,
		File: resp.FileContent,
	}, nil
}

type ProofResponse struct {
	Msg    string          `json:"msg"`
	Proofs []*api.TreeNode `json:"proofs"`
}

func GetMerkleProof(grpcClient api.MerkleTreeClient, fileIdx int) (*ProofResponse, error) {
	ctx := context.Background()
	resp, err := grpcClient.GetMerkleProof(
		ctx,
		&api.MerkleProofRequest{
			FileIndex: int64(fileIdx),
		},
	)

	if err != nil {
		util.ErrLog(err.Error())
		return nil, err
	}

	msg := fmt.Sprintf("merkle proofs for file%d generated successfully \n", fileIdx)
	return &ProofResponse{
		Msg:    msg,
		Proofs: resp.Proofs,
	}, nil
}

type VerifyRequest struct {
	RootHash []byte          `json:"root_hash"`
	FileIdx  int             `json:"file_idx"`
	File     []byte          `json:"file"`
	Proofs   []*api.TreeNode `json:"proofs"`
}

type VerifyResponse struct {
	Msg       string `json:"msg"`
	IsVerfied bool   `json:"is_Verified"`
}

func VerifyMerkleProof(grpcClient api.MerkleTreeClient, req VerifyRequest) (*VerifyResponse, error) {
	ctx := context.Background()
	resp, err := grpcClient.VerifyMerkleProof(
		ctx,
		&api.VerifyProofRequest{
			RootHash:  req.RootHash,
			FileIndex: int64(req.FileIdx),
			FileHash:  []byte(mt.CalcHash(req.File)),
			Proofs:    req.Proofs,
		},
	)

	if err != nil {
		util.ErrLog(err.Error())
		return nil, err
	}

	if !resp.IsVerified {
		return nil, mterr.ErrMerkleVerificationFail
	}

	msg := fmt.Sprintf("merkle verification for file%d is successful \n", req.FileIdx)
	return &VerifyResponse{
		Msg:       msg,
		IsVerfied: true,
	}, nil
}
