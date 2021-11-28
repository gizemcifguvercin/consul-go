package main

import (
	Config "github.com/gizemcifguvercin/consul-go/config"
	Tasks "github.com/gizemcifguvercin/consul-go/tasks"
)

var task Tasks.Tasks

func main() {
	cfg := Config.NewConfig(5, "http://localhost:8500", "go", "")
	task.Runner(&cfg)
}
