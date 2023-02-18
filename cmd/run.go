package cmd

import (
	"os"
	"path"
	"sync"

	"github.com/lnk2005/td_read/db"
	"github.com/lnk2005/td_read/global"
	infowriter "github.com/lnk2005/td_read/info_writer"
	"github.com/lnk2005/td_read/model"
	txtreader "github.com/lnk2005/td_read/txt_reader"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the reader to read data into database",
	Long:  `run the reader to read data into database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		basePath := "/data/"
		files, err := os.ReadDir(basePath)
		if err != nil {
			panic(err)
		}

		fc := make(chan string, 100)

		wc := make([]*chan *model.UserInfo, global.WRITER_NUM)
		for i := 0; i < global.WRITER_NUM; i++ {
			v := make(chan *model.UserInfo, 100)
			wc[i] = &v
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			for _, file := range files {
				fc <- path.Join(basePath, file.Name())
			}
			wg.Done()
		}()

		for i := 0; i < global.READER_NUM; i++ {
			wg.Add(1)
			go func() {
				r := txtreader.NewReader(&fc, wc)
				r.Run()
				wg.Done()
			}()
		}

		for i := 0; i < global.WRITER_NUM; i++ {
			wg.Add(5)
			for j := 0; j < 5; j++ {
				go func(index int) {
					db := db.GetDb(index)
					w := infowriter.NewInfoWriter(wc[index], db)
					w.Run()
					wg.Done()
				}(i)
			}
		}

		wg.Wait()
		close(fc)
		for i := 0; i < global.WRITER_NUM; i++ {
			close(*wc[i])
		}
		return nil
	},
}
