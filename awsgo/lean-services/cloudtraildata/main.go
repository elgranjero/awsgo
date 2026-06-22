package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cloudtraildata"
)

var fields_put_audit_events = []leanruntime.Field{
	{Name: "AuditEvents", Flag: "audit-events", Type: "[]types.AuditEvent", Required: true},
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ExternalId", Flag: "external-id", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"put-audit-events": {
			Name:   "put-audit-events",
			Fields: fields_put_audit_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAuditEventsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_audit_events, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAuditEvents(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cloudtraildata", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
