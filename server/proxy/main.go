package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"proxy/conf"
)

func main() {
	host := conf.HostCheck()
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
