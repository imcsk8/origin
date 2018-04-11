package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	hostflag := flag.String("host", "localhost:8000", "Route to connect")
	flag.IntVar(&requests, "requests", 10, "Number of requests to perform")
	flag.IntVar(&requests, "r", 10, "Number of requests to perform")
	flag.Parse()
	fmt.Printf("Testing connection trhough %s\n", *hostflag)
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())

	go http.ListenAndServe("localhost:8000", nil)
	for i := 0; i <= requests; i++ {
		if request("http://" + *hostflag) {
			fmt.Printf("(%d) Successful request\n", i)
		} else {
			fmt.Printf("(%d) Error in request\n", i)
		}
	}
	request("http://" + *hostflag + "/metrics")

}

// Response handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	log.Println("Test URL /")
}

// Request
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
