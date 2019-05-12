package main

import (
	"fmt"
	"github.com/coreos/etcd/pkg/idutil"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pjoc-team/base-service/pkg/logger"
	d "github.com/pjoc-team/pay-database-service/pkg/db"
	"github.com/pjoc-team/pay-database-service/pkg/model"
	"time"
	pb "github.com/pjoc-team/pay-proto/go"
)

var gen = idutil.NewGenerator(12, time.Now())

func main() {
	var db *gorm.DB
	var err error
	if db, err = d.InitDb(); err != nil {
		return
	}
	insert(db)
	insertTest(db)
	findAll(db)
	findOne(db)
	findByWhere(db)
	updateColumns(db)
	updateAllColumns(db)
	updateColumn(db)

}
func findByWhere(db *gorm.DB) {
	logger.Log.Infof("Find one...")

	order := &model.PayOrder{}
	order.AppId = "test"
	order.OutTradeNo = "tesz"
	orders := make([]model.PayOrder, 0)
	find := db.Where("out_trade_no <= ?", order.OutTradeNo).Find(&orders)
	if find.RecordNotFound() {
		fmt.Println("error: ", find.Error.Error())
	} else {
		fmt.Println("findByWhere result: ", orders)
	}

	response := &pb.PayOrderResponse{}
	response.PayOrders = make([]*pb.PayOrder, 0)
	if err := copier.Copy(&response.PayOrders, orders); err != nil {
		logger.Log.Error("Copy result error! error: %v", err.Error())
	} else {
		logger.Log.Infof("Found result: %v by query: %v", response, order)
	}
}

func insertTest(db *gorm.DB) {
	instance := model.PayOrder{}
	instance.GatewayOrderId = "test"
	instance.AppId = "test"
	instance.OutTradeNo = "test"
	instance.Remark = fmt.Sprintf("%d", gen.Next())

	if !db.Find(&instance).RecordNotFound() {
		fmt.Println("Found record: ", instance)
		return
	} else if create := db.Create(instance); create.Error != nil {
		fmt.Println("Create error with message: ", create.Error)
	}
}

func insert(db *gorm.DB) {
	instance := model.PayOrder{}
	instance.GatewayOrderId = fmt.Sprintf("%d", gen.Next())
	instance.AppId = fmt.Sprintf("%d", gen.Next())
	instance.OutTradeNo = fmt.Sprintf("%d", gen.Next())
	if create := db.Create(instance); create.Error != nil {
		fmt.Println("Create error with message: ", create.Error)
	}
}

func findAll(db *gorm.DB) {
	logger.Log.Infof("Find all...")
	orders := make([]model.PayOrder, 0)
	find := db.Find(&orders)
	if find.RecordNotFound() {
		fmt.Println("error: ", find.Error.Error())
	} else {
		fmt.Println("Find result: ", orders)
	}
}
func findOne(db *gorm.DB) {
	logger.Log.Infof("Find one...")

	order := &model.PayOrder{}
	order.GatewayOrderId = "test"
	orders := make([]model.PayOrder, 0)
	find := db.Find(&orders, order)
	if find.RecordNotFound() {
		fmt.Println("error: ", find.Error.Error())
	} else {
		fmt.Println("Find result: ", orders)
	}

	response := &pb.PayOrderResponse{}
	response.PayOrders = make([]*pb.PayOrder, 0)
	if err := copier.Copy(&response.PayOrders, orders); err != nil {
		logger.Log.Error("Copy result error! error: %v", err.Error())
	} else {
		logger.Log.Infof("Found result: %v by query: %v", response, order)
	}
}

func updateColumn(db *gorm.DB) {
	logger.Log.Infof("Update...")

	instance2 := model.PayOrder{}
	instance2.GatewayOrderId = "test"
	instance2.AppId = "t123123est"
	update := db.Model(instance2).Update(instance2)
	//save := db.Save(&instance2)
	fmt.Println("Update column result: ", update.Error)
}

func updateColumns(db *gorm.DB) {
	logger.Log.Infof("Update...")

	instance2 := model.PayOrder{}
	instance2.GatewayOrderId = "test"
	instance2.AppId = "test"
	instance2.OutTradeNo = "test"
	instance2.Remark = fmt.Sprintf("%d", gen.Next())
	update := db.Model(&instance2).Update("remark", instance2.Remark)
	//save := db.Save(&instance2)
	fmt.Println("Update result: ", update.Error)
}

func updateAllColumns(db *gorm.DB) {
	logger.Log.Infof("Update all...")

	instance2 := &model.PayOrder{}
	instance2.GatewayOrderId = "test"
	instance2.AppId = fmt.Sprintf("%d", gen.Next())
	instance2.OutTradeNo = fmt.Sprintf("%d", gen.Next())
	instance2.Remark = fmt.Sprintf("%d", gen.Next())
	//db.NewRecord(instance2)
	//update := db.Save(instance2)
	update := db.Model(instance2).Update(instance2)
	//save := db.Save(&instance2)
	fmt.Println("Update all result: ", update.Error)
}
