package main

import (
    "fmt"
    "github.com/shirou/gopsutil/host"
    "github.com/shirou/gopsutil/load"
)

func main() {
    infoStat, _ := host.Info()
    fmt.Printf("Total processes: %d\n", infoStat.Procs)

    miscStat, _ := load.Misc()
    fmt.Printf("Running processes: %d\n", miscStat.ProcsRunning)
}
