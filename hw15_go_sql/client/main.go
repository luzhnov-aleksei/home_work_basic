package main

import (
	"flag"
	"log"

	client "github.com/luzhnov-aleksei/home_work_basic/hw15_go_sql/client/pkg"
)

var (
	url    string
	path   string
	method string
	body   string
)

func init() {
	if flag.Lookup("url") == nil {
		flag.StringVar(&url, "url",
			"http://app:8080", "Server URL")
	}
	if flag.Lookup("path") == nil {
		flag.StringVar(&path, "path",
			"", "Path")
	}
	if flag.Lookup("method") == nil {
		flag.StringVar(&method, "method",
			"", "HTTP method")
	}
	if flag.Lookup("body") == nil {
		flag.StringVar(&body, "body",
			"", "Add body into request")
	}
}

func main() {
	flag.Parse()
	client := client.NewClient(url)

	msg, err := client.GetData(path, body, method)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(msg)
}
