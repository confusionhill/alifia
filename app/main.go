package main

import (
	"alifia/internal/balancer"
	"fmt"
	"log"
	"net/http"
)

func main() {
	rsc, err := loadResources()
	if err != nil {
		log.Fatal(err)
		return
	}
	serve, err := balancer.NewSimpleServer("https://be-staging-b6utdt2kwa-et.a.run.app")
	if err != nil {
		log.Fatal(err)
		return
	}
	servers := []balancer.Server{
		serve,
	}
	lb := balancer.NewLoadBalancer(rsc.db, "8000", servers)
	if lb == nil {
		fmt.Println("kok nil sih bgst!")
	}
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("sampe sini bosq")
		lb.ServeProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)
	fmt.Printf("server is running at :%s\n", lb.GetPort())
	http.ListenAndServe(":"+lb.GetPort(), nil)
}
