package main

import (
	"context"
	"log"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

type customMetricsExporter struct{}

func (ce *customMetricsExporter) ExportView(vd *view.Data) {
	log.Printf("vd.View: %+v\n%#v\n", vd.View, vd.Rows)
	for i, row := range vd.Rows {
		log.Printf("\tRow: %#d: %#v\n", i, row)
	}
	log.Printf("StrtTime: %s EndTime: %s\n\n", vd.Start.Round(0), vd.End.Round(0))
}

var keyMethod, _ = tag.NewKey("method")
var mLoops = stats.Int64("demo/loop_iterations", "The number of loop iterations", "1")
var loopCountView = &view.View{
	Measure: mLoops,
	Name: "demo/loop_iterations",
	Description: "Number of loop iterations",
	Aggregation: view.Count(),
	TagKeys:     []tag.Key{keyMethod},
}

func main() {
	log.SetFlags(0)
	if err := view.Register(loopCountView); nil != err {
		log.Fatalf("Failed to register loopCountView: %v", err)
	}
	view.RegisterExporter(new(customMetricsExporter))
	view.SetReportingPeriod(100 * time.Millisecond)
	ctx, _ := tag.New(context.Background(), tag.Upsert(keyMethod, "main"))
	for i := int64(0); i < 5; i++ {
		stats.Record(ctx, mLoops.M(i))
		<-time.After(10 * time.Millisecond)
	}
	<-time.After(500*time.Millisecond)
}
