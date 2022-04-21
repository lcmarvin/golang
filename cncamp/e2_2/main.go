package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":80", MyServer{})
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
}

type MyServer struct {
}

func (server MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		for _, m := range v {
			w.Header().Add(k, m)
		}
	}

	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	fmt.Printf("IP:%s, Status:%d", strings.Split(r.RemoteAddr, ":")[0], http.StatusOK)
}
