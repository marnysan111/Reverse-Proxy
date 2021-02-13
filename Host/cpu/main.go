package main

import (
	"cpu/db"
	"cpu/get"
)

func main() {
	db.DbInit()
	get.GetStatus()
}
