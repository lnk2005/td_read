package main

import (
	"fmt"

	"github.com/lnk2005/td_read/cmd"
	"github.com/lnk2005/td_read/db"
	"github.com/lnk2005/td_read/model"
)

func main() {
	cmd.Execute()

	err := db.CreateTables()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\r\n", model.GlobalConfig)
}
