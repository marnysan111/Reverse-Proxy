package main

import (
	"app/status"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/json", jsonCheck)
	http.Handle("/", r)
	http.ListenAndServe(":50000", nil)
}

func jsonCheck(w http.ResponseWriter, q *http.Request) {
	status := status.StatusCheck()
	fmt.Fprint(w, status)
}
