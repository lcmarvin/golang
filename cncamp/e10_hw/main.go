package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 定义指标
var m prometheus.Counter

func main() {
	// 创建指标
	m = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "total_delay_millis",
		Help: "The total milliseconds delayed",
	})
	// 注册指标
	prometheus.MustRegister(m)

	http.HandleFunc("/", copyHeaders)
	http.HandleFunc("/healthz/", healthz)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":80", nil)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	// 生成随机数睡眠
	n := rand.Intn(2000)
	d := time.Duration(n) * time.Millisecond
	fmt.Printf("Sleeping %s\n", d)
	time.Sleep(d)

	// 为指标设置值
	m.Add(float64(n))

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

	fmt.Printf("RemoteAddress:%s, Status:%d\n", r.RemoteAddr, http.StatusOK)
}
