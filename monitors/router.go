package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	requestDuration = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "router_function_duration_seconds",
		Help: "Router calls duration",
	})
)

func main() {
	var requests int

	prometheus.MustRegister(requestDuration)

	port := os.Getenv("ROUTER_METRICS_SERVICE_SERVICE_PORT")
	host := os.Getenv("ROUTER_METRICS_SERVICE_SERVICE_HOST")
	reqs := os.Getenv("ROUTER_METRICS_SERVICE_SERVICE_REQUESTS")
	interval := os.Getenv("ROUTER_METRICS_SERVICE_SERVICE_REQUESTS_INTERVAL")

	if port == "" {
		port = "8000"
	}

	if host == "" {
		host = "localhost"
	}

	if reqs == "" {
		reqs = "10"
	}

	if interval == "" {
		interval = "5"
	}

	requests, _ = strconv.Atoi(reqs)
	requestInterval, _ := strconv.Atoi(interval)

	fmt.Printf("Starting server on: %d:%d\nRequests: %s\nInterval: %s\n", host, port, reqs, interval)

	// Use environment instead of arguments
	if port != "" && host != "" {
		host = host + ":" + port
	}

	fmt.Printf("Testing connection trhough %s\n", host)
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())

	// Request ticker
	ticker := time.NewTicker(time.Duration(requestInterval) * time.Second)
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			handleRequests(requests, host)
		}
	}()

	http.ListenAndServe(host, nil)
}

func handleRequests(reqs int, host string) {
	hostname := "http://" + host
	for i := 0; i <= reqs; i++ {
		if request(hostname) {
			fmt.Printf("(%d) Successful request\n", i)
		} else {
			fmt.Printf("(%d) Error in request\n", i)
		}
	}
}

// Response handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	log.Println("Test URL /")
}

// Perform Request
func request(host string) bool {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(requestDuration.Set))
	defer timer.ObserveDuration()
	resp, err := http.Get(host)
	if err != nil {
		log.Printf("Error requesting: %s, %s\n", host, err)
		return false
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
	return true
}
