package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	db2 "github.com/pjoc-team/pay-database-service/pkg/db"
	"github.com/pjoc-team/pay-database-service/pkg/model"
)

func main() {
	if db, err := db2.InitDb(); err != nil {
		return
	} else {
		table := db.CreateTable(&model.PayOrder{})
		db.CreateTable(&model.PayOrderOk{})
		db.CreateTable(&model.Notice{})
		db.CreateTable(&model.NoticeOk{})
		db.Close()
		fmt.Println("create result: ", table)
	}
}
