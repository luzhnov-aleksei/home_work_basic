package main

import (
	"flag"
	"log"

	client "github.com/luzhnov-aleksei/home_work_basic/hw13_http/client/pkg"
)

var (
	url  string
	path string
	data string
)

func init() {
	if flag.Lookup("url") == nil {
		flag.StringVar(&url, "url",
			"http://localhost:10001", "Server URL")
	}

	if flag.Lookup("path") == nil {
		flag.StringVar(&path, "path",
			"", "Path")
	}

	if flag.Lookup("data") == nil {
		flag.StringVar(&data, "data",
			"", "Add data to server")
	}
}

func main() {
	flag.Parse()
	client := client.NewClient(url, path)
	if len(data) == 0 {
		msg, err := client.GetData()
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(msg)
	} else {
		resp, err := client.PostData(data)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(resp)
	}
}
