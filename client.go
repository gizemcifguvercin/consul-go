package consul_go

import (
	"github.com/hashicorp/consul/api"
)

type ConsulClient struct {
	APIClient    *api.Client
	QueryOptions *api.QueryOptions
	Config       *ConsulConfig
}

func NewConsulClient() *ConsulClient {
	return &ConsulClient{}
}

func (c *ConsulClient) Read() (result string) {
	kv := c.APIClient.KV()
	pair, _, err := kv.List(c.Config.Prefix, c.QueryOptions)
	if err != nil {
		panic(err)
	}
	return string(pair[0].Value)
}
