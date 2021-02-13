package status

import (
	"context"
	"log"

	"github.com/shirou/gopsutil/cpu"
)

type CPUStat struct {
	Hostname string  `json:"hostname"`
	Cpu0     float64 `json:"cpu0"`
	Cpu1     float64 `json:"cpu1"`
	Cpu2     float64 `json:"cpu2"`
	Cpu3     float64 `json:"cpu3"`
}

/* CPUの使用率を取得する */
func CpuCheck() CPUStat {
	ctx, _ := context.WithCancel(context.Background())
	cpuPercent, err := cpu.PercentWithContext(ctx, 0, true)
	if err != nil {
		log.Fatal(err)
	}
	var cpu CPUStat = CPUStat{
		Cpu0: cpuPercent[0],
		Cpu1: cpuPercent[1],
		Cpu2: cpuPercent[2],
		//Cpu3: cpuPercent[3],
	}
	return cpu
}
