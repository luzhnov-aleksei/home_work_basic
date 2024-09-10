package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type idJSON struct {
	ID string `json:"id"`
}

func Start(address string, port string) {
	url := fmt.Sprintf("%s:%s", address, port)
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(Handle)
	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%s", address, port),
		Handler:           finalHandler,
		ReadHeaderTimeout: 3 * time.Second,
	}

	mux.Handle("/", Logging(finalHandler))
	log.Printf("Listening on %s...", url)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Listening error")
	}
}

func Handle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fmt.Fprint(w, GetData(req))
	case http.MethodPost:
		defer req.Body.Close()
		resp, err := UpdateData(req)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprint(w, err)
			log.Println(err)
		}
		fmt.Fprint(w, resp)
	default:
		w.WriteHeader(405)
	}
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s %s",
			req.Method, req.UserAgent(), req.RemoteAddr, time.Since(start))
	})
}

func UpdateData(req *http.Request) (string, error) {
	var data []idJSON
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return "", err
	}

	var builder strings.Builder

	for _, v := range data {
		builder.WriteString(v.ID)
		builder.WriteString("; ")
	}

	return builder.String(), nil
}

func GetData(req *http.Request) string {
	return fmt.Sprintf("Request method: %s; RequestURI: %s", req.Method, req.RequestURI)
}
