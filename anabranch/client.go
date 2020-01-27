package balancer

import (
	"log"
	"net"
	"net/url"
	"sync"
	"time"
)

type client struct {
	host         string
	status       bool
	requestCount uint64
}

var mutex = &sync.RWMutex{}

func NewClient(host string) *client {
	newClient := client{
		host:         host,
		status:       true,
		requestCount: 0,
	}
	return &newClient
}

func (c *client) isAlive() bool {
	u, _ := url.Parse(c.host)
	conn, err := net.DialTimeout("tcp", u.Host, 2*time.Second)
	if err != nil {
		log.Println("client is unreachable: ", err)
		return false
	}
	log.Println("client is up: ", c.host)
	_ = conn.Close()
	return true
}

func (c *client) markStatus(status bool) {
	mutex.Lock()
	c.status = status
	mutex.Unlock()
}

func (c *client) getStatus() bool {
	mutex.RLock()
	status := c.status
	mutex.RUnlock()
	return status
}
