package status

import (
	"context"
	"log"

	"github.com/shirou/gopsutil/cpu"
)

func CpuCheck() []float64 {
	ctx, _ := context.WithCancel(context.Background())
	cpuPercent, err := cpu.PercentWithContext(ctx, 0, true)
	if err != nil {
		log.Fatal(err)
	}
	return cpuPercent
}
