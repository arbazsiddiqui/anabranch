package balancer

import (
	"net/http/httputil"
)

//load balancer implemented on top of built in httpUtil.ReverseProxy
type lb struct {
	*httputil.ReverseProxy
	cp *clientPool
}

//NewLoadBalancer : Creates a new LoadBalancer
func NewLoadBalancer(hosts []string, strategy string, healthCheckInterval int, healthCheckType string, addRequestId bool) *lb {
	cp := NewClientPool(hosts, strategy, healthCheckInterval, addRequestId, healthCheckType)
	lb := &lb{
		cp: cp,
	}
	lb.ReverseProxy = &httputil.ReverseProxy{
		//this function is called before sending request to downstream services
		Director: cp.Director,
		//this function is called before sending the response back to upstream services
		ModifyResponse: cp.ModifyResponse,
	}
	if healthCheckType == "passive" {
		//starts passive health check if opted
		go lb.StartPassiveHeathCheck()
	}
	return lb
}
