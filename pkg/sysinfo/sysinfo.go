package sysinfo

import (
	"clusters/config"
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	psdisk "github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	psnet "github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"net"
	"strings"
	"time"
)

type Processes struct {
	Pid     int32
	Name    string
	Username string
	Status  string
	Cpuuse  float64
	Memuse  float32
	Cmdline string
}

type Disks struct {
	Partion string
	Percent float64
}

func LocalIp() (local string, err error) {
	consuladdr := strings.Split(config.Configs.Consul, ":")[0]
	netadapters, _ := psnet.Interfaces()
	for _, item := range netadapters {
		for _, ip := range item.Addrs {
			ipp, subnet, _ := net.ParseCIDR(ip.Addr)
			if subnet.Contains(net.ParseIP(consuladdr)) {
				local = ipp.String()
				return
			}
		}
	}
	err = errors.New("Unable to find the IP in the same network segment as the consult service.")
	return
}

func GetPartion() (pardevice []string, err error) {
	par, err := psdisk.Partitions(true)
	for _, item := range par {
		pardevice = append(pardevice, item.Device)
	}
	return
}

func GetHostBrief() (info *host.InfoStat, err error) {
	return host.Info()
}

func GetProcessList() (processes []Processes, rerr error) {
	if tasks, err := process.Processes(); err != nil {
		rerr = err
		return
	} else {
		for _, item := range tasks {
			temp := Processes{}
			temp.Pid = item.Pid
			temp.Name, _ = item.Name()
			temp.Cpuuse, _ = item.CPUPercent()
			temp.Memuse, _ = item.MemoryPercent()
			temp.Status, _ = item.Status()
			temp.Cmdline, _ = item.Cmdline()
			temp.Username,_ = item.Username()
			processes = append(processes,temp)
		}
	}
	return
}

func ProcessKill(pid int32) (rerr error)  {
	if task,err := process.NewProcess(pid);err == nil{
		rerr = task.Kill()
	}else{
		rerr = err
	}
	return
}

func ProcessSuspend(pid int32) (rerr error) {
	if task,err := process.NewProcess(pid);err == nil{
		rerr = task.Suspend()
	}else{
		rerr = err
	}
	return
}

func ProcessResume(pid int32) (rerr error) {
	if task,err := process.NewProcess(pid);err == nil{
		rerr = task.Resume()
	}else{
		rerr = err
	}
	return
}

func GetCpuPercent() (percent float64) {
	percent = 0
	temp, _ := cpu.Percent(0, false)
	percent = temp[0]
	return
}

func GetDiskPercent() (disks []Disks,rerr error){
	if pars,err := GetPartion();err == nil{
		for _,dev := range pars{
			if par,err := psdisk.Usage(dev);err == nil{
				temp:=Disks{}
				temp.Partion = dev
				temp.Percent = par.UsedPercent
				disks = append(disks,temp)
			}
		}
	}else{
		rerr = err
	}
	return
}

func GetMemPercent() (percent float64) {
	percent = 0
	temp, _ := mem.VirtualMemory()
	percent = temp.UsedPercent
	return
}

//获取单位时间内指定网卡收发速率（单位时间：1S，单位MB）
func GetNetAdapterRate(adapter string) (inRate, outRate float64) {
	curnetio, _ := psnet.IOCounters(true)
	time.Sleep(time.Second)
	latnetio, _ := psnet.IOCounters(true)
	for i, item := range curnetio {
		if item.Name != adapter {
			continue
		}
		inRate = float64(latnetio[i].BytesRecv-curnetio[i].BytesRecv) / 1024.0 / 1024.0
		outRate = float64(latnetio[i].BytesSent-curnetio[i].BytesSent) / 1024.0 / 1024.0
	}
	return
}

func GetUUID()string{
	if info,err := host.Info();err == nil{
		return info.HostID
	}else{
		return fmt.Sprintf("UUID-%s:%d",config.Configs.LocalIP,config.Configs.Port)
	}
}