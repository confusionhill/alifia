package main

import (
	"alifia/internal/balancer"
	"fmt"
	"log"
	"net/http"
)

func main() {
	serve, err := balancer.NewSimpleServer("https://www.google.com")
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	servers := []balancer.Server{
		serve,
	}

	lb := balancer.NewLoadBalancer("8000", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.ServeProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at 'localhost:%s'\n", lb.GetPort())
	http.ListenAndServe(":"+lb.GetPort(), nil)
}
