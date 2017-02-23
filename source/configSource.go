package source

import (
	"github.com/hashicorp/consul/api"
)

type ConfigSource interface {
	Configure(config *interface{}) error
}


type consulConfigSource struct {
	ConfigSource
	client *api.Client
	configKeys []string
}

func (src consulConfigSource) Configure(config *interface{}) error{
	return nil
}

func Consul(client *api.Client, configKeys []string) *ConfigSource {
	return nil
	//config := consulConfigSource{
	//
	//}
	//return config
}