package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/scliang-strive/webServerTools/internal/db"
)

type User struct {
	*gorm.Model
	Username string
	Sex      string
	Age      string
}

func (*User) TableName() string {
	return "user"
}

func CreateTable() {
	err := db.BossDB()[db.Mysql].CreateTable(&User{}).Error
	if err != nil {
		fmt.Println("create table error:", err.Error())
	}
}

func ExportAllData() string {
	var s []string
	err := db.BossDB()[db.Mysql].Raw("select sex from user").Pluck("sex", &s).Error
	if err != nil {
		return ""
	}
	if len(s) == 0 {
		return ""
	}
	return s[0]
}
