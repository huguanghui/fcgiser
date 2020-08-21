package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fenv := fcgi.ProcessEnv(req)
	for k, v := range fenv {
		fmt.Printf("%s :%s\n", k, v)
	}
	fmt.Println(req.Method)
	fmt.Println(req.URL.RequestURI())
	resp.Write([]byte("<h1>hello, 世界</h1>\n<p>huguanghui</p>"))
}

func main() {
	args := os.Args
	fmt.Println(args)
	fmt.Println("hello xmake!")
	listener, _ := net.Listen("tcp", "127.0.0.1:9001")
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)
}
