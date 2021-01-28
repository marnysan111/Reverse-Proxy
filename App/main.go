package main

import (
	"app/status"
	"encoding/json"
	"fmt"
	"log"
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
	data := status.StatusCheck()
	status, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(status))
}
