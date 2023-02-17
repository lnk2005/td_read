package global

import (
	"sync/atomic"
)

const (
	READER_NUM = 200
	WRITER_NUM = 6
)

var (
	READER_EXIT_FLAG = atomic.Int32{}
)
