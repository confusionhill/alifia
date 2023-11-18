package main

import (
	"alifia/internal/balancer"
	"fmt"
	"log"
	"net/http"
)

func main() {
	serve, err := balancer.NewSimpleServer("https://be-staging-b6utdt2kwa-et.a.run.app")
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	servers := []balancer.Server{
		serve,
	}
	lb := balancer.NewLoadBalancer("8000", servers)
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
