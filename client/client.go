package client

import (
	"github.com/gizemcifguvercin/consul-go/config"
	"github.com/hashicorp/consul/api"
)

type Client struct {
	ConsulClient *api.Client
	QueryOptions *api.QueryOptions
	Config       *config.Config
}

func NewClient(c *config.Config) Client {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = c.BaseUrl
	client, err := api.NewClient(consulConfig)

	if err != nil {
		panic("Client couldn't create !")
	}

	queryOptions := &api.QueryOptions{}
	return Client{ConsulClient: client, QueryOptions: queryOptions, Config: c}
}

func (c *Client) Read() (result string) {
	kv := c.ConsulClient.KV()
	pair, _, err := kv.List(c.Config.Prefix, c.QueryOptions)
	if err != nil {
		panic(err)
	}
	return string(pair[0].Value)
}
