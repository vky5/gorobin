package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/vky5/gorobin/internal/balancer"
)


type Proxy struct {
	balancer *balancer.RoundRobin // we are storing the info of the RoundRobin in balancer
}

func NewProxy(rr *balancer.RoundRobin) *Proxy {
	return &Proxy{balancer: rr}
}

func (p *Proxy) Handler(w http.ResponseWriter, r *http.Request){
	target:=p.balancer.Next() // it will get the nex server
	
	targetUrl, err := url.Parse(target)

	if err!=nil{
		http.Error(w, "Invalid target url", http.StatusInternalServerError)
		return
	}


	proxy := httputil.NewSingleHostReverseProxy(targetUrl) // Create a reverse proxy for the selected backend server
	
	proxy.ServeHTTP(w, r) // Forward the incoming request to the target server and write the response back to the client
	// this handles sending the request to the targetURL and then sending the response back to the client automatically
}


/*
Purpose
- Accept incoming http request
- forwardd them to the next server from the round robin
- return the response to the original client (basically like a reverse proxy here)
*/




