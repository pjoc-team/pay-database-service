package main

import (
	"flag"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pjoc-team/base-service/pkg/logger"
	s "github.com/pjoc-team/base-service/pkg/service"
	"github.com/pjoc-team/pay-database-service/pkg/db"
	"github.com/pjoc-team/pay-database-service/pkg/service"
)

var (
	listenAddr            = flag.String("listen-addr", ":8080", "HTTP listen address.")
	configURI             = flag.String("c", "config.yaml", "uri to load config")
	tlsEnable             = flag.Bool("tls", false, "enable tls")
	logLevel              = flag.String("log-level", "debug", "logger level")
	logFormat             = flag.String("log-format", "text", "text or json")
	caCert                = flag.String("ca-cert", s.WithConfigDir("ca.pem"), "Trusted CA certificate.")
	tlsCert               = flag.String("tls-cert", s.WithConfigDir("cert.pem"), "TLS server certificate.")
	tlsKey                = flag.String("tls-key", s.WithConfigDir("key.pem"), "TLS server key.")
	serviceName           = flag.String("s", "", "Service name in service discovery.")
	registerServiceToEtcd = flag.Bool("r", true, "Register service to etcd.")
	etcdPeers             = flag.String("etcd-peers", "", "Etcd peers. example: 127.0.0.1:2379,127.0.0.1:12379")
)

func main() {
	var dbConn *gorm.DB
	var err error
	if dbConn, err = db.InitDb(); err != nil {
		logger.Log.Errorf("Failed to init db! error: %v", err.Error())
		return
	}
	defer dbConn.Close()
	flag.Parse()
	svc := s.InitService(*listenAddr,
		*configURI,
		*tlsEnable,
		*logLevel,
		*logFormat,
		*caCert,
		*tlsCert,
		*tlsKey,
		*serviceName,
		*registerServiceToEtcd,
		*etcdPeers,
		"")
	service.Init(svc, dbConn)
}
