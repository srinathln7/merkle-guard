package server

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	api "github.com/srinathln7/merkle_gaurd/api/v1/proto"
	mt "github.com/srinathln7/merkle_gaurd/internal/merkle"
	mterr "github.com/srinathln7/merkle_gaurd/lib/err"
	"github.com/srinathln7/merkle_gaurd/lib/util"
)

type grpcServer struct {
	api.UnimplementedMerkleTreeServer

	files      [][]byte
	merkleTree *mt.MerkleTree
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
	grpcServer, err := NewgrpcServer()
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

// newgrpcServer: creates a grpc server and registers the service to that server
func NewgrpcServer() (*grpc.Server, error) {
	gsrv := grpc.NewServer()
	srv := &grpcServer{}
	api.RegisterMerkleTreeServer(gsrv, srv)
	return gsrv, nil
}

func (s *grpcServer) Upload(ctx context.Context, req *api.UploadRequest) (
	*api.UploadResponse, error) {
	merkleTree, err := mt.BuildMerkleTree(req.Files)
	if err != nil {
		return nil, err
	}
	s.files = req.Files
	s.merkleTree = merkleTree
	merkleRoot := merkleTree.GetMerkleRoot()

	util.ServerLog("Resulting merkle tree after the client uploaded all the files")
	merkleTree.PrintTreeInfo()
	return &api.UploadResponse{MerkleRootHash: []byte(merkleRoot.Hash)}, nil
}

func (s *grpcServer) Download(ctx context.Context, req *api.DownloadRequest) (
	*api.DownloadResponse, error) {

	fileIdx := int(req.FileIndex)
	if fileIdx < 0 || fileIdx >= len(s.files) {
		return nil, mterr.ErrIndexOutOfBound
	}

	return &api.DownloadResponse{FileContent: s.files[fileIdx]}, nil
}

func (s *grpcServer) GetMerkleProof(ctx context.Context, req *api.MerkleProofRequest) (
	*api.MerkleProofResponse, error) {

	util.ServerLog("running GetMerkleProof ")
	fileIdx := int(req.FileIndex)
	if fileIdx < 0 || fileIdx >= len(s.files) {
		return nil, mterr.ErrIndexOutOfBound
	}

	merkleProofs, err := s.merkleTree.GenerateMerkleProof(fileIdx)
	if err != nil {
		util.ErrLog(err.Error())
		return nil, err
	}

	// Here, we use the unsafe package to perform a direct type conversion from []*merkle.TreeNode to []*api.TreeNode. This method avoids iterating
	//through each element of the slice, making it more efficient. However, it's important to note that the use of unsafe package should be handled with
	//caution as it bypasses Go's type safety mechanisms. Here we ensure that the types are truly compatible before using this approach.

	// util.ServerLog("starting unsafe type conversion")
	//var proofs []*api.TreeNode = *(*[]*api.TreeNode)(unsafe.Pointer(&merkleProofs))
	var proofs []*api.TreeNode
	for _, proof := range merkleProofs {
		apiProof := &api.TreeNode{
			Hash:     proof.Hash,
			LeftIdx:  int64(proof.LeftIdx),
			RightIdx: int64(proof.RightIdx),
		}
		if proof.Left != nil {
			apiProof.Left = &api.TreeNode{
				Hash:     proof.Left.Hash,
				LeftIdx:  int64(proof.Left.LeftIdx),
				RightIdx: int64(proof.Left.RightIdx),
			}
		}
		if proof.Right != nil {
			apiProof.Right = &api.TreeNode{
				Hash:     proof.Right.Hash,
				LeftIdx:  int64(proof.Right.LeftIdx),
				RightIdx: int64(proof.Right.RightIdx),
			}
		}
		proofs = append(proofs, apiProof)
	}

	return &api.MerkleProofResponse{Proofs: proofs}, nil
}

func (s *grpcServer) VerifyMerkleProof(ctx context.Context, req *api.VerifyProofRequest) (
	*api.VerifyProofResponse, error) {

	fileIdx := int(req.FileIndex)
	if fileIdx < 0 || fileIdx >= len(s.files) {
		return nil, mterr.ErrIndexOutOfBound
	}

	file := s.files[fileIdx]
	if string(req.FileHash) != mt.CalcHash(file) {
		return nil, mterr.ErrFileHashMisMatch
	}

	// var merkleProofs []*mt.TreeNode = *(*[]*mt.TreeNode)(unsafe.Pointer(&req.Proofs))
	var merkleProofs []*mt.TreeNode
	for _, proof := range req.Proofs {
		merkleProof := &mt.TreeNode{
			Hash:     proof.Hash,
			LeftIdx:  int(proof.LeftIdx),
			RightIdx: int(proof.RightIdx),
		}

		if proof.Left != nil {
			merkleProof.Left = &mt.TreeNode{
				Hash:     proof.Left.Hash,
				LeftIdx:  int(proof.Left.LeftIdx),
				RightIdx: int(proof.Left.RightIdx),
			}
		}

		if proof.Right != nil {
			merkleProof.Right = &mt.TreeNode{
				Hash:     proof.Right.Hash,
				LeftIdx:  int(proof.Right.LeftIdx),
				RightIdx: int(proof.Right.RightIdx),
			}
		}

		merkleProofs = append(merkleProofs, merkleProof)
	}

	isVerified, err := s.merkleTree.VerifyMerkleProof(string(req.RootHash), string(req.FileHash), fileIdx, merkleProofs)
	if err != nil {
		return nil, err
	}
	return &api.VerifyProofResponse{IsVerified: isVerified}, nil
}
