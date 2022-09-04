package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/scliang-strive/webServerTools/config"
	"log"
)

var (
	connection map[string]*gorm.DB
)

const (
	Mysql      = "mysql"
	PostgreSql = "postgres"
)

func NewConnection(cfg config.Database) {
	if cfg.Driver == "mysql" {
		dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
			cfg.UserName,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Driver)
		db, err := gorm.Open(cfg.Driver, dbUri)
		if err != nil {
			log.Print(err.Error())
		}
		connection[cfg.Driver] = db
	} else if cfg.Driver == "postgres" {
		dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			cfg.Host,
			cfg.Port,
			cfg.UserName,
			cfg.Database,
			cfg.Password)
		db, err := gorm.Open(cfg.Driver, dbUri)
		if err != nil {
			log.Print(err.Error())
		}
		connection[cfg.Driver] = db
	}
	return
}

func BossDB() map[string]*gorm.DB {
	return connection
}
