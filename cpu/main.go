package main

import (
	"cpu/manage"
	"cpu/write"
	"net/http"
)

func main() {
	ip := "192.168.56.1"
	port := "50000"
	write.Write(ip, port)

	http.HandleFunc("/", manage.Top)
	http.ListenAndServe(":8080", nil)
}
