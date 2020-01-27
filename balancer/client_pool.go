package balancer

import (
	"errors"
	"github.com/rs/xid"
	"log"
	"net/http"
	"net/url"
	"sync/atomic"
)

//A pool of all hosts and other configurations
type clientPool struct {
	cp                  []*client
	strategy            string
	currentReqClient    uint64
	healthCheckInterval int
	addRequestId        bool
	healthCheckType     string
}

//NewClientPool creates a new client pool
func NewClientPool(hosts []string, strategy string, healthCheckInterval int, addRequestId bool, healthCheckType string) *clientPool {
	var clients []*client
	for _, host := range hosts {
		clients = append(clients, NewClient(host))
	}
	return &clientPool{
		cp:                  clients,
		strategy:            strategy,
		healthCheckInterval: healthCheckInterval,
		currentReqClient:    0,
		addRequestId:        addRequestId,
		healthCheckType:     healthCheckType,
	}
}

//GetAvailableClient gets a new client available to forward request to
//based on strategy and health check
func (clientPool *clientPool) GetAvailableClient() (*client, error) {

	if clientPool.strategy == "roundRobin" {
		var nextClient *client
		var circularLength = len(clientPool.cp) + int(clientPool.currentReqClient)
		for i := int(clientPool.currentReqClient); i < circularLength; i++ {
			j := i % len(clientPool.cp)
			if clientPool.cp[j].getStatus() {
				nextClient = clientPool.cp[j]
				atomic.AddUint64(&clientPool.currentReqClient, 1)
				if int(clientPool.currentReqClient) == len(clientPool.cp) { //end of hosts
					clientPool.currentReqClient = 0
				}
				break
			}
		}
		return nextClient, nil
	}

	if clientPool.strategy == "leastConnections" {
		var leastConnectionClient *client
		numberOfConnection := 100000
		for _, c := range clientPool.cp {
			if int(c.requestCount) < numberOfConnection && c.getStatus() {
				numberOfConnection = int(c.requestCount)
				leastConnectionClient = c
			}
		}
		return leastConnectionClient, nil
	}

	return &client{}, errors.New("strategy not supported")
}

func (clientPool *clientPool) Director(req *http.Request) {
	client, _ := clientPool.GetAvailableClient()
	//increment the number of request the host is serving
	atomic.AddUint64(&client.requestCount, 1)
	log.Println(client)
	u, _ := url.Parse(client.host)
	//add a unique xid to request-id header
	if clientPool.addRequestId {
		guid := xid.New()
		req.Header.Set("Request-Id", guid.String())
	}

	req.URL.Scheme = u.Scheme
	req.URL.Host = u.Host
}

func (clientPool *clientPool) ModifyResponse(res *http.Response) error {
	for _, c := range clientPool.cp {
		u, _ := url.Parse(c.host)
		if u.Host == res.Request.URL.Host {
			//decrement the number of request the host is serving
			atomic.AddUint64(&c.requestCount, ^uint64(c.requestCount-1))
			log.Println(c)
			break
		}
	}
	return nil
}
