package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/forecastquery/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"query-forecast", "query-what-if-forecast"},
		OperationSet: map[string]bool{"query-forecast": true, "query-what-if-forecast": true},
		OperationInputs: map[string][]string{
			"query-forecast":         {"EndDate", "Filters", "ForecastArn", "NextToken", "StartDate"},
			"query-what-if-forecast": {"EndDate", "Filters", "NextToken", "StartDate", "WhatIfForecastArn"},
		},
		OperationInputTypes: map[string]map[string]string{
			"query-forecast":         {"EndDate": "*string", "Filters": "map[string]string", "ForecastArn": "*string", "NextToken": "*string", "StartDate": "*string"},
			"query-what-if-forecast": {"EndDate": "*string", "Filters": "map[string]string", "NextToken": "*string", "StartDate": "*string", "WhatIfForecastArn": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"query-forecast":         {"Filters", "ForecastArn"},
			"query-what-if-forecast": {"Filters", "WhatIfForecastArn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("forecastquery", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
