package db

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lnk2005/td_read/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDb(index int) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		model.GlobalConfig.Postgres.Host,
		model.GlobalConfig.Postgres.Port,
		model.GlobalConfig.Postgres.User,
		model.GlobalConfig.Postgres.Pass,
		strings.Join([]string{"info", strconv.FormatInt(int64(index), 10)}, "_"))

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return DB
}
