package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/sagemakermetrics"
)

var fields_batch_get_metrics = []leanruntime.Field{
	{Name: "MetricQueries", Flag: "metric-queries", Type: "[]types.MetricQuery", Required: true},
}

var fields_batch_put_metrics = []leanruntime.Field{
	{Name: "MetricData", Flag: "metric-data", Type: "[]types.RawMetricData", Required: true},
	{Name: "TrialComponentName", Flag: "trial-component-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-metrics": {
			Name:   "batch-get-metrics",
			Fields: fields_batch_get_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetMetricsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_metrics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetMetrics(ctx, input)
			},
		},
		"batch-put-metrics": {
			Name:   "batch-put-metrics",
			Fields: fields_batch_put_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchPutMetricsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_put_metrics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchPutMetrics(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("sagemakermetrics", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
