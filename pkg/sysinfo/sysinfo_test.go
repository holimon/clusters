package sysinfo

import (
	"fmt"
	"testing"
)

func TestGetHostBrief(t *testing.T) {
	GetHostBrief()
}

func TestGetDiskPercent(t *testing.T) {
	if per,err := GetDiskPercent(); err==nil {
		fmt.Println(per)
	}
}

func TestGetCpuPercent(t *testing.T) {
	if per := GetCpuPercent(); per != 0 {
		fmt.Println(per)
	}
}

func TestGetProcessList(t *testing.T) {
	GetProcessList()
}

func TestProcessSuspend(t *testing.T) {
	ProcessSuspend(732)
}

func TestProcessResume(t *testing.T) {
	ProcessResume(732)
}