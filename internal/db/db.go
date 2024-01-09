package db

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/scliangx/webServerTools/config"
	"github.com/sirupsen/logrus"
)

var (
	connection = make(map[string]*gorm.DB)
)

const (
	Mysql      = "mysql"
	PostgreSql = "postgres"
)

func NewConnection(dbConfig []config.Database) error {
	for _,cfg := range  dbConfig{
		if cfg.Driver == "mysql" {
			dbUri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
				cfg.Username,
				cfg.Password,
				cfg.Host,
				cfg.Port,
				cfg.Database)
			db, err := gorm.Open(cfg.Driver, dbUri)
			if err != nil {
				logrus.Error("init mysql database failed. [ERROR:]",err.Error())
			}
			connection[Mysql] = db
		} else if cfg.Driver == "postgres" {
			dbUri := fmt.Sprintf("host=%s port=%d dbname=%s user=%s  sslmode=disable password=%s",
				cfg.Host,
				cfg.Port,
				cfg.Database,
				cfg.Username,
				cfg.Password)
			db, err := gorm.Open(cfg.Driver, dbUri)
			if err != nil {
				logrus.Error("init postgres database failed. [ERROR:]",err.Error())
			}
			connection[PostgreSql] = db
		}
	}
	return nil
}

func BossDB() map[string]*gorm.DB {
	return connection
}
