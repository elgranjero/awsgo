package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/signerdata"
)

var fields_get_revocation_status = []leanruntime.Field{
	{Name: "CertificateHashes", Flag: "certificate-hashes", Type: "[]string", Required: true},
	{Name: "JobArn", Flag: "job-arn", Type: "*string", Required: true},
	{Name: "PlatformId", Flag: "platform-id", Type: "*string", Required: true},
	{Name: "ProfileVersionArn", Flag: "profile-version-arn", Type: "*string", Required: true},
	{Name: "SignatureTimestamp", Flag: "signature-timestamp", Type: "*time.Time", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-revocation-status": {
			Name:   "get-revocation-status",
			Fields: fields_get_revocation_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRevocationStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_revocation_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRevocationStatus(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("signerdata", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
