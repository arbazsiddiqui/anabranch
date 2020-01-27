package main

import (
	"anabranch/anabranch"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type Config struct {
	Port                int      `json:"port"`
	Hosts               []string `json:"hosts"`
	Strategy            string   `json:"strategy"`
	HealthCheckInterval int      `json:"health_check_interval"`
	AddRequestId        bool     `json:"add_request_id"`
	HealthCheckType     string   `json:"health_check_type"`
}

func main() {
	configPath := flag.String("config", "./config.json", "path to config")
	flag.Parse()
	file, _ := os.Open(*configPath)
	
	var config Config
	_ = json.NewDecoder(file).Decode(&config)
	http.ListenAndServe(fmt.Sprintf(":%v", config.Port), anabranch.NewLoadBalancer(
		config.Hosts,
		config.Strategy,
		config.HealthCheckInterval,
		config.HealthCheckType,
		config.AddRequestId))
}
