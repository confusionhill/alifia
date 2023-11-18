package balancer

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server interface {
	// Address returns the address with which to access the server
	Address() string

	// IsAlive returns true if the server is alive and able to serve requests
	IsAlive() bool

	// Serve uses this server to process the request
	Serve(rw http.ResponseWriter, req *http.Request)
}

type SimpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func (s *SimpleServer) Address() string { return s.addr }

func (s *SimpleServer) IsAlive() bool { return true }

func (s *SimpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

func NewSimpleServer(addr string) (*SimpleServer, error) {
	serverUrl, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	return &SimpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}, nil
}
