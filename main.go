package main

import (
	"anabranch/balancer"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type config struct {
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

	var config config
	_ = json.NewDecoder(file).Decode(&config)
	fmt.Println("Anabranch running on port :", config.Port)
	//send the parsed config to create a new load balancer
	http.ListenAndServe(fmt.Sprintf(":%v", config.Port), balancer.NewLoadBalancer(
		config.Hosts,
		config.Strategy,
		config.HealthCheckInterval,
		config.HealthCheckType,
		config.AddRequestId))
}
