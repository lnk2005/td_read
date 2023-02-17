package txtreader

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/lnk2005/td_read/db"
	"github.com/lnk2005/td_read/global"
	"github.com/lnk2005/td_read/model"
)

type Reader struct {
	Source *chan string
	Send   []*chan *model.UserInfo
}

func (r *Reader) Run() {
	// 延迟加载
	time.Sleep(time.Second * 1)

	for filename := range *r.Source {
		r.Read(filename)
	}

	global.READER_EXIT_FLAG.Add(1)
}

func (r *Reader) Read(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// t.Log(scanner.Text())
		tokens := strings.Split(scanner.Text(), " - ")
		if len(tokens) != 5 {
			log.Printf("error parsing tokens from %s\r\n", scanner.Text())
			return
		}

		var info model.UserInfo
		for _, token := range tokens {
			ts := strings.Split(token, ": ")
			if len(ts) != 2 {
				log.Printf("error parsing ts from %s\r\n", token)
				return
			}

			switch ts[0] {
			case "Email":
				info.Email = strings.ToLower(ts[1])
			case "Name":
				info.Name = ts[1]
			case "ScreenName":
				info.ScreenName = ts[1]
			case "Created At":
				info.CreatedAt = ParserubyTimeToTimeStamp(ts[1])
			}
		}
		info.Token = info.GetToken()
		r.send(&info)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func (r *Reader) send(info *model.UserInfo) {
	index := db.GetDbIndex(info.Token)
	(*r.Send[index]) <- info
}

func ParserubyTimeToTimeStamp(rbt string) int64 {
	t, err := time.Parse(time.RubyDate, rbt)
	if err != nil {
		return int64(0)
	}

	return t.Unix()
}

func NewReader(source *chan string, send []*chan *model.UserInfo) *Reader {
	return &Reader{
		Source: source,
		Send:   send,
	}
}
