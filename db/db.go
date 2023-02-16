package db

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lnk2005/td_read/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db_token = "abcdefghijklmnopqrstuvwxyz0123456789"
	index    = 6
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
	for i := 0; i < index+1; i++ {
		createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", strings.Join([]string{"info", strconv.FormatInt(int64(i), 10)}, "_"))
		DB.Exec(createDatabaseCommand)
	}

	return nil
}

func CheckTables() error {
	return nil
}

func getDbIndex(meta string) int {
	if strings.Contains(db_token, meta) {
		return strings.Index(db_token, meta) % 6
	}

	return 6
}
