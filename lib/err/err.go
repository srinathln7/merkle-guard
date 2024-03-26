package err

import "errors"

var (
	ErrEmptyFile       = errors.New("empty file passed")
	ErrIndexOutOfBound = errors.New("specified leaf index is out-of-bound")
	ErrEmptyRoot       = errors.New("empty tree found")
	ErrEmptyNode       = errors.New("empty node passed")
)
