package err

import "errors"

var (
	ErrEmptyFile              = errors.New("empty file passed")
	ErrIndexOutOfBound        = errors.New("specified file(leaf) index is out-of-bound")
	ErrEmptyRoot              = errors.New("empty tree found")
	ErrEmptyNode              = errors.New("empty node passed")
	ErrMerkleRootHashMisMatch = errors.New("merkle root hash mis-match")
	ErrFileHashMisMatch       = errors.New("file hash mis-match")
	ErrMerkleVerificationFail = errors.New("merkle tree verification failed")
	ErrLeafDoesNotExist       = errors.New("leaf node (file) does not exist")
	ErrConversion             = errors.New("type conversion not successful")
)
