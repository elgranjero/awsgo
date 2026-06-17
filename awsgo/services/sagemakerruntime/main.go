package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/sagemakerruntime/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"invoke-endpoint", "invoke-endpoint-async", "invoke-endpoint-with-response-stream"},
		OperationSet: map[string]bool{"invoke-endpoint": true, "invoke-endpoint-async": true, "invoke-endpoint-with-response-stream": true},
		OperationInputs: map[string][]string{
			"invoke-endpoint":                      {"Accept", "Body", "ContentType", "CustomAttributes", "EnableExplanations", "EndpointName", "InferenceComponentName", "InferenceId", "SessionId", "TargetContainerHostname", "TargetModel", "TargetVariant"},
			"invoke-endpoint-async":                {"Accept", "ContentType", "CustomAttributes", "EndpointName", "Filename", "InferenceId", "InputLocation", "InvocationTimeoutSeconds", "RequestTTLSeconds", "S3OutputPathExtension"},
			"invoke-endpoint-with-response-stream": {"Accept", "Body", "ContentType", "CustomAttributes", "EndpointName", "InferenceComponentName", "InferenceId", "SessionId", "TargetContainerHostname", "TargetVariant"},
		},
		OperationInputTypes: map[string]map[string]string{
			"invoke-endpoint":                      {"Accept": "*string", "Body": "[]byte", "ContentType": "*string", "CustomAttributes": "*string", "EnableExplanations": "*string", "EndpointName": "*string", "InferenceComponentName": "*string", "InferenceId": "*string", "SessionId": "*string", "TargetContainerHostname": "*string", "TargetModel": "*string", "TargetVariant": "*string"},
			"invoke-endpoint-async":                {"Accept": "*string", "ContentType": "*string", "CustomAttributes": "*string", "EndpointName": "*string", "Filename": "*string", "InferenceId": "*string", "InputLocation": "*string", "InvocationTimeoutSeconds": "*int32", "RequestTTLSeconds": "*int32", "S3OutputPathExtension": "*string"},
			"invoke-endpoint-with-response-stream": {"Accept": "*string", "Body": "[]byte", "ContentType": "*string", "CustomAttributes": "*string", "EndpointName": "*string", "InferenceComponentName": "*string", "InferenceId": "*string", "SessionId": "*string", "TargetContainerHostname": "*string", "TargetVariant": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"invoke-endpoint":                      {"Body", "EndpointName"},
			"invoke-endpoint-async":                {"EndpointName", "InputLocation"},
			"invoke-endpoint-with-response-stream": {"Body", "EndpointName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("sagemakerruntime", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
