package json

import (
	"encoding/json"
)

type Status struct {
	Id       int      `json:"id"`
	HostName string   `json:"hostname"`
	CPU      []Detail `json:"detail"`
	IP       string   `json:"ip"`
	PORT     int      `json:"port"`
}

type Detail struct {
	Cpu1 float64 `json:"cpu1"`
	Cpu2 float64 `json:"cpu2"`
	Cpu3 float64 `json:"cpu3"`
}

func JsonConv(cpuPercent []float64) string {
	data := []Status{
		Status{
			Id:       1,
			HostName: "rasPi 1",
			CPU: []Detail{
				{
					Cpu1: cpuPercent[0],
					Cpu2: cpuPercent[1],
					Cpu3: cpuPercent[2],
				},
			},
			PORT: 50000,
		},
	}
	//fmt.Println(data)
	test, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	//var st Status
	//fmt.Println(string(test[1]), reflect.TypeOf(string(test[1])))
	//for i := 0; i <=
	//fmt.Println(string(test))
	return string(test)
}
