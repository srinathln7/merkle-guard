# Merkle-Gaurd



## Requirements

* golang (v1.22)
* protoc compiler (v23.3 or higher)
* docker (v20.10.21 or higher)
* docker-compose (v20.20.2 or higher)
* VSCode or any other suitable IDE

## Project Structure

Refer [here](https://github.com/srinathLN7/merkle-gaurd/blob/main/OVERVIEW.md) for the complete overview of the protocol and the project structure.


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

### Examples

```
./mg upload -d "./sample/upload" -O "./sample"

./mg download -i 0 -o "./sample/download"

./mg getMerkleProofs -i 0 -o "./sample/merkle-proofs"

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

### Unit Tests

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



