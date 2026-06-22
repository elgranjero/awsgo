package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/forecastquery"
)

var fields_query_forecast = []leanruntime.Field{
	{Name: "EndDate", Flag: "end-date", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "map[string]string", Required: true},
	{Name: "ForecastArn", Flag: "forecast-arn", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartDate", Flag: "start-date", Type: "*string", Required: false},
}

var fields_query_what_if_forecast = []leanruntime.Field{
	{Name: "EndDate", Flag: "end-date", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "map[string]string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartDate", Flag: "start-date", Type: "*string", Required: false},
	{Name: "WhatIfForecastArn", Flag: "what-if-forecast-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"query-forecast": {
			Name:   "query-forecast",
			Fields: fields_query_forecast,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.QueryForecastInput{}
				if _, err := leanruntime.ApplyInput(input, fields_query_forecast, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.QueryForecast(ctx, input)
			},
		},
		"query-what-if-forecast": {
			Name:   "query-what-if-forecast",
			Fields: fields_query_what_if_forecast,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.QueryWhatIfForecastInput{}
				if _, err := leanruntime.ApplyInput(input, fields_query_what_if_forecast, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.QueryWhatIfForecast(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("forecastquery", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
