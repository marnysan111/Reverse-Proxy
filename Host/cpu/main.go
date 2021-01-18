package main

import (
	"cpu/manage"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	/*
		ip := "192.168.56.1"
		port := "50000"
		write.Write(ip, port)
	*/

	var addr = flag.String("addr", ":50000", "アプリケーションのアドレス")
	http.HandleFunc("/", manage.Top)

	fmt.Println("Webサーバを開始します。ポート:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
