package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/sagemakermetrics/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-get-metrics", "batch-put-metrics"},
		OperationSet: map[string]bool{"batch-get-metrics": true, "batch-put-metrics": true},
		OperationInputs: map[string][]string{
			"batch-get-metrics": {"MetricQueries"},
			"batch-put-metrics": {"MetricData", "TrialComponentName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-get-metrics": {"MetricQueries": "[]types.MetricQuery"},
			"batch-put-metrics": {"MetricData": "[]types.RawMetricData", "TrialComponentName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"batch-get-metrics": {"MetricQueries"},
			"batch-put-metrics": {"MetricData", "TrialComponentName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("sagemakermetrics", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
