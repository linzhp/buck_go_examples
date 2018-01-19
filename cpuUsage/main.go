package main

import (
	"fmt"
	"time"

	"code.uber.internal/go-common.git/x/log"
	"github.com/shirou/gopsutil/cpu"
)

func main() {
	usage, err := cpu.Percent(10*time.Second, true /*perCPU*/)
	if err != nil || len(usage) == 0 {
		log.Errorf("failed to get cpu usage: %v", err)
	} else {
		fmt.Println(usage)
	}
}
