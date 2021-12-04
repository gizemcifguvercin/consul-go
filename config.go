package consul_go

type ConsulConfig struct {
	Interval int
	BaseUrl  string
	Prefix   string
	ACL      string
}

func NewConsulConfig(interval int, baseUrl string, prefix string, acl string) ConsulConfig {
	return ConsulConfig{Interval: interval, BaseUrl: baseUrl, Prefix: prefix, ACL: acl}
}
