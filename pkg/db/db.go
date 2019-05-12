package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pjoc-team/base-service/pkg/logger"
	logger2 "github.com/pjoc-team/pay-database-service/pkg/logger"
	"github.com/pjoc-team/etcd-config/config"
	"github.com/pjoc-team/pay-database-service/pkg/conf"
	"time"
)

func InitDb() (db *gorm.DB, err error) {
	dbConfig := &conf.DbConfig{Dialect: "sqlite3", DbUrl: "tmp/gorm.dbConn", MaxIdleConnections: 100, MaxOpenConnections: 100, ConnectionMaxLifetimeSeconds: 14400}
	config.Init(config.URL("file://conf/mysql.yaml"), config.WithDefault(dbConfig))
	//unmarshal := yaml.UnmarshalFromFile("conf/dbConn.yaml", dbConfig)
	fmt.Println(dbConfig)
	if db, err = gorm.Open(dbConfig.Dialect, dbConfig.DbUrl); err != nil {
		logger.Log.Errorf("Failed to init dbConn! dialect: %s url: %s error: %s", dbConfig.Dialect, dbConfig.DbUrl, err.Error())
		return
	} else {
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		db.DB().SetMaxIdleConns(dbConfig.MaxIdleConnections)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		db.DB().SetMaxOpenConns(dbConfig.MaxOpenConnections)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		//db.DB().SetConnMaxLifetime(time.Hour)
		db.DB().SetConnMaxLifetime(time.Duration(dbConfig.ConnectionMaxLifetimeSeconds) * time.Second)

		// print model logs
		db.LogMode(true)

		// rewrite log
		logrusLogger := &logger2.LogrusLogger{}
		db.SetLogger(logrusLogger)
		logger.Log.Infof("Succeed connect db: %v", db)
		return
	}
}
