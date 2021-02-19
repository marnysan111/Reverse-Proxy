package write

import (
	"cpu/db"
	"cpu/status"
	"fmt"
	"io/ioutil"
)

/* 送信先のホストを決めるための評価 */
func HostGet() {
	DB, err := db.DbCommect()
	if err != nil {
		fmt.Println(err)
		return
	}
	var hostlist []status.HostList
	dbClose, err := DB.DB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbClose.Close()
	//var host status.HostList
	var id int
	var score float64
	DB.Select("id,hostname").Find(&hostlist)
	for _, i := range hostlist {
		var s float64
		cpu, err := db.CpuCheckOne(i.Hostname)
		if err != nil {
			fmt.Println(err)
			return
		}
		vir, err := db.VirCheckOne(i.Hostname)
		if err != nil {
			fmt.Println(err)
			return
		}
		swa, err := db.SwaCheckOne(i.Hostname)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, j := range cpu {
			s += j.Cpu0 + j.Cpu1 + j.Cpu2 + j.Cpu3
		}
		s = s / 4
		for _, j := range vir {
			s += j.UsedPercent
		}
		for _, j := range swa {
			s += j.UsedPercent
		}
		if score >= s || score == 0 {
			id = i.Id
			score = s
		}
		//fmt.Println(i.Hostname, s)
	}
	host, err := db.HostCheckOne(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = WriteEnv(host.IpAdd, host.Forward)
	if err != nil {
		fmt.Println(err)
		return
	}
}

/* Host/proxyのenvに書き込み */
func WriteEnv(ip string, port string) error {
	var lines = []string{"HOST_IP=" + ip + "\nHOST_PORT=" + port}
	b := []byte{}
	for _, line := range lines {
		ll := []byte(line)
		for _, l := range ll {
			b = append(b, l)
		}
	}
	err := ioutil.WriteFile("../proxy/.env", b, 0666)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	return nil
}
