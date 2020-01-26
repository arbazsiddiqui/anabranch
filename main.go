package main

import (
	"lb/balancer"
	"net/http"
)

func main() {
	var hosts = []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}
	http.ListenAndServe(":9001", balancer.NewLoadBalancer(hosts, "roundRobin", 5, "none", true))
}
