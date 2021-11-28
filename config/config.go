package config

type Config struct {
	Interval int
	BaseUrl  string
	Prefix   string
	ACL      string
}

func NewConfig(interval int, baseUrl string, prefix string, acl string) Config {
	return Config{Interval: interval, BaseUrl: baseUrl, Prefix: prefix, ACL: acl}
}
