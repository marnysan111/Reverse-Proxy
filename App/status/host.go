package status

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/shirou/gopsutil/host"
)

type HostInfoStat struct {
	Hostname             string `json:"hostname"`
	Uptime               uint64 `json:"uptime"`
	BootTime             uint64 `json:"bootTime"`
	Procs                uint64 `json:"procs"`           // number of processes
	OS                   string `json:"os"`              // ex: freebsd, linux
	Platform             string `json:"platform"`        // ex: ubuntu, linuxmint
	PlatformFamily       string `json:"platformFamily"`  // ex: debian, rhel
	PlatformVersion      string `json:"platformVersion"` // version of the complete OS
	KernelVersion        string `json:"kernelVersion"`   // version of the OS kernel (if available)
	KernelArch           string `json:"kernelArch"`      // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
	VirtualizationSystem string `json:"virtualizationSystem"`
	VirtualizationRole   string `json:"virtualizationRole"` // guest or host
	HostID               string `json:"hostid"`             // ex: uuid
}

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
