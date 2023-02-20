package db

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lnk2005/td_read/global"
	"github.com/lnk2005/td_read/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateTables() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", model.GlobalConfig.Postgres.Host, model.GlobalConfig.Postgres.Port, model.GlobalConfig.Postgres.User, model.GlobalConfig.Postgres.Pass)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	defer func() {
		db, _ := DB.DB()
		if db != nil {
			db.Close()
		}
	}()

	// 实际建库 + 1，用于处理异常情况，虽然可能性几乎没有
	for i := 0; i < len(global.DB_TOKEN); i++ {
		createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", strings.Join([]string{global.DB_BASE_NAME, strconv.FormatInt(int64(i), 10)}, "_"))
		DB.Exec(createDatabaseCommand)
	}

	return nil
}

func CheckTables() error {
	return nil
}

func GetDbIndex(meta string) int {
	return strings.Index(global.DB_TOKEN, meta[0:1])
}
