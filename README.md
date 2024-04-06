# Merkle-Gaurd

Merkle Guard aims at providing a secure and efficient way to manage and verify file integrity using Merkle trees. It leverages  segment trees for building Merkle trees, utilizes gRPC protocol for communication between clients and servers, and offers a CLI application for easy client access. The project is designed to be deployable using Docker Compose, ensuring ease of deployment and scalability.

Click [here](https://github.com/srinathln7/merkle-gaurd/blob/main/OVERVIEW.md) to watch the demo presentation. 

## Requirements

* golang (v1.22)
* protoc compiler (v23.3 or higher)
* docker (v20.10.21 or higher)
* docker-compose (v20.20.2 or higher)
* VSCode or any other suitable IDE

## Project Structure

Refer [here](https://github.com/srinathLN7/merkle-gaurd/blob/main/OVERVIEW.md) for the challenge description and the project structure.

## Features

- **Merkle Tree Construction**: The project employs [segment trees](https://en.wikipedia.org/wiki/Segment_tree) for constructing Merkle trees efficiently. Merkle trees provide a cryptographic hash-based data structure that allows for efficient verification of large datasets. Building merkle trees from scratch and generating merkle proofs for a given file index forms the crux of the project. For more information about this refer [here](https://github.com/srinathln7/merkle-gaurd/tree/main/internal/merkle).  

- **Protocol Buffers**: We define a Protocol Buffers (protobuf) file to specify the structure of messages exchanged between the client and server. Protocol Buffers offer a language-agnostic and efficient way to serialize structured data. Click [here](https://github.com/srinathln7/merkle-gaurd/tree/main/api/v1/proto) to learn more.

- **gRPC Communication**: gRPC is utilized as the underlying communication protocol between the client and server. This ensures efficient and secure communication between components. Click here to access more info about the grpc [server](https://github.com/srinathln7/merkle-gaurd/tree/main/internal/server) and [client](https://github.com/srinathln7/merkle-gaurd/tree/main/internal/client)

- **CLI Application**: Merkle Guard provides a CLI application for clients to interact with the server. This CLI application offers commands for uploading files, downloading files, generating Merkle proofs, and verifying file integrity. See [here](https://github.com/srinathln7/merkle-gaurd/tree/main/cmd) for more infomation.

- **Docker Compose Deployment**: The project includes Docker Compose configuration for easy deployment and scaling. Docker Compose allows for the deployment of the entire application stack with a single command, simplifying the deployment process.


## Usage

1. Clone the repository:

```
git clone https://github.com/username/merkle-gaurd.git
```

2. Change into the project directory:

```
cd merkle-gaurd
```

### Build Binaries

```
 
 go build -o mg .

```

### Start grpc server

```
./mg --server

```

### Start grpc client

```
./mg upload -d <files_dir> -O <merkle_root_hash_path>

./mg download -i <file_idx> -o <download_path_file_dir>

./mg getMerkleProofs -i <file_idx> -o <merkle_proof_path_dir>

./mg verifyMerkleProofs -r <merkle_root_hash_path> -f <file_dir> -i <file_idx> -p <merkle_proof_path_dir> 
```

### Example Usage

```

#  Upload from the client
./mg upload -d "./sample/upload" -O "./sample"

# Delete the uploaded files from the client's disk
rm -rf ./sample/upload

# Download from the server
./mg download -i 0 -o "./sample/download"

# Extract the merkle proof for file0 from the server
./mg getMerkleProofs -i 0 -o "./sample/merkle-proofs"

# Verify the merkle proof for file0 from the server
./mg verifyMerkleProofs -r "./sample" -f "./sample/download" -i 0 -p "./sample/merkle-proofs"
```

## Run with Docker

To run using Docker, ensure that Docker is installed on your machine and follow these steps:

1. Build the Docker images and containers:

```

cd deploy/local

docker compose up

```

If you encounter issues with building the containers due to IP address overlap, it is likely caused by conflicting IP addresses in the network. To resolve this, you can change the subnet address used for the containers to ensure uniqueness. By selecting a different subnet address, you can avoid conflicts and successfully build the containers.

2. Enter the Docker `local-merkle-gaurd-client` container:

```
docker exec -it local-merkle-gaurd-client sh
```

Repeat steps under the **Examples** section to test for various test sceanarios. 

3. Stop and remove the Docker containers:

```
docker compose down
```

## Testing


To run all the test files in this project, run the following command in your local development terminal:

```
make test
```

Upon running this command, you should see all the test cases passing, ensuring the proper functioning of all components within our project. Successful test results indicate that the application is operating as expected and meeting the desired requirements. 


## API Documentation

For the API documentation, refer to the [docs](https://github.com/srinathln7/merkle-gaurd/tree/main/docs) directory containing individual API documentation about the gRPC server and client APIs.

Alternatively, if you wish to build your own docs, run:

```
godoc 
```

and navigate to http://localhost:6060/pkg/github.com/srinathln7/merkle_gaurd/internal/?m=all in your browser. You will find the links to all three packages: server, client, and merkle.

## Improvements

To enhance the merkle-gaurd protocol, the following improvements can be implemented:

* Implement Mutual TLS Authentication:
  - Introduce Mutual TLS-based authentication between the gRPC server and client to establish a secure and trusted communication channel. This ensures that both parties can verify each other's identities and encrypt the data exchanged during communication.

* Deployment Scripts for Cloud:
  - Develop deployment scripts to automate the process of deploying the gRPC server and client containers to the cloud platform such as AWS, Azure, GCP etc. This streamlines the deployment process and facilitates scalability and reliability.

* Database integration:
  -  Integrate a database into the gRPC server to persist file information, upload proofs, and other relevant data across server restarts. By leveraging a database, , the server can store and retrieve file metadata efficiently, ensuring data integrity and reliability. This enhancement allows the server to maintain stateful information, enabling seamless continuation of operations even after server restarts or failures. 

* Introduce custom gRPC error messages:
  - Utilize packages like `google.golang.org/genproto/googleapis/rpc/errdetails` and `google.golang.org/grpc/status` to create and propagate custom error messages throughout the gRPC communication layer. This enhancement enhances error handling and provides more informative feedback to clients, aiding in debugging and troubleshooting. Custom error messages can convey specific details about the encountered issues, improving the overall user experience and facilitating faster resolution of errors.
  
## License

This project is licensed under the [MIT License](LICENSE).
