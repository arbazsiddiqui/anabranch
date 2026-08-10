package balancer

import (
	"testing"
)

func BenchmarkRoundRobin_GetAvailableClient(b *testing.B) {
	hosts := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}
	pool := NewClientPool(hosts, "roundRobin", 5, false, "none")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.GetAvailableClient()
	}
}

func BenchmarkLeastConnections_GetAvailableClient(b *testing.B) {
	hosts := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}
	pool := NewClientPool(hosts, "leastConnections", 5, false, "none")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.GetAvailableClient()
	}
}
