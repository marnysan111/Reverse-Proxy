package main

import (
	"app/check"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	/*
		st := make(chan []float64)
		go check.CheckStatus(st)
		status := <-st
		fmt.Println("status is", status)
	*/
	//go check.CheckStatus()
	r := mux.NewRouter()
	r.HandleFunc("/json", jsonCheck)
	http.Handle("/", r)
	http.ListenAndServe(":50000", nil)
}

func jsonCheck(w http.ResponseWriter, q *http.Request) {
	status := check.CheckStatus()
	fmt.Fprint(w, status)
}
