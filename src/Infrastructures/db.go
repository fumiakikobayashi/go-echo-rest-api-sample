package Infrastructures

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func SetupDB() *gorm.DB {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// 接続
	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err.Error())
	}

	// デバッグの設定
	debugMode, err := strconv.ParseBool(os.Getenv("DB_DEBUG_MODE"))
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(debugMode)

	return db
}
