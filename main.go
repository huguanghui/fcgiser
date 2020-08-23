package main

import (
	"net"
	"net/http"
	"net/http/fcgi"

	"github.com/huguanghui/fcgiser/api"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	//	fenv := fcgi.ProcessEnv(req)
	//	for k, v := range fenv {
	//		fmt.Printf("%s :%s\n", k, v)
	//	}
	resp.Header().Set("Content-Type", "text/json;charset=utf-8")
	opt := api.CmdRuner.Lookup(req.URL.RequestURI())
	if opt != nil {
		fun := opt.Lookup(req.Method)
		if fun != nil {
			fun.Opt(resp, req)
		}
	}
}

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:9001")
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)
}
