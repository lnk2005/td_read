package infowriter

import (
	"github.com/lnk2005/td_read/global"
	"github.com/lnk2005/td_read/model"
	"gorm.io/gorm"
)

type InfoWriter struct {
	Source *chan *model.UserInfo
	DB     *gorm.DB
}

func (w *InfoWriter) write(info *model.UserInfo) {
	w.DB.AutoMigrate(info)
	w.DB.Create(info)
}

func (w *InfoWriter) Run() {
	for {
		for info := range *w.Source {
			w.write(info)
		}

		if global.READER_EXIT_FLAG.Load() == global.READER_NUM {
			break
		}
	}
}

func NewInfoWriter(source *chan *model.UserInfo, db *gorm.DB) *InfoWriter {
	return &InfoWriter{
		Source: source,
		DB:     db,
	}
}
