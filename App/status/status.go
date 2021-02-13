package status

type AllInfo struct {
	Host   HostInfoStat      `json:"host"`
	CPU    CPUStat           `json:"cpu"`
	VirMem VirtualMemoryStat `json:"VirMem"`
	SwaMem SwapMemoryStat    `json:"SwaMem"`
}

/* 取得した情報を一つにまとめる */
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
