package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/marketplacemetering"
)

var fields_batch_meter_usage = []leanruntime.Field{
	{Name: "ProductCode", Flag: "product-code", Type: "*string", Required: false},
	{Name: "UsageRecords", Flag: "usage-records", Type: "[]types.UsageRecord", Required: true},
}

var fields_meter_usage = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ProductCode", Flag: "product-code", Type: "*string", Required: true},
	{Name: "Timestamp", Flag: "timestamp", Type: "*time.Time", Required: true},
	{Name: "UsageAllocations", Flag: "usage-allocations", Type: "[]types.UsageAllocation", Required: false},
	{Name: "UsageDimension", Flag: "usage-dimension", Type: "*string", Required: true},
	{Name: "UsageQuantity", Flag: "usage-quantity", Type: "*int32", Required: false},
}

var fields_register_usage = []leanruntime.Field{
	{Name: "Nonce", Flag: "nonce", Type: "*string", Required: false},
	{Name: "ProductCode", Flag: "product-code", Type: "*string", Required: true},
	{Name: "PublicKeyVersion", Flag: "public-key-version", Type: "*int32", Required: true},
}

var fields_resolve_customer = []leanruntime.Field{
	{Name: "RegistrationToken", Flag: "registration-token", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-meter-usage": {
			Name:   "batch-meter-usage",
			Fields: fields_batch_meter_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchMeterUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_meter_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchMeterUsage(ctx, input)
			},
		},
		"meter-usage": {
			Name:   "meter-usage",
			Fields: fields_meter_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MeterUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_meter_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MeterUsage(ctx, input)
			},
		},
		"register-usage": {
			Name:   "register-usage",
			Fields: fields_register_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterUsage(ctx, input)
			},
		},
		"resolve-customer": {
			Name:   "resolve-customer",
			Fields: fields_resolve_customer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResolveCustomerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resolve_customer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResolveCustomer(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("marketplacemetering", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
