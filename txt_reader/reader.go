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

var (
	EmailToken      = "Email: "
	NameToken       = " - Name: "
	ScreenNameToken = " - ScreenName: "
	FollowersToken  = " - Followers: "
	CreatedToken    = " - Created At: "
)

type Reader struct {
	Source *chan string
	Send   []*chan *model.UserInfoTmp
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
	defer func() {
		f.Close()
		// os.Remove(filename)
	}()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var info model.UserInfo
		meta := scanner.Text()

		emailIndex := strings.Index(meta, EmailToken)
		nameIndex := strings.Index(meta, NameToken)
		screenNameIndex := strings.Index(meta, ScreenNameToken)
		followersIndex := strings.Index(meta, FollowersToken)
		createdIndex := strings.Index(meta, CreatedToken)

		if emailIndex == -1 || nameIndex == -1 || screenNameIndex == -1 || followersIndex == -1 || createdIndex == -1 {
			log.Println(meta, emailIndex, nameIndex, screenNameIndex, followersIndex, createdIndex)
			continue
		}

		info.Email = strings.ToLower(meta[emailIndex+len(EmailToken) : nameIndex])
		info.Name = meta[nameIndex+len(NameToken) : screenNameIndex]
		info.ScreenName = meta[screenNameIndex+len(ScreenNameToken) : followersIndex]
		info.CreatedAt = ParserubyTimeToTimeStamp(meta[createdIndex+len(CreatedToken):])

		infoTmp := model.UserInfoTmp{
			Info:  info,
			Token: info.GetToken(),
		}

		r.send(&infoTmp)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func (r *Reader) send(info *model.UserInfoTmp) {
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

func NewReader(source *chan string, send []*chan *model.UserInfoTmp) *Reader {
	return &Reader{
		Source: source,
		Send:   send,
	}
}
