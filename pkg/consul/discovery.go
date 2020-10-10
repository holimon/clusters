package consul

import (
	"clusters/config"
	"clusters/pkg/logger"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
)

func ServiceDiscovery()(nodes []map[string]string,rerr error)  {
	consulcof := consulapi.DefaultConfig()
	consulcof.Address = config.Configs.Consul
	if client,err := consulapi.NewClient(consulcof);err == nil{
		services, _, err := client.Health().Service(config.Configs.ServiceName,"",true,nil)
		if err == nil {
			for _, item := range  services{
				logger.AppLogger.Info(fmt.Sprintf("Addr:%s Port:%d UID:%s",item.Service.Address, item.Service.Port,item.Service.ID))
				nodes = append(nodes, map[string]string{"uuid":item.Service.ID,"addr":fmt.Sprintf("%s:%d",item.Service.Address,item.Service.Port)})
			}
		}else {
			rerr = err
		}
	}else {
		rerr = err
	}
	return
}