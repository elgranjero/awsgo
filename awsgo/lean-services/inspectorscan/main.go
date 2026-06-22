package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/inspectorscan"
)

var fields_scan_sbom = []leanruntime.Field{
	{Name: "OutputFormat", Flag: "output-format", Type: "types.OutputFormat", Required: false},
	{Name: "Sbom", Flag: "sbom", Type: "document.Interface", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"scan-sbom": {
			Name:   "scan-sbom",
			Fields: fields_scan_sbom,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ScanSbomInput{}
				if _, err := leanruntime.ApplyInput(input, fields_scan_sbom, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ScanSbom(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("inspectorscan", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
