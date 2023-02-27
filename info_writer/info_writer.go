package infowriter

import (
	"time"

	"github.com/lnk2005/td_read/global"
	"github.com/lnk2005/td_read/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InfoWriter struct {
	Source *chan *model.UserInfoTmp
	DB     *gorm.DB
}

func (w *InfoWriter) write(info *[]*model.UserInfo) {
	w.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(info)
}

func (w *InfoWriter) Run() {
	tokenSet := make(map[string]bool)

	w.DB.AutoMigrate(&model.UserInfo{})

	infoSet := make([]*model.UserInfo, 1000)
	index := 0

	time.Sleep(time.Second * 10)
	for {

		for infoTmp := range *w.Source {
			if _, ok := tokenSet[infoTmp.Token]; ok {
				continue
			} else {
				tokenSet[infoTmp.Token] = true
			}

			if index == 1000 {
				w.write(&infoSet)
				infoSet = make([]*model.UserInfo, 1000)
				index = 0
			}

			infoSet[index] = &infoTmp.Info
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

func NewInfoWriter(source *chan *model.UserInfoTmp, db *gorm.DB) *InfoWriter {
	return &InfoWriter{
		Source: source,
		DB:     db,
	}
}
