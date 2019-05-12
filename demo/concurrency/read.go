package main

import (
	"fmt"
	"github.com/coreos/etcd/pkg/idutil"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pjoc-team/base-service/pkg/logger"
	"github.com/pjoc-team/pay-database-service/pkg/db"
	"github.com/pjoc-team/pay-database-service/pkg/model"
	"time"
)

func main() {
	var dbConn *gorm.DB
	var err error
	if dbConn, err = db.InitDb(); err != nil {
		return
	}
	bools := make(chan bool, 200)
	for i := 0; i < 200; i++ {
		go FindUser(dbConn, bools)
	}
	for {
		select {
		case <-bools:
			go FindUser(dbConn, bools)
		}
	}

}

var gen = idutil.NewGenerator(12, time.Now())

func FindUser(db *gorm.DB, bools chan bool) {
	defer func() {
		if err := recover(); err != nil {
			logger.Log.Error("db error...", err)
		}
	}()
	defer func() {
		bools <- true
	}()
	//b := &model.BasePayOrder{}
	//instance := &model.PayOrder{BasePayOrder: b}
	instance := model.PayOrder{}
	instance.GatewayOrderId = fmt.Sprintf("%v", gen.Next())
	instance.AppId = fmt.Sprintf("%v", gen.Next())
	instance.OutTradeNo = fmt.Sprintf("%v", gen.Next())
	if create := db.Create(instance); create.Error != nil {
		fmt.Println("Create error with message: ", create.Error)
	}
	find := db.Find(&model.PayOrder{})
	if rows, e := find.Rows(); e != nil {
		fmt.Println("error: ", e.Error())
	} else {
		//fmt.Println("Find result: ", rows)
		rows.Close()
	}
}
