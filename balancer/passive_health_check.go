package balancer

import (
	"log"
	"time"
)

func (clientPool *clientPool) passiveHeathCheck() {
	for _, c := range clientPool.cp {
		log.Println("Health check for host :", c.host)
		go c.markStatus(c.isAlive())
	}
}

func (lb *LB) StartPassiveHeathCheck() {
	t := time.NewTicker(time.Second * time.Duration(lb.cp.healthCheckInterval))
	for {
		select {
		case <-t.C:
			log.Println("Starting passive health check...")
			lb.cp.passiveHeathCheck()
			log.Println("Passive Health check completed")
		}
	}
}
