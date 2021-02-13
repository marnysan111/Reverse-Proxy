package status

import "gorm.io/gorm"

type IP struct {
	IP string `json:"IPADD"`
}

type AllInfo struct {
	Host    HostInfoStat      `json:"HOST"`
	CPU     CPUStat           `json:"CPU"`
	VirMem  VirtualMemoryStat `json:"VirMem"`
	SwapMem SwapMemoryStat    `json:"SwaMem"`
}

type HostInfoStat struct {
	gorm.Model
	Hostname   string `json:"hostname"`
	IpAdd      string `json:"IP"`
	Uptime     uint64 `json:"uptime"`
	Procs      uint64 `json:"procs"`      // number of processes
	OS         string `json:"os"`         // ex: freebsd, linux
	Platform   string `json:"platform"`   // ex: ubuntu, linuxmint
	KernelArch string `json:"kernelArch"` // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
}

type CPUStat struct {
	gorm.Model
	Hostname string  `json:"hostname"`
	Cpu0     float64 `json:"cpu0"`
	Cpu1     float64 `json:"cpu1"`
	Cpu2     float64 `json:"cpu2"`
	Cpu3     float64 `json:"cpu3"`
}

type VirtualMemoryStat struct {
	gorm.Model
	Hostname    string  `json:"hostname"`
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
	Free        uint64  `json:"free"`
	Active      uint64  `json:"active"`
	Buffers     uint64  `json:"buffers"`
	Cached      uint64  `json:"cached"`
}

type SwapMemoryStat struct {
	gorm.Model
	Hostname    string  `json:"hostname"`
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
}

type HostList struct {
	gorm.Model
	Hostname string `json:"hostname"`
	IpAdd    string `json:"IP"`
}
