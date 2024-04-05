# Merkle-Gaurd


## Usage

### Start server

```
go run main.go --server

```

### Client

```
go run main.go upload -d <files_dir> -O <merkle_root_hash_path>

go run main.go download -i <file_idx> -o <download_path_file_dir>

go run main.go getMerkleProofs -i <file_idx> -o <merkle_proof_path_dir>

go run main.go verifyMerkleProofs -r <merkle_root_hash_path> -f <file_dir> -i <file_idx> -p <merkle_proof_path_dir> 
```

### Examples

```
go run main.go upload -d "./sample/upload" -O "./sample"

go run main.go download -i 0 -o "./sample/download"

go run main.go getMerkleProofs -i 0 -o "./sample/merkle-proofs"

go run main.go verifyMerkleProofs -r "./sample" -f "./sample/download" -i 0 -p "./sample/merkle-proofs"
```