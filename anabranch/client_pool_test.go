package anabranch

import (
	"fmt"
	"testing"
)

func TestClientPool_RoundRobin_GetAvailableClient(t *testing.T) {
	var hosts = []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}
	testPool := NewClientPool(hosts, "roundRobin", 5, false)
	
	req1, _ := testPool.GetAvailableClient()
	if req1.host != "http://localhost:8080" {
		t.Errorf("Expected the host to be localhost:8080 but instead got %s!", req1.host)
	}
	req2, _ := testPool.GetAvailableClient()
	if req2.host != "http://localhost:8081" {
		t.Errorf("Expected the host to be localhost:8081 but instead got %s!", req2.host)
	}
	req3, _ := testPool.GetAvailableClient()
	if req3.host != "http://localhost:8082" {
		t.Errorf("Expected the host to be localhost:8082 but instead got %s!", req3.host)
	}
	currentReqClientAfterEnd := testPool.currentReqClient
	if currentReqClientAfterEnd != 0 {
		t.Errorf("Expected the host currentReqClient be 0 but instead got %d!", currentReqClientAfterEnd)
	}
	req4, _ := testPool.GetAvailableClient()
	if req4.host != "http://localhost:8080" {
		t.Errorf("Expected the host to be localhost:8080 but instead got %s!", req4.host)
	}
	currentReqClientAfterRestart := testPool.currentReqClient
	fmt.Println(currentReqClientAfterRestart)
	if currentReqClientAfterRestart != 1 {
		t.Errorf("Expected the host currentReqClient be 1 but instead got %d!", currentReqClientAfterRestart)
	}
}

func TestClientPool_LeastConnection_GetAvailableClient(t *testing.T) {
	var hosts = []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}
	testPool := NewClientPool(hosts, "leastConnections", 5, false)
	
	testPool.cp[0].requestCount = 5
	testPool.cp[1].requestCount = 3
	testPool.cp[2].requestCount = 4
	testPool.cp[1].status = false
	
	req1, _ := testPool.GetAvailableClient()
	if req1.host != "http://localhost:8082" {
		t.Errorf("Expected the host to be localhost:8082 but instead got %s!", req1.host)
	}
}

func TestClientPool_RoundRobin_GetAvailableClient_WithDeadClients(t *testing.T) {
	var hosts = []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}
	testPool := NewClientPool(hosts, "roundRobin", 5, false)
	
	req1, _ := testPool.GetAvailableClient()
	if req1.host != "http://localhost:8080" {
		t.Errorf("Expected the host to be localhost:8080 but instead got %s!", req1.host)
	}
	req2, _ := testPool.GetAvailableClient()
	if req2.host != "http://localhost:8081" {
		t.Errorf("Expected the host to be localhost:8081 but instead got %s!", req2.host)
	}
	testPool.cp[2].markStatus(false)
	
	req3, _ := testPool.GetAvailableClient()
	if req3.host != "http://localhost:8080" {
		t.Errorf("Expected the host to be localhost:8080 but instead got %s!", req3.host)
	}
	currentReqClientAfterEnd := testPool.currentReqClient
	if currentReqClientAfterEnd != 0 {
		t.Errorf("Expected the host currentReqClient be 0 but instead got %d!", currentReqClientAfterEnd)
	}
	testPool.cp[0].markStatus(false)
	
	req4, _ := testPool.GetAvailableClient()
	if req4.host != "http://localhost:8081" {
		t.Errorf("Expected the host to be localhost:8081 but instead got %s!", req4.host)
	}
	currentReqClientAfterRestart := testPool.currentReqClient
	fmt.Println(currentReqClientAfterRestart)
	if currentReqClientAfterRestart != 1 {
		t.Errorf("Expected the host currentReqClient be 1 but instead got %d!", currentReqClientAfterRestart)
	}
	
	req5, _ := testPool.GetAvailableClient()
	if req5.host != "http://localhost:8081" {
		t.Errorf("Expected the host to be localhost:8081 but instead got %s!", req5.host)
	}
}
