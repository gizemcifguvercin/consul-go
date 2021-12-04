# Consul K/V Store Implementation For Go

<img src="https://i.ibb.co/kGwcF7S/Screen-Shot-2021-12-04-at-18-56-46.png"/>
<br/>Enables Consul to be used as a configuration source in go applications
<br/>Dynamic Configuration with Consul's Key/Value Store Feature


## Installation


```bash
go get -u github.com/gizemcifguvercin/consul-go
```

## Usage


```html
package main

```

```html

import (
	Consul "github.com/gizemcifguvercin/consul-go"
	"github.com/jasonlvhit/gocron"
)

```

```html
func main() {
	consulConfig := Consul.NewConsulConfig(5, "http://localhost:8500", "go", "")
	consulWatcher := Consul.NewConsulWatcher(&consulConfig)

	appConfig := NewApplicationConfig()
	go func() {
		gocron.Every(uint64(consulConfig.Interval)).Seconds().Do(consulWatcher.Watch, &appConfig)
		<-gocron.Start()
	}()
}

```
