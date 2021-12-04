package consul_go

type AppConfig struct {
	Salary string `json:"Salary"`
	Title  string `json:"Title"`
}

func NewAppConfig() AppConfig {
	return AppConfig{}
}
