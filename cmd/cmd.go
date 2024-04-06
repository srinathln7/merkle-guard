package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/srinathln7/merkle_gaurd/internal/client"
	"github.com/srinathln7/merkle_gaurd/lib/util"
)

var (
	fileIdx     int
	filesDir    string
	fileDir     string
	rootHashDir string
	proofsDir   string
)

func SetupFlags() {
	RootCmd.PersistentFlags().IntVarP(&fileIdx, "fileIdx", "i", 0, "Index of the file")
	RootCmd.PersistentFlags().StringVarP(&filesDir, "uploadDir", "d", "", "Upload files directory")
	RootCmd.PersistentFlags().StringVarP(&rootHashDir, "merkleRootHash", "r", "", "Directory where the merkle root hash file is located")
	RootCmd.PersistentFlags().StringVarP(&rootHashDir, "merkleRootHashOP", "O", "", "Directory where the merkle root hash file is located")
	RootCmd.PersistentFlags().StringVarP(&fileDir, "downloadDir", "o", "", "File directory")
	RootCmd.PersistentFlags().StringVarP(&fileDir, "file", "f", "", "File directory")
	RootCmd.PersistentFlags().StringVarP(&proofsDir, "merkleProofs", "p", "", "Directory where the merkle proofs for the file is located")
	RootCmd.AddCommand(uploadCmd)
	RootCmd.AddCommand(downloadCmd)
	RootCmd.AddCommand(getMerkleProofsCmd)
	RootCmd.AddCommand(verifyMerkleProofsCmd)
}

var RootCmd = &cobra.Command{
	Use:   "merkle_gaurd",
	Short: "Merkle Gaurd CLI",
	Run: func(cmd *cobra.Command, args []string) {

		color.Yellow("************************ Welcome to Merkle-Gaurd CLI *****************")
		color.Yellow("Please use any of the following sub-commands 'upload', 'download', 'getMerkleProofs' or 'verifyMerkleProofs'")
		color.Yellow("To upload a set of files from the directory: go run main.go upload -d <files_dir> -O <merkle_root_hash_path>`")
		color.Yellow("To download a file for the given file index from the server to a specified path: `go run main.go download -i <file_idx> -o <download_path_file_dir>`")
		color.Yellow("To get merkle proofs for the given file index from the server: `go run main.go getMerkleProofs -i <file_idx> -o <merkle_proof_path_dir>`")
		color.Yellow("To verify merkle proofs for the given file `go run main.go verifyMerkleProofs -r <merkle_root_hash_path> -f <file_dir> -i <file_idx> -p <merkle_proof_path_dir>`")
		color.Yellow("To exit this terminal press CTRL+C")

		// Setup a signal handler to capture interrupt and termination signals
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGTERM)

		// when done exit the program gracefully
		<-done

	},
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a set of files",
	Run: func(cmd *cobra.Command, args []string) {
		grpcClient, err := client.SetupGRPCClient()
		if err != nil {
			log.Fatalf("error setting up grpc client %s", err.Error())
		}

		files, err := util.ReadFilesFromDir(filesDir)
		if err != nil {
			log.Fatal("error reading files from the directory:", err)
		}

		uploadResp, err := client.Upload(*grpcClient, files)
		if err != nil {
			log.Fatal("error during the client upload process:", err)
		}

		rootHashFile := filepath.Join(rootHashDir, os.Getenv("MERKLE_ROOT_FILE"))
		err = os.WriteFile(rootHashFile, []byte(uploadResp.RootHash), 0644)
		if err != nil {
			log.Fatal("error writing merkle root hash to the file:", err)
		}

		resJSON, err := json.Marshal(uploadResp)
		if err != nil {
			log.Fatal("error:", err)
		}

		color.Green(string(resJSON))
	},
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download the file corresponding to the specified file index to the specified file path",
	Run: func(cmd *cobra.Command, args []string) {
		grpcClient, err := client.SetupGRPCClient()
		if err != nil {
			log.Fatalf("error setting up grpc client %s", err.Error())
		}
		downloadRes, err := client.Download(*grpcClient, fileIdx)
		if err != nil {
			return
		}

		err = util.WriteFile(fileDir, os.Getenv("FILE_PREFIX")+strconv.Itoa(fileIdx)+os.Getenv("FILE_FORMAT"), string(downloadRes.File))
		if err != nil {
			log.Fatalf("error downloading file to the specified path: %v", err)
		}

		color.Green(fmt.Sprintf("File downloaded to the specified path: %s", fileDir))
	},
}

var getMerkleProofsCmd = &cobra.Command{
	Use:   "getMerkleProofs",
	Short: "Outputs the merkle proofs for the file corresponding to the specified file index",
	Run: func(cmd *cobra.Command, args []string) {
		grpcClient, err := client.SetupGRPCClient()
		if err != nil {
			log.Fatalf("error setting up grpc client %s", err.Error())
		}

		proofResp, err := client.GetMerkleProof(*grpcClient, fileIdx)
		if err != nil {
			return
		}

		resJSON, err := json.Marshal(proofResp)
		if err != nil {
			log.Fatal("error:", err)
		}

		err = util.WriteFile(fileDir, os.Getenv("FILE_PREFIX")+strconv.Itoa(fileIdx)+os.Getenv("FILE_FORMAT"), string(resJSON))
		if err != nil {
			log.Fatalf("error writing merkle proofs to the specified path: %v", err)
		}

		color.Green(string(resJSON))
	},
}

var verifyMerkleProofsCmd = &cobra.Command{
	Use:   "verifyMerkleProofs",
	Short: "Outputs the merkle proofs for the file corresponding to the specified file index",
	Run: func(cmd *cobra.Command, args []string) {
		grpcClient, err := client.SetupGRPCClient()
		if err != nil {
			log.Fatalf("error setting up grpc client %s", err.Error())
		}

		rootHashFile := filepath.Join(rootHashDir, os.Getenv("MERKLE_ROOT_FILE"))
		rootHash, err := os.ReadFile(rootHashFile)
		if err != nil {
			log.Fatalf("error reading merkle root hash from the client's disk path %s", rootHashFile)
		}

		filePath := filepath.Join(fileDir, os.Getenv("FILE_PREFIX")+strconv.Itoa(fileIdx)+os.Getenv("FILE_FORMAT"))
		file, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("error reading the file from the file path %s", filePath)
		}

		proofsFile := filepath.Join(proofsDir, os.Getenv("FILE_PREFIX")+strconv.Itoa(fileIdx)+os.Getenv("FILE_FORMAT"))
		proofsRespBytes, err := os.ReadFile(proofsFile)
		if err != nil {
			log.Fatalf("error reading the merkle proofs for the file from the specified path %s", proofsFile)
		}

		var proofResp client.ProofResponse
		err = json.Unmarshal(proofsRespBytes, &proofResp)
		if err != nil {
			log.Fatal(err.Error())
		}

		verifyResp, err := client.VerifyMerkleProof(*grpcClient, client.VerifyRequest{
			RootHash: rootHash,
			FileIdx:  fileIdx,
			File:     file,
			Proofs:   proofResp.Proofs,
		})
		if err != nil {
			return
		}

		resJSON, err := json.Marshal(verifyResp)
		if err != nil {
			log.Fatal("error:", err)
		}
		color.Green(string(resJSON))
	},
}
