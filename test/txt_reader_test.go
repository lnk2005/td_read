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
			tokens := strings.Split(scanner.Text(), " - ")
			if len(tokens) != 5 {
				t.Fatalf("error parsing tokens from %s", scanner.Text())
			}

			var info model.UserInfo
			for _, token := range tokens {
				ts := strings.Split(token, ": ")
				if len(ts) != 2 {
					t.Fatalf("error parsing ts from %s", token)
					break
				}

				switch ts[0] {
				case "Email":
					info.Email = strings.ToLower(ts[1])
				case "Name":
					info.Name = ts[1]
				case "ScreenName":
					info.ScreenName = ts[1]
				case "Created At":
					info.CreatedAt = txtreader.ParserubyTimeToTimeStamp(ts[1])
				}
			}
			info.Token = info.GetToken()
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
