package get

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetStatus() {
	host := GetHost()
	for _, h := range host {
		GetDo(h.IP)
	}
}

func GetHost() []IP {
	var IpAdd []IP
	bytes, err := ioutil.ReadFile("get/ip.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(bytes, &IpAdd)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//data, err := json.Marshal(IpAdd)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	/*
		fmt.Println(IpAdd)
		for _, p := range IpAdd {
			fmt.Println(p.IP)
		}
	*/
	return IpAdd
}

func GetDo(ipadd string) {
	r, err := http.Get("http://" + ipadd + "/json")
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var allinfo AllInfo
	if err := json.Unmarshal(body, &allinfo); err != nil {
		fmt.Println(err)
	}
	fmt.Println(allinfo)
}
