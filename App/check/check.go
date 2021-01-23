package check

import (
	"app/json"
	"context"
	"fmt"

	"github.com/shirou/gopsutil/cpu"
)

func CheckStatus() []float64 {
	/*
			v, _ := mem.VirtualMemory()
			fmt.Println(v)
			fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

		//a, _ := cpu.Info()
		b, _ := cpu.Times(false)
		fmt.Println(b)
	*/

	ctx, _ := context.WithCancel(context.Background())
	var cpuPercent []float64
	cpuPercent, _ = cpu.PercentWithContext(ctx, 0, true)
	//fmt.Println(cpuPercent)
	//st <- cpuPercent
	//fmt.Println(cpuPercent)

	//data := json.JsonConv(cpuPercent)
	data := json.JsonConv(cpuPercent)
	fmt.Println(data)
	return cpuPercent
}
