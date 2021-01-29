package status

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

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

func MemCheck() (VirtualMemoryStat, SwapMemoryStat) {
	var virtualmemory VirtualMemoryStat
	var swapmemory SwapMemoryStat
	vir, err := mem.VirtualMemory()
	swa, _ := mem.SwapMemory()
	if err != nil {
		fmt.Println("[Error]:", err)
	}

	jvir, err := json.Marshal(vir)
	if err := json.Unmarshal(jvir, &virtualmemory); err != nil {
		fmt.Println("[Error]:", err)
	}
	jswa, err := json.Marshal(swa)
	if err := json.Unmarshal(jswa, &swapmemory); err != nil {
		fmt.Println("[Error]:", err)
	}

	return virtualmemory, swapmemory
}
