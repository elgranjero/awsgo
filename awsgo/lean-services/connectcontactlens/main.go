package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/connectcontactlens"
)

var fields_list_realtime_contact_analysis_segments = []leanruntime.Field{
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"list-realtime-contact-analysis-segments": {
			Name:   "list-realtime-contact-analysis-segments",
			Fields: fields_list_realtime_contact_analysis_segments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRealtimeContactAnalysisSegmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_realtime_contact_analysis_segments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRealtimeContactAnalysisSegments(ctx, input)
				}
				var results []*svc.ListRealtimeContactAnalysisSegmentsOutput
				p := svc.NewListRealtimeContactAnalysisSegmentsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
	}
	if err := leanruntime.Execute("connectcontactlens", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
