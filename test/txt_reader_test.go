package test

import (
	"bufio"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/lnk2005/td_read/model"
	txtreader "github.com/lnk2005/td_read/txt_reader"
)

func TestTxtReader(t *testing.T) {
	basePath := "/data/"
	files, err := os.ReadDir(basePath)
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		t.Log(file.Name())
		f, err := os.Open(path.Join(basePath, file.Name()))
		if err != nil {
			t.Fatal(err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			t.Log(scanner.Text())

			var info model.UserInfo
			meta := scanner.Text()

			emailIndex := strings.Index(meta, txtreader.EmailToken)
			nameIndex := strings.Index(meta, txtreader.NameToken)
			screenNameIndex := strings.Index(meta, txtreader.ScreenNameToken)
			followersIndex := strings.Index(meta, txtreader.FollowersToken)
			createdIndex := strings.Index(meta, txtreader.CreatedToken)

			info.Email = meta[emailIndex+len(txtreader.EmailToken) : nameIndex]
			info.Name = meta[nameIndex+len(txtreader.NameToken) : screenNameIndex]
			info.ScreenName = meta[screenNameIndex+len(txtreader.ScreenNameToken) : followersIndex]
			info.CreatedAt = txtreader.ParserubyTimeToTimeStamp(meta[createdIndex+len(txtreader.CreatedToken):])

			// info.Token = info.GetToken()
			t.Logf("%+v", info)

			break
		}

		if err := scanner.Err(); err != nil {
			t.Fatal(err)
		}

		f.Close()
		break
	}
}
