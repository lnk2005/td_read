package test

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/lnk2005/td_read/model"
	homedir "github.com/mitchellh/go-homedir"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestDb(t *testing.T) {
	config := model.DbConfig{
		Host: "127.0.0.1",
		Port: "5432",
		User: "postgres",
		Pass: "mysecretpassword",
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", config.Host, config.Port, config.User, config.Pass)
	DB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	for i := 0; i < 10; i++ {
		createDatabaseCommand := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", strings.Join([]string{"info", strconv.FormatInt(int64(i), 10)}, "_"))
		DB.Exec(createDatabaseCommand)
	}
}

func TestDropDatabase(t *testing.T) {
	config := model.DbConfig{
		Host: "127.0.0.1",
		Port: "5432",
		User: "postgres",
		Pass: "mysecretpassword",
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", config.Host, config.Port, config.User, config.Pass)
	DB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	for i := 0; i < 10; i++ {
		createDatabaseCommand := fmt.Sprintf("DROP DATABASE IF EXISTS %s", strings.Join([]string{"info", strconv.FormatInt(int64(i), 10)}, "_"))
		DB.Exec(createDatabaseCommand)
	}
}

func TestHomeDir(t *testing.T) {
	t.Log(homedir.Dir())
	t.Log(os.Getwd())
}
