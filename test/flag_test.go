package test

import (
	"fmt"
	"testing"

	"github.com/lnk2005/td_read/global"
	"github.com/lnk2005/td_read/model"
)

func TestAtomicInt32(t *testing.T) {
	t.Log(global.READER_EXIT_FLAG)

	global.READER_EXIT_FLAG.Add(1)
	t.Log(global.READER_EXIT_FLAG)
}

func TestCreatChan(t *testing.T) {
	wc := make([]*chan *model.UserInfo, 6)
	t.Logf("%+v", wc)
}

func TestMap(t *testing.T) {
	tokenSet := make(map[string]bool, 10000)
	t.Log(fmt.Printf("%+v", tokenSet))
}
