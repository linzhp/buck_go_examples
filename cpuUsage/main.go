package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	usage, err := cpu.Percent(10*time.Second, true /*perCPU*/)
	if err != nil || len(usage) == 0 {
		fmt.Printf("failed to get cpu usage: %v\n", err)
	} else {
		fmt.Println(usage)
	}
}
