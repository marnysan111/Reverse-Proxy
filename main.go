package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"proxy/conf"
	"proxy/write"
)

func main() {
	ip := "192.168.56.1"
	port := "3000"
	write.Write(ip, port)
	host := conf.Check()
	director := func(request *http.Request) {
		request.URL.Scheme = "http"
		request.URL.Host = host.IP + ":" + host.Port
	}
	rp := &httputil.ReverseProxy{Director: director}
	server := http.Server{
		Addr:    ":80",
		Handler: rp,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(director)
		log.Fatal(err.Error())
	}
}
