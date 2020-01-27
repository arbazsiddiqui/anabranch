package balancer

import (
	"net/http/httputil"
)

type lb struct {
	*httputil.ReverseProxy
	cp *clientPool
}

//NewLoadBalancer : Creates a new LoadBalancer
func NewLoadBalancer(hosts []string, strategy string, healthCheckInterval int, healthCheckType string, addRequestId bool) *lb {
	cp := NewClientPool(hosts, strategy, healthCheckInterval, addRequestId)
	lb := &lb{
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
