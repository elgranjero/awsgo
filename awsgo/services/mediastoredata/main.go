package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/mediastoredata/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"delete-object", "describe-object", "get-object", "list-items", "put-object"},
		OperationSet: map[string]bool{"delete-object": true, "describe-object": true, "get-object": true, "list-items": true, "put-object": true},
		OperationInputs: map[string][]string{
			"delete-object":   {"Path"},
			"describe-object": {"Path"},
			"get-object":      {"Path", "Range"},
			"list-items":      {"MaxResults", "NextToken", "Path"},
			"put-object":      {"Body", "CacheControl", "ContentType", "Path", "StorageClass", "UploadAvailability"},
		},
		OperationInputTypes: map[string]map[string]string{
			"delete-object":   {"Path": "*string"},
			"describe-object": {"Path": "*string"},
			"get-object":      {"Path": "*string", "Range": "*string"},
			"list-items":      {"MaxResults": "*int32", "NextToken": "*string", "Path": "*string"},
			"put-object":      {"Body": "io.Reader", "CacheControl": "*string", "ContentType": "*string", "Path": "*string", "StorageClass": "types.StorageClass", "UploadAvailability": "types.UploadAvailability"},
		},
		OperationInputRequired: map[string][]string{
			"delete-object":   {"Path"},
			"describe-object": {"Path"},
			"get-object":      {"Path"},
			"list-items":      {},
			"put-object":      {"Body", "Path"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("mediastoredata", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
