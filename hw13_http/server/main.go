package main

import (
	"flag"

	server "github.com/luzhnov-aleksei/home_work_basic/hw13_http/server/pkg"
)

var (
	address string
	port    string
)

func init() {
	if flag.Lookup("address") == nil {
		flag.StringVar(&address, "address",
			"", "Server address")
	}

	if flag.Lookup("port") == nil {
		flag.StringVar(&port, "port",
			"10001", "Server port")
	}
}

func main() {
	flag.Parse()
	server.Start(address, port)
}
