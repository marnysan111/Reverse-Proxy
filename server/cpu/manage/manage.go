package manage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type Status struct {
	Id       int    `json:"id"`
	HostName string `json:"hostname"`
	CPU      string `json:"cpu"`
	IP       string `json:"ip"`
	PORT     int    `json:"port"`
}

func Top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./manage/template/index.html")
	if err != nil {
		fmt.Println(err)
		log.Fatalf("template ERROR")
	}
	method := r.Method
	fmt.Println("[Method]:", method)

	if method == "POST" {
		Check(r)

	}

	t.Execute(w, nil)
}

func Check(r *http.Request) {
	defer r.Body.Close()
	var status Status
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &status)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[ID]", status.Id)
	JsonWrite(status)
}

func JsonWrite(status Status) {
	data := map[string]interface{}{
		"id":       status.Id,
		"HostName": status.HostName,
	}
	fmt.Println(data)
}
