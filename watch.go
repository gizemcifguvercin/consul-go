package consul_go

import (
	"encoding/json"
	"github.com/hashicorp/consul/api"
	"github.com/jasonlvhit/gocron"
)

type ConsulWatcher struct {
	consulClient *ConsulClient
}

func NewConsulWatcher(c *ConsulConfig) ConsulWatcher {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = c.BaseUrl
	consulConfig.Token = c.ACL
	client, err := api.NewClient(consulConfig)

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

func (w *ConsulWatcher) Watch(appConfig *AppConfig) {
	data := w.consulClient.Read()
	json.Unmarshal([]byte(data), &appConfig)
}

func (w *ConsulWatcher) StartWatch(consulConfig *ConsulConfig, appConfig *AppConfig) {
	go func() {
		gocron.Every(uint64(consulConfig.Interval)).Seconds().Do(w.Watch, &appConfig)
		<-gocron.Start()
	}()
}
