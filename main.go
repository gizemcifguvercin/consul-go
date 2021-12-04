package main

import (
	"github.com/jasonlvhit/gocron"
)

func main() {
	//our custom configs
	appConfig := NewAppConfig()

	cfg := NewConsulConfig(5, "http://localhost:8500", "go", "")
	consulWatcher := NewConsulWatcher(&cfg)

	gocron.Every(uint64(cfg.Interval)).Seconds().Do(consulWatcher.Watch, &appConfig)
	<-gocron.Start()
}
