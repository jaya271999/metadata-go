package main

import (
	"log"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

stat, err := linuxproc.ReadStat("/proc/stat")
if err != nil {
	log.Fatal("stat read fail")
}

for _, s := range stat.CPUStats {
	s.User
	s.Nice
	s.System
	s.Idle
	s.IOWait
}
