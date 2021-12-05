package consul_go

import (
	"github.com/hashicorp/consul/api"
	"github.com/jasonlvhit/gocron"
)

type ConsulWatcher struct {
	consulClient *ConsulClient
}

func NewConsulWatcher(c *ConsulConfig) ConsulWatcher {
	consulApiConfig := api.DefaultConfig()
	consulApiConfig.Address = c.BaseUrl
	consulApiConfig.Token = c.ACL
	client, err := api.NewClient(consulApiConfig)

	if err != nil {
		panic("Client couldn't create !")
	}

	queryOptions := &api.QueryOptions{}

	consulClient := NewConsulClient()
	consulClient.APIClient = client
	consulClient.QueryOptions = queryOptions
	consulClient.Config = c
	return ConsulWatcher{consulClient: consulClient}
}

func (w *ConsulWatcher) Watch(c chan string) {
	data := w.consulClient.Read()
	c <- data
}

func (w *ConsulWatcher) StartWatch(c chan string) {
	gocron.Every(uint64(w.consulClient.Config.Interval)).Seconds().Do(w.Watch, c)
	<-gocron.Start()
}
