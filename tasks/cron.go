package tasks

import (
	"github.com/gizemcifguvercin/consul-go/client"
	"github.com/gizemcifguvercin/consul-go/config"
	"github.com/gizemcifguvercin/consul-go/consul"
	"github.com/jasonlvhit/gocron"
)

type Tasks struct {
}

func (t *Tasks) Runner(config *config.Config) {
	client := client.NewClient(config)
	consul := consul.NewConsul(&client)
	gocron.Every(uint64(config.Interval)).Seconds().Do(consul.Watch)
	<-gocron.Start()
}
