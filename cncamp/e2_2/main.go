package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", copyHeaders)
	http.HandleFunc("/healthz/", healthz)
	http.ListenAndServe(":80", nil)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	copyHeaders(w, r)
	io.WriteString(w, "200")
}

func copyHeaders(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		for _, m := range v {
			w.Header().Add(k, m)
		}
	}

	const VERSION = "VERSION"
	version := os.Getenv(VERSION)
	w.Header().Set(VERSION, version)

	fmt.Printf("RemoteAddress:%s, Status:%d", r.RemoteAddr, http.StatusOK)
}
