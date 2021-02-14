package get

import (
	"cpu/db"
	"cpu/status"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/* jsonのIPを取得してGETする関数を呼び出す */
func GetStatus() {
	host, err := GetHost()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, h := range host {
		data := GetDo(h.IP, h.PORT)
		db.DbInsert(data)
	}
}

/* jsonからGETするマシンのIPを取得する */
func GetHost() ([]status.IP, error) {
	var IpAdd []status.IP
	bytes, err := ioutil.ReadFile("get/ip.json")
	if err != nil {
		return IpAdd, err
	}
	err = json.Unmarshal(bytes, &IpAdd)
	if err != nil {
		return IpAdd, err
	}
	return IpAdd, nil
}

/* Getする関数 */
func GetDo(ipadd string, port string) status.AllInfo {
	var allinfo status.AllInfo
	r, err := http.Get("http://" + ipadd + ":" + port + "/json")
	if err != nil {
		fmt.Println(err)
		return allinfo
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return allinfo
	}
	if err := json.Unmarshal(body, &allinfo); err != nil {
		fmt.Println(err)
		return allinfo
	}

	allinfo.CPU.Hostname = allinfo.Host.Hostname
	allinfo.VirMem.Hostname = allinfo.Host.Hostname
	allinfo.SwapMem.Hostname = allinfo.Host.Hostname
	allinfo.Host.IpAdd = ipadd
	allinfo.Host.Port = port

	return allinfo
}
