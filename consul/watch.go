package consul

import (
	"github.com/gizemcifguvercin/consul-go/client"
)

type Consul struct {
	client *client.Client
}

func NewConsul(client *client.Client) Consul {
	return Consul{client: client}
}

func (consul *Consul) Watch() (data string) {
	data = consul.client.Read()
	return data
}
