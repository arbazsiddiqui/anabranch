# anabranch
[![Go Report Card](https://goreportcard.com/badge/github.com/arbazsiddiqui/anabranch)](https://goreportcard.com/report/github.com/arbazsiddiqui/anabranch)

> A simple HTTP load balancer and reverse proxy written in Go.

## Features
* Round Robin
* Least Connection
* Passive health check at configurable interval
* Request Id injection
* Config support using json



## Usage
1) Clone the repo using `git clone git@github.com:arbazsiddiqui/anabranch.git && cd anabranch`
2) Use the existing binary `anabranch` or build a new one using `go build`
3) Change the existing config to your needs to create a new one.
4) Run the load balancer using `./anabranch -config=./config.json`


## Config

### port
The port to the run the load balancer on.

### hosts
A list of hosts forwards requests too.

### strategy
Algorithm for load balancing. Can be `roundRobin` or `leastConnections`.

### health_check_type 
Defines the type of health check you want to perform on your host. Can be `passive` or `none` (for no health check). Active health check is WIP.

### health_check_interval
Time interval in seconds to perform health check on hosts.

### add_request_id
Boolean to add a unique request id in header of every request to downstream hosts. Header name : `Request-Id`

## Todo 
- [ ] Add support for retires 
- [ ] More test coverage 
- [ ] Add support for active health check
