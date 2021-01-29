package get

type IP struct {
	IP string `json:"IPADD"`
}

type AllInfo struct {
	Host   HostInfoStat      `json:"host"`
	CPU    []float64         `json:"cpu"`
	VirMem VirtualMemoryStat `json:"VirMem"`
	SwaMem SwapMemoryStat    `json:"SwaMem"`
}

type HostInfoStat struct {
	Hostname   string `json:"hostname"`
	Uptime     uint64 `json:"uptime"`
	Procs      uint64 `json:"procs"`      // number of processes
	OS         string `json:"os"`         // ex: freebsd, linux
	Platform   string `json:"platform"`   // ex: ubuntu, linuxmint
	KernelArch string `json:"kernelArch"` // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
}

type VirtualMemoryStat struct {
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
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
}
