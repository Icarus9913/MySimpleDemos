package main

import (
	"context"
	"log"
	"net/http"
	"simplePrometheus/demos/tool"

	"contrib.go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
)

var filHeight = stats.Int64("filecoin_height", "current height of filecoin", stats.UnitDimensionless)

var (
	FilHeightView = &view.View{
		Name:        "filecoin_height",
		Description: "current height of filecoin",
		Measure:     filHeight,
		Aggregation: view.LastValue(),
	}
)

func main() {
	ctx := context.Background()

	exporter, err := prometheus.NewExporter(prometheus.Options{})
	if nil != err {
		panic(err)
	}

	err = view.Register(FilHeightView)
	if nil != err {
		panic(err)
	}
	go func() {
		for {
			stats.Record(ctx, filHeight.M(int64(tool.Height())))
		}
	}()

	http.Handle("/metrics", exporter)
	log.Fatal(http.ListenAndServe(":8989", nil))
}
