package consul

import (
	"clusters/config"
	"clusters/pkg/logger"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"time"
)

func CheckRegistration()  {
	consulcof := consulapi.DefaultConfig()
	consulcof.Address = config.Configs.Consul
	client,err := consulapi.NewClient(consulcof)
	if err != nil{
		return
	}
	for{
		if _,_,err := client.Agent().Service(config.Configs.Uid,nil);err != nil{
			logger.AppLogger.Info("Possible unregistered detected.")
			//registration
			registration := new(consulapi.AgentServiceRegistration)
			registration.ID = config.Configs.Uid
			registration.Name = config.Configs.ServiceName
			registration.Address = config.Configs.LocalIP
			registration.Port = config.Configs.Port
			registration.Check = &consulapi.AgentServiceCheck{
				HTTP:fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/api/v1/check"),
				Timeout:"1s",
				Interval:"5s",
				DeregisterCriticalServiceAfter:"60s",
			}
			if err = client.Agent().ServiceRegister(registration);err != nil {
				logger.AppLogger.Error("Service register err.")
			}
		}
		time.Sleep(time.Second * 5)
	}
}