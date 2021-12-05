# Consul K/V Store Implementation For Go

<img src="https://i.ibb.co/kGwcF7S/Screen-Shot-2021-12-04-at-18-56-46.png"/>
<br/>Enables Consul to be used as a configuration source in go applications
<br/>Dynamic Configuration with Consul's Key/Value Store Feature


## Installation


```bash
go get -u github.com/gizemcifguvercin/consul-go
```

## Create Your Own Config 

Match with the data model on Consul config json
</br>Example:

<img src="https://i.ibb.co/ScnJtQZ/Screen-Shot-2021-12-04-at-22-22-41.png"/>

```html

package main

type ApplicationConfig struct {
	Salary string `json:"Salary"`
	Title  string `json:"Title"`
}

func NewApplicationConfig() ApplicationConfig {
	return ApplicationConfig{}
}


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

</br>You can set some of Consul Config props with this order below;
       </br> <b>Interval</b> int
       </br> <b>BaseUrl</b>  string
       </br> <b>Prefix</b>   string
       </br> <b>ACL</b>      string

```html
This project is still developing. Feel free to give feedback or send pull request
```
