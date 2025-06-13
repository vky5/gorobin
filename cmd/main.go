package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vky5/gorobin/internal/balancer"
	"github.com/vky5/gorobin/internal/config"
	"github.com/vky5/gorobin/internal/proxy"
)

func main() {
	// getting the list of the backend servers from the config file

	cfg, err := config.LoadConfig("../config.yaml")

	// hamdle the error case
	if err!=nil{
		log.Fatalf("failed to load the config: %v", err)
	}

	// creating the round robin balancer
	rr := balancer.NewRoundRobin(cfg.Servers)

	// creating a reverse proxy using the balancer
	proxyHandler := proxy.NewProxy(rr)

	// set up the http handler
	http.HandleFunc("/", proxyHandler.Handler) // whenever the load balancer receive the request it will run the handler func

	// start the server
	fmt.Println(("Load balancer running at http://localhost:5000"))

	if err := http.ListenAndServe(cfg.Port, nil); err != nil {
		log.Fatalf("Server failed %v", err)
	}

}
