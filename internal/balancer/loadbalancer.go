package balancer

import (
	"database/sql"
	"fmt"
	"net/http"
)

type LoadBalancer struct {
	db              *sql.DB
	port            string
	roundRobinCount int
	servers         []Server
}

func NewLoadBalancer(db *sql.DB, port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func (lb *LoadBalancer) ChangeServers(servers []Server) {
	lb.servers = servers
}

func (lb *LoadBalancer) GetPort() string {
	return lb.port
}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++

	return server
}

func (lb *LoadBalancer) ServeProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.servers[0] //lb.getNextAvailableServer()
	fmt.Printf("forwarding request to address %q\n", targetServer.Address())

	// could delete pre-existing X-Forwarded-For header to prevent IP spoofing
	targetServer.Serve(rw, req)
}
