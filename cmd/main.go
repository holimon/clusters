package main

import (
	"clusters/config"
	"clusters/internal/service"
	"clusters/pkg/consul"
	"clusters/pkg/logger"
	"clusters/pkg/sysinfo"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func usage() {
	fmt.Fprintf(os.Stderr, `Copyright reserved.
Usage: 
  clagent -server [-consul ip:port] [-bind port]
  clagent -client [-consul ip:port] [-bind port]
Options:
`)
	flag.PrintDefaults()
}

func realmain()  {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	logger.AppLogger.Info("start.")
	go consul.CheckRegistration()
	go service.Service()
	<-sig
	logger.AppLogger.Info("exit.")
}

func main() {
	flag.BoolVar(&(config.Configs.Client),"client",false,"Client mode.")
	flag.BoolVar(&(config.Configs.Server),"server",false,"Server mode.")
	flag.IntVar(&(config.Configs.Port),"bind",9090,"Sevice listen port.")
	flag.StringVar(&(config.Configs.Consul),"consul","127.0.0.1:8500","Consul sevice config.")
	flag.StringVar(&(config.Configs.WWW),"www","dist", "www folder path")
	flag.StringVar(&(config.Configs.ServiceName),"name","NodeAgent","Service registration name.")
	h:=flag.Bool("h",false,"Usage.")
	flag.Usage = usage
	flag.Parse()
	if(*h || ((!config.Configs.Client) && (!config.Configs.Server))){
		flag.Usage()
		os.Exit(-1)
	}
	if local,err := sysinfo.LocalIp();err==nil{
		config.Configs.LocalIP = local
		config.Configs.Uid = sysinfo.GetUUID()
		realmain()
	}else {
		panic(err)
	}
}
