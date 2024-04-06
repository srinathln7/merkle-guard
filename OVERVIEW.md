# Project Overview

Imagine a client has a large set of potentially small files {F0, F1, …, Fn} and wants to upload them to a server and then delete its local copies. The client wants, however, to later download an arbitrary file from the server and be convinced that the file is correct and is not corrupted in any way (in transport, tampered with by the server, etc.).

You should implement the client, the server and a Merkle tree to support the above (we expect you to implement the Merkle tree rather than use a library, but you are free to use a library for the underlying hash functions).

The client must compute a single Merkle tree root hash and keep it on its disk after uploading the files to the server and deleting its local copies. The client can request the i-th file Fi and a Merkle proof Pi for it from the server. The client uses the proof and compares the resulting root hash with the one it persisted before deleting the files - if they match, file is correct.


## Project structure

```
.
├── api
│   └── v1
│       └── proto
│           ├── merkle_grpc.pb.go
│           ├── merkle.pb.go
│           ├── merkle.proto
│           └── README.md
├── cmd
│   ├── cmd.go
│   └── README.md
├── deploy
│   └── local
│       ├── docker-compose.yml
│       ├── Dockerfile.client
│       └── Dockerfile.server
├── docs
│   ├── client.html
│   ├── index.html
│   ├── mt.html
│   └── server.html
├── go.mod
├── go.sum
├── internal
│   ├── client
│   │   ├── client.go
│   │   └── README.md
│   ├── merkle
│   │   ├── merkle.go
│   │   ├── merkle_test.go
│   │   └── README.md
│   ├── server
│   │   ├── README.md
│   │   └── server.go
│   └── tests
│       ├── client_test.go
│       ├── README.md
│       └── server_test.go
├── lib
│   ├── err
│   │   └── err.go
│   └── util
│       ├── README.md
│       └── util.go
├── LICENSE
├── main.go
├── MAIN.md
├── Makefile
├── OVERVIEW.md
├── README.md
└── sample
    └── upload
        ├── file0.txt
        ├── file1.txt
        ├── file2.txt
        └── file3.txt

17 directories, 38 files
```





