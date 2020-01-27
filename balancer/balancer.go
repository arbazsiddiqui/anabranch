package balancer

import (
	"net/http/httputil"
)

type LB struct {
	*httputil.ReverseProxy
	cp *clientPool
}

func NewLoadBalancer(hosts []string, strategy string, healthCheckInterval int, healthCheckType string, addRequestId bool) *LB {
	cp := NewClientPool(hosts, strategy, healthCheckInterval, addRequestId)
	lb := &LB{
		cp: cp,
	}
	lb.ReverseProxy = &httputil.ReverseProxy{
		Director:       cp.Director,
		ModifyResponse: cp.ModifyResponse,
	}
	if healthCheckType == "passive" {
		go lb.StartPassiveHeathCheck()
	}
	return lb
}
