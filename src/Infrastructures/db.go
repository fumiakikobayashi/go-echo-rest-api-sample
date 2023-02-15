package Infrastructures

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() *gorm.DB {
	DBMS := "mysql"
	USER := "user"
	PASS := "password"
	HOST := "tcp(db:3306)"
	DBNAME := "go-test"
	config := USER + ":" + PASS + "@" + HOST + "/" + DBNAME

	db, err = gorm.Open(DBMS, config)
	if err != nil {
		// 例外処理
		fmt.Println("接続の失敗")
	}
	return db
}

func GetDb() *gorm.DB {
	return db
}
