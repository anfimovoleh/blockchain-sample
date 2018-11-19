package core

import "errors"

var (
	ErrNotValidIndex    = errors.New("block index is not valid")
	ErrNotValidHash     = errors.New("block hash is not valid")
	ErrNotValidPrevHash = errors.New("block prev hash is not valid")
	ErrEmptyBlock       = errors.New("block is empty")
)
