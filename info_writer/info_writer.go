package infowriter

import (
	"time"

	"github.com/lnk2005/td_read/global"
	"github.com/lnk2005/td_read/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InfoWriter struct {
	Source *chan *model.UserInfo
	DB     *gorm.DB
}

func (w *InfoWriter) write(info *[]*model.UserInfo) {
	w.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(info)
}

func (w *InfoWriter) Run() {
	w.DB.AutoMigrate(&model.UserInfo{})

	infoSet := make([]*model.UserInfo, 1000)
	index := 0

	time.Sleep(time.Second * 10)
	for {

		for info := range *w.Source {
			if index == 1000 {
				w.write(&infoSet)
				infoSet = make([]*model.UserInfo, 1000)
				index = 0
			}

			infoSet[index] = info
			index += 1
		}

		if global.READER_EXIT_FLAG.Load() == global.READER_NUM {
			break
		}
	}

	if index != 0 {
		w.write(&infoSet)
	}
}

func NewInfoWriter(source *chan *model.UserInfo, db *gorm.DB) *InfoWriter {
	return &InfoWriter{
		Source: source,
		DB:     db,
	}
}
