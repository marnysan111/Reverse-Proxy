package status

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

type VirtualMemoryStat struct {
	Total          uint64  `json:"total"`
	Available      uint64  `json:"available"`
	Used           uint64  `json:"used"`
	UsedPercent    float64 `json:"usedPercent"`
	Free           uint64  `json:"free"`
	Active         uint64  `json:"active"`
	Inactive       uint64  `json:"inactive"`
	Wired          uint64  `json:"wired"`
	Laundry        uint64  `json:"laundry"`
	Buffers        uint64  `json:"buffers"`
	Cached         uint64  `json:"cached"`
	Writeback      uint64  `json:"writeback"`
	Dirty          uint64  `json:"dirty"`
	WritebackTmp   uint64  `json:"writebacktmp"`
	Shared         uint64  `json:"shared"`
	Slab           uint64  `json:"slab"`
	SReclaimable   uint64  `json:"sreclaimable"`
	SUnreclaim     uint64  `json:"sunreclaim"`
	PageTables     uint64  `json:"pagetables"`
	SwapCached     uint64  `json:"swapcached"`
	CommitLimit    uint64  `json:"commitlimit"`
	CommittedAS    uint64  `json:"committedas"`
	HighTotal      uint64  `json:"hightotal"`
	HighFree       uint64  `json:"highfree"`
	LowTotal       uint64  `json:"lowtotal"`
	LowFree        uint64  `json:"lowfree"`
	SwapTotal      uint64  `json:"swaptotal"`
	SwapFree       uint64  `json:"swapfree"`
	Mapped         uint64  `json:"mapped"`
	VMallocTotal   uint64  `json:"vmalloctotal"`
	VMallocUsed    uint64  `json:"vmallocused"`
	VMallocChunk   uint64  `json:"vmallocchunk"`
	HugePagesTotal uint64  `json:"hugepagestotal"`
	HugePagesFree  uint64  `json:"hugepagesfree"`
	HugePageSize   uint64  `json:"hugepagesize"`
}

type SwapMemoryStat struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
	Sin         uint64  `json:"sin"`
	Sout        uint64  `json:"sout"`
	PgIn        uint64  `json:"pgin"`
	PgOut       uint64  `json:"pgout"`
	PgFault     uint64  `json:"pgfault"`
	PgMajFault  uint64  `json:"pgmajfault"`
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
