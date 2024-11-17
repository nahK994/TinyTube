package db

import "sync"

var (
	files      = map[int]string{}
	fileIDSeq  = 1
	fileIDLock sync.Mutex
)
