package config

type StructConfig struct {
	LogFile     string
	Client      bool
	Server      bool
	Consul      string
	Port        int
	ServiceName string
	Uid         string
	LocalIP     string
	WWW 		string
}

var Configs StructConfig = StructConfig{LogFile:"ClusterNode.log"}