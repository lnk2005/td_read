package global

import (
	"sync/atomic"
)

const (
	READER_NUM   = 100
	INDEX        = 9
	WRITER_NUM   = 9
	DB_BASE_NAME = "info_new" // 结尾不加 "_"
)

var (
	READER_EXIT_FLAG = atomic.Int32{}
)
