package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/eksauth"
)

var fields_assume_role_for_pod_identity = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Token", Flag: "token", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"assume-role-for-pod-identity": {
			Name:   "assume-role-for-pod-identity",
			Fields: fields_assume_role_for_pod_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssumeRoleForPodIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assume_role_for_pod_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssumeRoleForPodIdentity(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("eksauth", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
