package main

import (
	"flag"
	"net/http"

	"github.com/elephanter/nginx_exporter/nginx_export"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/log"
)

const (
	namespace = "nginx" // For Prometheus metrics.
)

var (
	listeningAddress = flag.String("telemetry.address", ":9113", "Address on which to expose metrics.")
	metricsEndpoint  = flag.String("telemetry.endpoint", "/metrics", "Path under which to expose metrics.")
	nginxScrapeURI   = flag.String("nginx.scrape_uri", "http://localhost/nginx_status", "URI to nginx stub status page")
	insecure         = flag.Bool("insecure", true, "Ignore server certificate if using https")
)

func main() {
	flag.Parse()

	exporter := nginx_export.NewExporter(*nginxScrapeURI, *insecure)
	prometheus.MustRegister(exporter)

	log.Printf("Starting Server: %s", *listeningAddress)
	http.Handle(*metricsEndpoint, prometheus.Handler())
	log.Fatal(http.ListenAndServe(*listeningAddress, nil))
}
