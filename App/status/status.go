package status

type AllInfo struct {
	Host   HostInfoStat      `json:"host"`
	CPU    []float64         `json:"cpu"`
	VirMem VirtualMemoryStat `json:"VirMem"`
	SwaMem SwapMemoryStat    `json:"SwaMem"`
}

func StatusCheck() AllInfo {
	mem, swa := MemCheck()
	host := HostCheck()
	cpu := CpuCheck()
	allinfo := AllInfo{
		Host:   host,
		CPU:    cpu,
		VirMem: mem,
		SwaMem: swa,
	}
	return allinfo
}
