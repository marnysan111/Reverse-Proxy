package form

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Status struct {
	Id       string `json:"id"`
	HostName string `json:"hostname"`
	CPU      string `json:"cpu"`
	IP       string `json:"ip"`
	PORT     int    `json:"port"`
}

func PostForm() {
	var status Status = Status{
		Id:       "2",
		HostName: "aaa",
		CPU:      "100",
	}
	fmt.Println(status)
	values, err := json.Marshal(Status{
		Id:       "1",
		HostName: "test",
		CPU:      "1",
	})
	fmt.Println(bytes.NewBuffer(values))
	_, err = http.Post("http://192.168.1.10:50000", "application/json", bytes.NewBuffer(values))
	if err != nil {
		log.Fatal(err)
	}

}
