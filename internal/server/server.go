package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ServerInterface interface {
	// Address returns the address with which to access the server
	Address() string

	// IsAlive returns true if the server is alive and able to serve requests
	IsAlive() bool

	// Serve uses this server to process the request
	Serve(rw http.ResponseWriter, req *http.Request)
}

type Server struct {
	adress string
	proxy  *httputil.ReverseProxy
}

func NewServer(addr string) (*Server, error) {
	serverUrl, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	return &Server{
		adress: addr,
		proxy:  httputil.NewSingleHostReverseProxy(serverUrl),
	}, nil
}
