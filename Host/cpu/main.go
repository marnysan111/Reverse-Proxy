package main

import (
	"cpu/db"
	"cpu/get"
	"cpu/manage"
	"cpu/write"

	"github.com/gin-gonic/gin"
)

func main() {
	/* ラズパイの情報vを取得する */
	db.DbInit()

	//for {}
	get.GetStatus()
	write.HostGet()
	// ip, port := ConnectCheck()
	// write.WriteEnv()

	/* 管理用ページの構成 */
	r := gin.Default()
	r.LoadHTMLGlob("./manage/template/*.html")

	r.GET("/", manage.Index)
	r.GET("/detail/:id", manage.Detail)

	r.Run(":8080")
}
