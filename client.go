package consul_go

import (
	"log"

	"github.com/hashicorp/consul/api"
)

type ConsulClient struct {
	APIClient    *api.Client
	QueryOptions *api.QueryOptions
	Config       *ConsulConfig
	OnLoad       bool
}

func NewConsulClient() *ConsulClient {
	return &ConsulClient{OnLoad: true}
}

func (c *ConsulClient) Read() (result string) {
	defer func() {
		if err := recover(); err != nil {
			if c.OnLoad {
				log.Println("OnLoad Error!", err)
			} else {
				log.Println("OnWatch Error!", err)
			}
		}
	}()
	kv := c.APIClient.KV()
	pair, _, err := kv.List(c.Config.Prefix, c.QueryOptions)
	if err != nil {
		panic(err)
	}
	c.OnLoad = false
	return string(pair[0].Value)
}
