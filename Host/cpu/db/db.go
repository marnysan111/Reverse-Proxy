package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbCommect() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	USER := os.Getenv("USER")
	PASS := os.Getenv("DBPASS")
	PROTOCOL := "tcp(127.0.0.1:3306)"
	TABLE := "proxy"

	connect := USER + ":" + PASS + "@" + PROTOCOL + "/" + TABLE + "?charset=utf8mb4&parseTime=True&loc=Local"
	return connect
}

//const connect = "root:Maanii01!@tcp(127.0.0.1:3306)/proxy?charset=utf8mb4&parseTime=True&loc=Local"

func DbInit() {
	connect := DbCommect()
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println(err)
	}
	db.Table("HostStat").AutoMigrate(&HostInfoStat{})
	db.Table("VirtualStat").AutoMigrate(&VirtualMemoryStat{})
	db.Table("SwapStat").AutoMigrate(&SwapMemoryStat{})
}

//https://www.dbonline.jp/mysql/type/index8.html これは？
