package server

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	api "github.com/srinathln7/merkle_gaurd/api/v1/proto"
	"google.golang.org/grpc"

	mt "github.com/srinathln7/merkle_gaurd/internal/merkle"
)

type grpcServer struct {
	api.UnimplementedMerkleTreeServer

	mt.MerkleTree
}

func RunServer() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
		return
	}

	grpcServerAddr := os.Getenv("SERVER_ADDRESS")
	listener, err := net.Listen("tcp", grpcServerAddr)
	if err != nil {
		log.Fatalf("failed to dial grpc server: %v", err)
		return
	}

	// Create a new gRPC server and register the service
	grpcServer, err := newgrpcServer()
	if err != nil {
		log.Fatalf("failed to create gRPC server: %v", err)
	}

	// Listen on the specified grpc server port

	log.Printf("grpc server listening on: %s\n", listener.Addr().String())

	// Start the gRPC server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}

// NewGRPCServer: creates a grpc server and registers the service to that server
func newgrpcServer() (*grpc.Server, error) {
	gsrv := grpc.NewServer()
	srv := &grpcServer{}
	api.RegisterMerkleTreeServer(gsrv, srv)
	return gsrv, nil
}

func (s *grpcServer) Upload(ctx context.Context, req *api.UploadRequest) (
	*api.UploadResponse, error) {

	return &api.UploadResponse{}, nil
}

func (s *grpcServer) Download(ctx context.Context, req *api.DownloadRequest) (
	*api.DownloadResponse, error) {

	return &api.DownloadResponse{}, nil
}

func (s *grpcServer) GetMerkleProof(ctx context.Context, req *api.MerkleProofRequest) (
	*api.MerkleProofResponse, error) {

	return &api.MerkleProofResponse{}, nil
}

func (s *grpcServer) VerifyMerkleProof(ctx context.Context, req *api.VerifyProofRequest) (
	*api.VerifyProofResponse, error) {

	return &api.VerifyProofResponse{}, nil
}
