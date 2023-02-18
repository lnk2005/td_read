package global

import (
	"sync/atomic"
)

const (
	READER_NUM = 2000
	WRITER_NUM = 6
)

var (
	READER_EXIT_FLAG = atomic.Int32{}
)
