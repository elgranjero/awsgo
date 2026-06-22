package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/sagemakerruntime"
)

var fields_invoke_endpoint = []leanruntime.Field{
	{Name: "Accept", Flag: "accept", Type: "*string", Required: false},
	{Name: "Body", Flag: "body", Type: "[]byte", Required: true},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "CustomAttributes", Flag: "custom-attributes", Type: "*string", Required: false},
	{Name: "EnableExplanations", Flag: "enable-explanations", Type: "*string", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "InferenceComponentName", Flag: "inference-component-name", Type: "*string", Required: false},
	{Name: "InferenceId", Flag: "inference-id", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
	{Name: "TargetContainerHostname", Flag: "target-container-hostname", Type: "*string", Required: false},
	{Name: "TargetModel", Flag: "target-model", Type: "*string", Required: false},
	{Name: "TargetVariant", Flag: "target-variant", Type: "*string", Required: false},
}

var fields_invoke_endpoint_async = []leanruntime.Field{
	{Name: "Accept", Flag: "accept", Type: "*string", Required: false},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "CustomAttributes", Flag: "custom-attributes", Type: "*string", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "Filename", Flag: "filename", Type: "*string", Required: false},
	{Name: "InferenceId", Flag: "inference-id", Type: "*string", Required: false},
	{Name: "InputLocation", Flag: "input-location", Type: "*string", Required: true},
	{Name: "InvocationTimeoutSeconds", Flag: "invocation-timeout-seconds", Type: "*int32", Required: false},
	{Name: "RequestTTLSeconds", Flag: "request-ttl-seconds", Type: "*int32", Required: false},
	{Name: "S3OutputPathExtension", Flag: "s3-output-path-extension", Type: "*string", Required: false},
}

var fields_invoke_endpoint_with_response_stream = []leanruntime.Field{
	{Name: "Accept", Flag: "accept", Type: "*string", Required: false},
	{Name: "Body", Flag: "body", Type: "[]byte", Required: true},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "CustomAttributes", Flag: "custom-attributes", Type: "*string", Required: false},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "InferenceComponentName", Flag: "inference-component-name", Type: "*string", Required: false},
	{Name: "InferenceId", Flag: "inference-id", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
	{Name: "TargetContainerHostname", Flag: "target-container-hostname", Type: "*string", Required: false},
	{Name: "TargetVariant", Flag: "target-variant", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"invoke-endpoint": {
			Name:   "invoke-endpoint",
			Fields: fields_invoke_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeEndpoint(ctx, input)
			},
		},
		"invoke-endpoint-async": {
			Name:   "invoke-endpoint-async",
			Fields: fields_invoke_endpoint_async,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeEndpointAsyncInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_endpoint_async, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeEndpointAsync(ctx, input)
			},
		},
		"invoke-endpoint-with-response-stream": {
			Name:   "invoke-endpoint-with-response-stream",
			Fields: fields_invoke_endpoint_with_response_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeEndpointWithResponseStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_endpoint_with_response_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeEndpointWithResponseStream(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("sagemakerruntime", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
