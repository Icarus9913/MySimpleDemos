package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var addr = flag.String("listen-address", ":9001", "The address to listen on for HTTP requests.")

var rpcDurations = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "rpc_durations_seconds",
		Help:       "RPC latency distributions.",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"service"},
)

var up = prometheus.NewDesc(
	"ldap_up",
	"watch ladap info",
	[]string{"name", "class"},
	nil,
)

var rpc = prometheus.NewDesc(
	"rpc_request_total",
	"rpc request total",
	[]string{"rpc_func", "rpc_source"},
	nil,
)

var version = prometheus.MustNewConstMetric(
	prometheus.NewDesc("ldap_exporter_version", "ldap exporter version", []string{"type", "build"}, nil),
	prometheus.GaugeValue,
	1.0,
	"cust", "2020-09-08",
)

// MyLDAPCollector myLDAPCollector
type MyLDAPCollector struct {
	up      *prometheus.Desc	//标记exporter运行
	rpc     *prometheus.Desc	//一个简单的total模式的metrics定义
	version *prometheus.Desc	//记录exporter版本
}

//Describe describe
func (c MyLDAPCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
	ch <- c.rpc
	ch <- c.version
}

//Collect collect
func (c MyLDAPCollector) Collect(ch chan<- prometheus.Metric) {
	upMetrics := prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 1, "dalong", "demo")
	rpcRequest := prometheus.MustNewConstMetric(rpc, prometheus.CounterValue, 1000, "login", "a")
	ch <- version
	ch <- upMetrics
	ch <- rpcRequest
}

func init() {
	myLDAPCollector := MyLDAPCollector{
		up:      up,
		rpc:     rpc,
		version: version.Desc(),
	}
	// Add MyCus Collector
	prometheus.MustRegister(myLDAPCollector)
	// Add rpc summary collector
	prometheus.MustRegister(rpcDurations)
	// Add Go module build info.
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
}

func main() {
	flag.Parse()
	server := http.NewServeMux()
	server.HandleFunc("/", func(response http.ResponseWriter, Request *http.Request) {
		indexpage := `<html>
                <body>
                   <h1>ldap exporter</h1>
                   <p><a href="/metrics">metrics</a></p>
                </body>
          </html>`
		response.Write([]byte(indexpage))
	})
	server.HandleFunc("/api", func(response http.ResponseWriter, Request *http.Request) {
		rpcDurations.WithLabelValues("mydemo").Observe(22)
		response.Write([]byte("dalongdemo"))
	})
	server.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(*addr, server)
}
