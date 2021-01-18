package main

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	/*
			v, _ := mem.VirtualMemory()
			fmt.Println(v)
			fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

		//a, _ := cpu.Info()
		b, _ := cpu.Times(false)
		fmt.Println(b)
	*/

	ctx, _ := context.WithCancel(context.Background())
	fmt.Println(ctx)
	var cpuPercent []float64
	for {
		cpuPercent, _ = cpu.PercentWithContext(ctx, 0, true)
		fmt.Println(cpuPercent)
		time.Sleep(time.Second * 2)
	}
}
