package status

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/shirou/gopsutil/host"
)

type HostInfoStat struct {
	Hostname   string `json:"hostname"`
	Uptime     uint64 `json:"uptime"`
	Procs      uint64 `json:"procs"`      // number of processes
	OS         string `json:"os"`         // ex: freebsd, linux
	Platform   string `json:"platform"`   // ex: ubuntu, linuxmint
	KernelArch string `json:"kernelArch"` // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
}

/* ホスト情報を取得する */
func HostCheck() HostInfoStat {
	hostStatus, err := host.Info()
	if err != nil {
		log.Println("[hostStatus]:", err)
	}
	var host HostInfoStat
	data, err := json.Marshal(hostStatus)
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(data, &host)
	if err != nil {
		fmt.Print(err)
	}
	return host
}
